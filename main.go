package main

import (
	"fmt"
    "os"
    "github.com/slack-go/slack"
)

func main(){

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5750792226230-5757513992034-Z5scjBR85FmD9xTijsSc5xDC")
	os.Setenv("CHANNEL_ID", "C05N6FKEBM3")
	api:= slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params:= slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)

	}	
}