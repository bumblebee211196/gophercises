package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseLines(t *testing.T) {
	
	// Test cases
	cases := [] struct {
		name string
		input [][]string
		expect []problem
	}{
		{"TestCase", [][]string{
			{"1+1", "2"},
			{"2+2", "4"},
		}, []problem{
			{"1+1", "2"},
			{"2+2", "4"},
		}},
		{"TestCase", [][]string{}, []problem{}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %d", c.name, i + 1), func(t *testing.T) {
			actual := parseLines(c.input)
			if !reflect.DeepEqual(c.expect, actual) {
				t.Errorf("Error for input %v. \n%v != %v", c.input, c.expect, actual)
			}
		})
	}
}