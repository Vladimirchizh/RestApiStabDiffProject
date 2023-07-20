package queries

import (
	"math/rand"
	"time"
)

type Query struct {
	ID             int       `json:"id"`
	Prompt         string    `json:"prompt"`
	NegativePrompt string    `json:"negative_prompt"`
	Seed           string    `json:"seed"`
	NumberOfImages string    `json:"number_of_images"`
	Timestamp      time.Time `json:"timestamp"`
}

func NewQuery(prompt string, negative string, seed string, nOfImages string) *Query {
	return &Query{
		ID:             rand.Intn(1000000),
		Prompt:         prompt,
		NegativePrompt: negative,
		Seed:           seed,
		NumberOfImages: nOfImages,
		Timestamp:      time.Now().UTC(),
	}
}
