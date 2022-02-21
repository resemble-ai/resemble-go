# resemble.ai API

[resemble.ai](https://resemble.ai) is a state-of-the-art natural voice cloning and synthesis provider. Best of all, the platform is accessible by using our public API! Sign up [here](https://app.resemble.ai) to get an API token!

This repository hosts a Golong library for convenient usage of the [Resemble API](https://docs.resemble.ai).

# Quick start 

```golang 
package main

import (
	"fmt"
	"time"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
    client := resemble.NewClient("your_api_key")

    projects, err := client.Project.All(1)
	if err != nil {
		panic(err)
	}
    fmt.Println(projects.items)

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
}

```

## Example 
All example usage can be found in v2/example directory. before run it, you need set the environment variables: 

```bash
export TEST_API_KEY=your_api_key
export TEST_VOICE_UUID=your_voice_uuid
export TEST_PROJECT_UUID=your_project_uuid
export TEST_STREAM_URL="https://your-stream-url"
export TEST_CALLBACK_URL="https://webhook.site/"
```