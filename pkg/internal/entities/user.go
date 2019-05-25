package entities

type UserService interface {
	// User
	CreateUser(u *User) error
	GetUserByID(id uint) (*User, error)
	UpdateUser(u *User) error
}

type User struct {
	ID                 uint   `json:"id"`
	Email              string `json:"email"`
	PasswordHash       string `json:"password_hash"`
	GithubAuthToken    string `json:"github_auth_token"`
	BitbucketAuthToken string `json:"bitbucket_auth_token"`
	GitlabAuthToken    string `json:"gitlab_auth_token"`
	GoogleAuthToken    string `json:"google_auth_token"`
}
