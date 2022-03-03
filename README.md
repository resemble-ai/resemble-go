# resemble.ai API

[resemble.ai](https://resemble.ai) is a state-of-the-art natural voice cloning and synthesis provider. Best of all, the platform is accessible by using our public API! Sign up [here](https://app.resemble.ai) to get an API token!

This repository hosts a Golong library for convenient usage of the [Resemble API](https://docs.resemble.ai).

# Quick start 

```golang 
package main

import (
	"fmt"
	"time"

	"github.com/resemble-ai/resemble-go/v2"
	"github.com/resemble-ai/resemble-go/v2/request"
)

func main() {
    client := resemble.NewClient("your_api_key")

    projects, err := client.Project.All(1)
	if err != nil {
		panic(err)
	}
    fmt.Println(projects.Items)

    project, err := client.Project.Create(request.Payload{
        "name":             "Project 1",
        "description":      "project description",
        "is_public":        false,
        "is_collaborative": false,
        "is_archived":      false,
    })
    if err != nil {
        panic(err)
    }
    fmt.Println(project.Success)
}

```

## Streaming
The Streaming API is currently in beta and is not available to all users. Please reach out to team@resemble.ai to inquire more.

Streaming example:
```golang
client := resemble.NewClient("your_api_key")
client.SetSyncServerUrl("your_resemble_synthesis_server_url") # Extra configuration required for streaming

cMeta, cChunk, cDone, cErr := client.Clip.Stream(request.Payload{
    "voice_uuid":   "[voiceUUID]",
    "project_uuid": "[projectUUID]",
    "data":         "This is a streaming test.",
})

for {
    select {
    // receive error. print error then exit
    case err := <-cErr:
        log.Fatal(err)
    // receive metadata
    case meta := <-cMeta:
        b, _ := json.Marshal(meta)
        //print as json object
        fmt.Println(string(b))
    // receive chunk
    case chunk := <-cChunk:
        fmt.Println("chunk data", len(chunk))
    // receive done signal. exit
    case <-cDone:
        return
    }
}


```
# Development
The library files are located in `v2/`

# Testing 
```bash
$ cd v2/ 
export TEST_API_KEY=your_api_key
export TEST_VOICE_UUID=your_voice_uuid
export TEST_PROJECT_UUID=your_project_uuid
export TEST_STREAM_URL="https://your-stream-url"
export TEST_CALLBACK_URL="https://webhook.site/"
$ go test -v ./...
```

# Example 
All example usage can be found in v2/example directory. before run it, you need set the environment variables: 

```bash
export TEST_API_KEY=your_api_key
export TEST_VOICE_UUID=your_voice_uuid
export TEST_PROJECT_UUID=your_project_uuid
export TEST_STREAM_URL="https://your-stream-url"
export TEST_CALLBACK_URL="https://webhook.site/"
```