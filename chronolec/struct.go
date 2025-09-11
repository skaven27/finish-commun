package chronolec

type PassageChronolec struct {
	Transponder  int    `json:"transponder"`
	Time         int    `json:"time"` // HH:MM'ss"mmm
	LoopIP       string `json:"loopid"`
	Power        int    `json:"power"`
	NumberInLoop int    `json:"numberinloop"`
	Battery      int    `json:"battery"` // Battery power   0=100-75%   1=75-50%    2=50-25%    3=< 25%
	Checksum     int    `json:"checksum"`
}

type StatusChronolec struct {
	Hour             int `json:"h"`
	Minutes          int `json:"m"`
	Seconds          int `json:"s"`
	NoiseStart       int `json:"noises"`
	NoisePit         int `json:"noisep"`
	SensitivityStart int `json:"sstart"`
	SensitivityPit   int `json:"spit"`
}
