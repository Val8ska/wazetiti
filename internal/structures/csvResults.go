package structures

type CsvData struct {
	DateTime       string  `json:"dateTime"`
	Lat            float64 `json:"lat"`
	Lng            float64 `json:"lng"`
	Category       string  `json:"category"`
	Country        string  `json:"country,omitempty" csv:"country"`
	NThumbsUp      int     `json:"nThumbsUp,omitempty" csv:"nThumbsUp"`
	Reliability    int     `json:"reliability" csv:"reliability"`
	Type           string  `json:"type" csv:"type"`
	UUID           string  `json:"uuid" csv:"uuid"`
	Subtype        string  `json:"subtype" csv:"subtype"`
	Street         string  `json:"street,omitempty" csv:"street"`
	AdditionalInfo string  `json:"additionalInfo"`
	ID             string  `json:"id"`
	Text           string  `json:"text,omitempty" csv:"text"`
	City           string  `json:"city,omitempty" csv:"city"`
	Confidence     int     `json:"confidence" csv:"confidence"`
}
