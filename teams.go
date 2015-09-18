package ghapi

type TeamPayload struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Slug string `json:"slug"`
	// "pull", "push", "admin"
	Permission      string `json:"permission"`
	Url             string `json:"url"`
	MembersUrl      string `json:"members_url"`
	RepositoriesUrl string `json:"repositories_url"`
	// "secret" - only visible to organization owners and members of this team.
	// "closed" - visible to all members of this organization.
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
}
