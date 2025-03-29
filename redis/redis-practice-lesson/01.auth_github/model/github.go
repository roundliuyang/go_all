package model

import "time"

type GithubToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type GithubUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node-id"`
	AvatarUrl         string    `json:"avatarUrl"`
	GravatarID        string    `json:"gravatarID"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"htmlUrl"`
	FollowersURL      string    `json:"followersURL"`
	FollowingURL      string    `json:"followingUrl"`
	GistsURL          string    `json:"gistsURL"`
	StarredURL        string    `json:"starredURL"`
	SubscriptionURL   string    `json:"subscriptionURL"`
	OrganizationsURL  string    `json:"organizationsURL"`
	ReposURL          string    `json:"reposURL"`
	EventsURL         string    `json:"eventsURL"`
	ReceivedEventsURL string    `json:"receivedEventsURL"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"siteAdmin"`
	Name              string    `json:"name"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             *string   `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"publicRepos"`
	PublicGists       int       `json:"publicGists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	Token             string    `json:"-"`
}
