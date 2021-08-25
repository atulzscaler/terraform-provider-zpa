package applicationsegment

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig         = "/mgmtconfig/v1/admin/customers/"
	appSegmentEndpoint = "/application"
)

type ApplicationSegmentResource struct {
	ID                   string            `json:"id,omitempty"`
	DomainNames          []string          `json:"domainNames,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Description          string            `json:"description,omitempty"`
	Enabled              bool              `json:"enabled,omitempty"`
	PassiveHealthEnabled bool              `json:"passiveHealthEnabled,omitempty"`
	DoubleEncrypt        bool              `json:"doubleEncrypt,omitempty"`
	ConfigSpace          string            `json:"configSpace,omitempty"`
	Applications         string            `json:"applications,omitempty"`
	BypassType           string            `json:"bypassType,omitempty"`
	HealthCheckType      string            `json:"healthCheckType,omitempty"`
	IsCnameEnabled       bool              `json:"isCnameEnabled,omitempty"`
	IpAnchored           bool              `json:"ipAnchored,omitempty"`
	HealthReporting      string            `json:"healthReporting,omitempty"`
	IcmpAccessType       string            `json:"icmpAccessType,omitempty"`
	SegmentGroupId       int               `json:"segmentGroupId,string"`
	SegmentGroupName     string            `json:"segmentGroupName,omitempty"`
	CreationTime         int               `json:"creationTime,string"`
	ModifiedBy           string            `json:"modifiedBy,omitempty"`
	ModifiedTime         int               `json:"modifiedTime,string"`
	TcpPortRanges        []interface{}     `json:"tcpPortRanges,omitempty"`
	UdpPortRanges        []interface{}     `json:"udpPortRanges,omitempty"`
	ClientlessApps       []ClientlessApps  `json:"clientlessApps,omitempty"`
	ServerGroups         []AppServerGroups `json:"serverGroups,omitempty"`
	DefaultIdleTimeout   int32             `json:"defaultIdleTimeout,string,omitempty"`
	DefaultMaxAge        int32             `json:"defaultMaxAge,string,omitempty"`
}
type ClientlessApps struct {
	AllowOptions        bool   `json:"allowOptions"`
	AppId               int    `json:"appId,string"`
	ApplicationPort     int    `json:"applicationPort,string"`
	ApplicationProtocol string `json:"applicationProtocol"`
	CertificateId       int    `json:"certificateId,string"`
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
	Name                string `json:"name,omitempty"`
	Path                string `json:"path,omitempty"`
	Portal              bool   `json:"portal,omitempty"`
	TrustUntrustedCert  bool   `json:"trustUntrustedCert"`
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

func (service *Service) Get(applicationId string) (*ApplicationSegmentResource, *http.Response, error) {
	v := new(ApplicationSegmentResource)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appSegmentEndpoint, applicationId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(appSegment ApplicationSegmentResource) (*ApplicationSegmentResource, *http.Response, error) {
	v := new(ApplicationSegmentResource)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+appSegmentEndpoint, nil, appSegment, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(applicationId string, appSegmentRequest ApplicationSegmentResource) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appSegmentEndpoint, applicationId)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, appSegmentRequest, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (service *Service) Delete(applicationId string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appSegmentEndpoint, applicationId)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
