package server_model

type TaxBand struct {
	Start         float64 `json:"start"`
	End           float64 `json:"end"`
	PercentageTax float64 `json:"percentageTax"`
}
