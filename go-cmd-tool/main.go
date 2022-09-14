package main 

import (
	"fmt"
	"flag"
	"os"
)

func main() {

	//'videos get' command
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	//input for 'videos get' command
	getAll := getCmd.Bool("all", false, "Get all videos")
	getId := getCmd.String("id", "", "youtube video Id")

	//'videos add' command
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	//input for 'videos add' command
	addId := addCmd.String("id", "", "Youtube video Id")
	addTitle := addCmd.String("title","", "Youtube video Title")
	addUrl := addCmd.String("url", "", "Youtube video URL")
	addImageUrl := addCmd.String("imageurl", "", "Youtube video ImageURL")
	addDesc := addCmd.String("desc", "", "Youtube video description")



	if len(os.Args) < 2 {
		fmt.Println("expected get or add subcommands")
		os.Exit(1)
	}

	switch os.Args[1]{
		case "get":
			HandleGet(getCmd, getAll, getId)
		case "add":
			HandleAdd(addCmd, addId, addTitle, addUrl, addImageUrl, addDesc)
		default:
	}


}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("id is required or specify --all to fetch all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description\n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v\n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
	}

	if *id != "" {
		videos := getVideos()
		for _, video := range videos {
			if video.Id == *id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description\n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v\n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
			
		}
	}

	return
}

func ValiadateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string){
	if *id == "" || *title == "" || *url == "" || *imageurl == "" || *desc == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string){

	addCmd.Parse(os.Args[2:])

	ValiadateVideo(addCmd, id, title, url, imageurl, desc)

	video:= video{
		Id: *id,
		Title: *title,
		Url: *url,
		Imageurl: *imageurl,
		Description: *desc,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}