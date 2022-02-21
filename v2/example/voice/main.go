package main

import (
	"fmt"
	"time"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/example"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
	client := resemble.NewClient(example.LoadConfigByKey("TEST_API_KEY"))

	// create voice
	voice, err := client.Voice.Create(request.Payload{"name": "test voide"})
	if err != nil {
		panic(err)
	}
	fmt.Println(voice.Item.UUID)
	uuid := voice.Item.UUID

	time.Sleep(time.Second)
	// get voice by uuid
	voice, err = client.Voice.Get(uuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(voice.Item.Name)

	time.Sleep(time.Second)
	// get all voices
	voices, err := client.Voice.All(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", voices.Items)

	time.Sleep(time.Second)
	// update voice
	voice, err = client.Voice.Update(uuid, request.Payload{"name": "update voice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(voice.Item.Name)

	time.Sleep(time.Second)
	// delete voice
	message, err := client.Voice.Delete(uuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(message.Success)
}
