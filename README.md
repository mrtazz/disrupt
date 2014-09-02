# disrupt [![Build Status](https://travis-ci.org/mrtazz/disrupt?branch=master)](https://travis-ci.org/mrtazz/disrupt)

## Overview
Library to enable notifications via multiple providers. Because nothing is
more useful than being interrupted more easily. It's possible to register
multiple notifiers and when the `NotifyAll` call is run, all of them will be
notified.

## Usage

### Library
disrupt can be used as a library from other applications like this:

```
  disrupt.RegisterPushoverNotifier(*pushover_token, *pushover_user)
  disrupt.RegisterTwilioNotifier(*twilio_sid, *twilio_token, *twilio_from, *twilio_to, *sms_max_char)
  disrupt.NotifyAll(title, content)
```

### CLI
disrupt also comes with a simple CLI notifier that takes input via STDIN. The
first line is used as the notification title and all subsequent lines are sent
as the message body. It uses config.ini style parsing for easy scripting and
also allows params to be passed in on the command line:
```
% ./cli-notifier -help
  -debug=false: set debug mode
  -help=false: show help
  -maxchar=160: max number of characters to use for SMS
  -pushover=false: send notifications via pushover
  -pushover_token="": API token for pushover
  -pushover_user="": pushover user key
  -twilio=false: send notifications via twilio
  -twilio_from="": from address for twilio
  -twilio_sid="": twilio sid
  -twilio_to="": address to send to via twilio
  -twilio_token="": twilio token

% cat ~/.config/cli-notifier/config.ini
pushover_token = 999999999999999999
pushover_user  = uuuuuuuuuuuuuuu
pushover       = true

twilio_sid     = ooooooooooooooooooooooo
twilio_token   = 88888888888888888888
twilio_from    = +11111111
twilio_to      = +2222222
twilio         = false
maxchar        = 160
```

## Installation
go get it like it's hot.

```
go get github.com/mrtazz/disrupt
```

