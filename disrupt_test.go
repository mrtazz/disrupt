package disrupt

import (
	"testing"
)

type MockNotifier struct {
}

func NewMockNotifier() *MockNotifier {
	return &MockNotifier{}
}

func (m *MockNotifier) Notify(title string, messages string) error {
	return nil
}

func (m *MockNotifier) Description() string {
	return ""
}

func TestNotifyAll(t *testing.T) {
	p := NewMockNotifier()

	RegisterGenericNotifier(p)
	err := NotifyAll("foo", "bla")

	if err != nil {
		t.Errorf("something went wrong: %s", err.Error())
	}
}
