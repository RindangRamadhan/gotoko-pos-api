package category

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/category/getcategorydetail"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategoryDetailHandler godoc
// @Summary   Get category detail
// @Tags      Category
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     id   path      int  true  "Category Id"
// @Success   200  {object}  response.BodySuccess{data=getcategorydetail.InportResponse}
// @Failure   500  {object}  response.BodyFailure{}
// @Router    /categories/{id} [get]
func GetCategoryDetailHandler(inport getcategorydetail.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getcategorydetail.InportRequest
		req.CategoryId, _ = strconv.ParseInt(c.Param("id"), 10, 64)

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
			response.WriteError(c, "failed when executing get category detail interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", resp)
	}
}
