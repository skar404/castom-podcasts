package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DownloadView struct {
	UUID string `json:"uuid" form:"uuid" query:"uuid"`
}
type GetAudioView struct {
	Link string `json:"link" form:"link" query:"link"`
}

func Download(c echo.Context) error {
	d := DownloadView{
		UUID: "uuid",
	}
	return c.JSON(http.StatusOK, d)
}

func GetAudio(c echo.Context) error {
	d := GetAudioView{
		Link: "http://test/uuid.mp3",
	}
	return c.JSON(http.StatusOK, d)
}

func main() {
	e := echo.New()

	v1 := e.Group("/api/v1/")
	{
		v1.POST("download", Download)
		v1.GET("audio/:uuid", GetAudio)
	}

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
	//download.GetVideo("https://www.youtube.com/watch?v=DxdJ_-T9ZVw")

}
