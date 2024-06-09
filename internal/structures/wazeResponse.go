package structures

type WazeResponse struct {
	Alerts          []Alerts `json:"alerts"`
	EndTimeMillis   int64    `json:"endTimeMillis"`
	StartTimeMillis int64    `json:"startTimeMillis"`
	StartTime       string   `json:"startTime"`
	EndTime         string   `json:"endTime"`
	Jams            []Jams   `json:"jams"`
}
