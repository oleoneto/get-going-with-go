package greetings

import (
	"testing"
	"regexp"
)


func TestHelloName(t *testing.T) {
	name := "Leo"
	expected := regexp.MustCompile(`\b`+name+`\b`)

	msg, err := Hello("Leo")

	if !expected.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Leo") = %q, %v, want match for %#q, nil`, msg, err, expected)
	}
}


func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
