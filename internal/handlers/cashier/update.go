package cashier

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/cashier/updatecashier"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateCashierHandler godoc
// @Summary   Update existing cashier
// @Tags      Cashier
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     id       path      int                          true  "Cashier Id"
// @Param     payload  body      updatecashier.InportRequest  true  "Payload"
// @Success   200      {object}  response.BodySuccess{}
// @Failure   500      {object}  response.BodyFailure{}
// @Router    /cashiers/{id} [put]
func UpdateCashierHandler(inport updatecashier.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req updatecashier.InportRequest

		// ? Binding Request
		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		req.CashierId, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing update cashier interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", nil)
	}
}
