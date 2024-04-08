package main

import (
	"backend-trainee-assignment-2024/config"
	"errors"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	host       string
	healthPath string
	basePath   string

	attempts = 20
)

//goland:noinspection HttpUrlsUsage
func TestMain(m *testing.M) {
	cfg := config.MustConfig().HTTP

	host = "localhost:"
	dockerHost, exists := os.LookupEnv("HOST")
	if exists {
		host = dockerHost
	}

	host += cfg.Port
	healthPath = "http://" + host + "/healthz"
	basePath = "http://" + host

	err := healthCheck(attempts)
	if err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	log.Printf("Integration tests: host %s is available", host)

	code := m.Run()
	os.Exit(code)
}

func healthCheck(attempts int) error {
	var (
		err  error
		resp *http.Response
	)

	for attempts > 0 {
		resp, err = http.Get(healthPath)
		if err == nil {
			break
		}
		if attempts%3 == 0 {
			log.Printf("Integration tests: url %s is not available, attempts left: %d", healthPath, attempts)
		}

		time.Sleep(time.Second * 2)
		attempts--
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("")
}
