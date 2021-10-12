package main

import (
	"context"
	"fmt"
	"io/ioutil"
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
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run sign.go <file> <path/to/key>")
		os.Exit(1)
	}

	blobPath := os.Args[1]
	keyPath := os.Args[2]

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
		PassFunc:     cosigncli.GetPass,
		KeyRef:       keyPath,
	}

	sig, err := cosigncli.SignBlobCmd(context.Background(), opt, blobPath, false, "")
	if err != nil {
		fmt.Println("error occured in signing: ", err.Error())
	}

	err = ioutil.WriteFile("./signature", sig, 0644)
	if err != nil {
		fmt.Println("error occured in creating signature file: ", err.Error())
	}
}
