package openWeatherMap

import (
	"bytes"
	"io"
	"testing"
)

func getMockData() io.Reader {
	response := `{"coord": { "lon": 139,"lat": 35},
  "weather": [
    {
      "id": 800,
      "main": "Clear",
      "description": "clear sky",
      "icon": "01n"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 289.92,
    "pressure": 1009,
    "humidity": 92,
    "temp_min": 288.71,
    "temp_max": 290.93
  },
  "wind": {
    "speed": 0.47,
    "deg": 107.538
  },
  "clouds": {
    "all": 2
  },
  "dt": 1560350192,
  "sys": {
    "type": 3,
    "id": 2019346,
    "message": 0.0065,
    "country": "JP",
    "sunrise": 1560281377,
    "sunset": 1560333478
  },
  "timezone": 32400,
  "id": 1851632,
  "name": "Shuzenji",
  "cod": 200
}`

	r := bytes.NewReader([]byte(response))
	return r
}

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIKey: ""}
	weather, err := openWeatherMap.responseParse(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 1851632 {
		t.Errorf("id is 1851632, not %d\n", weather.ID)
	}
}
