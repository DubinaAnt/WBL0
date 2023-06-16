package handler

import (
	"WBL0/internal/model"
	"WBL0/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("internal/web/templates/*")

	api := router.Group("/api/v1")
	{
		order := api.Group("/order")
		{
			order.GET("/search", h.searchOrder)
			order.GET("/", h.getOrderByUID)
		}
	}
	return router
}

func (h *Handler) getOrderByUID(c *gin.Context) {
	var order model.Order

	uid := c.Query("uid")
	if uid == "" {
		c.HTML(http.StatusBadRequest, "errorForm.html", gin.H{
			"UID": uid,
		})
		return
	}

	order, err := h.service.GetOrder(uid)
	if err != nil {
		c.HTML(http.StatusBadRequest, "errorForm.html", gin.H{
			"UID": uid,
		})
		return
	}

	c.HTML(http.StatusOK, "orderForm.html", gin.H{
		"UID":         order.UID,
		"TrackNumber": order.TrackNumber,
		"Address":     order.Delivery.Address,
		"Amount":      order.Payment.Amount,
	})
}

func (h *Handler) searchOrder(c *gin.Context) {
	c.HTML(http.StatusOK, "searchForm.html", gin.H{})
}
