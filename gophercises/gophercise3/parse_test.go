package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"golang.org/x/net/html"
)

var (
	emptyLinks []Link
)

func TestParse(t *testing.T) {
	// Test cases
	cases := [] struct {
		name	string
		file 	string
		expect 	[]Link
	}{
		{
			"Testcase", 
			"sample1.html", 
			[]Link{
				{Href:"/other-page", Text:"A link to another page"},
				{Href:"/other-page", Text:"A link to second page"},
			},
		},
		{
			"Testcase", 
			"sample2.html", 
			[]Link{
				{Href:"https://www.twitter.com/joncalhoun", Text:"Check me out on twitter"},
				{Href:"https://github.com/gophercises", Text:"Gophercises is on Github!"},
			},
		},
		{
			"Testcase", 
			"sample3.html", 
			[]Link{
				{Href:"#", Text:"Login"},
				{Href:"/lost", Text:"Lost? Need help?"},
				{Href:"https://twitter.com/marcusolsson", Text:"@marcusolsson"},
			},
		},
		{
			"Testcase", 
			"sample4.html", 
			[]Link{
				{Href:"/dog-cat", Text:"dog cat"},
			},
		},
		{
			"Testcase", 
			"empty.html", 
			emptyLinks,
		},
	}
	// Start test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %d", c.name, i + 1), func(t *testing.T) {
			htmlData, _ := os.Open(c.file)
			actual, _ := Parse(htmlData)
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v, \n%#v != %#v", c.file, c.expect, actual)
			}
		})
	}
}

func TestBuildLink(t *testing.T) {
	// Test cases
	cases := [] struct {
		name 	string
		href 	string
		text 	string
		expect  Link
	}{
		{
			"Testcase",
			"href1",
			"text1",
			Link{"href1", "text1"},
		},
		{
			"Testcase",
			"",
			"",
			Link{"", ""},
		},
	}
	// Start test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %d", c.name, i + 1), func(t *testing.T) {
			actual := buildLink(c.href, c.text)
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v %v, \n%#v != %#v", c.href, c.text, c.expect, actual)
			}
		})
	}
}

func TestGetLinkNodes(t *testing.T) {
	// Test cases
	cases := [] struct {
		name 	string
		file 	string
		expect 	int
	}{
		{
			"Testcase",
			"sample1.html",
			2,
		},
		{
			"Testcase",
			"sample2.html",
			2,
		},
		{
			"Testcase",
			"sample3.html",
			3,
		},
		{
			"Testcase",
			"sample4.html",
			1,
		},
	}
	// Start test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.name, i + 1), func(t *testing.T) {
			htmlData, _ := os.Open(c.file)
			nodes, _ := html.Parse(htmlData)
			actual := len(getLinkNodes(nodes))
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v, \n%#v != %#v", c.file, c.expect, actual)
			}
		})
	}
}

func TestGetLinkNodeText(t *testing.T) {
	// Test cases
	cases := [] struct {
		name 	string
		file 	string
		expect 	[]string
	}{
		{
			"Testcase",
			"sample1.html",
			[]string{"A link to another page", "A link to second page"},
		},
		{
			"Testcase",
			"sample2.html",
			[]string{"Check me out on twitter", "Gophercises is on Github!"},
		},
		{
			"Testcase",
			"sample3.html",
			[]string{"Login", "Lost? Need help?", "@marcusolsson"},
		},
		{
			"Testcase",
			"sample4.html",
			[]string{"dog cat"},
		},
	}
	// Start test
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.name, i + 1), func(t *testing.T) {
			actual := make([]string, 0)
			htmlData, _ := os.Open(c.file)
			node, _ := html.Parse(htmlData)
			linkNodes := getLinkNodes(node)
			for _, node := range linkNodes {
				actual = append(actual, getLinkNodeText(node))
			}
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v, \n%#v != %#v", c.file, c.expect, actual)
			}
		})
	}
}

func TestGetLinkNodeHref(t *testing.T) {
	// Test cases
	cases := [] struct {
		name 	string
		file 	string
		expect 	[]string
	}{
		{
			"Testcase",
			"sample1.html",
			[]string{"/other-page", "/other-page"},
		},
		{
			"Testcase",
			"sample2.html",
			[]string{"https://www.twitter.com/joncalhoun", "https://github.com/gophercises"},
		},
		{
			"Testcase",
			"sample3.html",
			[]string{"#", "/lost", "https://twitter.com/marcusolsson"},
		},
		{
			"Testcase",
			"sample4.html",
			[]string{"/dog-cat"},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s %v", c.name, i + 1), func(t *testing.T) {
			actual := make([]string, 0)
			htmlData, _ := os.Open(c.file)
			node, _ := html.Parse(htmlData)
			linkNodes := getLinkNodes(node)
			for _, node := range linkNodes {
				actual = append(actual, getLinkNodeHref(node))
			}
			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("Error in output for input %v, \n%#v != %#v", c.file, c.expect, actual)
			}
		})
	}
}
