package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	c, err := pubsub.NewClient(ctx, "laqiiz-test01")
	if err != nil {
		log.Fatal(err)
	}

	res := c.Topic("avrotopic2").Publish(ctx, &pubsub.Message{
		Data: []byte("dummy payload"),
	})
	if _, err := res.Get(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("success publish")

}
