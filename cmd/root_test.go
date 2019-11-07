package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var wrongParamsOrderTest = []struct {
	in     string
	expect error
}{
	{"1s1m1h", WrongOrderError},
	{"1s1h", WrongOrderError},
	{"1m1h", WrongOrderError},
	{"1s1m", WrongOrderError},
}

func TestWrongOrderOfParams(t *testing.T) {

	for _, tt := range wrongParamsOrderTest {
		_, err := sanitizeUserInput(tt.in)
		assert.EqualError(t, tt.expect, err.Error(), "Expected:"+tt.expect.Error())
	}
}

var correctParamsOrderTestWithWaitTime = []struct {
	input    string
	indexes  []int
	waitTime int
}{
	{"1h", []int{1, -1, -1}, 3600},
	{"1h30m", []int{1, 4, -1}, 5400},
	{"1h4m2s", []int{1, 3, 5}, 3842},
	{"1h40s", []int{1, -1, 4}, 3640},
	{"40m", []int{-1, 2, -1}, 2400},
	{"15m02s", []int{-1, 2, 5}, 902},
	{"20s", []int{-1, -1, 2}, 20},
}

func TestCorrectOrderOfParams(t *testing.T) {

	for _, tt := range correctParamsOrderTestWithWaitTime {
		tab, err := sanitizeUserInput(tt.input)
		assert.Nil(t, err)
		assert.Equal(t, tt.indexes, tab)

	}
}

func TestIndexesToWaitTime(t *testing.T) {
	for _, tt := range correctParamsOrderTestWithWaitTime {
		waitTime, err := indexesToWaitTime(tt.input, tt.indexes)
		assert.Nil(t, err)
		assert.Equal(t, tt.waitTime, waitTime)
	}
}
