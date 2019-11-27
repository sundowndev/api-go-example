/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package functional

import (
	"log"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sundowndev/api-go-example/src/services"
	"github.com/sundowndev/api-go-example/test/fixture"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const RegexToken = `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`

var BaseUrl string
var Tokens Token

func getToken() {
	var body map[string]interface{}
	var err error

	var user = fixture.DefaultUser()
	var data = url.Values{"email": {user["email"]}, "password": {user["password"]}}

	if body, err = RequestApiJson("POST", BaseUrl+"/auth", data, false); err != nil {
		log.Fatal(err)
	}

	Tokens.RefreshToken = body["refresh_token"].(string)
	Tokens.AccessToken = body["access_token"].(string)
}

func setupDB() {
	var user = fixture.DefaultUser()

	CreateUser(user)
}

func removeDB() {
	services.Db.Exec("DROP TABLE users;")
	services.Db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME") + ";")
}

// init
func init() {
	// Initialize environment variable contained in .env
	if err := godotenv.Load("../../.env.testing"); err != nil {
		log.Fatal("No .env.testing file found")
	}

	BaseUrl = os.Getenv("HOST") + ":" + os.Getenv("PORT")
}

// TestUserPost
func TestMain(m *testing.M) {
	var ret int

	// Set gin mode to test
	gin.SetMode(gin.TestMode)

	// Initialize Database
	services.Database()

	// Initialize Jwt
	services.Jwt()

	// Setup DB
	setupDB()

	// Get token for test request
	getToken()

	ret = m.Run()

	// Remove DB
	removeDB()

	os.Exit(ret)
}
