package controllers

import (
	"API/src/Orders/application"
)

type OrderRequest struct {
    UserID int32   `json:"user_id"`
    Amount float32 `json:"amount"`
    Status string  `json:"status"`
}

type OrderController struct {
    getAllOrders    *application.GetAllOrders
    getOrderByID    *application.GetOrderByID
	deleteOrderByID *application.DeleteOrder
	createOrder     *application.CreateOrder
    updateOrder     *application.UpdateOrder
}

func NewOrderController( 
	createOrder *application.CreateOrder, updateOrder *application.UpdateOrder, getOrderByID *application.GetOrderByID,
	getAllOrders *application.GetAllOrders, deleteOrderByID *application.DeleteOrder,
	) *OrderController { return &OrderController{ getOrderByID: getOrderByID, getAllOrders: getAllOrders, updateOrder: updateOrder,
		createOrder: createOrder, deleteOrderByID: deleteOrderByID,}
    }



