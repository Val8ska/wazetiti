package extract

import (
	"strconv"
	"wazetiti/internal/structures"
)

func WazeJams(wazeResponse structures.WazeResponse) (data []structures.CsvData) {
	for _, jam := range wazeResponse.Jams {
		data = append(data, structures.CsvData{
			DateTime:    wazeResponse.StartTime,
			Lng:         jam.Line[0].X,
			Lat:         jam.Line[0].Y,
			Category:    "jam",
			Country:     jam.Country,
			Type:        jam.Type,
			UUID:        strconv.Itoa(jam.UUID),
			Street:      jam.Street,
			ID:         strconv.Itoa(jam.ID),
			City:       jam.City,
		})
	}
	return data
}
