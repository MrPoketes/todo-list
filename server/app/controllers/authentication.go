package controllers

import (
	"context"
	"encoding/base64"
	"net/http"

	"strings"

	"github.com/MrPoketes/todo-list/server/app"
	"github.com/MrPoketes/todo-list/server/app/models"
	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	*revel.Controller
}

// Decodes request data to username and password
func getCredentials(data string) (username, password string, err error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", "", err
	}
	strData := strings.Split(string(decodedData), ":")
	username = strData[0]
	password = strData[1]
	return
}

func (c Authentication) Register() revel.Result {
	response := models.Response{Status: http.StatusUnauthorized, Data: nil, Message: "User could not be authorized"}
	if auth := c.Request.Header.Get("Authorization"); auth != "" {
		// Get user data
		username, password, err := getCredentials(strings.Split(auth, " ")[1])
		// Validate user data
		if err := validation(username, password, err, c); err != nil {
			response.Message = "Username and/or password are missing"
			return c.RenderJSON(response)
		}

		// Find the user record
		_, collErr := getCollectionResult(username)
		// If the user is not found, then it's a new user
		if collErr == mongo.ErrNoDocuments {
			// Encrypt password, generate a new user and insert it into db
			encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			newUser := models.User{Username: username, Password: string(encryptedPassword), Items: make([]models.Items, 0)}
			_, insertErr := app.Collection.InsertOne(context.TODO(), newUser)
			if insertErr != nil {
				panic(err)
			}
			response.Status = http.StatusOK
			response.Message = "Successfully registered"
			return c.RenderJSON(newUser)
		}
		response.Status = http.StatusConflict
		response.Message = "User already exists"
	}
	return c.RenderJSON(response)
}

func (c Authentication) Login() revel.Result {
	response := models.Response{Status: http.StatusUnauthorized, Data: nil, Message: "User could not be authorized"}
	if auth := c.Request.Header.Get("Authorization"); auth != "" {
		// Get user data
		username, password, err := getCredentials(strings.Split(auth, " ")[1])
		// Validate user data
		if err := validation(username, password, err, c); err != nil {
			response.Message = "Username and/or password are missing"
			return c.RenderJSON(response)
		}

		// Find the user record
		result, collErr := getCollectionResult(username)
		// Check if the user exists and if the passwords match
		if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil || collErr == mongo.ErrNoDocuments {
			response.Message = "Incorrect username or password"
			return c.RenderJSON(response)
		}
		response.Status = http.StatusOK
		response.Message = "Successfully loged in"
		response.Data = result
		return c.RenderJSON(response)
	}

	response.Status = http.StatusConflict
	response.Message = "User already exists"
	return c.RenderJSON(response)
}

// Validation function to check for required values
func validation(username, password string, err error, c Authentication) error {
	if err != nil {
		return err
	}
	c.Validation.Required(username)
	c.Validation.Required(password)

	if c.Validation.HasErrors() {
		return err
	}
	return nil
}

// Finds a record from the collection
func getCollectionResult(username string) (result models.User, collErr error) {
	collErr = app.Collection.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&result)
	return
}
