package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	c, err := pubsub.NewClient(ctx, "<YOUR PROJECT ID>")
	if err != nil {
		log.Fatal(err)
	}
	topic := c.Topic("avrotopic")

	data := []string{
		`{"StringField":"hello", "FloatField":123.45, "BooleanField":true}`,
		`{"StringField":"world", "FloatField":0, "BooleanField":false}`,
		`{"NGField":"dummy"}`,
	}

	for _, v := range data {
		res := topic.Publish(ctx, &pubsub.Message{
			Data: []byte(v),
		})
		if _, err := res.Get(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("success publish", v)
	}

}
