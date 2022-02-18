package main

import (
	"fmt"
	"time"

	"github.com/ashadi-cc/resemble/v2"
	"github.com/ashadi-cc/resemble/v2/example"
	"github.com/ashadi-cc/resemble/v2/request"
)

func main() {
	client := resemble.NewClient(example.LoadConfigByKey("api_key"))

	// create project
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

	uuid := project.Item.UUID

	time.Sleep(time.Second)
	// get project
	project, err = client.Project.Get(uuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(project.Item.Name)

	time.Sleep(time.Second)
	// update project
	project, err = client.Project.Update(uuid, request.Payload{
		"name":             "Project 1",
		"description":      "project update description",
		"is_public":        false,
		"is_collaborative": false,
		"is_archived":      false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(project.Item.Description)

	time.Sleep(time.Second)
	// get all projects
	projects, err := client.Project.All(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", projects.Items)

	time.Sleep(time.Second)
	// delete project
	message, err := client.Project.Delete(uuid)
	if err != nil {
		panic(err)
	}
	fmt.Println(message.Success)
}
