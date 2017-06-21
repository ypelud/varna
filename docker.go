package main

// DockerPushData ...
type DockerPushData struct {
	PushedAt int    `json:"pushed_at"`
	Tag      string `json:"tag"`
	Pusher   string `json:"pusher"`
}

// DockerRepository ...
type DockerRepository struct {
	Status          string `json:"status"`
	Description     string `json:"description"`
	IsTrusted       bool   `json:"is_trusted"`
	FullDescription string `json:"full_description"`
	Repo            string `json:"repo_url"`
	Owner           string `json:"owner"`
	IsOfficial      bool   `json:"is_official"`
	IsPrivate       bool   `json:"is_private"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	StarCount       int    `json:"star_count"`
	CommentCount    int    `json:"comment_count"`
	DateCreated     int    `json:"date_created"`
	Dockerfile      string `json:"dockerfile"`
	RepoName        string `json:"repo_name"`
}

// DockerRelease ...
type DockerRelease struct {
	PushedData DockerPushData   `json:"push_data"`
	Callback   string           `json:"callback_url"`
	Repository DockerRepository `json:"repository"`
}
