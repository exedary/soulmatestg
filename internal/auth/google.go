package auth

import (
	"context"
	"log"
	"os"

	"github.com/exedary/soulmates/internal/domain/person"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

type authCodeDto struct {
	authCode string `json:"code"`
}

var config = &oauth2.Config{
	ClientID:     os.Getenv("g_services_client_id"),
	ClientSecret: os.Getenv("g_services_client_secret"),
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
	Endpoint:     google.Endpoint,
}

func SignInWithGoogle(ctx context.Context, personRepository person.Repository, stateHash string) string {
	return config.AuthCodeURL(stateHash)
}

func ProcessGoogleCallback(ctx context.Context, personRepository person.Repository, authCode string) {
	tok, err := config.Exchange(ctx, authCode)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := idtoken.Validate(ctx, tok.Extra("id_token").(string), config.ClientID)
	googleIdentifier := payload.Claims["sub"].(string)
	googleEmail := payload.Claims["email"].(string)
	googleGivenName := payload.Claims["given_name"].(string)

	existedPerson, err := personRepository.GetByExternalId(ctx, googleIdentifier)

	if existedPerson != nil {

	} else {
		person := person.NewPerson(googleGivenName, googleEmail, googleIdentifier)
		id, err := personRepository.Create(ctx, person)
		log.Printf("%s", id)
		if err != nil {

		}
	}

}
