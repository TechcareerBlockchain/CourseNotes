package utils

type Message struct {
	Sender  string `json:"sender" bson:"sender"`
	Message string `json:"message"`
}
