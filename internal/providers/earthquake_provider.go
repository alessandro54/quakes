package providers

import (
	"encoding/json"
	"fmt"
	"github.com/alessandro54/quakes/internal/models"
	"io"
	"net/http"
)

const baseUrl = "https://ultimosismo.igp.gob.pe/ultimo-sismo/ajaxb/"

type Response struct {
	Data []models.Earthquake `json:"data"`
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
		earthquakes = append(earthquakes, item)
	}

	return earthquakes
}
