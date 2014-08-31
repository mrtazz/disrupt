package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mrtazz/disrupt"
	"github.com/rakyll/globalconf"
	"io"
	"os"
)

var Debug bool

func main() {
	var (
		pushover_token = flag.String("pushover_token", "", "API token for pushover")
		pushover_user  = flag.String("pushover_user", "", "pushover user key")
		use_pushover   = flag.Bool("pushover", false, "send notifications via pushover")
		twilio_sid     = flag.String("twilio_sid", "", "twilio sid")
		twilio_token   = flag.String("twilio_token", "", "twilio token")
		twilio_from    = flag.String("twilio_from", "", "from address for twilio")
		twilio_to      = flag.String("twilio_to", "", "address to send to via twilio")
		use_twilio     = flag.Bool("twilio", false, "send notifications via twilio")
		sms_max_char   = flag.Int("maxchar", 160, "max number of characters to use for SMS")
		show_help      = flag.Bool("help", false, "show help")
		debug          = flag.Bool("debug", false, "set debug mode")
	)

	conf, _ := globalconf.New("cli-notifier")
	conf.ParseAll()

	if *show_help {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: cli-notifier\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	disrupt.Debug = *debug

	stdinReader := bufio.NewReader(os.Stdin)

	var content string
	title, err := stdinReader.ReadString('\n')

	if err == nil {
		for {
			line, err := stdinReader.ReadString('\n')

			if err == io.EOF {
				break
			}

			content = content + "\n" + line
		}
	}

	if *use_pushover {
		disrupt.RegisterPushoverNotifier(*pushover_token, *pushover_user)
	}
	if *use_twilio {
		disrupt.RegisterTwilioNotifier(*twilio_sid, *twilio_token, *twilio_from, *twilio_to, *sms_max_char)
	}

	disrupt.NotifyAll(title, content)
}
