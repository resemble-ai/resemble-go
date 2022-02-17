package main

import (
	"fmt"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/example"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
	client := resemble.NewClient(example.LoadConfigByKey("api_key"))

	voiceUUID := example.LoadConfigByKey("voice_uuid")
	projectUUID := example.LoadConfigByKey("project_uuid")
	body := "This is a streaming test."
	streamUrl := example.LoadConfigByKey("stream_url")

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
