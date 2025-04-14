package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Aqil"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Aqil")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Aqil") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
