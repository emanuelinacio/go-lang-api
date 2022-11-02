package event_type

import (
	"encoding/json"
	"errors"
)

type ObExtras struct {
	Extras string
}

func NewObExtras(Extras string) (ObExtras, error) {
	o := new(ObExtras)
	o.Extras = Extras

	if o.IsValidExtras() {
		return *o, nil
	}

	return ObExtras{}, errors.New("Invalid json format")
}

func (o ObExtras) IsValidExtras() bool {

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(o.Extras), &dat)

	return err == nil
}
