package browseraccess

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	browserAccessEndpoint = "/application"
)

type BrowserAccess struct {
	ID                   string            `json:"id,omitempty"`
	SegmentGroupId       string            `json:"segmentGroupId,omitempty"`
	SegmentGroupName     string            `json:"segmentGroupName,omitempty"`
	BypassType           string            `json:"bypassType,omitempty"`
	ConfigSpace          string            `json:"configSpace,omitempty"`
	DomainNames          []string          `json:"domainNames,omitempty"`
	Name                 string            `json:"name"`
	Description          string            `json:"description,omitempty"`
	Enabled              bool              `json:"enabled,omitempty"`
	PassiveHealthEnabled bool              `json:"passiveHealthEnabled,omitempty"`
	DoubleEncrypt        bool              `json:"doubleEncrypt,omitempty"`
	HealthCheckType      string            `json:"healthCheckType,omitempty"`
	IsCnameEnabled       bool              `json:"isCnameEnabled,omitempty"`
	IpAnchored           bool              `json:"ipAnchored,omitempty"`
	HealthReporting      string            `json:"healthReporting,omitempty"`
	CreationTime         string            `json:"creationTime,string,omitempty"`
	ModifiedBy           string            `json:"modifiedBy,omitempty"`
	ModifiedTime         string            `json:"modifiedTime,string,omitempty"`
	TcpPortRanges        []interface{}     `json:"tcpPortRanges,omitempty"`
	UdpPortRanges        []interface{}     `json:"udpPortRanges,omitempty"`
	ClientlessApps       []ClientlessApps  `json:"clientlessApps"`
	AppServerGroups      []AppServerGroups `json:"serverGroups"`
}

type ClientlessApps struct {
	AllowOptions        bool   `json:"allowOptions,omitempty"`
	AppId               string `json:"appId,omitempty"`
	ApplicationPort     string `json:"applicationPort,omitempty"`
	ApplicationProtocol string `json:"applicationProtocol,omitempty"`
	CertificateId       string `json:"certificateId,omitempty"`
	CertificateName     string `json:"certificateName,omitempty"`
	Cname               string `json:"cname,omitempty"`
	CreationTime        string `json:"creationTime,omitempty"`
	Description         string `json:"description,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Enabled             bool   `json:"enabled,omitempty"`
	Hidden              bool   `json:"hidden,omitempty"`
	ID                  string `json:"id,omitempty"`
	LocalDomain         string `json:"localDomain,omitempty"`
	ModifiedBy          string `json:"modifiedBy,omitempty"`
	ModifiedTime        string `json:"modifiedTime,omitempty"`
	Name                string `json:"name"`
	Path                string `json:"path,omitempty"`
	TrustUntrustedCert  bool   `json:"trustUntrustedCert"`
}

type AppServerGroups struct {
	ID string `json:"id,omitempty"`
}

func (service *Service) Get(id string) (*BrowserAccess, *http.Response, error) {
	v := new(BrowserAccess)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(browserAccess BrowserAccess) (*BrowserAccess, *http.Response, error) {
	v := new(BrowserAccess)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+browserAccessEndpoint, nil, browserAccess, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, browserAccess BrowserAccess) (*http.Response, error) {
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
