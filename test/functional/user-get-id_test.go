/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package functional

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/sundowndev/api-go-example/test/fixture"
)

// TestGetId
func TestGetId(t *testing.T) {
    var body map[string]interface{}
    var item map[string]interface{}
    var err error
    var user = fixture.DefaultUser()

    if body, err = RequestApiJson("GET", BaseUrl + "/users/1", nil, true); err != nil {
        t.Fatal(err)
    }

    item = body["item"].(map[string]interface{})
    assert.Equal(t, float64(1), item["ID"])
    assert.Equal(t, user["name"], item["name"])
    assert.Equal(t, user["email"], item["email"])
    assert.Nil(t, item["password"])
    assert.NotNil(t, item["CreatedAt"])
    assert.NotNil(t, item["UpdatedAt"])
}
