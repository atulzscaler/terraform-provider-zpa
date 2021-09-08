package policysetglobal

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

/*
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

// Get the authentication policy and all rules for a Client Forwarding Policy Rule
func (service *Service) GetBypass() (*PolicySet, *http.Response, error) {
	v := new(PolicySet)
	relativeURL := fmt.Sprintf(mgmtConfig + service.Client.Config.CustomerID + "/policySet/bypass")
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
