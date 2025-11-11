package domain

// Instance = one group (no nested groups).
// Placeholder struct definitions for future PRs.

type Instance struct {
	ID          string
	Name        string
	Description string
	NumberE164  string
	VercelURL   string

	// TODO: Policies
	Members []Member
}

type Member struct {
	Phone string
	Name  string
}
