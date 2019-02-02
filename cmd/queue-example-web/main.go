package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	log = logrus.WithField("cmd", "queue-example-web")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.WithField("PORT", port).Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	log.Println(http.ListenAndServe(":"+port, nil))
}
