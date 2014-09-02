package pushover

import (
	"testing"
)

func TestDescription(t *testing.T) {
	p := NewPushoverNotifier("foo", "bla")

	expDescription := "Pushover send to user: bla"

	if p.Description() != expDescription {
		t.Errorf("wrong answer, expected '%s' and got %s", expDescription, p.Description())
	}
}
