package goadl_rt

import (
	"errors"
	"fmt"
)

func SplitKV(ba []byte) (key string, val []byte, err error) {
	var state, ks, ke, vs int
	for i, b := range ba {
		switch {
		case state == 0 && asciiSpace[b] != 0:
			continue
		case state == 0 && b == '{':
			state = 1
			continue
		case state == 0 && b == '"':
			state = 20
			ks = i
			continue
		case state == 20 && b == '"':
			state = 21
			ke = i
			continue
		case state == 20:
			continue
		case state == 1 && b == '"':
			state = 2
			ks = i
			continue
		case state == 2 && b == '"':
			state = 3
			ke = i
			continue
		case state == 3 && asciiSpace[b] != 0:
			continue
		case state == 3 && b == ':':
			state = 4
			continue
		case state == 4 && asciiSpace[b] != 0:
			continue
		case state == 4 && b == '"':
			vs = i
			state = 5
			break
		case state == 4 && b >= '0' && b <= '9':
			vs = i
			state = 6
			break
		case state == 4 && b == 'n':
			vs = i
			state = 7
			break
		case state == 4 && (b == 't' || b == 'T'):
			vs = i
			state = 8
			break
		case state == 4 && (b == 'f' || b == 'F'):
			vs = i
			state = 9
			break
		case state == 4 && b == '{':
			vs = i
			state = 10
			break
		case state == 4 && b == '[':
			vs = i
			state = 11
			break
		}
	}
	if state == 1 {
		return "", nil, errors.New("only open { found")
	}
	if state == 20 {
		return "", nil, errors.New("key (only) not terminated")
	}
	if state == 2 {
		return "", nil, errors.New("key not terminated")
	}
	if state == 3 {
		return "", nil, errors.New("colon not found")
	}
	if state == 3 {
		return "", nil, errors.New("start value not found")
	}
	if state == 21 {
		return string(ba[ks+1 : ke]), nil, nil
	}
	// state 4 or state 5
	var estate int
	for i := len(ba) - 1; i >= vs; i-- {
		b := ba[i]
		switch {
		case estate == 0 && asciiSpace[b] != 0:
			continue
		case estate == 0 && b == '}':
			estate = 1
			continue
		case estate == 1 && asciiSpace[b] != 0:
			continue
		case estate == 1 && state == 5 && b == '"':
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		case estate == 1 && state == 6 && b >= '0' && b <= '9':
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		case estate == 1 && state == 7 && b == 'l':
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		case estate == 1 && state == 8 && (b == 'e' || b == 'E' || b == 't' || b == 'T'):
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		case estate == 1 && state == 9 && (b == 'e' || b == 'E' || b == 'F' || b == 'f'):
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		case estate == 1 && state == 10 && b == '}':
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		case estate == 1 && state == 11 && b == ']':
			return string(ba[ks+1 : ke]), ba[vs : i+1], nil
		default:
			return "", nil, fmt.Errorf("unexpected end char '%s' estate:%d state:%d", string(b), estate, state)
		}
	}
	return "", nil, fmt.Errorf("UNEXPECTED PATH")
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
