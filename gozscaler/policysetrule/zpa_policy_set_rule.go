package policysetrule

/*
import (
	"fmt"
	"net/http"
	"strings"
)

const (
	mgmtConfig = "/mgmtconfig/v1/admin/customers/"
)

type PolicyRule struct {
	Action             string               `json:"action"`
	ActionID           int                  `json:"actionId,string"`
	CreationTime       int                  `json:"creationTime,string"`
	CustomMsg          string               `json:"customMsg"`
	Description        string               `json:"description"`
	ID                 int                  `json:"id,string"`
	ModifiedBy         int                  `json:"modifiedBy,string"`
	ModifiedTime       int                  `json:"modifiedTime,string"`
	Name               string               `json:"name"`
	Operator           string               `json:"operator"`
	PolicySetID        int                  `json:"policySetId,string"`
	PolicyType         int                  `json:"policyType,string"`
	Priority           int                  `json:"priority,string"`
	ReauthIdleTimeout  int                  `json:"reauthIdleTimeout,string"`
	ReauthTimeout      int                  `json:"reauthTimeout,string"`
	RuleOrder          int                  `json:"ruleOrder,string"`
	ZpnCbiProfileID    int                  `json:"zpnCbiProfileId,string"`
	Conditions         []Conditions         `json:"conditions"`
	AppServerGroups    []AppServerGroups    `json:"appServerGroups"`
	AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups"`
}

type Conditions struct {
	CreationTime int        `json:"creationTime,string"`
	ID           int        `json:"id,string"`
	ModifiedBy   int        `json:"modifiedBy,string"`
	ModifiedTime int        `json:"modifiedTime,string"`
	Negated      bool       `json:"negated"`
	Operands     []Operands `json:"operands"`
	Operator     string     `json:"operator"`
}
type Operands struct {
	CreationTime int    `json:"creationTime,string"`
	ID           int    `json:"id,string"`
	IdpID        int    `json:"idpId,string"`
	LHS          string `json:"lhs"`
	ModifiedBy   int    `json:"modifiedBy,string"`
	ModifiedTime int    `json:"modifiedTime,string"`
	Name         string `json:"name"`
	ObjectType   string `json:"objectType"`
	RHS          string `json:"rhs"`
}

type AppServerGroups struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
}

type AppConnectorGroups struct {
	ID           int            `json:"id,string"`
	Name         string         `json:"name"`
	Connectors   []Connectors   `json:"connectors"`
	ServerGroups []ServerGroups `json:"serverGroups"`
}

type Connectors struct {
	Name string `json:"name"`
	ID   int    `json:"id,string"`
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

func (service *Service) Create(customerId int64, policySetId int64, rule PolicyRule) (PolicyRule, *http.Response, error) {
	v := new(PolicyRule)
	localVarPath := a.client.cfg.BasePath + "/mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", fmt.Sprintf("%v", customerId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policySetId"+"}", fmt.Sprintf("%v", policySetId), -1)

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
*/
