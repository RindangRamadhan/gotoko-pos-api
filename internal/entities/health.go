package entities

type (
	HealthChekResponse struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		MySQL   bool   `json:"mysql"`
	}
)
