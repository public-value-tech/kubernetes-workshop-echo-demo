package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	livenessCheck  = true
	readinessCheck = true
)

func main() {
	http.HandleFunc("/readiness", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("readiness check")
		if readinessCheck {
			fmt.Fprintln(w, "OK")
		} else {
			http.Error(w, "not ready", http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/liveness", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("liveness check")
		if livenessCheck {
			fmt.Fprintln(w, "OK")
		} else {
			http.Error(w, "not alive", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /")
		fmt.Fprintln(w, "Hello. Go to /env or /configmap to see something")
	})

	http.HandleFunc("/unready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /unready")
		readinessCheck = false
		fmt.Fprintln(w, "set service to unready")
	})

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /ready")
		readinessCheck = true
		fmt.Fprintln(w, "set service to ready")
	})

	http.HandleFunc("/unalive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /unalive")
		livenessCheck = false
		fmt.Fprintln(w, "set service to unalive")
	})

	http.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /alive")
		livenessCheck = true
		fmt.Fprintln(w, "set service to alive")
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /env")
		fmt.Fprintf(w, "Hello from /env: %q", os.Getenv("ECHO_MESSAGE"))
	})

	http.HandleFunc("/configmap", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /configmap")
		all, err := os.ReadFile("/etc/configmap/message")
		if err != nil {
			fmt.Fprintf(w, "Error reading configmap: %q", err)
			return
		}
		fmt.Fprintf(w, "Hello from /configmap: %q", all)
	})

	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request to /secret")
		all, err := os.ReadFile("/etc/secret/message")
		if err != nil {
			fmt.Fprintf(w, "Error reading secret: %q", err)
			return
		}
		fmt.Fprintf(w, "Hello from /secret: %q", all)
	})

	log.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
