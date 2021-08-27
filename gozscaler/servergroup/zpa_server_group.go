package servergroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig          = "/mgmtconfig/v1/admin/customers/"
	serverGroupEndpoint = "/serverGroup"
)

type ServerGroup struct {
	ID               string `json:"id"`
	Enabled          bool   `json:"enabled"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	IpAnchored       bool   `json:"ipAnchored"`
	ConfigSpace      string `json:"configSpace"`
	DynamicDiscovery bool   `json:"dynamicDiscovery"`
	CreationTime     int32  `json:"creationTime,string"`
	ModifiedBy       string `json:"modifiedBy"`
	ModifiedTime     int32  `json:"modifiedTime,string"`
	// AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups"`
	// Servers            []ApplicationServer  `json:"servers"`
	// Applications       []Applications       `json:"applications"`
}

/*
type Applications struct {
	ID   int64  `json:"id,string,omitempty"`
	Name string `json:"name,omitempty"`
}
type AppConnectorGroups struct {
	Citycountry           string            `json:"cityCountry,omitempty"`
	CountryCode           string            `json:"countryCode,omitempty"`
	CreationTime          int32             `json:"creationTime,string,omitempty"`
	Description           string            `json:"description,omitempty"`
	DnsqueryType          string            `json:"dnsQueryType,omitempty"`
	Enabled               bool              `json:"enabled,omitempty"`
	GeolocationId         int64             `json:"geoLocationId,string,omitempty"`
	ID                    int64             `json:"id,string,omitempty"`
	Latitude              string            `json:"latitude,omitempty"`
	Location              string            `json:"location,omitempty"`
	Longitude             string            `json:"longitude,omitempty"`
	ModifiedBy            int64             `json:"modifiedBy,string,omitempty"`
	ModifiedTime          int32             `json:"modifiedTime,string,omitempty"`
	Name                  string            `json:"name"`
	SiemAppconnectorGroup bool              `json:"siemAppConnectorGroup,omitempty"`
	UpgradeDay            string            `json:"upgradeDay,omitempty"`
	UpgradeTimeinSecs     string            `json:"upgradeTimeInSecs,omitempty"`
	VersionProfileId      int64             `json:"versionProfileId,string,omitempty"`
	AppServerGroups       []AppServerGroups `json:"serverGroups,omitempty"`
	Connectors            []Connectors      `json:"connectors,omitempty"`
}

type Connectors struct {
	ApplicationStartTime     int64             `json:"applicationStartTime,string,omitempty"`
	AppConnectorGroupId      string            `json:"appConnectorGroupId,omitempty"`
	AppConnectorGroupName    string            `json:"appConnectorGroupName,omitempty"`
	ControlChannelStatus     string            `json:"controlChannelStatus,omitempty"`
	CreationTime             int32             `json:"creationTime,string,omitempty"`
	CtrlBrokerName           string            `json:"ctrlBrokerName,omitempty"`
	CurrentVersion           string            `json:"currentVersion,omitempty"`
	Description              string            `json:"description,omitempty"`
	Enabled                  bool              `json:"enabled,omitempty"`
	ExpectedUpgradeTime      int64             `json:"expectedUpgradeTime,string,omitempty"`
	ExpectedVersion          string            `json:"expectedVersion,omitempty"`
	Fingerprint              string            `json:"fingerprint,omitempty"`
	ID                       int64             `json:"id,string,omitempty"`
	IpAcl                    []string          `json:"ipAcl,omitempty"`
	IssuedCertId             int64             `json:"issuedCertId,string,omitempty"`
	LastBrokerConnecttime    int64             `json:"lastBrokerConnectTime,string,omitempty"`
	LastBrokerDisconnectTime int64             `json:"lastBrokerDisconnectTime,string,omitempty"`
	LastUpgradeTime          int64             `json:"lastUpgradeTime,string,omitempty"`
	Latitude                 float64           `json:"latitude,string,omitempty"`
	Location                 string            `json:"location,omitempty"`
	Longitude                float64           `json:"longitude,string,omitempty"`
	ModifiedBy               int64             `json:"modifiedBy,string,omitempty"`
	ModifiedTime             int32             `json:"modifiedTime,omitempty"`
	Name                     string            `json:"name"`
	Platform                 string            `json:"platform,omitempty"`
	PreviousVersion          string            `json:"previousVersion,omitempty"`
	PrivateIp                string            `json:"privateIp,omitempty"`
	PublicIp                 string            `json:"publicIp,omitempty"`
	SigningCert              map[string]string `json:"signingCert,omitempty"`
	UpgradeAttempt           int32             `json:"upgradeAttempt,string,omitempty"`
	UpgradeStatus            string            `json:"upgradeStatus,omitempty"`
}

type AppServerGroups struct {
	ConfigSpace      string `json:"configSpace,omitempty"`
	CreationTime     int32  `json:"creationTime,string,omitempty"`
	Description      string `json:"description,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	ID               int64  `json:"id,string,omitempty"`
	DynamicDiscovery bool   `json:"dynamicDiscovery,omitempty"`
	ModifiedBy       int64  `json:"modifiedBy,string,omitempty"`
	ModifiedTime     int32  `json:"modifiedTime,string,omitempty"`
	Name             string `json:"name"`
}

type ApplicationServer struct {
	Address           string   `json:"address,omitempty"`
	AppServerGroupIds []string `json:"appServerGroupIds,omitempty"`
	ConfigSpace       string   `json:"configSpace,omitempty"`
	CreationTime      int32    `json:"creationTime,string,omitempty"`
	Description       string   `json:"description,omitempty"`
	Enabled           bool     `json:"enabled,omitempty"`
	ID                int64    `json:"id,string,omitempty"`
	ModifiedBy        int64    `json:"modifiedBy,string,omitempty"`
	ModifiedTime      int32    `json:"modifiedTime,string,omitempty"`
	Name              string   `json:"name"`
}
*/
func (service *Service) Get(groupId string) (*ServerGroup, *http.Response, error) {
	v := new(ServerGroup)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, groupId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(serverGroup ServerGroup) (*ServerGroup, *http.Response, error) {
	v := new(ServerGroup)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, nil, serverGroup, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(groupId string, serverGroup ServerGroup) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, groupId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, serverGroup, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(groupId string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, groupId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
