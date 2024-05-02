package goadl

import (
	"encoding/json"

	"github.com/golang/glog"
)

type Void struct{}

func (obj *Void) UnmarshalJSON(b []byte) error {
	return nil
}

func (obj Void) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

type Maybe[T any] struct {
	Nothing *Void `json:"nothing"`
	Just    *T    `json:"just"`
}

func (obj *Maybe[T]) UnmarshalJSON(b []byte) error {
	// fmt.Printf("1. %s\n", string(b))
	key, val, err := SplitKV(b)
	// fmt.Printf("2. '%s' '%s' '%v'\n", key, string(val), err)
	if err != nil {
		return err
	}
	switch key {
	case "nothing":
		obj.Nothing = &Void{}
	case "just":
		var just T
		err = json.Unmarshal(val, &just)
		if err != nil {
			glog.Errorf("just Unmarshal %v", err)
			return err
		}
		obj.Just = &just
	}
	return nil
}
