package handler

import (
	"WBL0/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
