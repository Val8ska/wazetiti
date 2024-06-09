package save

import (
	"log"
	"os"
	"path/filepath"
	"wazetiti/internal/structures"

	"github.com/gocarina/gocsv"
)

func ToCsvFile(data []structures.CsvData, resDir string) {
	f, err := os.Create(filepath.Join(resDir, "wazeInfo.csv"))
	if err != nil {
		log.Println("[WAZE realtime] os create file", err)
	}
	defer f.Close()
	gocsv.MarshalFile(&data, f)
	log.Println("[WAZE realtime]", "new waze data saved as csv")
}
