/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package functional

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUserPost
func TestUserPost(t *testing.T) {
	var body map[string]interface{}
	var item map[string]interface{}
	var err error
	var data = url.Values{"name": {"test1"}, "email": {"example@test1.com"}, "password": {"12345678"}}

	if body, err = RequestApiJson("POST", BaseUrl+"/users", data, true); err != nil {
		t.Fatal(err)
	}

	item = body["item"].(map[string]interface{})
	assert.Equal(t, data.Get("name"), item["name"])
	assert.Equal(t, data.Get("email"), item["email"])
	assert.Nil(t, item["password"])
	assert.NotNil(t, item["ID"])
	assert.NotNil(t, item["CreatedAt"])
	assert.NotNil(t, item["UpdatedAt"])
}
