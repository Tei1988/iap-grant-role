package auth

import (
	"net/http"

	"github.com/tei1988/iap-grant-role/common"
)

type IAuthProvider interface {
	Authenticate(*http.Request) (common.EmailAddress, error)
}
