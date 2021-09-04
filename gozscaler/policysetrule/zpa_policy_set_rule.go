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
	ActionID           int64                `json:"actionId,string,omitempty"`
	BypassDefaultRule  bool                 `json:"bypassDefaultRule,omitempty"`
	CustomMsg          string               `json:"customMsg,omitempty"`
	Description        string               `json:"description,omitempty"`
	ID                 int64                `json:"id,string,omitempty"`
	Name               string               `json:"name,omitempty"`
	Operator           string               `json:"operator,omitempty"`
	PolicySetID        int64                `json:"policySetId,string"`
	PolicyType         int32                `json:"policyType,string,omitempty"`
	Priority           int32                `json:"priority,string,omitempty"`
	ReauthDefaultRule  bool                 `json:"reauthDefaultRule,omitempty"`
	ReauthIdleTimeout  int32                `json:"reauthIdleTimeout,string,omitempty"`
	ReauthTimeout      int32                `json:"reauthTimeout,string,omitempty"`
	RuleOrder          int32                `json:"ruleOrder,string,omitempty"`
	Conditions         []Conditions         `json:"conditions,omitempty"`
	AppServerGroups    []AppServerGroups    `json:"appServerGroups,omitempty"`
	AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups,omitempty"`
}

type Conditions struct {
	ID       int64      `json:"id,string,omitempty"`
	Negated  bool       `json:"negated,omitempty"`
	Operands []Operands `json:"operands,omitempty"`
	Operator string     `json:"operator,omitempty"`
}
type Operands struct {
	ID         int64  `json:"id,string,omitempty"`
	IdpID      int64  `json:"idpId,string,omitempty"`
	LHS        string `json:"lhs,omitempty"`
	ObjectType string `json:"objectType,omitempty"`
	RHS        string `json:"rhs,omitempty"`
	Name       string `json:"name,omitempty"`
}

type AppServerGroups struct {
	ID int64 `json:"id,string,omitempty"`
}
type AppConnectorGroups struct {
	ID int64 `json:"id,string,omitempty"`
}

// GET --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Get(policySetID, ruleId int64) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	url := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%d/rule/%d", policySetID, ruleId)
	resp, err := service.Client.NewRequestDo("GET", url, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// POST --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule
func (service *Service) Create(rule *PolicyRule) (*PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%d/rule", rule.PolicySetID)
	resp, err := service.Client.NewRequestDo("POST", path, nil, &rule, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// PUT --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Update(policySetID, ruleId int64, policySetRule *PolicyRule) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%d/rule/%d", policySetID, ruleId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, policySetRule, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// DELETE --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) Delete(policySetID, ruleId int64) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%d/rule/%d", policySetID, ruleId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
