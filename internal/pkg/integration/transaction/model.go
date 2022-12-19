package transaction

type BaseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type (
	GetOrderRequest struct {
		ID []int `json:"order_id"`
	}

	GetOrderResponse struct {
		BaseResponse
		Data []Order `json:"data"`
	}

	GetOrderDetailRequest struct {
		ID int `json:"order_id"`
	}

	GetOrderDetailResponse struct {
		BaseResponse
		Data OrderDetail `json:"data"`
	}

	Customer struct {
		Name       string `json:"name"`
		Address    string `json:"address"`
		Phone      string `json:"phone"`
		Latitude   string `json:"latitude"`
		Longtitude string `json:"longtitude"`
	}

	Product struct {
		Name      string  `json:"name"`
		Qty       int     `json:"qty"`
		UnitPrice float64 `json:"unit_price"`
	}
)
