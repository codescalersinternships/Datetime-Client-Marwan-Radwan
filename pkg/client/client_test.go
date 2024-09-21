package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDateTimeJson(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		expectedDateTime := "2023-10-01T12:00:00Z"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/datetime/json" {
				t.Fatalf("Expected path to be /datetime/json, got %s", r.URL.Path)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"datetime":"` + expectedDateTime + `"}`))
		}))
		defer server.Close()

		client := NewClient(server.URL)
		dateTime, err := client.GetDateTime(true)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if dateTime != expectedDateTime {
			t.Fatalf("Expected %s, got %s", expectedDateTime, dateTime)
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		client := NewClient(server.URL)
		_, err := client.GetDateTime(true)
		if err == nil {
			t.Fatalf("Expected an error, got none")
		}
	})
}

func TestGetDateTimePlain(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		expectedDateTime := "2023-10-01T12:00:00Z"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/datetime/plain" {
				t.Fatalf("Expected path to be /datetime/plain, got %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedDateTime))
		}))
		defer server.Close()

		client := NewClient(server.URL)
		dateTime, err := client.GetDateTime(false)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if dateTime != expectedDateTime {
			t.Fatalf("Expected %s, got %s", expectedDateTime, dateTime)
		}
	})

	t.Run("Invalid Request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		client := NewClient(server.URL)
		_, err := client.GetDateTime(false)
		if err == nil {
			t.Fatalf("Expected an error, got none")
		}
	})
}
