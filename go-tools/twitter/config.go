package main

import (
	"github.com/g8rswimmer/go-twitter/v2"
	"net/http"
)

const TOKEN = "AAAAAAAAAAAAAAAAAAAAAJFMtQEAAAAAu5pA0mpCPQNHYYGnj0gzJmkd66Q%3DSvohga9PryG8KeOnKIGa2J9wQtzzbHxOuHvcs5fKE19n60BhP9"

var client = &twitter.Client{
	Authorizer: authorize{
		Token: TOKEN,
	},
	Client: http.DefaultClient,
	Host:   "https://api.twitter.com",
}
