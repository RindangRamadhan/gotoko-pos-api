package hub

type BaseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type (
	GetHubRequest struct {
		ID int `json:"hub_id"`
	}

	GetHubResponse struct {
		BaseResponse
		Data Hub `json:"data"`
	}
)
