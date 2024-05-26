// Code generated by goadlc v3 - DO NOT EDIT.
package types

import (
	"fmt"
)

type Either[T1 any, T2 any] struct {
	Branch EitherBranch[T1, T2]
}

type EitherBranch[T1 any, T2 any] interface {
	isEitherBranch()
}

func (*Either[T1, T2]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "left":
		return &Either_Left[T1]{}, nil
	case "right":
		return &Either_Right[T2]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type Either_Left[T1 any] struct {
	V T1 `branch:"left"`
}
type Either_Right[T2 any] struct {
	V T2 `branch:"right"`
}

func (Either_Left[T1]) isEitherBranch()  {}
func (Either_Right[T2]) isEitherBranch() {}

func Make_Either_left[T1 any, T2 any](v T1) Either[T1, T2] {
	return Either[T1, T2]{
		Either_Left[T1]{v},
	}
}

func Make_Either_right[T1 any, T2 any](v T2) Either[T1, T2] {
	return Either[T1, T2]{
		Either_Right[T2]{v},
	}
}

func Handle_Either[T1 any, T2 any, T any](
	_in EitherBranch[T1, T2],
	left func(left T1) T,
	right func(right T2) T,
	_default func() T,
) T {
	switch _b := _in.(type) {
	case Either_Left[T1]:
		if left != nil {
			return left(_b.V)
		}
	case Either_Right[T2]:
		if right != nil {
			return right(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("code gen error unhandled branch in : Either")
}

func HandleWithErr_Either[T1 any, T2 any, T any](
	_in EitherBranch[T1, T2],
	left func(left T1) (T, error),
	right func(right T2) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.(type) {
	case Either_Left[T1]:
		if left != nil {
			return left(_b.V)
		}
	case Either_Right[T2]:
		if right != nil {
			return right(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("code gen error unhandled branch in : Either")
}

type Map[K any, V any] []MapEntry[K, V]

type MapEntry[K any, V any] struct {
	Key   K `json:"k"`
	Value V `json:"v"`
}

func New_MapEntry[K any, V any](
	key K,
	value V,
) MapEntry[K, V] {
	return MapEntry[K, V]{
		Key:   key,
		Value: value,
	}
}

func Make_MapEntry[K any, V any](
	key K,
	value V,
) MapEntry[K, V] {
	ret := MapEntry[K, V]{
		Key:   key,
		Value: value,
	}
	return ret
}

type Maybe[T any] struct {
	Branch MaybeBranch[T]
}

type MaybeBranch[T any] interface {
	isMaybeBranch()
}

type Maybe_Nothing struct {
	V struct{} `branch:"nothing"`
}
type Maybe_Just[T any] struct {
	V T `branch:"just"`
}

func (Maybe_Nothing) isMaybeBranch() {}
func (Maybe_Just[T]) isMaybeBranch() {}

func Make_Maybe_nothing[T any]() Maybe[T] {
	return Maybe[T]{
		Maybe_Nothing{struct{}{}},
	}
}
func Make_Maybe_just[T any](v T) Maybe[T] {
	return Maybe[T]{
		Maybe_Just[T]{v},
	}
}

func (*Maybe[T]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "nothing":
		return &Maybe_Nothing{}, nil
	case "just":
		return &Maybe_Just[T]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

func Handle_Maybe[T any, T2 any](
	_in MaybeBranch[T],
	nothing func(nothing struct{}) T2,
	just func(just T) T2,
	_default func() T2,
) T2 {
	switch _b := _in.(type) {
	case Maybe_Nothing:
		if nothing != nil {
			return nothing(_b.V)
		}
	case Maybe_Just[T]:
		if just != nil {
			return just(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Maybe")
}

func HandleWithErr_Maybe[T any, T2 any](
	_in MaybeBranch[T],
	nothing func(nothing struct{}) (T2, error),
	just func(just T) (T2, error),
	_default func() (T2, error),
) (T2, error) {
	switch _b := _in.(type) {
	case Maybe_Nothing:
		if nothing != nil {
			return nothing(_b.V)
		}
	case Maybe_Just[T]:
		if just != nil {
			return just(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Maybe")
}

type Pair[T1 any, T2 any] struct {
	V1 T1 `json:"v1"`
	V2 T2 `json:"v2"`
}

func New_Pair[T1 any, T2 any](
	v1 T1,
	v2 T2,
) Pair[T1, T2] {
	return Pair[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

func Make_Pair[T1 any, T2 any](
	v1 T1,
	v2 T2,
) Pair[T1, T2] {
	ret := Pair[T1, T2]{
		V1: v1,
		V2: v2,
	}
	return ret
}

type Result[T any, E any] struct {
	Branch ResultBranch[T, E]
}

type ResultBranch[T any, E any] interface {
	isResultBranch()
}

func (*Result[T, E]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "ok":
		return &Result_Ok[T]{}, nil
	case "error":
		return &Result_Error[E]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type Result_Ok[T any] struct {
	V T `branch:"ok"`
}
type Result_Error[E any] struct {
	V E `branch:"error"`
}

func (Result_Ok[T]) isResultBranch()    {}
func (Result_Error[E]) isResultBranch() {}

func Make_Result_ok[T any, E any](v T) Result[T, E] {
	return Result[T, E]{
		Result_Ok[T]{v},
	}
}

func Make_Result_error[T any, E any](v E) Result[T, E] {
	return Result[T, E]{
		Result_Error[E]{v},
	}
}

func Handle_Result[T any, E any, T2 any](
	_in ResultBranch[T, E],
	ok func(ok T) T2,
	error func(error E) T2,
	_default func() T2,
) T2 {
	switch _b := _in.(type) {
	case Result_Ok[T]:
		if ok != nil {
			return ok(_b.V)
		}
	case Result_Error[E]:
		if error != nil {
			return error(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("code gen error unhandled branch in : Result")
}

func HandleWithErr_Result[T any, E any, T2 any](
	_in ResultBranch[T, E],
	ok func(ok T) (T2, error),
	error func(error E) (T2, error),
	_default func() (T2, error),
) (T2, error) {
	switch _b := _in.(type) {
	case Result_Ok[T]:
		if ok != nil {
			return ok(_b.V)
		}
	case Result_Error[E]:
		if error != nil {
			return error(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("code gen error unhandled branch in : Result")
}

type Set[T any] []T
