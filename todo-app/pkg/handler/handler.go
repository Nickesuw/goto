package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h *Handler) InitRoutes() *gin.Engine{
	router:=gin.New()

	auth:=router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)

	}
	api:= router.Group("/api")
	{
		lists:= api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getItemsById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.DeleteList)

			items:=lists.Group(":id/items")
			{
				items.POST("/", h.createList)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemsById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.DeleteItem)
			}
		}
	}
	return router
}
