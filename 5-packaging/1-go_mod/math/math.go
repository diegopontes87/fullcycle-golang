package math

type Math struct {
	NumberA int
	NumberB int
}

func (m Math) Add() int {
	return m.NumberA + m.NumberB
}
