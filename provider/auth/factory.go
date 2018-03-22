package auth

import (
	"fmt"
	"log"

	"github.com/tei1988/iap-grant-role/common"
)

func AuthProviderFactory(config common.ProviderConfig) IAuthProvider {
	var ap IAuthProvider
	switch config.Name {
	case "iap":
		ap = NewIAPAuthProvider(config.Options)
	default:
		log.Fatal(fmt.Sprintf("%s is not registered.", config.Name))
	}
	return ap
}
