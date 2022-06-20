package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "20")
}

// stopped here - https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#complete-the-scaffolding
