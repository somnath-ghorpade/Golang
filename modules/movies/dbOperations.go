package movies

import (
	"GinWebServer/common"
	"GinWebServer/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadMovies(c *gin.Context) {

	client, ctx, cancelFunc, err := common.GetMongoConnection()
	if err != nil {
		log.Error("ERROR:Failed to get connection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancelFunc()
	collection := client.Database("Movies").Collection("MovieDetails")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Error("ERROR:Failed to get collection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res []bson.M
	for cur.Next(ctx) {
		var res1 bson.M
		err := cur.Decode(&res1)
		if err != nil {
			log.Error("ERROR:Failed to decode result", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		res = append(res, res1)
	}
	c.IndentedJSON(200, res)
}

func FetchMovieById(c *gin.Context) {

	client, ctx, cancelFunc, err := common.GetMongoConnection()
	if err != nil {
		log.Error("ERROR:Failed to get connection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancelFunc()
	collection := client.Database("Movies").Collection("MovieDetails")
	//fetch movie by movieId "id" : "tt0379225"
	var movieDetails models.MovieDetails
	err = c.ShouldBind(&movieDetails)
	// err = c.ShouldBindJSON(&movieDetails)
	if err != nil {
		log.Error("validation failed", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	filter := bson.M{"id": movieDetails.Id}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error("ERROR:Failed to get collection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res []bson.M
	for cur.Next(ctx) {
		var res1 bson.M
		err := cur.Decode(&res1)
		if err != nil {
			log.Error("ERROR:Failed to decode result", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		res = append(res, res1)
	}
	c.IndentedJSON(200, res)
}

// using Query Params
func FetchMovieByIdQuery(c *gin.Context) {

	client, ctx, cancelFunc, err := common.GetMongoConnection()
	if err != nil {
		log.Error("ERROR:Failed to get connection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancelFunc()
	collection := client.Database("Movies").Collection("MovieDetails")
	//fetch movie by movieId "id" : "tt0379225"

	Id := c.Query("id")
	filter := bson.M{"id": Id}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Error("ERROR:Failed to get collection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res []bson.M
	for cur.Next(ctx) {
		var res1 bson.M
		err := cur.Decode(&res1)
		if err != nil {
			log.Error("ERROR:Failed to decode result", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		res = append(res, res1)
	}
	c.IndentedJSON(200, res)
}

func InsertMovie(c *gin.Context) {
	client, ctx, cancelFunc, err := common.GetMongoConnection()
	if err != nil {
		log.Error("ERROR:Failed to get connection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancelFunc()
	collection := client.Database("Movies").Collection("MovieDetails")

	var movieDetails models.MovieDetails
	err = c.ShouldBind(&movieDetails)
	// err = c.ShouldBindJSON(&movieDetails)
	if err != nil {
		log.Error("validation failed", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var movie models.Movie
	err = movie.ValidateMovie()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = collection.InsertOne(ctx, movieDetails)
	if err != nil {
		log.Error("ERROR:Failed to insert movie", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, "Movie Inserted success")
}

func UpdateMovie(c *gin.Context) {
	client, ctx, cancelFunc, err := common.GetMongoConnection()
	if err != nil {
		log.Error("ERROR:Failed to get connection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancelFunc()
	collection := client.Database("Movies").Collection("MovieDetails")

	var movieDetails models.MovieDetails
	err = c.ShouldBind(&movieDetails)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// update operation
	selctor := bson.M{"id": movieDetails.Id}
	updator := bson.M{"$set": bson.M{"movie": movieDetails.Movie, "sessions": movieDetails.Sessions}}

	result, err := collection.UpdateOne(ctx, selctor, updator)
	if err != nil {
		log.Error("failed to updatw movies", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(200, result)
}

func DeleteMovieByIdQuery(c *gin.Context) {
	client, ctx, cancelFunc, err := common.GetMongoConnection()
	if err != nil {
		log.Error("ERROR:Failed to get connection", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer cancelFunc()
	collection := client.Database("Movies").Collection("MovieDetails")
	Id := c.Query("id")
	filter := bson.M{"id": Id}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("Failed to delete movie", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
