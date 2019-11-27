/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package unit

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/sundowndev/api-go-example/src/validators"
)

// TestShouldSuccessValidatorUserCreate
func TestShouldSuccessValidatorUserCreate(t *testing.T) {
	var c *gin.Context
	var data = url.Values{"name": {"test"}, "email": {"example@test.com"}, "password": {"12345678"}}

	_, c = CreateContext(data)

	validators.ValidateUserCreate(c)

	assert.Equal(t, false, c.IsAborted())
}

// TestShouldFailStringToIntValidatorUserCreate
func TestShouldFailEmailBadFormatValidatorUserCreate(t *testing.T) {
	var res *httptest.ResponseRecorder
	var err error
	var c *gin.Context
	var result map[string]interface{}
	var data = url.Values{"name": {"test"}, "email": {"exampletest.com"}, "password": {"12345678"}}

	res, c = CreateContext(data)

	validators.ValidateUserCreate(c)

	if err = json.Unmarshal(res.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}

	msg := result["error"].(string)

	assert.Equal(t, true, c.IsAborted())
	assert.Equal(t, "Key: 'userCreate.Email' Error:Field validation for 'Email' failed on the 'email' tag", msg)
}
