package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubOAuth struct {
	config *oauth2.Config
}

type GitHubUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func NewGitHubOAuth() *GitHubOAuth {
	ghci := os.Getenv("GITHUB_CLIENT_ID")
	ghcs := os.Getenv("GITHUB_CLIENT_SECRET")
	gru := os.Getenv("GITHUB_REDIRECT_URL")
	if ghci == "" || ghcs == "" || gru == "" {
		panic("github oauth vars cannot be empty")
	}
	config := &oauth2.Config{
		ClientID:     ghci,
		ClientSecret: ghcs,
		Endpoint:     github.Endpoint,
		RedirectURL:  gru,
		Scopes:       []string{"user:email", "read:user"},
	}
	return &GitHubOAuth{config: config}
}

func (g *GitHubOAuth) GetAuthURL(state string) string {
	return g.config.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (g *GitHubOAuth) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return g.config.Exchange(ctx, code)
}

func (g *GitHubOAuth) GetGitHubUser(ctx context.Context, token *oauth2.Token) (*GitHubUser, error) {
	client := g.config.Client(ctx, token)

	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned status %d", resp.StatusCode)
	}
	var user GitHubUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err

	}

	return &user, nil
}
