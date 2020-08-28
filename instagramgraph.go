package instagramgraph

import (
	"fmt"
	"strings"

	types "github.com/Leapforce-nl/go_types"
	fb "github.com/huandu/facebook/v2"
	"golang.org/x/oauth2"
	oauth2fb "golang.org/x/oauth2/facebook"
)

const apiName string = "InstagramGraph"

// GoogleAdminDirectory stores GoogleAdminDirectory configuration
//
type InstagramGraph struct {
	baseURL string
	//accessToken string
	session *fb.Session
}

// methods
//
func NewInstagramGraph(baseURL string, clientID string, clientSecret string, scopes []string, accessToken string, isLive bool) (*InstagramGraph, error) {
	ig := InstagramGraph{}
	ig.baseURL = baseURL
	//ig.accessToken = accessToken

	// Get Facebook access token.
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/oauth/redirect",
		Scopes:       scopes,
		Endpoint:     oauth2fb.Endpoint,
	}

	token := oauth2.Token{}
	token.AccessToken = accessToken
	//token.Expiry, _ = time.Parse("2006-01-02", "2020-10-01")

	// Create a client to manage access token life cycle.
	client := conf.Client(oauth2.NoContext, &token)

	// Use OAuth2 client with session.
	_session := &fb.Session{
		Version:    "v2.4",
		HttpClient: client,
	}
	_session.SetDebug(fb.DEBUG_OFF)

	ig.session = _session

	return &ig, nil
}

func (ig *InstagramGraph) Validate() error {
	if ig.baseURL == "" {
		return &types.ErrorString{fmt.Sprintf("%s baseURL not provided", apiName)}
	}

	if !strings.HasSuffix(ig.baseURL, "/") {
		ig.baseURL = ig.baseURL + "/"
	}

	return nil
}
