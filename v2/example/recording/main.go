package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/example"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
	client := resemble.NewClient(example.LoadConfigByKey("TEST_API_KEY"))

	// create voice
	voice, err := client.Voice.Create(request.Payload{"name": "voice for test recording"})
	if err != nil {
		panic(err)
	}
	voiceUUId := voice.Item.UUID
	fmt.Println(voiceUUId)

	recordingFile, err := filepath.Abs("./spec_sample_audio.wav")
	if err != nil {
		panic(err)
	}

	// create recording
	time.Sleep(time.Second)
	r, err := client.Recording.Create(voiceUUId, recordingFile, request.Payload{
		"name":      "test recording",
		"text":      "transcription",
		"is_active": true,
		"emotion":   "neutral",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", r)

	// update recording
	time.Sleep(time.Second)
	r, err = client.Recording.Update(voiceUUId, r.Item.UUID, request.Payload{
		"name":      "test update recording",
		"text":      "new transcription",
		"is_active": true,
		"emotion":   "neutral",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", r)

	// get all recordings
	time.Sleep(time.Second)
	recordings, err := client.Recording.All(voiceUUId, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", recordings)

	// get recording by uuid
	time.Sleep(time.Second)
	r, err = client.Recording.Get(voiceUUId, r.Item.UUID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", r)

	// delete recording
	time.Sleep(time.Second)
	m, err := client.Recording.Delete(voiceUUId, r.Item.UUID)
	if err != nil {
		panic(err)
	}
	fmt.Println(m.Success)
}
