package quotegen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Quote struct {
	Id           string   `json:"_id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Length       int      `json:"length"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
	TimeStamp    int64    `json:"timestamp"`
}

func fetchQuotes(num int) []Quote {
	url := fmt.Sprintf("http://api.quotable.io/quotes/random?limit=%d", num)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal().Err(err).Msg("Error fetching quotes")
	}
	defer res.Body.Close()
	var quotes []Quote
	err = json.NewDecoder(res.Body).Decode(&quotes)
	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing quotes")
	}
	return quotes
}

func Generator(done <-chan any) <-chan any {
	quoteStream := make(chan any)
	go func() {
		defer close(quoteStream)
		for {
			quotecache := fetchQuotes(10)
			for _, quote := range quotecache {
				quote.TimeStamp = time.Now().UnixMilli()
				select {
				case <-done:
					return
				case quoteStream <- quote:
				}
			}
		}
	}()
	log.Info().Msg("Quote Generator started")
	return quoteStream

}
