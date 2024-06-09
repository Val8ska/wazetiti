package save

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"
	"wazetiti/internal/structures"
)

func ToJsonFile(wazeResponse structures.WazeResponse, resDir string) {
	thisIsNow := time.Now().Format("20060102150405")
	file, _ := json.MarshalIndent(wazeResponse, "", " ")

	err := os.WriteFile(filepath.Join(resDir, thisIsNow+"_wazeData.json"), file, 0644)
	if err!=nil{
		log.Println("[WAZE realtime] save as json file", err)
	}
	time.Sleep(time.Duration(time.Second * 1))
	log.Println("[WAZE realtime]", "new waze response saved as json")
}
