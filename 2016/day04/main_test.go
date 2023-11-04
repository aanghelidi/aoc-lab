package main

import "testing"

func TestEncryptedName(t *testing.T) {
	testName := "qzmt-zixmtkozy-ivhz-343"
	expected := "very encrypted name"
	actual := decryptName(testName)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
