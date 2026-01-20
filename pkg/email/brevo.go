package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type brevoEmail struct {
	Sender      brevoContact   `json:"sender"`
	To          []brevoContact `json:"to"`
	Subject     string         `json:"subject"`
	HTMLContent string         `json:"htmlContent"`
}

type brevoContact struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// Client handles email sending via Brevo API
type Client struct {
	apiKey       string
	contactEmail string
}

// NewClient creates a new Brevo email client
func NewClient(apiKey, contactEmail string) *Client {
	return &Client{
		apiKey:       apiKey,
		contactEmail: contactEmail,
	}
}

// APIKey returns the configured API key
func (c *Client) APIKey() string {
	return c.apiKey
}

// ContactEmail returns the configured contact email
func (c *Client) ContactEmail() string {
	return c.contactEmail
}

// SendContactNotification sends a contact form notification email
func (c *Client) SendContactNotification(name, email, phone, service, message string) error {
	return SendContactNotification(c.apiKey, c.contactEmail, name, email, phone, service, message)
}

func SendContactNotification(apiKey, toEmail, name, email, phone, service, message string) error {
	subject := fmt.Sprintf("New Contact Form Submission from %s", name)

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #16a34a; color: white; padding: 20px; text-align: center; }
        .content { padding: 20px; background: #f9f9f9; }
        .field { margin-bottom: 15px; }
        .label { font-weight: bold; color: #16a34a; }
        .value { margin-top: 5px; }
        .message { background: white; padding: 15px; border-left: 4px solid #16a34a; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>New Contact Form Submission</h1>
        </div>
        <div class="content">
            <div class="field">
                <div class="label">Name:</div>
                <div class="value">%s</div>
            </div>
            <div class="field">
                <div class="label">Email:</div>
                <div class="value"><a href="mailto:%s">%s</a></div>
            </div>
            <div class="field">
                <div class="label">Phone:</div>
                <div class="value">%s</div>
            </div>
            <div class="field">
                <div class="label">Service Interested In:</div>
                <div class="value">%s</div>
            </div>
            <div class="field">
                <div class="label">Message:</div>
                <div class="message">%s</div>
            </div>
        </div>
    </div>
</body>
</html>
`, name, email, email, phone, service, message)

	payload := brevoEmail{
		Sender: brevoContact{
			Email: "noreply@cloverbeltconstructionwi.com",
			Name:  "Cloverbelt Construction Website",
		},
		To: []brevoContact{
			{Email: toEmail},
		},
		Subject:     subject,
		HTMLContent: htmlContent,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal email payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.brevo.com/v3/smtp/email", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("brevo API returned status %d", resp.StatusCode)
	}

	return nil
}
