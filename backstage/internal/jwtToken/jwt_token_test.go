package jwtToken

import (
	"fmt"
	"testing"
	"time"

	"github.com/brahma-adshonor/gohook"
)

func myTime() time.Time {
	return time.Date(2022, 1, 1, 0, 0, 0, 0, &time.Location{})
}

func TestFooBasic(t *testing.T) {
	fmt.Println(time.Now())
	gohook.Hook(time.Now, myTime, nil)
	fmt.Println(time.Now())
	//assert.Equal(t, expect, actual)
}
