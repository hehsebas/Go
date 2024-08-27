package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Sebastian"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Sebastian")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Sebastian")= %q, %v, requiere coincidencia para %#q, nil`, msg, err, want)
	}
}
func TestEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("")= %q, %v, requiere  "", error`, msg, err)
	}
}
