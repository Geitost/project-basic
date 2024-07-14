package api

import (
	"fmt"
	"mime/multipart"
	"sort"

	"github.com/gofiber/fiber/v2"
	ort "github.com/yalue/onnxruntime_go"
	"www.autoblox.xyz/server/structs"
	"www.autoblox.xyz/server/utils"
)

var classes = []string{
	"bef",
	"vef",
	"let",
	"tom",
	"che",
	"oni",
	"si-fri",
	"si-moz",
	"si-oni",
	"dr-sod",
	"dr-app",
	"dr-sha",
}

var infoClasses = []string{
	"1",
	"2",
	"l",
	"m",
	"s",
}

func BloxburgCashierRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// -------------
		//	Get screens
		// -------------
		screen1, err := ctx.FormFile("screen1")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid, screen1 please try again",
			})
		}

		screen2, err := ctx.FormFile("screen2")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid, screen2 please try again",
			})
		}

		screen3, err := ctx.FormFile("screen3")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid, screen3 please try again",
			})
		}

		// --------------------------------
		// Convert each one into file type
		// --------------------------------
		img1File, err := screen1.Open()
		if err != nil {
			fmt.Println("Error while opening file, ", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid, image please try again",
			})
		}
		defer img1File.Close()

		img2File, err := screen2.Open()
		if err != nil {
			fmt.Println("Error while opening file, ", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid, image please try again",
			})
		}
		defer img2File.Close()

		img3File, err := screen3.Open()
		if err != nil {
			fmt.Println("Error while opening file, ", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid, image please try again",
			})
		}
		defer img3File.Close()

		// -------------------------------
		// Get predictions for each image
		// -------------------------------
		//var wg sync.WaitGroup
		//wg.Add(2)

		c := make(chan structs.ChannelMessage)
		c2 := make(chan structs.ChannelMessage)
		c3 := make(chan structs.ChannelMessage)

		go func() {
			getPredictions(&img1File, c)
			//wg.Done()
		}()

		go func() {
			getPredictions(&img2File, c2)
			//wg.Done()
		}()

		go func() {
			getPredictions(&img3File, c3)
			//wg.Done()
		}()

		//wg.Wait()

		// Collect Data from channels
		screen1Data := <-c
		if len(screen1Data.Message) != 0 {
			fmt.Println("Error while getting predictions, ", screen1Data.Message)
			return ctx.Status(screen1Data.StatusCode).JSON(fiber.Map{
				"success": false,
				"message": screen1Data.Message,
			})
		}
		close(c)

		screen2Data := <-c2
		if len(screen1Data.Message) != 0 {
			fmt.Println("Error while getting predictions, ", screen1Data.Message)
			return ctx.Status(screen1Data.StatusCode).JSON(fiber.Map{
				"success": false,
				"message": screen1Data.Message,
			})
		}
		close(c2)

		screen3Data := <-c3
		if len(screen1Data.Message) != 0 {
			fmt.Println("Error while getting predictions, ", screen1Data.Message)
			return ctx.Status(screen1Data.StatusCode).JSON(fiber.Map{
				"success": false,
				"message": screen1Data.Message,
			})
		}
		close(c3)

		// -----------------------------
		// Process data for each screen
		// -----------------------------

		data := make(map[string]string)

		for key, value := range screen1Data.Data {
			data[key] = value
		}

		for key, value := range screen2Data.Data {
			data[key] = value
		}

		for key, value := range screen3Data.Data {
			data[key] = value
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"data":    data,
		})
	}
}

func getPredictions(imgFile *multipart.File, c chan structs.ChannelMessage) {

	// Get 2d array of rgb values for image
	imgRgb, resizedImgWidth, resizedImgHeight, imgW, imgH, err := utils.PrepareImage(imgFile, 640, 640)
	if err != nil {
		fmt.Println("Error while preparing image, ", err)
		c <- structs.ChannelMessage{
			Data:       nil,
			StatusCode: fiber.StatusBadGateway,
			Message:    "Invalid, image please try again",
		}
		return
	}

	output, err := runMainModel(imgRgb)
	if err != nil {
		fmt.Println("Error while running model, ", err)
		c <- structs.ChannelMessage{
			Data:       nil,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Oops! Something went wrong. Please try again later",
		}
		return
	}

	itemsData := processOutput(output, resizedImgWidth, resizedImgHeight, imgW, imgH)

	// Stores the rgb value of each item in the image [[float32], [float32], [float32]]
	itemRgbs := [][]float32{}

	// Prepare each item in the image for info detection
	for i := 0; i < len(itemsData); i++ {
		x := uint(itemsData[i][0].(float64))
		y := uint(itemsData[i][1].(float64))
		w := uint(itemsData[i][2].(float64) - itemsData[i][0].(float64))
		h := uint(itemsData[i][3].(float64) - itemsData[i][1].(float64))

		imgRgb, _, _, _, _, err := utils.CropAndPrepareImage(imgFile, x, y, w, h, 224, 224)
		if err != nil {
			fmt.Println("Error while preparing image, ", err)
			c <- structs.ChannelMessage{
				Data:       nil,
				StatusCode: fiber.StatusBadGateway,
				Message:    "Invalid, image please try again",
			}
			return
		}

		itemRgbs = append(itemRgbs, imgRgb)
	}

	// {"bef":"1", "che":"1", let:"1"}
	data := make(map[string]string)

	for i := 0; i < len(itemRgbs); i++ {
		output, err := runInfoModel(itemRgbs[i])
		if err != nil {
			fmt.Println("Error while running model, ", err)

			c <- structs.ChannelMessage{
				Data:       nil,
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Oops! Something went wrong. Please try again later",
			}
			return
		}

		highest := 0

		// Filter the index with highest confidence
		for i := 1; i < len(output); i++ {
			if output[i] > output[highest] {
				highest = i
			}
		}

		data[itemsData[i][4].(string)] = infoClasses[highest]
	}

	c <- structs.ChannelMessage{
		Data:       data,
		StatusCode: fiber.StatusOK,
		Message:    "",
	}
}

func runMainModel(input []float32) ([]float32, error) {
	// Create the input layer shape and get the tensor ready with it
	inputShape := ort.NewShape(1, 3, 640, 640)
	inputTensor, err := ort.NewTensor(inputShape, input)
	if err != nil {
		return nil, fmt.Errorf("error while creating input layer, %s", err.Error())
	}

	// Create the output layer shape and get the tensor ready with it
	outputShape := ort.NewShape(1, 16, 8400)
	outputTensor, err := ort.NewEmptyTensor[float32](outputShape)
	if err != nil {

		return nil, fmt.Errorf("error while creating output layer, %s", err.Error())
	}

	// Initialize the model
	model, err := ort.NewSession[float32]("models/bloxburg/cashier/main.onnx",
		[]string{"images"}, []string{"output0"},
		[]*ort.Tensor[float32]{inputTensor}, []*ort.Tensor[float32]{outputTensor})

	if err != nil {
		return nil, fmt.Errorf("error while loading model, %s", err.Error())
	}

	// Run model
	err = model.Run()

	// Destroy model
	defer model.Destroy()

	if err != nil {
		return nil, fmt.Errorf("error while running model, %s", err.Error())
	}

	return outputTensor.GetData(), nil
}

func runInfoModel(input []float32) ([]float32, error) {
	// Create the input layer shape and get the tensor ready with it
	inputShape := ort.NewShape(1, 3, 224, 224)
	inputTensor, err := ort.NewTensor(inputShape, input)
	if err != nil {
		return nil, fmt.Errorf("error while creating input layer, %s", err.Error())
	}

	// Create the output layer shape and get the tensor ready with it
	outputShape := ort.NewShape(1, 5)
	outputTensor, err := ort.NewEmptyTensor[float32](outputShape)
	if err != nil {
		return nil, fmt.Errorf("error while creating output layer, %s", err.Error())
	}

	// Initialize the model
	model, err := ort.NewSession[float32]("models/bloxburg/cashier/info.onnx",
		[]string{"images"}, []string{"output0"},
		[]*ort.Tensor[float32]{inputTensor}, []*ort.Tensor[float32]{outputTensor})

	if err != nil {
		return nil, fmt.Errorf("error while loading model, %s", err.Error())
	}

	// Run model
	err = model.Run()

	// Destroy model
	defer model.Destroy()

	if err != nil {
		return nil, fmt.Errorf("error while running model, %s", err.Error())
	}

	return outputTensor.GetData(), nil
}

func processOutput(output []float32, resizedImgWidth, resizedImgHeight, imgW, imgH int) [][]interface{} {

	// Resulting bounding boxes
	boxes := [][]interface{}{}

	// Loop through each row or detection (8400 is the max/default for yoloV8)
	for index := 0; index < 8400; index++ {
		// Initialize classId
		classId, prob := 0, float32(0.0)

		// Loop through each Column or each class idx there are 12 of them in this model
		for col := 0; col < 12; col++ {
			// Get probability and the classId for that detection
			if output[8400*(col+4)+index] > prob {
				prob = output[8400*(col+4)+index]
				classId = col
			}
		}

		// Make sure probability is greater then 70%
		if prob < 0.7 {
			continue
		}

		// Get the label
		label := classes[classId]

		// Get the center x and center y
		xc := output[index]
		yc := output[8400+index]

		// Get height and width
		w := output[2*8400+index]
		h := output[3*8400+index]

		// Calculate top left and resize to original
		x1 := (xc - w/2) / float32(resizedImgWidth) * float32(imgW)
		y1 := (yc - h/2) / float32(resizedImgHeight) * float32(imgH)

		// Calculate bottom right and resize to original
		x2 := (xc + w/2) / float32(resizedImgWidth) * float32(imgW)
		y2 := (yc + h/2) / float32(resizedImgHeight) * float32(imgH)

		// Append to boxes array
		boxes = append(boxes, []interface{}{float64(x1), float64(y1), float64(x2), float64(y2), label, prob})
	}

	// Reorganize array so the greatest probity is at the first index (Getting ready for NMS)
	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i][5].(float32) < boxes[j][5].(float32)
	})

	// Non-maximum Suppression
	result := [][]interface{}{}
	for len(boxes) > 0 {
		result = append(result, boxes[0])
		tmp := [][]interface{}{}
		for _, box := range boxes {
			// If they are intersecting less then 70% continue
			if utils.Iou(boxes[0], box) < 0.7 {
				tmp = append(tmp, box)
			}
		}
		boxes = tmp
	}

	return result
}
