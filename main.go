package main

import (
	"log"
	"os"

	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	hum "github.com/grokify/gotilla/net/httputilmore"
	"github.com/jessevdk/go-flags"

	ro "github.com/grokify/oauth2more/ringcentral"
)

type CliOptions struct {
	EnvFile string `short:"e" long:"env" description:"Env filepath"`
	Url     string `short:"u" long:"url" description:"URL" required:"true"`
}

func main() {
	opts := CliOptions{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadDotEnvFirst(opts.EnvFile, os.Getenv("ENV_PATH"), "./.env")
	if err != nil {
		log.Fatal(err)
	}

	fmtutil.PrintJSON(opts)

	httpClient, err := ro.NewClientPassword(
		ro.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		},
		ro.PasswordCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD")})

	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Get(opts.Url)
	if err != nil {
		log.Fatal(err)
	}

	if err := hum.PrintResponse(resp, true); err != nil {
		log.Fatal(err)
	}
}
