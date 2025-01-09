package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var urlStore = struct {
	sync.RWMutex
	data map[string]string
}{data: make(map[string]string)}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil || requestData.URL == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var shortKey string
	for {
		shortKey = GenerateRandomString(6)
		urlStore.RLock()
		_, exists := urlStore.data[shortKey]
		urlStore.RUnlock()
		if !exists {
			break
		}
	}

	urlStore.Lock()
	urlStore.data[shortKey] = requestData.URL
	urlStore.Unlock()
	shortURL := "http://localhost:8080/" + shortKey
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/")

	urlStore.RLock()
	originalURL, exists := urlStore.data[shortKey]
	urlStore.RUnlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", RedirectHandler)
	http.HandleFunc("/shorten", ShortenURLHandler)
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
