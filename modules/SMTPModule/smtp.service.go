package SMTPModule

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type ISMTPService interface {
	From(f string) ISMTPService
	To(f []string) ISMTPService
	SendVerifyMail()
}

type smtpService struct {
	sender   string
	receipts []string
}

// constructor
func SMTPService() ISMTPService {
	return &smtpService{}
}

func (s *smtpService) From(f string) ISMTPService {
	s.sender = f
	return s
}

func (s *smtpService) To(r []string) ISMTPService {
	s.receipts = r
	return s
}

func (s *smtpService) SendVerifyMail() {

	fmt.Println(s.sender, s.receipts)
	addr := "sandbox.smtp.mailtrap.io:2525"
	host := "sandbox.smtp.mailtrap.io"
	auth := smtp.PlainAuth("", "71a2c12d073174", "ad75825bd4642c", host)
	dat, err := os.ReadFile("modules/SMTPModule/templates/confirmation.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte(
		"From: jdow@example.com\r\n" +
			"To: deneme@example.com\r\n" +
			"Subject: Test mail\r\n" +
			"Content-Type: multipart/alternative; boundary=\"boundary-string\"\r\n\r\n" +
			"--boundary-string\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n" +
			"Content-Transfer-Encoding: quoted-printable\n" +
			"Content-Disposition: inline\r\n\r\n" +
			string(dat) + "\r\n" + // HTML BODY
			"--boundary-string--",
	)

	err = smtp.SendMail(addr, auth, s.sender, s.receipts, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}
