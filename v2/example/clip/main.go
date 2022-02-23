package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/example"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
	client := resemble.NewClient(example.LoadConfigByKey("TEST_API_KEY"))
	client.SetSyncServerUrl(example.LoadConfigByKey("TEST_STREAM_URL"))

	// get values from environment variable
	voiceUUID := example.LoadConfigByKey("TEST_VOICE_UUID")
	projectUUID := example.LoadConfigByKey("TEST_PROJECT_UUID")
	// callbackUrl := example.LoadConfigByKey("TEST_CALLBACK_URL")

	// // create sync
	// clip, err := client.Clip.CreateSync(projectUUID, request.Payload{
	// 	"body":       "this is test.",
	// 	"voice_uuid": voiceUUID,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(clip.Success)

	// // craete async clip
	// time.Sleep(time.Second)
	// clip2, err := client.Clip.CreateAsync(projectUUID, callbackUrl, request.Payload{
	// 	"voice_uuid": voiceUUID,
	// 	"body":       "test async.",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v \n", clip2)

	// // get clip by uuid
	// time.Sleep(time.Second)
	// clip, err = client.Clip.Get(projectUUID, clip.Item.UUID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(clip.Success)

	// // get all clips
	// time.Sleep(time.Second)
	// clips, err := client.Clip.All(projectUUID, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v \n", clips.Items)

	// // update clip
	// time.Sleep(time.Second)
	// clip, err = client.Clip.UpdateAsync(projectUUID, clip.Item.UUID, callbackUrl, request.Payload{"body": "test update.", "voice_uuid": voiceUUID})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(clip.Success)

	// // delete clip
	// time.Sleep(time.Second)
	// m, err := client.Clip.Delete(projectUUID, clip.Item.UUID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(m.Success)

	// stream clip
	time.Sleep(time.Second)
	body := "This is a streaming test."
	cMeta, cChunk, cDone, cErr := client.Clip.Stream(request.Payload{
		"voice_uuid":   voiceUUID,
		"project_uuid": projectUUID,
		"data":         body,
	})

	for {
		select {
		// receive error. print error then exit
		case err := <-cErr:
			log.Fatal(err)
		// receive metadata
		case meta := <-cMeta:
			fmt.Println(meta)
		// receive chunk
		case chunk := <-cChunk:
			_ = chunk
			fmt.Println("chunk")
		// receive done signal. exit
		case <-cDone:
			return
		}
	}
}
