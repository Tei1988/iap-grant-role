package role

import (
	"log"

	"github.com/tei1988/iap-grant-role/common"
)

func NewPassThroughRoleProvider(_ map[string]interface{}) IRoleProvider {
	log.Println("A pass-through role provider is initialized.")
	return passThroughRoleProvider{}
}

type passThroughRoleProvider struct {
}

func (rp passThroughRoleProvider) FindRole(e *common.EmailAddress) (common.Role, error) {
	return common.Role{}, nil
}
