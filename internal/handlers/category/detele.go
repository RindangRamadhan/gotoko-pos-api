package category

import (
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/category/deletecategory"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteCategoryHandler godoc
// @Summary   Delete category
// @Tags      Category
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     id   path      int  true  "Category Id"
// @Success   200  {object}  response.BodySuccess{}
// @Failure   500  {object}  response.BodyFailure{}
// @Router    /categories/{id} [delete]
func DeleteCategoryHandler(inport deletecategory.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req deletecategory.InportRequest
		req.CategoryId, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing get category detail interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", nil)
	}
}
