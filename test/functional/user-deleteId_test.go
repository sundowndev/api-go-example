/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package functional

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDeleteId
func TestDeleteId(t *testing.T) {
	var body map[string]interface{}
	var item map[string]interface{}
	var err error

	var user = CreateUser(map[string]string{
		"name":     "DeletedUser",
		"email":    "DeletedUser@example.com",
		"password": "12345678",
	})

	if body, err = RequestApiJson("DELETE", fmt.Sprintf("%s/users/%d", BaseUrl, user.ID), nil, true); err != nil {
		t.Fatal(err)
	}

	item = body["item"].(map[string]interface{})
	assert.Equal(t, float64(user.ID), item["ID"])
	assert.Equal(t, user.Name, item["name"])
	assert.Equal(t, user.Email, item["email"])
	assert.Nil(t, item["password"])
	assert.NotNil(t, item["CreatedAt"])
	assert.NotNil(t, item["UpdatedAt"])
	// assert.NotNil(t, item["DeletedAt"])
}
