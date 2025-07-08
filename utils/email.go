package utils

import "log"

func SendPasswordResetEmail(to, subject, body string) error {
	log.Printf("DEBUG: sending email to %s with subject '%s' and body:\n%s\n", to, subject, body)
	// Implement here with net/smtp

	return nil
}
