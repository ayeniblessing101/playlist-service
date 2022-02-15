package temperature

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// GetTemperature retrieves the temperature based on the location or longitude/latitude
func GetTemperature(param string) ([]byte, error) {
	fmt.Println(os.Getenv("OPEN_WEATHER_API_ID"))
	result := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s", param, os.Getenv("OPEN_WEATHER_API_ID"), "metric")
	resp, err := http.Get(result)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
