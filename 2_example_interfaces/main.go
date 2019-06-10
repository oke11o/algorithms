package example_interfaces

import (
	"fmt"
	"testing"
)

type MMYY interface {
	Ge() error
}

type NNYY interface {
	Se() error
}

type ttt struct {
}

func (t *ttt) Ge() error {
	return nil
}

func (t *ttt) Se() error {
	return nil
}

func Test_tmnp(t *testing.T) {
	s := &ttt{}
	tmp(s)
	tmp2(s)

	s2 := tmpCreate()
	tmp2(s2)

	s3, ok := s2.(NNYY)
	if ok {
		tmp(s3)
	}
	s4, ok := s2.(fmt.Stringer)
	if ok {
		fmt.Printf("%s", s4)
	}

	fmt.Printf("%v", s4)
}

func tmp(t NNYY) error {
	return t.Se()
}
func tmp2(t MMYY) error {
	return t.Ge()
}
func tmpCreate() MMYY {
	return &ttt{}
}