package cashier

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/cashier/getcashierdetail"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCashierDetailHandler godoc
// @Summary   Get cashier detail
// @Tags      Cashier
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     id   path      int  true  "Cashier Id"
// @Success   200  {object}  response.BodySuccess{data=getcashierdetail.InportResponse}
// @Failure   500  {object}  response.BodyFailure{}
// @Router    /cashiers/{id} [get]
func GetCashierDetailHandler(inport getcashierdetail.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getcashierdetail.InportRequest
		req.CashierId, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		// ? Binding Request
		if err := request.BindParam(c, &req); err != nil {
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
			response.WriteError(c, "failed when executing get cashier detail interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", resp)
	}
}
