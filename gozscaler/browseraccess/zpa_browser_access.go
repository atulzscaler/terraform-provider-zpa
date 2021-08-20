package browseraccess

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	browserAccessEndpoint = "/application"
)

type BrowserAccessRequest struct {
	SegmentGroupId   string            `json:"segmentGroupId"`
	SegmentGroupName string            `json:"segmentGroupName"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Enabled          bool              `json:"enabled"`
	HealthReporting  string            `json:"healthReporting"`
	IpAnchored       bool              `json:"ipAnchored"`
	DoubleEncrypt    bool              `json:"doubleEncrypt"`
	BypassType       string            `json:"bypassType"`
	IsCnameEnabled   bool              `json:"isCnameEnabled"`
	DomainNames      []string          `json:"domainNames"`
	AppServerGroups  []AppServerGroups `json:"serverGroups"`
	ClientlessApps   []ClientlessApps  `json:"clientlessApps"`
	TcpPortRanges    []interface{}     `json:"tcpPortRanges"`
	UdpPortRanges    []interface{}     `json:"udpPortRanges"`
}

type BrowserAccessResponse struct {
	ID                   string            `json:"id"`
	SegmentGroupId       string            `json:"segmentGroupId"`
	SegmentGroupName     string            `json:"segmentGroupName"`
	BypassType           string            `json:"bypassType"`
	ConfigSpace          string            `json:"configSpace"`
	DomainNames          []string          `json:"domainNames"`
	Name                 string            `json:"name"`
	Description          string            `json:"description"`
	Enabled              bool              `json:"enabled"`
	PassiveHealthEnabled bool              `json:"passiveHealthEnabled"`
	DoubleEncrypt        bool              `json:"doubleEncrypt"`
	HealthCheckType      string            `json:"healthCheckType"`
	IsCnameEnabled       bool              `json:"isCnameEnabled"`
	IpAnchored           bool              `json:"ipAnchored"`
	HealthReporting      string            `json:"healthReporting"`
	CreationTime         int               `json:"creationTime,string"`
	ModifiedBy           string            `json:"modifiedBy"`
	ModifiedTime         int               `json:"modifiedTime,string"`
	TcpPortRanges        []interface{}     `json:"tcpPortRanges"`
	UdpPortRanges        []interface{}     `json:"udpPortRanges"`
	ClientlessApps       []ClientlessApps  `json:"clientlessApps"`
	AppServerGroups      []AppServerGroups `json:"serverGroups"`
}

type ClientlessApps struct {
	AllowOptions        bool   `json:"allowOptions"`
	AppId               int64  `json:"appId,string"`
	ApplicationPort     int32  `json:"applicationPort,string"`
	ApplicationProtocol string `json:"applicationProtocol"`
	CertificateId       int64  `json:"certificateId,string"`
	CertificateName     string `json:"certificateName"`
	Cname               string `json:"cname"`
	CreationTime        int32  `json:"creationTime,string"`
	Description         string `json:"description"`
	Domain              string `json:"domain"`
	Enabled             bool   `json:"enabled"`
	Hidden              bool   `json:"hidden"`
	ID                  int64  `json:"id,string"`
	LocalDomain         string `json:"localDomain"`
	ModifiedBy          int64  `json:"modifiedBy,string"`
	ModifiedTime        int32  `json:"modifiedTime,string"`
	Name                string `json:"name"`
	Path                string `json:"path"`
	TrustUntrustedCert  bool   `json:"trustUntrustedCert"`
}

type AppServerGroups struct {
	// ConfigSpace      string `json:"configSpace"`
	// CreationTime     int32  `json:"creationTime,string"`
	// Description      string `json:"description"`
	// Enabled          bool   `json:"enabled"`
	ID int64 `json:"id,string"`
	// DynamicDiscovery bool   `json:"dynamicDiscovery"`
	// ModifiedBy       int64  `json:"modifiedBy,string"`
	// ModifiedTime     int32  `json:"modifiedTime,string"`
	// Name             string `json:"name"`
}

func (service *Service) Get(id string) (*BrowserAccessResponse, *http.Response, error) {
	v := new(BrowserAccessResponse)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]BrowserAccessResponse, *http.Response, error) {
	v := new([]BrowserAccessResponse)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(browserAccess BrowserAccessRequest) (*BrowserAccessResponse, *http.Response, error) {
	v := new(BrowserAccessResponse)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, nil, browserAccess, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, browserAccess BrowserAccessRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, id)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, browserAccess, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
