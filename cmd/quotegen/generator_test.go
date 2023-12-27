package quotegen

import (
	"testing"
)

func TestFetchQuotes(t *testing.T) {
	quotes := fetchQuotes(10)
	if len(quotes) != 10 {
		t.Errorf("Expected 10 quotes, got %d", len(quotes))
	}
}

func TestTee(t *testing.T) {

}
