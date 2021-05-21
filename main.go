package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	event := cloudevents.NewEvent()
	event.SetID("example-uuid-32943bac6fea")
	event.SetSource("example/uri")
	event.SetDataContentEncoding("application/json") // The data encoding type.
	event.SetSubject("hello")
	event.SetTime(time.Now()) // The time when the event happened.
	event.SetType("example.type")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(event); err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		panic(err)
	}
	log.Println("JSON")
	log.Println(string(b))

	e := cloudevents.NewEvent()
	if err := json.Unmarshal(b, &e); err != nil {
		panic(err)
	}
	log.Println("EVENT")
	log.Println(e)
}
