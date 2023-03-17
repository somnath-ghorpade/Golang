package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

type MovieDetails struct {
	Id       string    `json:"id" bson:"id"`
	Movie    Movie     `json:"movie" bson:"movie" validate:"required"`
	Sessions []Session `json:"sessions" bson:"sessions"`
}
type Session struct {
	Id    string `json:"id" bson:"id" binding:"required"`
	Time  string `json:"time" bson:"time" binding:"required"`
	Seats int    `json:"seats" bson:"seats" binding:"required"`
}

type Movie struct {
	Title    string `json:"title" bson:"Title" `
	Rated    string `json:"rating"  bson:"Rated"`
	Genre    string `json:"genre"  bson:"Genre"`
	Director string `json:"director"  bson:"Director"`
	Actors   string `json:"cast"  bson:"Actors" validate:"phone"`
	Plot     string `json:"plot"  bson:"Plot"`
	Poster   string `json:"poster"  bson:"Poster"`
	Released string `json:"releaseDate"  bson:"Released"`
	Runtime  string `json:"running"  bson:"Runtime"`
	ImdbId   string `json:"imdbid"  bson:"imdbID"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("phone", validatePhonenumber)
}

func (movie *Movie) ValidateMovie() error {

	err := validate.Struct(movie)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			return err
		}
		return err
	}
	return nil
}

func validatePhonenumber(v validator.FieldLevel) bool {
	p, err := phonenumbers.Parse(v.Field().String(), "US")
	if err != nil {
		return false
	}
	return phonenumbers.IsPossibleNumber(p)
}
