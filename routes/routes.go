package routes

import (
	"GinWebServer/middleware"
	"GinWebServer/modules/movies"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) {

	o := g.Group("/o")
	// /o/checkServerStatus
	o.GET("/checkServerStatus", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Status": "OK",
		})
	})
	o.GET("/fetchMovies", movies.ReadMovies)
	o.POST("/fetchMovieById", movies.FetchMovieById)
	o.GET("/fetchMovieByIdQuery", movies.FetchMovieByIdQuery)

	r := g.Group("/r")
	r.Use(middleware.ValidateUser)
	r.GET("/deleteMovieByIdQuery", movies.DeleteMovieByIdQuery)

	// /o/insertMovie
	o.POST("/insertMovie", movies.InsertMovie)
	o.POST("/updateMovie", movies.UpdateMovie)

	c := g.Group("/c")

	// /r/c/deleteMovie
	c.DELETE("/deleteMovie", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Movie": "Deleted",
		})
	})
}
