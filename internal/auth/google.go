package auth

import (
	"context"
	"log"

	"github.com/exedary/soulmates/internal/config"
	"github.com/exedary/soulmates/internal/domain/person"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
	googleOauth "golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

var GoogleOauthModule = fx.Provide(New)

type google struct {
	oauthConfiguration *oauth2.Config
}

func New(config *config.Oauth) *google {
	return &google{oauthConfiguration: &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Endpoint:     googleOauth.Endpoint,
		Scopes:       config.Scopes,
	}}
}

func (google *google) SignInWithGoogle(ctx context.Context, personRepository person.Repository, stateHash string) string {
	return google.oauthConfiguration.AuthCodeURL(stateHash)
}

func (google *google) ProcessGoogleCallback(ctx context.Context, personRepository person.Repository, authCode string) {
	tok, err := google.oauthConfiguration.Exchange(ctx, authCode)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := idtoken.Validate(ctx, tok.Extra("id_token").(string), google.oauthConfiguration.ClientID)
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
