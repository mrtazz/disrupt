package twilio

import (
	"fmt"
	"github.com/mrtazz/disrupt/helpers"
	"github.com/sfreiberg/gotwilio"
)

type TwilioNotifier struct {
	Client  *gotwilio.Twilio
	From    string
	To      string
	MaxChar int
}

func NewTwilioNotifier(sid string, token string, from string, to string) *TwilioNotifier {
	t := TwilioNotifier{}
	t.Client = gotwilio.NewTwilioClient(sid, token)
	t.From = from
	t.To = to
	t.MaxChar = 160
	return &t
}

func (t *TwilioNotifier) Notify(title string, msg string) error {
	msg = fmt.Sprintf("%s\n%s", title, msg)
	helpers.PrintIfDebug("Sending twilio message...\n")
	if len(msg) > t.MaxChar {
		msg = msg[:t.MaxChar]
	}
	_, _, err := t.Client.SendSMS(t.From, t.To, msg, "", "")
	return err
}

func (t *TwilioNotifier) Description() string {
	return fmt.Sprintf("Twilio from: %s to %s.", t.From, t.To)
}
