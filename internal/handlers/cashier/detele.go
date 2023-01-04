package cashier

import (
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/cashier/deletecashier"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteCashierHandler godoc
// @Summary   Delete cashier
// @Tags      Cashier
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     id   path      int  true  "Cashier Id"
// @Success   200  {object}  response.BodySuccess{}
// @Failure   500  {object}  response.BodyFailure{}
// @Router    /cashiers/{id} [delete]
func DeleteCashierHandler(inport deletecashier.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req deletecashier.InportRequest
		req.CashierId, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing get cashier detail interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", nil)
	}
}
