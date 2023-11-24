package br

import (
	"enum"
)

type IfFunc[T enum.SumType, U any] func(T) U
type ElseFunc[T enum.SumType, U any] func(T) U
type CondFunc func() (enum.SumType, bool, error)

func Branch[S, T, U, V enum.SumType](cond CondFunc, ifFn IfFunc[S, T], elseFn ElseFunc[U, V]) enum.SumType {
	v, b, err := cond()
	if err != nil {
		return enum.NewSumTypeOne(err)
	}

	if b {
		return ifFn((interface{})(v).(S))
	}
	return elseFn((interface{})(v).(U))
}
