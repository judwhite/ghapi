package ghapi

type TeamPayload struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	// "pull", "push", "admin"
	Permission      string `json:"permission"`
	URL             string `json:"url"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
	// "secret" - only visible to organization owners and members of this team.
	// "closed" - visible to all members of this organization.
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
}
