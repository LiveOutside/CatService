package home

import (
	"cat_service/internal/data/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var mutex sync.RWMutex

type Repo interface {
	FetchCat() (models.CatImage, error)
}

type Service struct {
	apiURL         string
	cachedImageURL string
}

func NewService(url string) *Service {
	s := &Service{apiURL: url}

	if cat, err := FetchCat(s.apiURL); err == nil {
		s.cachedImageURL = cat.URL
	} else {
		log.Printf("Initial cat fetch failed: %v", err)
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			cat, err := FetchCat(s.apiURL)
			if err != nil {
				log.Printf("Failed to refresh cat image: %v", err)
				continue
			}
			mutex.Lock()
			s.cachedImageURL = cat.URL
			mutex.Unlock()
		}
	}()

	return s
}

func FetchCat(apiURL string) (models.CatImage, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return models.CatImage{}, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.CatImage{}, fmt.Errorf("read body failed %w", err)
	}

	var cats []models.CatImage
	if err := json.Unmarshal(body, &cats); err != nil {
		return models.CatImage{}, fmt.Errorf("unmarshal failed: %w", err)
	}

	if len(cats) == 0 {
		return models.CatImage{}, fmt.Errorf("no cats returned")
	}
	return cats[0], nil
}
