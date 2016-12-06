package ghapi

import "time"

type GitHubEventPayload struct {
	Repository RepositoryPayload `json:"repository"`
	Sender     User              `json:"sender"`
}

type PingEventPayload struct {
	GitHubEventPayload
	Zen    string `json:"zen"`
	HookID int    `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		ID     int      `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSSL string `json:"insecure_ssl"`
			Secret      string `json:"secret"`
			URL         string `json:"url"`
		} `json:"config"`
		UpdatedAt    time.Time `json:"updated_at"`
		CreatedAt    time.Time `json:"created_at"`
		URL          string    `json:"url"`
		TestURL      string    `json:"test_url"`
		PingURL      string    `json:"ping_url"`
		LastResponse struct {
			Code    *int   `json:"code"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"last_response"`
	} `json:"hook"`
}

type PullRequestEventPayload struct {
	GitHubEventPayload
	Action      PullRequestAction  `json:"action"`
	Number      int                `json:"number"`
	PullRequest PullRequestPayload `json:"pull_request"`
}

type IssueCommentEventPayload struct {
	GitHubEventPayload
	Action  string              `json:"action"`
	Issue   IssuePayload        `json:"issue"`
	Comment IssueCommentPayload `json:"comment"`
}

type Pusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PushEventPayload struct {
	GitHubEventPayload
	Ref     string  `json:"ref"`
	Before  string  `json:"before"`
	After   string  `json:"after"`
	Created bool    `json:"created"`
	Deleted bool    `json:"deleted"`
	Forced  bool    `json:"forced"`
	BaseRef *string `json:"base_ref"`
	Compare string  `json:"compare"`
	Commits []struct {
		ID        string `json:"id"`
		Distinct  bool   `json:"distinct"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []string `json:"added"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	} `json:"commits"`
	HeadCommit struct {
		ID        string `json:"id"`
		Distinct  bool   `json:"distinct"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []string `json:"added"`
		Removed  []string `json:"removed"`
		Modified []string `json:"modified"`
	} `json:"head_commit"`
}
