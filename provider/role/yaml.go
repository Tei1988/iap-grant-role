package role

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tei1988/iap-grant-role/common"
	yaml "gopkg.in/yaml.v2"
)

func NewYamlRoleProvider(config map[string]interface{}) IRoleProvider {
	filepath := config["path"].(string)

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s not found.", filepath))
	}
	var rp yamlRoleProvider
	err = yaml.Unmarshal(data, &rp)
	if err != nil {
		log.Fatal(fmt.Sprintf("wrong %s found", filepath), err)
	}

	return rp
}

type userRoleMap map[common.EmailAddress]common.Role

type yamlRoleProvider struct {
	UserRoleMap userRoleMap `yaml:"role"`
}

func (rp yamlRoleProvider) FindRole(e *common.EmailAddress) (common.Role, error) {
	return rp.UserRoleMap[*e], nil
}
