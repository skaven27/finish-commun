package passage

import (
	"encoding/json"
	"strings"
	"time"
)

type Passage struct {
	NumPassage            int           `json:"nNumPassage"`
	NTranspondeur         int           `json:"nTranspondeur"`
	DHPassage             DHDateTime    `json:"dhPassage"`
	DHPassageUTC          DHDateTimeUTC `json:"dhPassageUTC"`
	ForceSignal           int           `json:"nForceSignal"`
	STranspondeur         string        `json:"sTranspondeur"`
	Sport                 int           `json:"nSport"`
	NombreReceptionValide int           `json:"nNombreR\u00e9ceptionValide"`
	Canal                 int           `json:"nCanal"`
	Boucle                string        `json:"sBoucle"`
	NumBoucle             string        `json:"nNumBoucle"`
	EstPhotoCell          bool          `json:"bEstPhotoCell"`
	DeLaMemoire           bool          `json:"bDeLaM\u00e9moire"`
	Batterie              string        `json:"sBatterie"`
	BufAutreInfo          string        `json:"bufAutreInfo"`
	Cle                   string        `json:"sCle"`
	MACDecodeur           string        `json:"sMACDecodeur"`
	IPDecodeur            string        `json:"sIPD\u00e9codeur"`
	CheckSum              int           `json:"nCheckSum"`
	SAutreInfo            string        `json:"sAutreInfo"`
	FileNumber            int           `json:"nFileNumber"`
	PassingNumber         string        `json:"nPassingNumber"`
	DHRecu                DHDateTime    `json:"dhRecu"`
}

type DHDateTime time.Time
type DHDateTimeUTC time.Time

func (dh *DHDateTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05.000"`, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	*dh = DHDateTime(date)
	return nil
}

func (dh DHDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(dh))
}

// Maybe a Format function for printing your date
func (dh DHDateTime) Format(s string) string {
	t := time.Time(dh)
	return t.Format(s)
}

func (dh *DHDateTimeUTC) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05.000-0700"`, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	*dh = DHDateTimeUTC(date)
	return nil
}

func (dh DHDateTimeUTC) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(dh))
}

// Maybe a Format function for printing your date
func (dh DHDateTimeUTC) Format(s string) string {
	t := time.Time(dh)
	return t.Format(s)
}
