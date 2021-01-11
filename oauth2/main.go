package main

import (
	"github.com/go-chi/chi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
)

func main() {
	config := oauth2.Config{
		ClientID:     "4001a97e7452c32a4f22",
		ClientSecret: "9c1e5cb6f5f420591535b9c98ac5e076e24ae1a5",
		Endpoint:     github.Endpoint,
	}

	p := NewProvider(config)

	s := NewSessionManager(map[ProviderID]Provider{
		ProviderGitHub: p,
	})

	r := chi.NewRouter()

	r.Post("/oauth/github", s.Provider(ProviderGitHub).Redirect())
	r.Post("/oauth/github/receive", s.Provider(ProviderGitHub).Receive())

	http.ListenAndServe(":8080", r)
}
