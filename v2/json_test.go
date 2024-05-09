package goadl_test

import (
	"reflect"
	"testing"

	"github.com/helix-collective/goadl/v2"
)

func TestSplitKV(t *testing.T) {
	type args struct {
		ba []byte
	}
	tests := []struct {
		name    string
		args    []byte
		wantKey string
		wantVal []byte
		wantErr bool
	}{
		{"string", []byte(`"primitive"`), "primitive", nil, false},
		{"string string", []byte(`{"primitive":"Int32"}`), "primitive", []byte(`"Int32"`), false},
		{"string null", []byte(`{"primitive":null}`), "primitive", []byte(`null`), false},
		{"string true", []byte(`{"primitive":true}`), "primitive", []byte(`true`), false},
		{"string false", []byte(`{"primitive":false}`), "primitive", []byte(`false`), false},
		{"string false", []byte(`{"primitive":f}`), "primitive", []byte(`f`), false},
		{"string array", []byte(`{"primitive":[]}`), "primitive", []byte(`[]`), false},
		{"string obj", []byte(`{"primitive":{"a":1}}`), "primitive", []byte(`{"a":1}`), false},
		{"string num", []byte(`{"primitive":123}`), "primitive", []byte(`123`), false},
		{"string num with char return", []byte(`{
"primitive"
:
123
}`), "primitive", []byte(`123`), false},
		{"string num with char return & tab", []byte(`{
	"primitive"	
	:	
	123	
	}`), "primitive", []byte(`123`), false},
		{"string num with char return & tab", []byte("{\v\"primitive\"\n\r:\r\n\f\v123}"), "primitive", []byte(`123`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotVal, err := goadl.SplitKV(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitKV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("SplitKV() gotKey = %v, want %v", string(gotKey), string(tt.wantKey))
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("SplitKV() gotVal = %v, want %v", string(gotVal), string(tt.wantVal))
			}
		})
	}
}
