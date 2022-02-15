package main

import (
	"encoding/json"
	"github.com/ayeniblessing101/playlist-service/spotify"
	"github.com/ayeniblessing101/playlist-service/temperature"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

type TempRes struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading env variable", err)
	}
	router := gin.Default()
	router.GET("/:location", getPlaylist)
	router.Run(":8080")

}

func getPlaylist(c *gin.Context) {
	param := c.Param("location")
	if param == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Param cannot be empty"})
	}

	res, err := temperature.GetTemperature(param)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't get the temperature"})
	}

	var tempData TempRes
	err = json.Unmarshal(res, &tempData)
	if err != nil {
		log.Fatalln(err)
	}

	trackData, err := spotify.GetTracks(tempData.Main.Temp)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't get the tracks"})
	}
	c.JSON(http.StatusOK, trackData)
}
