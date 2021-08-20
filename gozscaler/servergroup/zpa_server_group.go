package servergroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig          = "/mgmtconfig/v1/admin/customers/"
	serverGroupEndpoint = "/serverGroup"
)

type ServerGroupRequest struct {
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	Enabled            bool                 `json:"enabled"`
	DynamicDiscovery   bool                 `json:"dynamicDiscovery"`
	IpAnchored         bool                 `json:"ipAnchored"`
	AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups,omitempty"`
	ApplicationServers []ApplicationServers `json:"servers,omitempty"`
	Applications       []Applications       `json:"applications,omitempty"`
}
type ServerGroupResponse struct {
	ID                 string               `json:"id"`
	Enabled            bool                 `json:"enabled"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	IpAnchored         bool                 `json:"ipAnchored"`
	ConfigSpace        string               `json:"configSpace"`
	DynamicDiscovery   bool                 `json:"dynamicDiscovery"`
	CreationTime       int32                `json:"creationTime,string"`
	ModifiedBy         string               `json:"modifiedBy"`
	ModifiedTime       int32                `json:"modifiedTime,string"`
	AppConnectorGroups []AppConnectorGroups `json:"appConnectorGroups,omitempty"`
	ApplicationServers []ApplicationServers `json:"servers,omitempty"`
	Applications       []Applications       `json:"applications,omitempty"`
}

type Applications struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}
type AppConnectorGroups struct {
	Citycountry  string `json:"cityCountry"`
	CountryCode  string `json:"countryCode"`
	CreationTime int    `json:"creationTime,string"`
	Description  string `json:"description"`
	DnsqueryType string `json:"dnsQueryType"`
	Enabled      bool   `json:"enabled"`
	//GeolocationId         int64                `json:"geoLocationId,string"`
	ID                    int64  `json:"id,string"`
	Latitude              string `json:"latitude"`
	Location              string `json:"location"`
	Longitude             string `json:"longitude"`
	ModifiedBy            int64  `json:"modifiedBy,string"`
	ModifiedTime          int32  `json:"modifiedTime,string"`
	Name                  string `json:"name"`
	SiemAppconnectorGroup bool   `json:"siemAppConnectorGroup"`
	UpgradeDay            string `json:"upgradeDay"`
	UpgradeTimeinSecs     string `json:"upgradeTimeInSecs"`
	VersionProfileId      int64  `json:"versionProfileId,string"`
	//	ApplicationServers    []ApplicationServers `json:"serverGroups"`
	Connectors []Connectors `json:"connectors"`
}

type Connectors struct {
	ApplicationStartTime     int64    `json:"applicationStartTime,string"`
	AppConnectorGroupId      string   `json:"appConnectorGroupId"`
	AppConnectorGroupName    string   `json:"appConnectorGroupName"`
	ControlChannelStatus     string   `json:"controlChannelStatus"`
	CreationTime             int32    `json:"creationTime,string"`
	CtrlBrokerName           string   `json:"ctrlBrokerName"`
	CurrentVersion           string   `json:"currentVersion"`
	Description              string   `json:"description"`
	Enabled                  bool     `json:"enabled"`
	ExpectedUpgradeTime      int64    `json:"expectedUpgradeTime,string"`
	ExpectedVersion          string   `json:"expectedVersion"`
	Fingerprint              string   `json:"fingerprint"`
	ID                       int64    `json:"id,string"`
	IpAcl                    []string `json:"ipAcl"`
	IssuedCertId             int64    `json:"issuedCertId,string"`
	LastBrokerConnecttime    int64    `json:"lastBrokerConnectTime,string"`
	LastBrokerDisconnectTime int64    `json:"lastBrokerDisconnectTime,string"`
	LastUpgradeTime          int64    `json:"lastUpgradeTime,string"`
	Latitude                 int      `json:"latitude,string"` // Swagger shows number($double)
	Location                 string   `json:"location"`
	Longitude                int      `json:"longitude,string"` // Swagger shows number($double)
	ModifiedBy               int64    `json:"modifiedBy,string"`
	ModifiedTime             int32    `json:"modifiedTime"`
	Name                     string   `json:"name"`
	Platform                 string   `json:"platform"`
	PreviousVersion          string   `json:"previousVersion"`
	PrivateIp                string   `json:"privateIp"`
	PublicIp                 string   `json:"publicIp"`
	//	SigningCert              SigningCert `json:"signingCert"` // May be unecessay for Terraform. Will re-assess in the future and implement as new feature
	UpgradeAttempt int32  `json:"upgradeAttempt,string"`
	UpgradeStatus  string `json:"upgradeStatus"`
}

/*type SigningCert struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
*/

type ApplicationServers struct {
	Address           string   `json:"address"`
	AppServerGroupIds []string `json:"appServerGroupIds"`
	ConfigSpace       string   `json:"configSpace"`
	CreationTime      int32    `json:"creationTime,string"`
	Description       string   `json:"description"`
	Enabled           bool     `json:"enabled"`
	ID                int64    `json:"id,string"`
	ModifiedBy        int64    `json:"modifiedBy,string"`
	ModifiedTime      int32    `json:"modifiedTime,string"`
	Name              string   `json:"name"`
}

func (service *Service) Get(id string) (*ServerGroupResponse, *http.Response, error) {
	v := new(ServerGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]ServerGroupResponse, *http.Response, error) {
	v := new([]ServerGroupResponse)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(serverGroupRequest ServerGroupRequest) (*ServerGroupResponse, *http.Response, error) {
	v := new(ServerGroupResponse)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, nil, serverGroupRequest, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, serverGroupRequest ServerGroupRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, id)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, serverGroupRequest, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+serverGroupEndpoint, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
