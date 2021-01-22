package facebookgraph

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	oauth2fb "golang.org/x/oauth2/facebook"

	go_fb "github.com/leapforce-libraries/go_facebookgraph/fb"
	ig "github.com/leapforce-libraries/go_facebookgraph/ig"

	fb "github.com/huandu/facebook/v2"
)

const APIName string = "FacebookGraph"

// GoogleAdminDirectory stores GoogleAdminDirectory configuration
//
type Service struct {
	session *fb.Session
}

func (service *Service) FacebookService() *go_fb.Service {
	return go_fb.NewService(service.session)
}

func (service *Service) InstagramService() *ig.Service {
	return ig.NewService(service.session)
}

// methods
//
func NewService(clientID string, clientSecret string, scopes []string, accessToken string) *Service {
	service := Service{}

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

	service.session = _session

	return &service
}

func InitToken(clientID string, clientSecret string, scopes []string) {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/oauth/redirect",
		Scopes:       scopes,
		Endpoint:     oauth2fb.Endpoint,
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Create a new redirect route
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		//
		// get authorization code
		//
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		code := r.FormValue("code")

		fmt.Println("code: ", code)

		token, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			return
		}

		fmt.Println("AccessToken: ", token.AccessToken)
		fmt.Println("RefreshToken: ", token.RefreshToken)
		fmt.Println("Expiry: ", token.Expiry)

		w.WriteHeader(http.StatusFound)

		return
	})

	http.ListenAndServe(":8080", nil)

	return
}
