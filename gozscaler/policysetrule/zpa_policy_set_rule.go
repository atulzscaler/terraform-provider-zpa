package policysetrule

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig = "/mgmtconfig/v1/admin/customers/"
)

type PolicyRule struct {
	Action                   string               `json:"action"`
	ActionID                 int                  `json:"actionId,string"`
	BypassDefaultRule        bool                 `json:"bypassDefaultRule"`
	CreationTime             int                  `json:"creationTime,string"`
	CustomMsg                string               `json:"customMsg"`
	Description              string               `json:"description"`
	ID                       int                  `json:"id,string"`
	IsolationDefaultRule     bool                 `json:"isolationDefaultRule"`
	ModifiedBy               int                  `json:"modifiedBy,string"`
	ModifiedTime             int                  `json:"modifiedTime,string"`
	Name                     string               `json:"name"`
	Operator                 string               `json:"operator"`
	PolicySetID              int                  `json:"policySetId,string"`
	PolicyType               int                  `json:"policyType,string"`
	Priority                 int                  `json:"priority,string"`
	ReauthDefaultRule        bool                 `json:"reauthDefaultRule"`
	ReauthIdleTimeout        int                  `json:"reauthIdleTimeout,string"`
	ReauthTimeout            int                  `json:"reauthTimeout,string"`
	RuleOrder                int                  `json:"ruleOrder,string"`
	ZpnCbiProfileID          int                  `json:"zpnCbiProfileId,string"`
	ZpnInspectionProfileId   int                  `json:"zpnInspectionProfileId,string"`
	ZpnInspectionProfileName string               `json:"zpnInspectionProfileName,string"`
	Conditions               []Conditions         `json:"conditions"`
	AppServerGroups          []AppServerGroups    `json:"appServerGroups"`
	AppConnectorGroups       []AppConnectorGroups `json:"appConnectorGroups"`
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
func (service *Service) Get(policySetID PolicyRule, ruleId string) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	url := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetID.PolicySetID, ruleId)
	resp, err := service.Client.NewRequestDo("GET", url, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// POST --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule
func (service *Service) Create(policySetID PolicyRule) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule", policySetID.PolicySetID)
	resp, err := service.Client.NewRequestDo("POST", path, nil, &v, nil)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// PUT --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Update(policySetID PolicyRule, ruleId string, policySetRule PolicyRule) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetID.PolicySetID, ruleId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, policySetRule, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// DELETE --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Delete(policySetID PolicyRule, ruleId string) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule/%v", policySetID.PolicySetID, ruleId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
