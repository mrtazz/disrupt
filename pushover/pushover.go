package pushover

import (
	"bitbucket.org/kisom/gopush/pushover"
	"fmt"
	"github.com/mrtazz/disrupt/helpers"
)

type PushoverNotifier struct {
	Identity pushover.Identity
	User     string
	Verbose  bool
}

func NewPushoverNotifier(token string, user string) *PushoverNotifier {
	p := PushoverNotifier{}
	p.Identity = pushover.Authenticate(token, user)
	p.User = user
	p.Verbose = false

	return &p
}

// Notify interface method
func (p *PushoverNotifier) Notify(title string, message string) error {
	pushover.Verbose = p.Verbose
	helpers.PrintIfDebug("Sending pushover message...\n")
	if len(message) > 512 {
		message = message[:512]
	}
	pushover.Notify_titled(p.Identity, message, title)
	return nil
}

// Description interface method
func (p *PushoverNotifier) Description() string {
	return fmt.Sprintf("Pushover send to user: %s", p.User)
}
