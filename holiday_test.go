package holidayjp

import (
	"testing"
	"time"
)

func TestFalse(t *testing.T) {

	if b, _ := Holiday(time.Date(1991, 3, 15, 0, 0, 0, 0, time.UTC)); b {
		t.Errorf("My birthday was not a holiday!")
	}
}

func TestTrue(t *testing.T) {
	b, d := Holiday(time.Date(2018, 4, 30, 0, 0, 0, 0, time.UTC))
	if !b {
		t.Errorf("Should be a holiday! showa day observation!")
	} else {
		t.Log(d)
	}
}
