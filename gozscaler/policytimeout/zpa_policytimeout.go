package policytimeout

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig = "/mgmtconfig/v1/admin/customers/"
)

type PolicySet struct {
	CreationTime string  `json:"creationTime,omitempty"`
	Description  string  `json:"description,omitempty"`
	Enabled      bool    `json:"enabled"`
	ID           string  `json:"id,omitempty"`
	ModifiedBy   string  `json:"modifiedBy,omitempty"`
	ModifiedTime string  `json:"modifiedTime,omitempty"`
	Name         string  `json:"name,omitempty"`
	PolicyType   string  `json:"policyType,omitempty"`
	Rules        []Rules `json:"rules"`
}

type Rules struct {
	Action                   string       `json:"action,omitempty"`
	ActionID                 string       `json:"actionId,omitempty"`
	BypassDefaultRule        bool         `json:"bypassDefaultRule,omitempty"`
	CreationTime             string       `json:"creationTime,omitempty"`
	CustomMsg                string       `json:"customMsg,omitempty"`
	Description              string       `json:"description,omitempty"`
	ID                       string       `json:"id,omitempty"`
	IsolationDefaultRule     bool         `json:"isolationDefaultRule,omitempty"`
	ModifiedBy               string       `json:"modifiedBy,omitempty"`
	ModifiedTime             string       `json:"modifiedTime,omitempty"`
	Name                     string       `json:"name,omitempty"`
	Operator                 string       `json:"operator,omitempty"`
	PolicySetID              string       `json:"policySetId,omitempty"`
	PolicyType               string       `json:"policyType,omitempty"`
	Priority                 string       `json:"priority,omitempty"`
	ReauthDefaultRule        bool         `json:"reauthDefaultRule,omitempty"`
	ReauthIdleTimeout        string       `json:"reauthIdleTimeout,omitempty"`
	ReauthTimeout            string       `json:"reauthTimeout,omitempty"`
	RuleOrder                string       `json:"ruleOrder,omitempty"`
	ZpnCbiProfileID          string       `json:"zpnCbiProfileId,omitempty"`
	ZpnInspectionProfileId   string       `json:"zpnInspectionProfileId,omitempty"`
	ZpnInspectionProfileName string       `json:"zpnInspectionProfileName,omitempty"`
	Conditions               []Conditions `json:"conditions,omitempty"`
}
type Conditions struct {
	CreationTime string      `json:"creationTime,omitempty"`
	ID           string      `json:"id,omitempty"`
	ModifiedBy   string      `json:"modifiedBy,omitempty"`
	ModifiedTime string      `json:"modifiedTime,omitempty"`
	Negated      bool        `json:"negated,omitempty"`
	Operands     *[]Operands `json:"operands,omitempty"`
	Operator     string      `json:"operator,omitempty"`
}
type Operands struct {
	CreationTime string `json:"creationTime,omitempty"`
	ID           string `json:"id,omitempty"`
	IdpID        string `json:"idpId,omitempty"`
	LHS          string `json:"lhs,omitempty"`
	ModifiedBy   string `json:"modifiedBy,omitempty"`
	ModifiedTime string `json:"modifiedTime,omitempty"`
	Name         string `json:"name,omitempty"`
	ObjectType   string `json:"objectType,omitempty"`
	RHS          string `json:"rhs,omitempty"`
}
type AppServerGroups struct {
	ConfigSpace      string `json:"configSpace,omitempty"`
	CreationTime     string `json:"creationTime,omitempty"`
	Description      string `json:"description,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	ID               string `json:"id,omitempty"`
	DynamicDiscovery bool   `json:"dynamicDiscovery,omitempty"`
	ModifiedBy       string `json:"modifiedBy,omitempty"`
	ModifiedTime     string `json:"modifiedTime,omitempty"`
	Name             string `json:"name,omitempty"`
}

// Get the authentication policy and all rules for a Timeout policy rule
func (service *Service) GetReauth() (*PolicySet, *http.Response, error) {
	v := new(PolicySet)
	relativeURL := fmt.Sprintf(mgmtConfig + service.Client.Config.CustomerID + "/policySet/reauth")
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*
// POST --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule
func (service *Service) CreateReath(rule *PolicySet) (*PolicySet, *http.Response, error) {
	v := new(PolicySet)
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%s/rule", rule.PolicySetID)
	resp, err := service.Client.NewRequestDo("POST", path, nil, &rule, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

// PUT --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) UpdateReath(policySetID, ruleId string, policySetRule *PolicySet) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%s/rule/%s", policySetID, ruleId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, policySetRule, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// DELETE --> mgmtconfig​/v1​/admin​/customers​/{customerId}​/policySet​/{policySetId}​/rule​/{ruleId}
func (service *Service) DeleteReath(policySetID, ruleId string) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%s/rule/%s", policySetID, ruleId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
*/
