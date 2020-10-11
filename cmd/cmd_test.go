package cmd

import "testing"

func TestFindString(t *testing.T) {
	expected := newStringToken(0, "hello")
	actual, _, err := FindString("\"hello\"", 0)
	if err != nil {
		t.Errorf("Got unexpected error %s", err)
	} else if actual != expected {
		t.Errorf("Expected `%v` to equal `%v`", actual, expected)
	}

	expected = newStringToken(0, " hello")
	actual, _, err = FindString("\" hello\"", 0)
	if err != nil {
		t.Errorf("Got unexpected error %s", err)
	} else if actual != expected {
		t.Errorf("Expected `%v` to equal `%v`", actual, expected)
	}

	expected = newStringToken(0, " hello world! There are many words in this string!")
	actual, _, err = FindString("\" hello world! There are many words in this string!\"", 0)
	if err != nil {
		t.Errorf("Got unexpected error %s", err)
	} else if actual != expected {
		t.Errorf("Expected `%v` to equal `%v`", actual, expected)
	}
}
