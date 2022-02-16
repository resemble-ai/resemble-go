package main

import (
	"fmt"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
	client := resemble.NewClient("<your_api_token>")

	voiceUUID := "<your_voice_id>"
	projectUUID := "<your_project_uuid>"
	body := "Data is ready."
	streamUrl := "<your_sync_server_url>"

	fmt.Println("Example streaming clip")
	clipStream, err := client.Clip.Stream(streamUrl, request.Payload{
		"voice_uuid":   voiceUUID,
		"project_uuid": projectUUID,
		"data":         body,
	})
	if err != nil {
		panic(err)
	}

	for clip := range clipStream {
		if clip.Err != nil {
			panic(err)
		}
		fmt.Println("chunk data", clip.Chunk)
	}
}
