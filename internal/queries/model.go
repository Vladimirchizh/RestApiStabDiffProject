package queries

type Query struct {
	ID             string `json:"id" bson:"_id.omitempty"`
	Isbn           string `json:"isbin" bson:"isbn"`
	Prompt         string `json:"prompt" bson:"prompt"`
	NegativePrompt string `json:"negative_prompt" bson:"negative_prompt"`
	Seed           string `json:"seed" bson:"seed"`
	NumberOfImages string `json:"number_of_images" bson:"number_of_images"`
}
type CreateQuery struct {
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
	Seed           string `json:"seed"`
	NumberOfImages string `json:"number_of_images"`
}
