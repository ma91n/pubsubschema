package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/linkedin/goavro/v2"
)

var codec, _ = goavro.NewCodec(`
   {
	 "type" : "record",
	 "name" : "Avro",
	 "fields" : [
	   {"name" : "StringField", "type" : "string"},
	   {"name" : "FloatField", "type" : "float"},
	   {"name" : "BooleanField", "type" : "boolean"}
	 ]
	}
`)

func main() {
	ctx := context.Background()
	c, err := pubsub.NewClient(ctx, "laqiiz-test01")
	if err != nil {
		log.Fatal(err)
	}
	topic := c.Topic("avrotopic2") // 先程作成したトピック

	data := []map[string]interface{}{
		{"StringField": "hello", "FloatField": 123.45, "BooleanField": true},
		{"NGField": "dummy"},
	}

	for _, v := range data {
		binary, err := codec.BinaryFromNative(nil, v)
		if err != nil {
			log.Fatal("codec.BinaryFromNative", err)
		}

		res := topic.Publish(ctx, &pubsub.Message{
			Data: binary,
		})
		if _, err := res.Get(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("success publish", v)
	}

}
