// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main provides a CLI for generating a JWT.
package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/oauth2/google"
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

const (
	pemPrivateKeyType  = "PRIVATE KEY"
	jwtHeaderKeyID     = "kid"
	jwtHeaderType      = "typ"
	jwtHeaderTypeValue = "JWT"
	audience           = "cisco-auth"
)

var (
	saLocation = flag.String("service-account", "/keys/sa.json", "Location of the Service Account JSON file obtained from GCP")
)

// Data models the dynamic data payload.
type Data struct {
	TenantName string `json:"tenant_name"`
	TenantID   string `json:"tenant_id"`
	POCEmail   string `json:"poc_email"`
}

// Claims models the Payload of the JWT.
type Claims struct {
	Audience string           `json:"aud,omitempty"`
	IssuedAt *jwt.NumericDate `json:"iat,omitempty"` // value will contain the numeric Unix timestamp at which the token was generated
	Issuer   string           `json:"iss,omitempty"` // value will contain Service Account's Email address
	Subject  string           `json:"sub,omitempty"` // value will contain Service Account's Email address
	Data     Data             `json:"data"`
}

func main() {
	flag.Parse()

	data, err := ioutil.ReadFile(*saLocation)
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(data, "" /*scope (unneeded, so blank)*/)
	if err != nil {
		log.Fatal(err)
	}

	pemData, _ := pem.Decode(conf.PrivateKey)
	if pemData == nil {
		log.Fatal("Could not parse Private Key")
	}
	if pemData.Type != pemPrivateKeyType {
		log.Fatal("Could not parse Private Key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(pemData.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{
			Algorithm: jose.RS256,
			Key:       privateKey,
		},
		&jose.SignerOptions{
			ExtraHeaders: map[jose.HeaderKey]interface{}{
				jwtHeaderKeyID: conf.PrivateKeyID,
				jwtHeaderType:  jwtHeaderTypeValue,
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	token, err := jwt.Signed(signer).Claims(Claims{
		Audience: audience,
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Issuer:   conf.Email,
		Subject:  conf.Email,
		Data: Data{
			TenantName: "tenant_name_here",
			TenantID:   "tenant_id_here",
			POCEmail:   "poc_email_here",
		},
	}).CompactSerialize()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Token generated:\n\n=== BEGIN TOKEN ===\n%s\n=== END TOKEN ===\n", token)
}
