package athletes

type Athletes struct {
	ID          string        `json:"id"`
	Surname     string        `json:"surname"`
	FirstName   string        `json:"firstname"`
	MiddleName  string        `json:"middlename"`
	Age         int           `json:"age"`
	Sexe        int           `json:"sexe"`
	Transponder []Transponder `json:"transponder"`
	Min         int           `json:"min"`
	Max         int           `json:"max"`
	FGColor     int           `json:"fgcolor"`
	BGColor     int           `json:"bgcolor"`
	LapRecorded []RaceRecord  `json:"lapsrecorded"`
	RaceRecord  []RaceRecord  `json:"racerecord"`
}

type Transponder struct {
	ID string `json:"id"`
}

type RaceRecord struct {
	RaceID    string `json:"raceid"`
	Distance  int    `json:"distance"`
	Time      int    `json:"time"`
	Timestamp string `json:"timestamp"`
}

func NewAthlete() *Athletes {
	return &Athletes{}
}
