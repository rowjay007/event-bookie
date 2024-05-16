package utils

import (
	"net/smtp"
)

// SendEmail sends an email using SMTP
func SendEmail(to, subject, body string) error {
    // SMTP server configuration
    smtpServer := "smtp.example.com"
    smtpPort := "587"
    smtpUsername := "username"
    smtpPassword := "password"

    // Message
    msg := "From: " + smtpUsername + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n" +
        body

    // Send email
    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
    err := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpUsername, []string{to}, []byte(msg))
    return err
}
