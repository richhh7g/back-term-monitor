package resend_client

type SendEmailRequest struct {
	From    string            `json:"from"`    // The email address of the sender
	To      []string          `json:"to"`      // The email addresses of the recipients
	Text    string            `json:"text"`    // The plain text content of the email
	Html    string            `json:"html"`    // The HTML content of the email
	Subject string            `json:"subject"` // The subject of the email
	Headers map[string]string `json:"headers"` // The headers of the email
}

type SendEmailResponse struct {
	ID string `json:"id"` // The ID of the email
}
