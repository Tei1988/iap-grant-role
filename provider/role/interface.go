package role

import "github.com/tei1988/iap-grant-role/common"

type IRoleProvider interface {
	FindRole(*common.EmailAddress) (common.Role, error)
}
