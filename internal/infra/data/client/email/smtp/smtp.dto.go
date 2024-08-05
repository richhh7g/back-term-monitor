package smtp_client

type SendParams struct {
	To      []string
	From    string
	Subject string
	Text    string
	Html    string
}
