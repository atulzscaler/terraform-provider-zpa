package policysetrule

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig = "/mgmtconfig/v1/admin/customers/"
)

type PolicyRule struct {
	Action             string               `json:"action"`
	ActionID           int64                `json:"actionId,string"`
	CreationTime       int32                `json:"creationTime,string"`
	CustomMsg          string               `json:"customMsg"`
	Description        string               `json:"description"`
	ID                 int64                `json:"id,string"`
	ModifiedBy         int64                `json:"modifiedBy,string"`
	ModifiedTime       int32                `json:"modifiedTime,string"`
	Name               string               `json:"name"`
	Operator           string               `json:"operator"`
	PolicySetID        int64                `json:"policySetId,string"`
	PolicyType         int32                `json:"policyType,string"`
	Priority           int32                `json:"priority,string"`
	ReauthIdleTimeout  int32                `json:"reauthIdleTimeout,string"`
	ReauthTimeout      int32                `json:"reauthTimeout,string"`
	RuleOrder          int32                `json:"ruleOrder,string"`
	ZpnCbiProfileID    int64                `json:"zpnCbiProfileId,string"`
	Conditions         []Conditions         `json:"conditions"`
	AppServerGroups    []AppServerGroups    `json:"appServerGroups"`
	AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups"`
}

type Conditions struct {
	CreationTime int32      `json:"creationTime,string"`
	ID           int64      `json:"id,string"`
	ModifiedBy   int64      `json:"modifiedBy,string"`
	ModifiedTime int32      `json:"modifiedTime,string"`
	Negated      bool       `json:"negated"`
	Operands     []Operands `json:"operands"`
	Operator     string     `json:"operator"`
}
type Operands struct {
	CreationTime int32  `json:"creationTime,string"`
	ID           int64  `json:"id,string"`
	IdpID        int64  `json:"idpId,string"`
	LHS          string `json:"lhs"`
	ModifiedBy   int64  `json:"modifiedBy,string"`
	ModifiedTime int32  `json:"modifiedTime,string"`
	Name         string `json:"name"`
	ObjectType   string `json:"objectType"`
	RHS          string `json:"rhs"`
}

type AppServerGroups struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

type AppConnectorGroups struct {
	ID           int64          `json:"id,string"`
	Name         string         `json:"name"`
	Connectors   []Connectors   `json:"connectors"`
	ServerGroups []ServerGroups `json:"serverGroups"`
}

type Connectors struct {
	Name string `json:"name"`
	ID   int64  `json:"id,string"`
}

type ServerGroups struct {
	Name string `json:"name"`
	ID   int64  `json:"id,string"`
}

func (service *Service) Get(policySetId, ruleId string) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	url := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("GET", url, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(policySetId string) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule", policySetId, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(policySetId, ruleId string, policySetRule PolicyRule) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, policySetRule, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(policySetId, ruleId string) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
