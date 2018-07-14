package mail

import (
	"bytes"
	"html/template"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// Info ..
type Info struct {
	Repo string
	Time time.Time
	Size uint64
}

// SendMail ..
func (info Info) SendMail() (err error) {
	var tmpl *template.Template
	var bodyBuf = new(bytes.Buffer)

	tmpl, err = template.New("mail").Parse(TemplateEmail)
	if err != nil {
		return
	}

	if err = tmpl.Execute(bodyBuf, info); err != nil {
		return
	}

	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", "test@tosiney.com", viper.GetString("Setting.Name"))
	msg.SetHeader("To", viper.GetString("Mail.EMail"))
	msg.SetHeader("Subject", viper.GetString("Setting.Name"))
	msg.SetBody("text/html", bodyBuf.String())

	dialer := gomail.NewDialer(viper.GetString("Mail.SMTP"), viper.GetInt("Mail.Port"), viper.GetString("Mail.Username"), viper.GetString("Mail.Password"))

	err = dialer.DialAndSend(msg)

	return
}
