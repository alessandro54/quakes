package providers

import (
	"encoding/json"
	"fmt"
	"github.com/alessandro54/quakes/internal/models"
	"io"
	"net/http"
)

const baseUrl = "https://ultimosismo.igp.gob.pe/ultimo-sismo/ajaxb/"

type EarthquakeResponse struct {
	Magnitude    float32 `json:"magnitud"`
	Intensity    string  `json:"intensidad"`
	Latitude     float32 `json:"latitud"`
	Longitude    float32 `json:"longitud"`
	Depth        float32 `json:"profundidad"`
	Reference    string  `json:"referencia"`
	LocalDate    string  `json:"fecha_local"`
	LocalTime    string  `json:"hora_local"`
	ReportNumber int     `json:"numero_reporte"`
}

type Response struct {
	Data []EarthquakeResponse `json:"data"`
}

func ByYear(year int) []models.Earthquake {
	var earthquakes []models.Earthquake
	res, err := http.Get(baseUrl + fmt.Sprint(year))

	if err != nil {
		fmt.Printf("Error creando la solicitud: %v\n", err)
		return nil
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error leyendo la respuesta: %v\n", err)
		return nil
	}

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error decodificando la respuesta: %v\n", err)
		return nil
	}

	for _, item := range response.Data {
		earthquakes = append(earthquakes, convertToEarthquake(item))
	}

	return earthquakes
}

func convertToEarthquake(earthquake EarthquakeResponse) models.Earthquake {
	return models.Earthquake{
		Magnitude:    earthquake.Magnitude,
		Intensity:    earthquake.Intensity,
		Latitude:     earthquake.Latitude,
		Longitude:    earthquake.Longitude,
		Depth:        earthquake.Depth,
		Reference:    earthquake.Reference,
		LocalDate:    earthquake.LocalDate,
		LocalTime:    earthquake.LocalTime,
		ReportNumber: earthquake.ReportNumber,
	}
}
