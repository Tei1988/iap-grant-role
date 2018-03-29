package role

import (
	"fmt"
	"log"

	"github.com/tei1988/iap-grant-role/common"
)

func RoleProviderFactory(config common.ProviderConfig) IRoleProvider {
	var rp IRoleProvider
	switch config.Name {
	case "yaml":
		rp = NewYamlRoleProvider(config.Options)
	case "pass-through":
		rp = NewPassThroughRoleProvider(config.Options)
	default:
		log.Fatal(fmt.Sprintf("%s is not registered.", config.Name))
	}
	return rp
}
