package entity

import "testing"

func TestSum(t *testing.T) {

	var s int

	s = Sum(1,1)
	if s != 2 {
		t.Error("Expected 2, got ", s)
	}

}