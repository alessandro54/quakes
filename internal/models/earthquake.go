package models

type Earthquake struct {
	Magnitude float32 `firestore:"magnitude" json:"magnitud"`
	Intensity string  `firestore:"intensity" json:"intensidad"`
	Latitude  float32 `firestore:"latitude" json:"latitud"`
	Longitude float32 `firebase:"longitude" json:"longitud"`
	Depth     float32 `firebase:"depth" json:"profundidad"`
	Reference string  `firebase:"reference" json:"referencia"`
	LocalDate string  `firebase:"localDate" json:"fecha_local"`
	LocalTime string  `firebase:"localTime" json:"hora_local"`
}
