package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"net/http"
	"sync"
	"time"
)

var (
	ErrLoginAttemptNotFound = errors.New("login attempt not found")
)

type Provider interface {
	Redirect() http.HandlerFunc
	Receive() http.HandlerFunc
	registerLoginAttempt(state uuid.UUID)
}

type LoginAttempt struct {
	Date      time.Time
	Code      string
	Completed bool
	Client    *http.Client
}

type provider struct {
	Config            oauth2.Config
	LoginAttempts     map[string]LoginAttempt
	LoginAttemptsLock sync.Mutex
}

func (p *provider) registerLoginAttempt(state uuid.UUID) {
	p.LoginAttemptsLock.Lock()
	defer p.LoginAttemptsLock.Unlock()

	p.LoginAttempts[state.String()] = LoginAttempt{Date: time.Now()}
}

func (p provider) Receive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		if len(code) == 0 {
			http.Error(w, fmt.Sprintf("%s: error parsing OAuth2 code", http.StatusText(http.StatusBadRequest)), http.StatusBadRequest)
			return
		}

		state, err := uuid.Parse(r.FormValue("state"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%s: error parsing OAuth2 state, error: %v", http.StatusText(http.StatusBadRequest), err), http.StatusBadRequest)
			return
		}

		err = p.markLoginAttemptCompleted(state, code)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s: error processing login attempt, error: %v", http.StatusText(http.StatusNotFound), err), http.StatusNotFound)
			return
		}

		token, err := p.Config.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s: error exchanging token, error: %v", http.StatusText(http.StatusInternalServerError), err), http.StatusInternalServerError)
			return
		}

		tokenSource := p.Config.TokenSource(r.Context(), token)

		p.setClient(state, oauth2.NewClient(r.Context(), tokenSource))
	}
}

func (p provider) Redirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := uuid.New()
		go p.registerLoginAttempt(state)
		http.Redirect(w, r, p.Config.AuthCodeURL(state.String(), oauth2.AccessTypeOnline), http.StatusSeeOther)
	}
}

func (p *provider) markLoginAttemptCompleted(state uuid.UUID, code string) error {
	p.LoginAttemptsLock.Lock()
	defer p.LoginAttemptsLock.Unlock()
	attempt, ok := p.LoginAttempts[state.String()]
	if !ok {
		return ErrLoginAttemptNotFound
	}

	attempt.Completed = true
	attempt.Code = code

	p.LoginAttempts[state.String()] = attempt

	return nil
}

func (p *provider) setClient(state uuid.UUID, client *http.Client) {
	p.LoginAttemptsLock.Lock()
	defer p.LoginAttemptsLock.Unlock()

	attempt, ok := p.LoginAttempts[state.String()]
	if !ok {
		return
	}
	attempt.Client = client
	p.LoginAttempts[state.String()] = attempt
}

func NewProvider(config oauth2.Config) Provider {
	return &provider{
		Config: config,
	}
}

type ProviderID string

const (
	ProviderGoogle ProviderID = "google"
	ProviderGitHub ProviderID = "github"
)

type ProviderGetter interface {
	Provider(provider ProviderID) Provider
}

type SessionManager interface {
	ProviderGetter
}

type sessionManager struct {
	providers map[ProviderID]Provider
}

func (s sessionManager) Provider(provider ProviderID) Provider {
	v, ok := s.providers[provider]
	if !ok {
		return nil
	}
	return v
}

func NewSessionManager(providers map[ProviderID]Provider) SessionManager {
	return &sessionManager{
		providers: providers,
	}
}
