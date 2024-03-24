package main

func main() {
	r.gin.Default

	userRoute := r.Group("/user")
	{
		userRoute.GET("/:id", userController.GetUserByID)
	}

	r.Run(":8080")

}
