package structures

type Jams struct {
	Country           string     `json:"country" csv:"country"`
	City              string     `json:"city" csv:"city`
	Line              []Location `json:"line"`
	SpeedKMH          int        `json:"speedKMH"`
	Type              string     `json:"type"`
	BlockingAlertID   int        `json:"blockingAlertID,omitempty"`
	BlockExpiration   int64      `json:"blockExpiration,omitempty"`
	UUID              int        `json:"uuid" csv:"uuid`
	EndNode           string     `json:"endNode,omitempty"`
	Speed             int        `json:"speed"`
	Segments          []Segments `json:"segments"`
	Street            string     `json:"street" csv:"street`
	ID                int        `json:"id"`
	BlockStartTime    int64      `json:"blockStartTime,omitempty"`
	BlockUpdate       int64      `json:"blockUpdate,omitempty"`
	Severity          int        `json:"severity" csv:"severity`
	Level             int        `json:"level"`
	BlockType         string     `json:"blockType,omitempty"`
	Length            int        `json:"length" csv:"length`
	TurnType          string     `json:"turnType"`
	BlockingAlertUUID string     `json:"blockingAlertUuid,omitempty"`
	RoadType          int        `json:"roadType"`
	Delay             int        `json:"delay"`
	BlockDescription  string     `json:"blockDescription,omitempty"`
	UpdateMillis      int64      `json:"updateMillis"`
	CauseAlert        CauseAlert `json:"causeAlert,omitempty"`
	PubMillis         int64      `json:"pubMillis"`
	StartNode         string     `json:"startNode,omitempty"`
}

type Segments struct {
	FromNode  int  `json:"fromNode"`
	ID        int  `json:"ID"`
	ToNode    int  `json:"toNode"`
	IsForward bool `json:"isForward"`
}

type CauseAlert struct {
	ReportBy                 string   `json:"reportBy"`
	Country                  string   `json:"country"`
	Inscale                  bool     `json:"inscale"`
	City                     string   `json:"city"`
	ReportRating             int      `json:"reportRating"`
	ReportByMunicipalityUser string   `json:"reportByMunicipalityUser"`
	Confidence               int      `json:"confidence"`
	Reliability              int      `json:"reliability"`
	Type                     string   `json:"type"`
	UUID                     int      `json:"uuid"`
	Speed                    int      `json:"speed"`
	ReportMood               int      `json:"reportMood"`
	RoadType                 int      `json:"roadType"`
	Magvar                   int      `json:"magvar"`
	Subtype                  string   `json:"subtype"`
	Street                   string   `json:"street"`
	AdditionalInfo           string   `json:"additionalInfo"`
	WazeData                 string   `json:"wazeData"`
	ReportDescription        string   `json:"reportDescription"`
	Location                 Location `json:"location"`
	ID                       string   `json:"id"`
	PubMillis                int64    `json:"pubMillis"`
}
