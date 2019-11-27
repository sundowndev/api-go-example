/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package functional

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"regexp"
	"testing"
)

// TestLogin
func TestLogin(t *testing.T) {
	var body map[string]interface{}
	var err error
	var data = url.Values{"email": {"example@test.com"}, "password": {"12345678"}}

	if body, err = RequestApiJson("POST", BaseUrl+"/auth", data, false); err != nil {
		t.Fatal(err)
	}

	assert.Regexp(t, regexp.MustCompile(RegexToken), body["access_token"])
	assert.Regexp(t, regexp.MustCompile(RegexToken), body["refresh_token"])
}
