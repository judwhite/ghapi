package ghapi

import (
	"strings"
	"testing"
)

func TestValidateOwnerName(t *testing.T) {
	cases := []struct {
		name string
		ok   bool
	}{
		{"", false},
		{".", false},
		{"..", false},
		{"-", false},
		{"_", false},
		{"-helloworld", false},
		{"helloworld-", false},
		{"hello--world", false},
		{"hello_world", false},
		{"ñ", false},
		{strings.Repeat("a", 40), false},

		{"a", true},
		{"n", true},
		{"helloworld", true},
		{"HelloWorld", true},
		{"Hello-World", true},
		{"0", true},
		{"9", true},
		{"a", true},
		{"z", true},
		{"00hello-world00", true},
		{"Owner-Hyphen", true},
		{strings.Repeat("a", 39), true},
	}

	for _, c := range cases {
		err := ValidateOwnerName(c.name)
		if c.ok {
			if err != nil {
				t.Errorf("'%s' want: OK got: %v", c.name, err)
			}
		} else {
			if err == nil {
				t.Errorf("'%s' want: err got: <nil>", c.name)
			}
		}
	}
}

func TestValidateRepoName(t *testing.T) {
	cases := []struct {
		name string
		ok   bool
	}{
		{"", false},
		{".", false},
		{"..", false},
		{".git", false},
		{".wiki", false},
		{"a.GIT", false},
		{"a.WIKI", false},
		{"hello/world", false},
		{"hello world", false},
		{strings.Repeat("a", 101), false},

		{".a", true},
		{"..a", true},
		{"a.", true},
		{"a..", true},
		{"...", true},
		{"-", true},
		{"_", true},
		{"0", true},
		{"9", true},
		{"a", true},
		{"z", true},
		{"--hello-world--", true},
		{"..hello.world..", true},
		{"__hello_world__", true},
		{"00hello0world00", true},
		{"Project-Hyphen", true},
		{strings.Repeat("a", 100), true},
	}

	for _, c := range cases {
		err := ValidateRepoName(c.name)
		if c.ok {
			if err != nil {
				t.Errorf("'%s' want: OK got: %v", c.name, err)
			}
		} else {
			if err == nil {
				t.Errorf("'%s' want: err got: <nil>", c.name)
			}
		}
	}
}
