package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	// := operator is a shortcut for declaring and initializing a variable in one line
	return message, nil /*nil (meaning no rror) as a second value in the successful return. That way, the caller can see that the function succeeded.
	 */
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names { // '_' is the index of the current item
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
