package main

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	ort "github.com/yalue/onnxruntime_go"
	"www.autoblox.xyz/server/api"
	database "www.autoblox.xyz/server/db"
	"www.autoblox.xyz/server/utils"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	db := database.GetDB()

	// Load onnx runtime
	ort.SetSharedLibraryPath("onnx/onnxruntime.dll")
	err := ort.InitializeEnvironment()
	if err != nil {
		log.Fatalln("Onnx runtime could not be initialized, ", err)
	}
	// Close the runtime after
	defer ort.DestroyEnvironment()

	// ONLY DEVELOPMENT
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env could not be loaded, ", err)
	}

	// Delete expired keys every 1 minute
	s.Every(5).Minutes().Do(utils.DeleteExpiredKeys, db)

	// Start the corns to delete expired keys
	s.StartAsync()

	// Start the api
	api.Start(db, os.Getenv("STRIPE_KEY"), os.Getenv("HCAPTCHA_KEY"), os.Getenv("HCAPTCHA_SECRET"))
}
