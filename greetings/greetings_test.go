package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	// given
	name := "Fabio"
	want := regexp.MustCompile(`\b` + name + `\b`)
	// when
	msg, err := Hello(name)
	// then
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Fabio") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmptyName(t *testing.T) {
	// given
	name := ""
	// when
	msg, err := Hello(name)
	// then
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
