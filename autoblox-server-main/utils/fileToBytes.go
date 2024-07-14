package utils

import (
	"bytes"
	"io"
	"mime/multipart"
)

func FileToBytes(file multipart.File) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return nil, err
	}

	// Set file cursor back to start
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err

	}

	return buf.Bytes(), nil
}
