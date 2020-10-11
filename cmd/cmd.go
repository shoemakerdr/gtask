package cmd

import (
	"fmt"
	"log"
)

// TOKENS
//    - ARG
//    - STRING

type token struct {
	id     string
	cursor int
	value  string
}

func newArgToken(cursor int, value string) token {
	return token{"ARG", cursor, value}
}

func newStringToken(cursor int, value string) token {
	return token{"STRING", cursor, value}
}

func isQuote(char rune) bool {
	return char == '"' || char == '\''
}

func isWhitespace(char rune) bool {
	// for now only handling spaces and tabs, but may need to handle
	// newline or carriage return at some point
	return char == ' ' || char == '\t'
}

func findArg(cmd string, cursor int) (t token, newCursor int, err error) {
	value := "hello"
	return newArgToken(cursor, value), cursor + 1, nil
}

// finds a string from a given command string
func FindString(cmd string, cursor int) (t token, newCursor int, err error) {
	// we expect a string should start with a quote
	originalCursor := cursor
	if !isQuote(rune(cmd[cursor])) {
		return token{}, cursor, fmt.Errorf("Expected beginning `\"` or `'` in %q", cmd)
	}
	// we can skip the cursor up if we find the quote
	cursor++

	vBytes := []byte{}
	// get all the individual bytes til we get another quote or end the string
	for cursor < len(cmd) && !isQuote(rune(cmd[cursor])) {
		vBytes = append(vBytes, cmd[cursor])
		cursor++
	}

	// we expect to end with a quote
	if !isQuote(rune(cmd[cursor])) {
		return token{}, cursor, fmt.Errorf("Expected ending `\"` or `'` in %q", cmd)
	}
	cursor++

	return newStringToken(originalCursor, string(vBytes)), cursor + 1, nil
}

// takes in a command string and returns a slice
// of tokens
func tokenize(cmd string) ([]token, error) {
	tokens := []token{}
	cursor := 0
	for cursor < len(cmd) {
		if isQuote(rune(cmd[cursor])) {
			token, updatedCursor, err := FindString(cmd, cursor+1)
			if err != nil {
				log.Fatalf("Encountered error trying to parse command `%s`: %s", cmd, err)
			}
			tokens = append(tokens, token)
			cursor = updatedCursor
		}
	}
	return tokens, nil
}

// cmd ::=
//     | ARG [ARG*]
//     | ARG string

func Parse(cmdStr string) ([]string, error) {
	args := []string{}
	return args, nil
}
