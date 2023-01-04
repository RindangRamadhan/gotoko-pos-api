package category

import (
	"gotoko-pos-api/common/request"
	"gotoko-pos-api/common/response"
	"gotoko-pos-api/common/validator"
	"gotoko-pos-api/internal/usecases/category/createcategory"

	"github.com/gin-gonic/gin"
)

// CreateCategoryHandler godoc
// @Summary   Create a new category
// @Tags      Category
// @Security  JWTBearer
// @Accept    json
// @Produce   json
// @Param     payload  body      createcategory.InportRequest  true  "Payload"
// @Success   200      {object}  response.BodySuccess{data=createcategory.InportResponse}
// @Failure   500      {object}  response.BodyFailure{}
// @Router    /categories [post]
func CreateCategoryHandler(inport createcategory.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createcategory.InportRequest

		// ? Binding Request
		if err := request.UnmarshalJSON(c, &req); err != nil {
			return
		}

		// ? Validate Request
		if err := validator.Validate(&req); err != nil {
			response.WriteError(c, "unable to validate payload", err)
			return
		}

		resp, err := inport.Execute(c.Copy().Request.Context(), req)
		if err != nil {
			response.WriteError(c, "failed when executing create category interactor", err)
			return
		}

		response.WriteSuccess(c, "Success", resp)
	}
}
