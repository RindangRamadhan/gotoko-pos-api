package cashier

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/cashier/getcashier"

	"github.com/gin-gonic/gin"
)

// GetCashierHandler godoc
// @Summary   Get list of cashiers
// @Tags      Cashier
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     skip   query     int  true  "Page Number"
// @Param     limit  query     int  true  "Limit Display"
// @Success   200    {object}  response.BodySuccess{data=getcashier.InportResponse}
// @Failure   500    {object}  response.BodyFailure{}
// @Router    /cashiers [get]
func GetCashierHandler(inport getcashier.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getcashier.InportRequest

		// ? Binding Request
		if err := request.BindParam(c, &req.CashierFilter); err != nil {
			response.WriteError(c, "failed binding request", err)
			return
		}

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing get cashier interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", resp)
	}
}
