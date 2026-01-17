package github

// ReviewPayload represents the incoming GitHub Webhook data
type ReviewPayload struct {
	Action      string     `json:"action"` // e.g., "submitted"
	Review      ReviewData `json:"review"`
	PullRequest PRData     `json:"pull_request"`
	Repository  RepoData   `json:"repository"`
	Sender      GitHubUser `json:"sender"` // The person who reviewed
}

type ReviewData struct {
	State   string `json:"state"`    // "changes_requested", "approved", "commented"
	Body    string `json:"body"`     // The main review comment
	HTMLURL string `json:"html_url"` // Link to the review
}

type PRData struct {
	Title   string     `json:"title"`
	Number  int        `json:"number"`
	User    GitHubUser `json:"user"` // The person who opened the PR
	HTMLURL string     `json:"html_url"`
}

type GitHubUser struct {
	Login string `json:"login"` // GitHub Username (e.g., "JuniorDev123")
}

type RepoData struct {
	FullName string `json:"full_name"` // e.g., "Association/Website-Backend"
}
