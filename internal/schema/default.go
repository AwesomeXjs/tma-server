package schema

type Base struct {
	Request string `json:"request" example:"/api/v1/endpoint"`
	Time    string `json:"time" example:"2023-08-01T00:00:00Z"`
	Title   string `json:"title" example:"success"`
}

type NoData struct {
	Base
	Data string `json:"data" example:"status message"`
}
