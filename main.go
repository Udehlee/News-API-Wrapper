package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Udehlee/News-API-Wrapper/cache"
	model "github.com/Udehlee/News-API-Wrapper/models"
)

func NewsInfo(query string) (model.NewsData, error) {
	apiKey := os.Getenv("YOUR_API_KEY")
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", query, apiKey)

	var newsData model.NewsData

	resp, err := http.Get(url)
	if err != nil {
		return newsData, fmt.Errorf("unable to fetch news info: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return newsData, fmt.Errorf("unable to read news body: %w", err)
	}

	if err := json.Unmarshal(body, &newsData); err != nil {
		return newsData, fmt.Errorf("unable to unmarshal json body: %w", err)
	}
	
	return newsData, nil
}

// NewsHandler retrieves cachedNews if found
// return news from Newsapi if no news is found
func NewsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	key := query

	cachedNews, err := cache.GetCacheNews(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Data found in cache, return it
	if cachedNews != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cachedNews)
		return
	}

	// Data not found in cache, fetch from API
	newsData, err := NewsInfo(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store the fetched data in the cache
	if err := cache.SetCache(key, newsData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newsData)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	cache.RedisConn()
	mux := http.NewServeMux()

	mux.HandleFunc("/", Index)
	mux.HandleFunc("GET /api/news", NewsHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("error connecting to port:", err)
	}
}
