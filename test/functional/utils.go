/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package functional

import (
    "encoding/json"
    "log"
    "net/http"
    "net/http/httptest"
    "net/url"
    "strings"

    "github.com/sundowndev/api-go-example/src/models"
    "github.com/sundowndev/api-go-example/src/routers"
    "github.com/sundowndev/api-go-example/src/services"
)

// RequestApiJson
func RequestApiJson(method, url string, data url.Values, auth bool) (map[string]interface{}, error) {
    var result map[string]interface{}
    var req *http.Request
    var res *httptest.ResponseRecorder
    var err error

    routerEngine := routers.RouterInitialize()

    if req, err = http.NewRequest(method, url, strings.NewReader(data.Encode())); err != nil {
        return result, err
    }

    if auth == true {
        req.Header.Add("Authorization", `Bearer ` + Tokens.AccessToken)
    }

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    res = httptest.NewRecorder()
    routerEngine.ServeHTTP(res, req)

    if err = json.Unmarshal(res.Body.Bytes(), &result); err != nil {
        return result, err
    }

    return result, err
}

// CreateUser
func CreateUser(data map[string]string) models.User {
    var user = models.User{
        Name: data["name"],
        Email: data["email"],
        Password: data["password"],
    }

    if err := services.Db.Create(&user).Error; err != nil {
        log.Fatal(err)
    }

    return user
}
