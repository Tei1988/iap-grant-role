package auth

import (
	"crypto/x509"
	"encoding/base64"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"github.com/tei1988/iap-grant-role/common"
)

func NewIAPAuthProvider(config map[string]interface{}) IAuthProvider {
	var ap iapAuthProvider
	mapstructure.Decode(config, &ap)
	return ap
}

type PublicKeyMap map[string]string

type iapAuthProvider struct {
	PublicKeyMap PublicKeyMap
}

func (ap iapAuthProvider) Authenticate(r *http.Request) (common.EmailAddress, error) {
	jwtTokens, _ := r.Header["X-Goog-Iap-Jwt-Assertion"]
	token, _ := jwt.Parse(jwtTokens[0], ap.findPublicKey)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		return common.EmailAddress(claims["email"].(string)), nil
	}
	return "", nil
}

func (ap iapAuthProvider) findPublicKey(t *jwt.Token) (interface{}, error) {
	key := ap.PublicKeyMap[t.Header["kid"].(string)]
	asn1, _ := base64.URLEncoding.DecodeString(key)
	publicKey, _ := x509.ParsePKIXPublicKey(asn1)
	return publicKey, nil
}
