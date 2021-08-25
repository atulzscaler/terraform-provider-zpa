package applicationsegment

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig         = "/mgmtconfig/v1/admin/customers/"
	appSegmentEndpoint = "/application"
)

type ApplicationSegmentRequest struct {
	//ID              string            `json:"id,string"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Enabled          bool              `json:"enabled"`
	HealthReporting  string            `json:"healthReporting"`
	IpAnchored       bool              `json:"ipAnchored"`
	DoubleEncrypt    bool              `json:"doubleEncrypt"`
	BypassType       string            `json:"bypassType"`
	IsCnameEnabled   bool              `json:"isCnameEnabled"`
	DomainNames      []string          `json:"domainNames"`
	AppServerGroups  []AppServerGroups `json:"serverGroups,omitempty"`
	TcpPortRanges    []interface{}     `json:"tcpPortRanges"` // Need to fix for conversion - json: cannot unmarshal string into Go struct field ApplicationSegmentResponse.tcpPortRanges of type int32
	UdpPortRanges    []interface{}     `json:"udpPortRanges"` // Need to fix for conversion - json: cannot unmarshal string into Go struct field ApplicationSegmentResponse.tcpPortRanges of type int32
	ClientlessApps   []ClientlessApps  `json:"clientlessApps"`
	SegmentGroupId   int               `json:"segmentGroupId,string"`
	SegmentGroupName string            `json:"segmentGroupName"`
}

type ApplicationSegmentResponse struct {
	ID                   string            `json:"id"`
	DomainNames          []string          `json:"domainNames"`
	Name                 string            `json:"name"`
	Description          string            `json:"description"`
	Enabled              bool              `json:"enabled"`
	PassiveHealthEnabled bool              `json:"passiveHealthEnabled"`
	DoubleEncrypt        bool              `json:"doubleEncrypt"`
	ConfigSpace          string            `json:"configSpace"`
	Applications         string            `json:"applications"`
	BypassType           string            `json:"bypassType"`
	HealthCheckType      string            `json:"healthCheckType"`
	IsCnameEnabled       bool              `json:"isCnameEnabled"`
	IpAnchored           bool              `json:"ipAnchored"`
	HealthReporting      string            `json:"healthReporting"`
	IcmpAccessType       string            `json:"icmpAccessType"`
	SegmentGroupId       int               `json:"segmentGroupId,string"`
	SegmentGroupName     string            `json:"segmentGroupName"`
	CreationTime         int               `json:"creationTime,string"`
	ModifiedBy           string            `json:"modifiedBy"`
	ModifiedTime         int               `json:"modifiedTime,string"`
	TcpPortRanges        []interface{}     `json:"tcpPortRanges"`
	UdpPortRanges        []interface{}     `json:"udpPortRanges"`
	ClientlessApps       []ClientlessApps  `json:"clientlessApps"`
	AppServerGroups      []AppServerGroups `json:"serverGroups,omitempty"`
	DefaultIdleTimeout   int32             `json:"defaultIdleTimeout,string"`
	DefaultMaxAge        int32             `json:"defaultMaxAge,string"`
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
	Name                string `json:"name"`
	Path                string `json:"path"`
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

func (service *Service) Get(applicationId string) (*ApplicationSegmentResponse, *http.Response, error) {
	v := new(ApplicationSegmentResponse)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appSegmentEndpoint, applicationId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Create(appSegment ApplicationSegmentRequest) (*ApplicationSegmentResponse, *http.Response, error) {
	v := new(ApplicationSegmentResponse)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+appSegmentEndpoint, nil, appSegment, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(applicationId string, appSegmentRequest ApplicationSegmentRequest) (*http.Response, error) {
	//v := new(ApplicationSegmentResponse)
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
