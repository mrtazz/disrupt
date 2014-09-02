package twilio

import (
	"testing"
)

func TestDescription(t *testing.T) {
	p := NewTwilioNotifier("xxx", "yyy", "123", "456")

	expDescription := "Twilio from: 123 to 456."

	if p.Description() != expDescription {
		t.Errorf("wrong answer, expected '%s' and got %s", expDescription, p.Description())
	}
}
