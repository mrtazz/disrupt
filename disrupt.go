// library to send notifications to users via different ways
package disrupt

import (
	"github.com/mrtazz/disrupt/helpers"
	"github.com/mrtazz/disrupt/pushover"
	"github.com/mrtazz/disrupt/twilio"
	"sync"
)

type Notifier interface {
	Notify(title string, message string) error
	Description() string
}

var Debug bool
var Notifiers []Notifier
var NotifierMutex *sync.Mutex

func init() {
	Notifiers = make([]Notifier, 0, 0)
	Debug = false
	helpers.Debug = Debug
	NotifierMutex = &sync.Mutex{}
}

// send a notification via all registered notifiers
//
// Parameters:
// title string - the title of the message
// message string - the body of the message
//
// Returns an error object if any notifier failed
func NotifyAll(title string, message string) error {

	var err error

	for _, notifier := range Notifiers {
		err = notifier.Notify(title, message)
	}

	return err
}

func RegisterPushoverNotifier(token string, user string) error {
	p := pushover.NewPushoverNotifier(token, user)
	p.Verbose = Debug
	NotifierMutex.Lock()
	defer NotifierMutex.Unlock()
	Notifiers = append(Notifiers, p)
	return nil
}

func RegisterTwilioNotifier(sid string, token string, from string, to string, maxchar int) error {
	t := twilio.NewTwilioNotifier(sid, token, from, to)
	t.MaxChar = maxchar
	NotifierMutex.Lock()
	defer NotifierMutex.Unlock()
	Notifiers = append(Notifiers, t)
	return nil
}
