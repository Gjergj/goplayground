package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestAddNumber(t *testing.T) {
	testCases := []struct {
		Name   string
		I      int
		J      int
		Result int
	}{
		{
			Name:   "Ok",
			I:      1,
			J:      1,
			Result: 2,
		},
		{
			Name:   "NOk",
			I:      2,
			J:      1,
			Result: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(tc.Name)
			time.Sleep(3 * time.Second)
			res := AddNumber(tc.I, tc.J)
			assert.Equal(t, res, tc.Result)
		})
	}
}