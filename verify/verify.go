package main

import (
	"context"
	"fmt"
	"os"

	cosigncli "github.com/sigstore/cosign/cmd/cosign/cli"
	fulcioclient "github.com/sigstore/fulcio/pkg/client"
)

const (
	rekorServerEnvKey     = "REKOR_SERVER"
	defaultRekorServerURL = "https://rekor.sigstore.dev"
	defaultOIDCIssuer     = "https://oauth2.sigstore.dev/auth"
	defaultOIDCClientID   = "sigstore"
	cosignPasswordEnvKey  = "COSIGN_PASSWORD"
)

func GetRekorServerURL() string {
	url := os.Getenv(rekorServerEnvKey)
	if url == "" {
		url = defaultRekorServerURL
	}
	return url
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run sign.go <file> <path/to/signature> <path/to/pubkey>")
		os.Exit(1)
	}

	msgPath := os.Args[1]
	sigPath := os.Args[2]
	keyPath := os.Args[3]

	sk := false
	idToken := ""
	rekorSeverURL := GetRekorServerURL()
	fulcioServerURL := fulcioclient.SigstorePublicServerURL

	opt := cosigncli.KeyOpts{
		Sk:           sk,
		IDToken:      idToken,
		RekorURL:     rekorSeverURL,
		FulcioURL:    fulcioServerURL,
		OIDCIssuer:   defaultOIDCIssuer,
		OIDCClientID: defaultOIDCClientID,
		KeyRef:       keyPath,
	}

	err := cosigncli.VerifyBlobCmd(context.Background(), opt, "", sigPath, msgPath)
	if err != nil {
		fmt.Println("error occured in verifying: ", err.Error())
	}
}
