package appconnectorgroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                = "/mgmtconfig/v1/admin/customers/"
	appConnectorGroupEndpoint = "/appConnectorGroup"
)

type AppConnectorGroup struct {
	CityCountry           string           `json:"cityCountry,omitempty"`
	CountryCode           string           `json:"countryCode,omitempty"`
	CreationTime          int32            `json:"creationTime,string"`
	Description           string           `json:"description,omitempty"`
	DNSQueryType          string           `json:"dnsQueryType,omitempty"`
	Enabled               bool             `json:"enabled,omitempty"`
	GeoLocationID         int64            `json:"geoLocationId,string,omitempty"`
	ID                    int64            `json:"id,string,omitempty"`
	Latitude              string           `json:"latitude,omitempty"`
	Location              string           `json:"location,omitempty"`
	Longitude             string           `json:"longitude,omitempty"`
	ModifiedBy            int64            `json:"modifiedBy,string"`
	ModifiedTime          int32            `json:"modifiedTime,string"`
	Name                  string           `json:"name"`
	SiemAppConnectorGroup bool             `json:"siemAppConnectorGroup,omitempty"`
	UpgradeDay            string           `json:"upgradeDay,omitempty"`
	UpgradeTimeInSecs     string           `json:"upgradeTimeInSecs,omitempty"`
	VersionProfileID      int64            `json:"versionProfileId,string,omitempty"`
	AppServerGroup        []AppServerGroup `json:"serverGroups,omitempty"`
	Connectors            []Connector      `json:"connectors,omitempty"`
}
type Connector struct {
	ApplicationStartTime     int64             `json:"applicationStartTime,string,omitempty"`
	AppConnectorGroupID      string            `json:"appConnectorGroupId,omitempty"`
	AppConnectorGroupName    string            `json:"appConnectorGroupName,omitempty"`
	ControlChannelStatus     string            `json:"controlChannelStatus,omitempty"`
	CreationTime             int32             `json:"creationTime,string,omitempty"`
	CtrlBrokerName           string            `json:"ctrlBrokerName,omitempty"`
	CurrentVersion           string            `json:"currentVersion,omitempty"`
	Description              string            `json:"description,omitempty"`
	Enabled                  bool              `json:"enabled,omitempty"`
	ExpectedUpgradeTime      int64             `json:"expectedUpgradeTime,string"`
	ExpectedVersion          string            `json:"expectedVersion,omitempty"`
	Fingerprint              string            `json:"fingerprint,omitempty"`
	ID                       int64             `json:"id,string,omitempty"`
	IpAcl                    []string          `json:"ipAcl,omitempty"`
	IssuedCertID             int64             `json:"issuedCertId,string,omitempty"`
	LastBrokerConnectTime    int64             `json:"lastBrokerConnectTime,string,omitempty"`
	LastBrokerDisconnectTime int64             `json:"lastBrokerDisconnectTime,string,omitempty"`
	LastUpgradeTime          int64             `json:"lastUpgradeTime,omitempty"`
	Latitude                 float64           `json:"latitude,omitempty"`
	Location                 string            `json:"location,omitempty"`
	Longitude                float64           `json:"longitude,omitempty"`
	ModifiedBy               int64             `json:"modifiedBy,string,omitempty"`
	ModifiedTime             int32             `json:"modifiedTime,string,omitempty"`
	Name                     string            `json:"name"`
	Platform                 string            `json:"platform,omitempty"`
	PreviousVersion          string            `json:"previousVersion,omitempty"`
	PrivateIp                string            `json:"privateIp,omitempty"`
	SigningCert              map[string]string `json:"signingCert,omitempty"`
	PublicIp                 string            `json:"publicIp,omitempty"`
	UpgradeAttempt           int32             `json:"upgradeAttempt,string,omitempty"`
	UpgradeStatus            string            `json:"upgradeStatus,omitempty"`
}
type AppServerGroup struct {
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

func (service *Service) Get(appConnectorGroupId string) (*AppConnectorGroup, *http.Response, error) {
	v := new(AppConnectorGroup)
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appConnectorGroupEndpoint, appConnectorGroupId)
	resp, err := service.Client.NewRequestDo("GET", path, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]AppConnectorGroup, *http.Response, error) {
	v := new([]AppConnectorGroup)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+appConnectorGroupEndpoint, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
