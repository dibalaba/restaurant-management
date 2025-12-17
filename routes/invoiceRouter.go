package routes

import ()

func Invoiceroutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoices", controller.GetFoods())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetFood())
	incomingRoutes.POST("/invoices", controller.Createinvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateFood())

}
