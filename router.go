package main

import (
	"github.com/gin-gonic/gin"
)

/*func RegisterRoutes() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/goodbye", GoodbyeHandler)
}*/

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	public := r.Group("/public")
	public.Use(CORSMiddleware())
	{
		public.GET("/hello", HelloHandler)
		public.GET("/goodbye", GoodbyeHandler)
	}

	// Routes requiring reCAPTCHA
	// secure := r.Group("/secure")
	/*secure.Use(RecaptchaMiddleware())
	{
		secure.POST("/submit", SubmitHandler)
	}*/

	return r
}
