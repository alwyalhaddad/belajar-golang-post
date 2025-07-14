package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendPasswordResetEmail(to, subject, body string) error {
	log.Printf("DEBUG: sending email to %s with subject '%s' and body:\n%s\n", to, subject, body)

	// Get credentials from environment variables
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("FROM_EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if from == "" || password == "" || smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("email credentials (FROM_EMAIL, FROM_EMAIL_PASSWORD, SMTP_HOST, SMTP_PORT) must be set as environment variables")
	}

	// Format email message
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject" + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text-plain; charset=\"UTF-8\";\n" + // Important for modern email
		"\r\n" + // Empty line to separates the header from the body
		body)

	// Authorization SMTP
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email from %s to %s via %s: %w", from, to, addr, err)
	}

	log.Printf("Email Successfuly sent to %s\n", to)
	return nil
}
