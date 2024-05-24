package request

type Request struct {
	Departure     string
	DepartureDate string
	Destination   string
	ArrivalDate   string
	Page          string
	PerPage       string
	SortDirection string
	Airlines      []string
}
