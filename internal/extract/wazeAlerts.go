package extract

import (
	"wazetiti/internal/structures"
)

func WazeAlerts(wazeResponse structures.WazeResponse) (data []structures.CsvData) {
	for _, alert := range wazeResponse.Alerts {
		data = append(data, structures.CsvData{
			DateTime:       wazeResponse.StartTime,
			Lng:            alert.Location.X,
			Lat:            alert.Location.Y,
			Category:       "alert",
			Country:        alert.Country,
			NThumbsUp:      alert.NThumbsUp,
			Reliability:    alert.Reliability,
			Type:           alert.Type,
			UUID:           alert.UUID,
			Subtype:        alert.Subtype,
			Street:         alert.Street,
			AdditionalInfo: alert.AdditionalInfo,
			ID:             alert.ID,
			Text:           alert.Text,
			City:           alert.City,
			Confidence:     alert.Confidence,
		})
	}
	return data
}
