package main

import (
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"net/http"
)

type Provider interface {
	Redirect() http.HandlerFunc
	Receive() http.HandlerFunc
}

type provider struct {
	Config oauth2.Config
}

func (p provider) Receive() http.HandlerFunc {
	panic("implement me")
}

func (p provider) Redirect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := uuid.New()
		http.Redirect(w, r, p.Config.AuthCodeURL(state.String(), oauth2.AccessTypeOnline), http.StatusTemporaryRedirect)
	}
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
