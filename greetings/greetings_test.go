package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName is a test function for the Hello function.
func TestHelloName(t *testing.T) {
	// Define a string with the name "Gladys".
	name := "Gladys"

	// Construct a regular expression that looks for the word "Gladys" as a whole word.
	// The `\b` specifies a word boundary, ensuring the name is matched as an entire word and not as a substring.
	want := regexp.MustCompile(`\b` + name + `\b`)

	// Call the Hello function with the string "Gladys" and store the result in msg and err.
	msg, err := Hello("Gladys")

	// Check if the result from Hello contains the name "Gladys" (using the regular expression) and that no error was returned.
	// If the name is not found or an error is returned, the test will fail.
	if !want.MatchString(msg) || err != nil {
		// t.Fatalf is used to format the error message and immediately stop the test.
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
