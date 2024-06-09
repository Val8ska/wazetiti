package search

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"wazetiti/internal/structures"
)

func WazeApi(options structures.Options) (wazeResponse structures.WazeResponse) {
	params := url.Values{
		"top":    {options.TopLeftLat},
		"bottom": {options.BottomRightLat},
		"left":   {options.TopLeftLng},
		"right":  {options.BottomRightLng},
		"env":    {"row"},
		"types":  {"alerts", "traffic"},
	}
	resp, err := http.Get("https://www.waze.com/live-map/api/georss?" + params.Encode())
	if err != nil {
		log.Println("[WAZE realtime] http get resp",err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[WAZE realtime] io readAll body",err)
	}

	json.Unmarshal(body, &wazeResponse)
	return wazeResponse
}
