package structures

type Comments struct {
}

type Alerts struct {
	Country                  string     `json:"country,omitempty" csv:"country"`
	NThumbsUp                int        `json:"nThumbsUp,omitempty" csv:"nThumbsUp"`
	ReportRating             int        `json:"reportRating"`
	ReportMillis             int64      `json:"reportMillis,omitempty"`
	ReportByMunicipalityUser string     `json:"reportByMunicipalityUser"`
	Reliability              int        `json:"reliability" csv:"reliability"`
	Type                     string     `json:"type" csv:"type"`
	UUID                     string     `json:"uuid" csv:"uuid"`
	Speed                    int        `json:"speed"`
	ReportMood               int        `json:"reportMood"`
	Subtype                  string     `json:"subtype" csv:"subtype"`
	Street                   string     `json:"street,omitempty" csv:"street"`
	AdditionalInfo           string     `json:"additionalInfo"`
	ID                       string     `json:"id"`
	Text                     string     `json:"text,omitempty" csv:"text"`
	NComments                int        `json:"nComments,omitempty"`
	Inscale                  bool       `json:"inscale"`
	Comments                 []Comments `json:"comments,omitempty"`
	Confidence               int        `json:"confidence" csv:"confidence"`
	NearBy                   string     `json:"nearBy,omitempty"`
	RoadType                 int        `json:"roadType,omitempty"`
	Magvar                   int        `json:"magvar"`
	WazeData                 string     `json:"wazeData"`
	Location                 Location   `json:"location" csv:"location"`
	PubMillis                int64      `json:"pubMillis"`
	IsThumbsUp               bool       `json:"isThumbsUp,omitempty"`
	City                     string     `json:"city,omitempty" csv:"city"`
	ReportBy                 string     `json:"reportBy,omitempty"`
	ReportDescription        string     `json:"reportDescription,omitempty"`
	Provider                 string     `json:"provider,omitempty"`
	ProviderID               string     `json:"providerId,omitempty"`
}
