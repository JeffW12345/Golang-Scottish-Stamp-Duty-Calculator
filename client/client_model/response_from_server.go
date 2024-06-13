package client_model

type ResponseFromServer struct {
	TaxDue float64 `json:"taxDue"`
	Message string `json:"message"`
}