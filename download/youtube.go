package download

import (
	"fmt"
	"os"

	"github.com/rylio/ytdl"
)

func GetVideo(link string) {
	vid, err := ytdl.GetVideoInfo(link)
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	file, _ := os.Create(vid.Title + ".mp4")
	defer file.Close()
	_ = vid.Download(vid.Formats[0], file)
}
