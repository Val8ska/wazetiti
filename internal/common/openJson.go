package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"wazetiti/internal/structures"
)

func OpenJson(jsonPath string) (wazeResponse structures.WazeResponse) {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Println("[WAZE realtime] os open jsonfile",err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &wazeResponse)
	return wazeResponse
}
