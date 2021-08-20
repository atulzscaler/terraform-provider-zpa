package policysetglobal

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig = "/mgmtconfig/v1/admin/customers/"
)

type PolicySet struct {
	CreationTime int32   `json:"creationTime,string"`
	Description  string  `json:"description"`
	Enabled      bool    `json:"enabled"`
	ID           int64   `json:"id,string"`
	ModifiedBy   int64   `json:"modifiedBy,string"`
	ModifiedTime int32   `json:"modifiedTime,string"`
	Name         string  `json:"name"`
	PolicyType   int32   `json:"policyType,string"`
	Rules        []Rules `json:"rules,omitempty"`
}

type Rules struct {
	Action            string        `json:"action"`
	ActionID          int64         `json:"actionId,string"`
	CreationTime      int32         `json:"creationTime,string"`
	CustomMsg         string        `json:"customMsg"`
	Description       string        `json:"description"`
	ID                int64         `json:"id,string"`
	ModifiedBy        int64         `json:"modifiedBy,string"`
	ModifiedTime      int32         `json:"modifiedTime,string"`
	Name              string        `json:"name"`
	Operator          string        `json:"operator"`
	PolicySetID       int64         `json:"policySetId,string"`
	PolicyType        int32         `json:"policyType,string"`
	Priority          int32         `json:"priority,string"`
	ReauthIdleTimeout int32         `json:"reauthIdleTimeout,string"`
	ReauthTimeout     int32         `json:"reauthTimeout,string"`
	RuleOrder         int32         `json:"ruleOrder,string"`
	ZpnCbiProfileID   int64         `json:"zpnCbiProfileId,string"`
	Conditions        *[]Conditions `json:"conditions"`
}
type Conditions struct {
	CreationTime int32       `json:"creationTime,string"`
	ID           int64       `json:"id,string"`
	ModifiedBy   int64       `json:"modifiedBy,string"`
	ModifiedTime int32       `json:"modifiedTime,string"`
	Negated      bool        `json:"negated"`
	Operands     *[]Operands `json:"operands"`
	Operator     string      `json:"operator"`
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

/*
type AppServerGroups struct {
	ConfigSpace      string `json:"configSpace"`
	CreationTime     int32  `json:"creationTime,string"`
	Description      string `json:"description"`
	Enabled          bool   `json:"enabled"`
	ID               int64  `json:"id,string"`
	DynamicDiscovery bool   `json:"dynamicDiscovery"`
	ModifiedBy       int64  `json:"modifiedBy,string"`
	ModifiedTime     int32  `json:"modifiedTime,string"`
	Name             string `json:"name"`
}

type Connectors struct {
	ApplicationStartTime     int      `json:"applicationStartTime"`
	AppConnectorGroupID      string   `json:"appConnectorGroupId"`
	AppConnectorGroupName    string   `json:"appConnectorGroupName"`
	ControlChannelStatus     string   `json:"controlChannelStatus"`
	CreationTime             int      `json:"creationTime"`
	CtrlBrokerName           string   `json:"ctrlBrokerName"`
	CurrentVersion           string   `json:"currentVersion"`
	Description              string   `json:"description"`
	Enabled                  bool     `json:"enabled"`
	ExpectedUpgradeTime      int      `json:"expectedUpgradeTime"`
	ExpectedVersion          string   `json:"expectedVersion"`
	Fingerprint              string   `json:"fingerprint"`
	ID                       int      `json:"id"`
	IPACL                    []string `json:"ipAcl"`
	IssuedCertID             int      `json:"issuedCertId"`
	LastBrokerConnectTime    int      `json:"lastBrokerConnectTime"`
	LastBrokerDisconnectTime int      `json:"lastBrokerDisconnectTime"`
	LastUpgradeTime          int      `json:"lastUpgradeTime"`
	Latitude                 int      `json:"latitude"`
	Location                 string   `json:"location"`
	Longitude                int      `json:"longitude"`
	ModifiedBy               int      `json:"modifiedBy"`
	ModifiedTime             int      `json:"modifiedTime"`
	Name                     string   `json:"name"`
	Platform                 string   `json:"platform"`
	PreviousVersion          string   `json:"previousVersion"`
	PrivateIP                string   `json:"privateIp"`
	PublicIP                 string   `json:"publicIp"`
	//	SigningCert              SigningCert `json:"signingCert"`
	UpgradeAttempt int    `json:"upgradeAttempt"`
	UpgradeStatus  string `json:"upgradeStatus"`
}

type AppServerGroups struct {
	ConfigSpace      string `json:"configSpace"`
	CreationTime     int32  `json:"creationTime,string"`
	Description      string `json:"description"`
	Enabled          bool   `json:"enabled"`
	ID               int64  `json:"id,string"`
	DynamicDiscovery bool   `json:"dynamicDiscovery"`
	ModifiedBy       int64  `json:"modifiedBy,string"`
	ModifiedTime     int32  `json:"modifiedTime,string"`
	Name             string `json:"name"`
}
type AppConnectorGroups struct {
	Connectors            []Connectors   `json:"connectors"`
	CityCountry           string         `json:"cityCountry"`
	CountryCode           string         `json:"countryCode"`
	CreationTime          int            `json:"creationTime"`
	Description           string         `json:"description"`
	DNSQueryType          string         `json:"dnsQueryType"`
	Enabled               bool           `json:"enabled"`
	GeoLocationID         int            `json:"geoLocationId"`
	ID                    int            `json:"id"`
	Latitude              string         `json:"latitude"`
	Location              string         `json:"location"`
	Longitude             string         `json:"longitude"`
	ModifiedBy            int            `json:"modifiedBy"`
	ModifiedTime          int            `json:"modifiedTime"`
	Name                  string         `json:"name"`
	ServerGroups          []ServerGroups `json:"serverGroups"`
	SiemAppConnectorGroup bool           `json:"siemAppConnectorGroup"`
	UpgradeDay            string         `json:"upgradeDay"`
	UpgradeTimeInSecs     string         `json:"upgradeTimeInSecs"`
	VersionProfileID      int            `json:"versionProfileId"`
}

*/

func (service *Service) Get() (*PolicySet, *http.Response, error) {
	v := new(PolicySet)
	relativeURL := fmt.Sprintf(mgmtConfig + service.Client.Config.CustomerID + "/policySet/global")
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(policySetId string) (*PolicySet, *http.Response, error) {
	v := new(PolicySet)
	relativeURL := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rule", policySetId)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(policySetId string, ruleId PolicySet) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rules/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(policySetId string, ruleId PolicySet) (*http.Response, error) {
	path := fmt.Sprintf(mgmtConfig+service.Client.Config.CustomerID+"/policySet/%v/rules/%v", policySetId, ruleId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
