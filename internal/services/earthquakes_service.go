package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://ultimosismo.igp.gob.pe/ultimo-sismo/ajaxb/"

type Earthquake struct {
	Magnitude float32 `json:"magnitud"`
	Intensity string  `json:"intensidad"`
	Latitude  float32 `json:"latitud"`
	Longitude float32 `json:"longitud"`
	Depth     float32 `json:"profundidad"`
	Reference string  `json:"referencia"`
	LocalDate string  `json:"fecha_local"`
	LocalTime string  `json:"hora_local"`
}

type Response struct {
	Data []Earthquake `json:"data"`
}

func ByYear(year int) string {
	res, err := http.Get(baseUrl + fmt.Sprint(year))

	if err != nil {
		fmt.Printf("Error creando la solicitud: %v\n", err)
		return ""
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error leyendo la respuesta: %v\n", err)
		return ""
	}

	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error decodificando la respuesta: %v\n", err)
		return ""
	}

	for _, item := range response.Data {
		fmt.Printf("Magnitud: %f\n", item.Magnitude)
	}

	return string(body)
}
