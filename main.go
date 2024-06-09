package main

import (
	"log"
	"os"
	"time"
	"wazetiti/internal/common"
	"wazetiti/internal/extract"
	"wazetiti/internal/save"
	"wazetiti/internal/search"
	"wazetiti/internal/structures"
)

var options = structures.Options{
	TopLeftLat:     "51.6084336",
	BottomRightLat: "51.4284816",
	TopLeftLng:     "-0.2465508",
	BottomRightLng: "0.0894058",
	Delay:          2,
	ResDir:         time.Now().Format("20060102150405") + "_wazeData",
	EndDate:        "2024/06/10",
}

func main() {
	var fullData []structures.CsvData
	os.Mkdir(options.ResDir, os.ModePerm)
	endTime, _ := time.Parse("2006/01/02", options.EndDate)
	loopSize := int(endTime.Sub(time.Now()).Minutes() / float64(options.Delay))
	if loopSize > 0 {
		for i := 1; i < loopSize; i++ {
			wazeResponse := search.WazeApi(options)
			save.ToJsonFile(wazeResponse, options.ResDir)
			wazeAlerts := extract.WazeAlerts(wazeResponse)
			wazeJams := extract.WazeJams(wazeResponse)
			fullData = append(fullData, wazeAlerts...)
			fullData = append(fullData, wazeJams...)
			fullData = common.RemoveDuplicatesAlerts(fullData)
			save.ToCsvFile(fullData, options.ResDir)
			if i != loopSize {
				time.Sleep(time.Duration(options.Delay) * time.Second)
			}
		}
	} else {
		log.Println("[WAZE realtime], error endTime is before now")
	}
}
