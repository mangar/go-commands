package send

import (
	"bytes"
	"fmt"
	"gmail/cmd/defaults"
	"io/ioutil"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

/**
	Modified from https://gist.github.com/jpillora/cb46d183eca0710d909a
	Thank you very much.
**/

const (
	SMTPServer = "smtp.gmail.com"
)

type Sender struct {
	User     string
	Password string
}

func Send() {
	sender := Sender{defaults.Email, defaults.Pwd}

	Receiver := strings.Split(defaults.To, ",")
	Subject := defaults.Subject
	message := getEmailContent()

	// bodyMessage := sender.writeHTMLEmail(Receiver, Subject, message)
	// sender.sendMail(Receiver, Subject, bodyMessage)

	for i := range strings.Split(defaults.To, ",") {
		bodyMessage := sender.writeHTMLEmail(Receiver[i], Subject, message)
		sender.sendMail(Receiver[i], Subject, bodyMessage)
	}

}

func (sender Sender) sendMail(Dest string, Subject, bodyMessage string) {

	// msg := "From: " + sender.User + "\n" +
	// 	"To: " + strings.Join(Dest, ",") + "\n" +
	// 	"Subject: " + Subject + "\n" + bodyMessage

	msg := "From: " + sender.User + "\n" +
		"To: " + Dest + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer), sender.User, []string{Dest}, []byte(msg))

	if err != nil {
		fmt.Printf("Ups.... SMTP error: %s", err)
		return
	}

	fmt.Println("Cheers! Mail sent successfully!")
}

func (sender Sender) writeEmail(dest string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	// receipient := dest
	// receipient := ""

	// for _, user := range dest {
	// 	receipient = receipient + user
	// }

	// header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) writeHTMLEmail(dest string, subject, bodyMessage string) string {
	return sender.writeEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *Sender) writePlainEmail(dest string, subject, bodyMessage string) string {
	return sender.writeEmail(dest, "text/plain", subject, bodyMessage)
}

func getEmailContent() string {
	b, err := ioutil.ReadFile(defaults.Content)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	return str
}
