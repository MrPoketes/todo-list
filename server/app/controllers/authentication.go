package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
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
	notAuthorized := models.Response{Status: http.StatusUnauthorized, Data: nil, Message: "User could not be authorized"}
	if auth := c.Request.Header.Get("Authorization"); auth != "" {
		username, password, err := getCredentials(strings.Split(auth, " ")[1])
		if err != nil {
			return c.RenderError(err)
		}
		c.Validation.Required(username)
		c.Validation.Required(password)

		if c.Validation.HasErrors() {
			notAuthorized.Message = "Username and/or password are missing"
			return c.RenderJSON(notAuthorized)
		}
		var result bson.M
		err = app.Collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			newUser := models.User{Username: username, Password: string(encryptedPassword), Items: make([]models.Items, 0)}
			result, err := app.Collection.InsertOne(context.TODO(), newUser)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
			return c.RenderJSON(newUser)
		}
		notAuthorized.Status = http.StatusConflict
		notAuthorized.Message = "User already exists"
	}

	return c.RenderJSON(notAuthorized)
}
