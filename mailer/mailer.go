package mailer

import (
	"github.com/hw-cs-reps/platform/config"

	"fmt"
	"net/smtp"
	"strings"
	"time"
)

// formatTo formats the recipients by separating them with commas for the email
// header.
func formatTo(to []string) string {
	var str strings.Builder
	for i, t := range to {
		str.WriteString("<" + t + ">")
		if i != len(to)-1 { // separate them by commas
			str.WriteString(", ")
		}
	}
	return str.String()
}

// Email sends an email to a specific email address.
func Email(to []string, title string, message string) (err error) {
	from := config.Config.EmailAddress
	full := "From: <" + from + ">\n" +
		"To: " + formatTo(to) + "\n" +
		"Date: " + time.Now().Format(time.RFC1123Z) + "\n" +
		"Subject: " + title + "\n\n" +
		message

	fmt.Println(full)

	err = smtp.SendMail(config.Config.EmailSMTPServer,
		smtp.PlainAuth("", from, config.Config.EmailPassword, strings.Split(config.Config.EmailSMTPServer, ":")[0]),
		from, to, []byte(full))

	return err
}
