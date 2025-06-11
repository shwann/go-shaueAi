package main

import (
	"context"
	"fmt"
	"github.com/shwann/go-shaueAi/shaueai"
	"os"
)

func main() {
	client := shaueai.NewClient(os.Getenv("SHAUEAI_API_KEY"), os.Getenv("SHAUEAI_BASE_URL"))

	respUrl, err := client.CreateImage(
		context.Background(),
		shaueai.ImageRequest{
			Prompt:         "Parrot on a skateboard performs a trick, cartoon style, natural light, high detail",
			Size:           shaueai.CreateImageSize256x256,
			ResponseFormat: shaueai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)
}
