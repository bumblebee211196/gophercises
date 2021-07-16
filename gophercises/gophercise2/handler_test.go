package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)


func TestParseYAML(t *testing.T) {
	// Test cases
	cases := [] struct {
		name 	string
		input 	[]byte
		expect 	[]pathToURL
		err 	error
	}{
		{
			"TestCase",
			[]byte(`
- path: /p1
  url: url1
- path: /p2
  url: url2
`), 
			[]pathToURL{
				{"/p1", "url1"},
				{"/p2", "url2"},
			},
			nil,
		},
		{
			"TestCase",
			[]byte(""),
			make([]pathToURL, 0),
			nil,
		},
		{
			"TestCase",
			[]byte(`
- he
`), 
			[]pathToURL{},
			errors.New("yaml: unmarshal errors:\n  line 2: cannot unmarshal !!str `he` into main.pathToURL"),
		},
		{
			"TestCase",
			[]byte(`
""
`), 
			make([]pathToURL, 0),
			errors.New("yaml: unmarshal errors:\n  line 2: cannot unmarshal !!str `` into []main.pathToURL"),
		},
	}

	// Start test
	for i, c := range cases {
		t.Run(
			fmt.Sprintf("%s %d", c.name, i + 1), func(t *testing.T) {
				actual, err := parseYAML(c.input)
				if c.err != nil && !reflect.DeepEqual(err.Error(), c.err.Error()) {
					t.Errorf("Error in error for input %v, \n%v != %v", c.input, err, c.err)
				}
				if !reflect.DeepEqual(actual, c.expect) {
					t.Errorf("Error in output for input %v, \n%#v != %#v", c.input, c.expect, actual)
				}
			})
	}
}

func TestParseJSON(t *testing.T) {
	// Test cases
	cases := [] struct {
		name	string
		input	[]byte
		expect	[]pathToURL
		err		error
	}{
		{
			"TestCase",
			[]byte(`
[
	{
		"path": "/p1",
		"url": "url1"
	},
	{
		"path": "/p2",
		"url": "url2"
	}
]
`), 
			[]pathToURL{
				{"/p1", "url1"},
				{"/p2", "url2"},
			},
			nil,
		},
		{
			"TestCase",
			[]byte(`
- he
`),
			make([]pathToURL, 0),
			nil,
		},
		{
			"TestCase",
			[]byte(`
""
`), 
			make([]pathToURL, 0),
			errors.New("json: cannot unmarshal string into Go value of type []main.pathToURL"),
		},
	}
	// Start test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %d", c.name, i + 1), func(t *testing.T) {
			actual, err := parseJSON(c.input)
			if c.err != nil && !reflect.DeepEqual(err.Error(), c.err.Error()) {
				t.Errorf("Error in error for input %v, \n%v != %v", c.input, err, c.err)
			}
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v, \n%#v != %#v", c.input, c.expect, actual)
			}
		})
	}
}

func TestBuildMap(t *testing.T) {
	// Test cases
	cases := [] struct {
		name	string
		input	[]pathToURL
		expect	map[string]string
	}{
		{
			"TestCase",
			[]pathToURL{},
			make(map[string]string),
		},
		{
			"TestCase",
			[]pathToURL{
				{"/p1", "url1"},
				{"/p2", "url2"},
			},
			map[string]string{"/p1": "url1", "/p2": "url2"},
		},
	}
	// Start Test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %d", c.name, i + 1), func(t *testing.T) {
			actual := buildMap(c.input)
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v, \n%#v != %#v", c.input, c.expect, actual)
			}
		})
	}
}
