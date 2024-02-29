package email

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
)

func SendEmailWithAttachment(to, subject, body, filePath string) error {
	// Sender and recipient email addresses
	from := "sender@example.com"
	password := "password" // Provide your email password here
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	// Read the PDF file content
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Message headers
	headers := map[string]string{
		"From":         from,
		"To":           to,
		"Subject":      subject,
		"MIME-version": "1.0",
		"Content-Type": `multipart/mixed; boundary="f34sFSD6"`,
	}

	// Message body
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body + "\r\n"

	// Add PDF attachment
	message += "--f34sFSD6\r\n"
	message += `Content-Type: application/pdf; name="attachment.pdf"` + "\r\n"
	message += "Content-Disposition: attachment; filename=\"attachment.pdf\"\r\n"
	message += "\r\n" + string(fileData) + "\r\n"
	message += "--f34sFSD6--"

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
