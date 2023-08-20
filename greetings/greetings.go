package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// := operator is a shortcut for declaring and initializing a variable in one line
	return message, nil /*nil (meaning no rror) as a second value in the successful return. That way, the caller can see that the function succeeded.
	 */
}
