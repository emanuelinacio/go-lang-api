package event_type

import (
	"encoding/json"
	"errors"
)

type ObExtras struct {
	Extras string
}

func (o ObExtras) Init(Extras string) (ObExtras, error) {
	o.Extras = Extras

	if o.IsValidExtras() {
		return o, nil
	}

	return o, errors.New("Invalid json format")
}

func (e ObExtras) IsValidExtras() bool {

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(e.Extras), &dat)

	return err == nil
}
