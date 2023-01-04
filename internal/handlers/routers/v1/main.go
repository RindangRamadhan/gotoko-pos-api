package v1

import (
	"gotoko-pos-api/internal"
	"gotoko-pos-api/internal/handlers/cashier"
	"gotoko-pos-api/internal/handlers/category"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type (
	RoutePath struct {
		Method   string
		Group    string
		Pattern  string
		Handlers []gin.HandlerFunc
	}
)

func Paths(container *internal.Container) []RoutePath {
	return []RoutePath{
		{
			Method: "GET", Group: "", Pattern: "/docs/*any", Handlers: []gin.HandlerFunc{
				ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)),
			},
		},
		{
			Method: "GET", Group: "", Pattern: "/cashiers", Handlers: []gin.HandlerFunc{
				// middleware.CrossServiceAuth(), uom.GetUomHandler(container.GetUomUsecase),
				cashier.GetCashierHandler(container.GetCashierUsecase),
			},
		},
		{
			Method: "GET", Group: "", Pattern: "/cashiers/:id", Handlers: []gin.HandlerFunc{
				cashier.GetCashierDetailHandler(container.GetCashierDetailUsecase),
			},
		},
		{
			Method: "POST", Group: "", Pattern: "/cashiers", Handlers: []gin.HandlerFunc{
				cashier.CreateCashierHandler(container.CreateCashierUsecase),
			},
		},
		{
			Method: "PUT", Group: "", Pattern: "/cashiers/:id", Handlers: []gin.HandlerFunc{
				cashier.UpdateCashierHandler(container.UpdateCashierUsecase),
			},
		},
		{
			Method: "DELETE", Group: "", Pattern: "/cashiers/:id", Handlers: []gin.HandlerFunc{
				cashier.DeleteCashierHandler(container.DeleteCashierUsecase),
			},
		},
		{
			Method: "GET", Group: "", Pattern: "/categories", Handlers: []gin.HandlerFunc{
				category.GetCategoryHandler(container.GetCategoryUsecase),
			},
		},
		{
			Method: "GET", Group: "", Pattern: "/categories/:id", Handlers: []gin.HandlerFunc{
				category.GetCategoryDetailHandler(container.GetCategoryDetailUsecase),
			},
		},
		{
			Method: "POST", Group: "", Pattern: "/categories", Handlers: []gin.HandlerFunc{
				category.CreateCategoryHandler(container.CreateCategoryUsecase),
			},
		},
		{
			Method: "PUT", Group: "", Pattern: "/categories/:id", Handlers: []gin.HandlerFunc{
				category.UpdateCategoryHandler(container.UpdateCategoryUsecase),
			},
		},
		{
			Method: "DELETE", Group: "", Pattern: "/categories/:id", Handlers: []gin.HandlerFunc{
				category.DeleteCategoryHandler(container.DeleteCategoryUsecase),
			},
		},
	}

}
