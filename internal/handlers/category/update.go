package category

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/category/updatecategory"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateCategoryHandler godoc
// @Summary   Update existing category
// @Tags      Category
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     id       path      int                           true  "Category Id"
// @Param     payload  body      updatecategory.InportRequest  true  "Payload"
// @Success   200      {object}  response.BodySuccess{}
// @Failure   500      {object}  response.BodyFailure{}
// @Router    /categories/{id} [put]
func UpdateCategoryHandler(inport updatecategory.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req updatecategory.InportRequest

		// ? Binding Request
		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		req.CategoryId, _ = strconv.ParseInt(c.Param("id"), 10, 64)

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing update category interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", nil)
	}
}
