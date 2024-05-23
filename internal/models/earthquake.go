package models

type Earthquake struct {
	Magnitude    float32 `firestore:"magnitude" json:"magnitude"`
	Intensity    string  `firestore:"intensity" json:"intensity"`
	Latitude     float32 `firestore:"latitude" json:"latitude"`
	Longitude    float32 `firestore:"longitude" json:"longitude"`
	Depth        float32 `firestore:"depth" json:"depth"`
	Reference    string  `firestore:"reference" json:"reference"`
	LocalDate    string  `firestore:"local_date" json:"local_date"`
	LocalTime    string  `firestore:"local_time" json:"local_time"`
	ReportNumber int     `firestore:"report_number" json:"report_number"`
}

type ByReportNumber []Earthquake

func (a ByReportNumber) Len() int      { return len(a) }
func (a ByReportNumber) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByReportNumber) Less(i, j int) bool {
	return a[i].ReportNumber > a[j].ReportNumber
}
