package facebookgraph

import (
	"fmt"
	"net/http"
	"os"

	fb "github.com/huandu/facebook/v2"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_fb "github.com/leapforce-libraries/go_facebookgraph/fb"
	ig "github.com/leapforce-libraries/go_facebookgraph/ig"
	"golang.org/x/oauth2"
	oauth2fb "golang.org/x/oauth2/facebook"
)

const (
	apiName     string = "FacebookGraph"
	redirectURL string = "http://localhost:8080/oauth/redirect"
)

type Service struct {
	session *fb.Session
}

type ServiceConfig struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	AccessToken  string
}

func (service *Service) FacebookService() *go_fb.Service {
	return go_fb.NewService(service.session)
}

func (service *Service) InstagramService() *ig.Service {
	return ig.NewService(service.session)
}

// methods
//
func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	conf := &oauth2.Config{
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       serviceConfig.Scopes,
		Endpoint:     oauth2fb.Endpoint,
	}

	token := oauth2.Token{}
	token.AccessToken = serviceConfig.AccessToken

	// Create a client to manage access token life cycle.
	client := conf.Client(oauth2.NoContext, &token)

	// Use OAuth2 client with session.
	_session := &fb.Session{
		Version:    "v10.0",
		HttpClient: client,
	}
	_session.SetDebug(fb.DEBUG_OFF)

	return &Service{_session}, nil
}

func InitToken(config ServiceConfig) *errortools.Error {

	conf := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes:       config.Scopes,
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
			errortools.CaptureFatal(err)
		}

		fmt.Println("AccessToken: ", token.AccessToken)
		fmt.Println("RefreshToken: ", token.RefreshToken)
		fmt.Println("Expiry: ", token.Expiry)

		w.WriteHeader(http.StatusFound)

		return
	})

	http.ListenAndServe(":8080", nil)

	return nil
}
