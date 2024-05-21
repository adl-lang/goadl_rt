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
	Nothing *Void `json:"nothing,omitempty"`
	Just    *T    `json:"just,omitempty"`
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

func (obj *Maybe[T]) MarshalJSON() ([]byte, error) {
	if obj.Nothing != nil {
		return []byte(`"nothing"`), nil
	}
	j, e := json.Marshal(obj.Just)
	if e != nil {
		return nil, e
	}
	b := make([]byte, len(j)+len([]byte(`{"just":}`)))
	copy(b, []byte(`{"just":`))
	copy(b[len([]byte(`{"just":`)):], j)
	b[len(b)-1] = '}'
	return b, nil
}
