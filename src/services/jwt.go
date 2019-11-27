/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package services

import (
	"log"
	"os"

	"crypto/ecdsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var PublicKeyAccess *ecdsa.PublicKey
var PrivateKeyAccess *ecdsa.PrivateKey
var PublicKeyRefresh *ecdsa.PublicKey
var PrivateKeyRefresh *ecdsa.PrivateKey

func Jwt() {
	var tmp []byte
	var err error
	var root = os.Getenv("ROOT") + "/source/services"

	if tmp, err = ioutil.ReadFile(root + "/jwt/access.public.pem"); err != nil {
		log.Fatal(err)
	}
	if PublicKeyAccess, err = jwt.ParseECPublicKeyFromPEM(tmp); err != nil {
		log.Fatal(err)
	}

	if tmp, err = ioutil.ReadFile(root + "/jwt/access.private.pem"); err != nil {
		log.Fatal(err)
	}
	if PrivateKeyAccess, err = jwt.ParseECPrivateKeyFromPEM(tmp); err != nil {
		log.Fatal(err)
	}

	if tmp, err = ioutil.ReadFile(root + "/jwt/refresh.public.pem"); err != nil {
		log.Fatal(err)
	}
	if PublicKeyRefresh, err = jwt.ParseECPublicKeyFromPEM(tmp); err != nil {
		log.Fatal(err)
	}

	if tmp, err = ioutil.ReadFile(root + "/jwt/refresh.private.pem"); err != nil {
		log.Fatal(err)
	}
	if PrivateKeyRefresh, err = jwt.ParseECPrivateKeyFromPEM(tmp); err != nil {
		log.Fatal(err)
	}
}
