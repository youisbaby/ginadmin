/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-07-13 19:45:35
 */
package router

import (
	"ginadmin/internal/controllers/api/user"

	"github.com/gin-gonic/gin"
)

func ApiRouter(apiRouter *gin.RouterGroup) {
	apiRouter.Use()
	{
		apiUserRouter := apiRouter.Group("user")
		{
			apiUserRouter.GET("/list", user.Uc.UserList)
		}
	}
}
