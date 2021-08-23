package appconnectorgroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                = "/mgmtconfig/v1/admin/customers/"
	appConnectorGroupEndpoint = "/appConnectorGroup"
)

type AppConnectorGroupRequest struct {
	AppConnectorGroupResponse []AppConnectorGroupResponse `json:"list"`
	Connectors                []Connectors                `json:"connectors"`
	CityCountry               string                      `json:"cityCountry,omitempty"`
	CountryCode               string                      `json:"countryCode,omitempty"`
	CreationTime              int32                       `json:"creationTime,string"`
	Description               string                      `json:"description,omitempty"`
	DNSQueryType              string                      `json:"dnsQueryType,omitempty"`
	Enabled                   bool                        `json:"enabled,omitempty"`
	// GeoLocationID             int64                       `json:"geoLocationId,string"`
	ID                    int              `json:"id,string"`
	Latitude              string           `json:"latitude,omitempty"`
	Location              string           `json:"location,omitempty"`
	Longitude             string           `json:"longitude,omitempty"`
	ModifiedBy            int64            `json:"modifiedBy,string"`
	ModifiedTime          int32            `json:"modifiedTime,string"`
	Name                  string           `json:"name"`
	AppServerGroup        []AppServerGroup `json:"serverGroups,omitempty"`
	SiemAppConnectorGroup bool             `json:"siemAppConnectorGroup,omitempty"`
	UpgradeDay            string           `json:"upgradeDay,omitempty"`
	UpgradeTimeInSecs     string           `json:"upgradeTimeInSecs,omitempty"`
	VersionProfileID      int64            `json:"versionProfileId,string"`
}

type AppConnectorGroupResponse struct {
	Connectors   *[]Connectors `json:"connectors"`
	CityCountry  string        `json:"cityCountry,omitempty"`
	CountryCode  string        `json:"countryCode,omitempty"`
	CreationTime int32         `json:"creationTime,string"`
	Description  string        `json:"description,omitempty"`
	DNSQueryType string        `json:"dnsQueryType,omitempty"`
	Enabled      bool          `json:"enabled,omitempty"`
	// GeoLocationID         int64            `json:"geoLocationId,string"`
	ID                    int              `json:"id,string"`
	Latitude              string           `json:"latitude,omitempty"`
	Location              string           `json:"location,omitempty"`
	Longitude             string           `json:"longitude,omitempty"`
	ModifiedBy            int64            `json:"modifiedBy,string"`
	ModifiedTime          int32            `json:"modifiedTime,string"`
	Name                  string           `json:"name"`
	AppServerGroup        []AppServerGroup `json:"serverGroups,omitempty"`
	SiemAppConnectorGroup bool             `json:"siemAppConnectorGroup,omitempty"`
	UpgradeDay            string           `json:"upgradeDay,omitempty"`
	UpgradeTimeInSecs     string           `json:"upgradeTimeInSecs,omitempty"`
	VersionProfileID      int64            `json:"versionProfileId,string"`
}
type Connectors struct {
	// ApplicationStartTime     int64        `json:"applicationStartTime,string"`
	// AppConnectorGroupID      string       `json:"appConnectorGroupId,omitempty"`
	// AppConnectorGroupName    string       `json:"appConnectorGroupName,omitempty"`
	// ControlChannelStatus     string       `json:"controlChannelStatus,omitempty"`
	CreationTime int32 `json:"creationTime,string"`
	// CtrlBrokerName           string       `json:"ctrlBrokerName,omitempty"`
	// CurrentVersion           string       `json:"currentVersion,omitempty"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled,omitempty"`
	// ExpectedUpgradeTime      int64        `json:"expectedUpgradeTime,string"`
	// ExpectedVersion          string       `json:"expectedVersion,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	ID          int    `json:"id,string"`
	// IpAcl                    []string     `json:"ipAcl,omitempty"`
	IssuedCertID int64 `json:"issuedCertId,string,omitempty"`
	// LastBrokerConnectTime    int64        `json:"lastBrokerConnectTime,string,omitempty"`
	// LastBrokerDisconnectTime int64        `json:"lastBrokerDisconnectTime,string,omitempty"`
	// LastUpgradeTime          int64        `json:"lastUpgradeTime,string,omitempty"`
	// Latitude                 int          `json:"latitude,string,omitempty"`
	// Location                 string       `json:"location,omitempty"`
	// Longitude                int          `json:"longitude,string,omitempty"`
	ModifiedBy   int64  `json:"modifiedBy,string,omitempty"`
	ModifiedTime int32  `json:"modifiedTime,string,omitempty"`
	Name         string `json:"name,omitempty"`
	// Platform                 string       `json:"platform"`
	// PreviousVersion          string       `json:"previousVersion"`
	// PrivateIp                string       `json:"privateIp"`
	// PublicIp                 string       `json:"publicIp"`
	UpgradeAttempt int32 `json:"upgradeAttempt,string"`
	// UpgradeStatus            string       `json:"upgradeStatus"`
}
type AppServerGroup struct {
	ConfigSpace      string `json:"configSpace"`
	CreationTime     int32  `json:"creationTime,string"`
	Description      string `json:"description"`
	Enabled          bool   `json:"enabled"`
	ID               int    `json:"id,string"`
	DynamicDiscovery bool   `json:"dynamicDiscovery"`
	ModifiedBy       int64  `json:"modifiedBy,string"`
	ModifiedTime     int32  `json:"modifiedTime,string"`
	Name             string `json:"name"`
}

func (service *Service) Get(connectorGroupId string) (*AppConnectorGroupRequest, *http.Response, error) {
	v := new(AppConnectorGroupRequest)
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appConnectorGroupEndpoint, connectorGroupId)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]AppConnectorGroupRequest, *http.Response, error) {
	v := new([]AppConnectorGroupRequest)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+appConnectorGroupEndpoint, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
