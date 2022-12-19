package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authentication(jwtSecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// authorizationHeader := c.GetHeader("Authorization")

		// if len(authorizationHeader) == 0 {
		// 	err := errors.New("authorization header is not provided")
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, &response.BodyFailure{
		// 		Status:  false,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }

		// fields := strings.Fields(authorizationHeader)
		// if len(fields) < 2 {
		// 	err := errors.New("invalid authorization header format")
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, &response.BodyFailure{
		// 		Status:  false,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }

		// authorizationType := strings.ToLower(fields[0])
		// if authorizationType != "bearer" {
		// 	err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, &response.BodyFailure{
		// 		Status:  false,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }

		// tokenMaker, err := token.NewJWTMaker(jwtSecretKey)

		// accessToken := fields[1]
		// _, err = tokenMaker.VerifyToken(accessToken)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, &response.BodyFailure{
		// 		Status:  false,
		// 		Message: err.Error(),
		// 	})
		// 	return
		// }

		c.Next()
	}
}

func StaticTokenAuth(staticTokenOpt ...string) gin.HandlerFunc {
	staticToken := os.Getenv("STATIC_TOKEN")
	if len(staticTokenOpt) > 0 {
		staticToken = staticTokenOpt[0]
	}
	return func(c *gin.Context) {
		xAppToken := c.Request.Header.Get("X-App-Token")
		if xAppToken != staticToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "unauthorized",
				"data":    nil,
			})
			return
		}
		c.Next()
	}
}
