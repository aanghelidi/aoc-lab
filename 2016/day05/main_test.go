package main

import "testing"

func TestRetrievePassword(t *testing.T) {
	testInput := "abc"
	expected := "18f47a30"
	actual := retrievePassword(testInput, false)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestRetrievePassword2(t *testing.T) {
	testInput := "abc"
	expected := "05ace8e3"
	actual := retrievePassword(testInput, true)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
