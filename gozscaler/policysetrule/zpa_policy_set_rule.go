package policysetrule

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig = "/mgmtconfig/v1/admin/customers/"
)

type PolicyRule struct {
	Action             string               `json:"action,omitempty"`
	ActionID           int                  `json:"actionId,string,omitempty"`
	CreationTime       int                  `json:"creationTime,string,omitempty"`
	CustomMsg          string               `json:"customMsg,omitempty"`
	Description        string               `json:"description,omitempty"`
	ID                 int                  `json:"id,string,omitempty"`
	ModifiedBy         int                  `json:"modifiedBy,string,omitempty"`
	ModifiedTime       int                  `json:"modifiedTime,string,omitempty"`
	Name               string               `json:"name"`
	Operator           string               `json:"operator,omitempty"`
	PolicySetID        int                  `json:"policySetId,string,omitempty"`
	PolicyType         int                  `json:"policyType,string,omitempty"`
	Priority           int                  `json:"priority,string,omitempty"`
	ReauthIdleTimeout  int                  `json:"reauthIdleTimeout,string,omitempty"`
	ReauthTimeout      int                  `json:"reauthTimeout,string,omitempty"`
	RuleOrder          int                  `json:"ruleOrder,string,omitempty"`
	ZpnCbiProfileID    int                  `json:"zpnCbiProfileId,string,omitempty"`
	Conditions         []Conditions         `json:"conditions,omitempty"`
	AppServerGroups    []AppServerGroups    `json:"appServerGroups,omitempty"`
	AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups,omitempty"`
}

type Conditions struct {
	CreationTime int32      `json:"creationTime,string,omitempty"`
	ID           int64      `json:"id,string,omitempty"`
	ModifiedBy   int64      `json:"modifiedBy,string,omitempty"`
	ModifiedTime int32      `json:"modifiedTime,string,omitempty"`
	Negated      bool       `json:"negated,omitempty"`
	Operands     []Operands `json:"operands,omitempty"`
	Operator     string     `json:"operator,omitempty"`
}
type Operands struct {
	CreationTime int32  `json:"creationTime,string,omitempty"`
	ID           int64  `json:"id,string,omitempty"`
	IdpID        int64  `json:"idpId,string,omitempty"`
	LHS          string `json:"lhs,omitempty"`
	ModifiedBy   int64  `json:"modifiedBy,string,omitempty"`
	ModifiedTime int32  `json:"modifiedTime,string,omitempty"`
	Name         string `json:"name"`
	ObjectType   string `json:"objectType,omitempty"`
	RHS          string `json:"rhs,omitempty"`
}

type AppServerGroups struct {
	ID   int64  `json:"id,string,omitempty"`
	Name string `json:"name"`
}

type AppConnectorGroups struct {
	ID           int64          `json:"id,string,omitempty"`
	Name         string         `json:"name"`
	Connectors   []Connectors   `json:"connectors,omitempty"`
	ServerGroups []ServerGroups `json:"serverGroups,omitempty"`
}

type Connectors struct {
	Name string `json:"name"`
	ID   int64  `json:"id,string,omitempty"`
}

type ServerGroups struct {
	Name string `json:"name"`
	ID   int64  `json:"id,string,omitempty"`
}

// GET --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Get(policySetId string, ruleId string) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	url := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("GET", url, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// POST --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule
func (service *Service) Create(policySetId PolicyRule) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule", policySetId)
	resp, err := service.Client.NewRequestDo("POST", path, nil, &v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// PUT --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Update(policySetId, ruleId string, policySetRule PolicyRule) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, policySetRule, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// DELETE --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Delete(policySetId, ruleId string) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
