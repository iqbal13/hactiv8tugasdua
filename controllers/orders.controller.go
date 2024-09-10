package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iqbal13/hactiv8tugasdua/config"
	"github.com/iqbal13/hactiv8tugasdua/models"
)

func GetOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetOrderByID(c *gin.Context) {
	var order models.Order
	if err := config.DB.Preload("Items").First(&order, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order Tidak Ditemukan!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrder(c *gin.Context) {
	var payload models.Order
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.CreatedAt = time.Now()
	config.DB.Create(&payload)
	c.JSON(http.StatusOK, gin.H{"data": payload, "messages": "Order Berhasil Disimpan", "success": true})
}

func UpdateOrder(c *gin.Context) {
	var order models.Order

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("order_id = ?", c.Param("id")).Delete(&models.Items{})

	for _, item := range order.Items {
		item.OrderId = id
		config.DB.Create(&item)
	}
	log.Println(order.Items)
	var existingOrder models.Order

	if err := config.DB.Preload("Items").First(&existingOrder, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	existingOrder.CustomerName = order.CustomerName
	existingOrder.OrderedAt = order.OrderedAt

	config.DB.Save(&existingOrder)

	c.JSON(http.StatusOK, gin.H{
		"data":     existingOrder,
		"messages": "Update data success",
		"success":  true,
	})
}

func DeleteOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("id")
	// if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Order Tidak Ditemukan"})
	// 	return
	// }

	if err := config.DB.Where("order_id = ?", id).Delete(&models.Items{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Delete Item"})
		return
	}

	if err := config.DB.Where("order_id = ?", id).Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Delete Order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messages": "Order Berhasil Dihapus", "success": true})
}
