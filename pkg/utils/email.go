package utils

import (
    "net/smtp"
)

// SendEmail sends an email
func SendEmail(to, subject, body string) error {
    // Set up authentication information.
    auth := smtp.PlainAuth("", "your-email@example.com", "your-email-password", "smtp.example.com")

    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    err := smtp.SendMail("smtp.example.com:587", auth, "your-email@example.com", []string{to}, []byte("Subject: "+subject+"\r\n"+body))
    if err != nil {
        return err
    }
    return nil
}
