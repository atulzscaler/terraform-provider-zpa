package segmentgroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig           = "/mgmtconfig/v1/admin/customers/"
	segmentGroupEndpoint = "/segmentGroup"
)

type SegmentGroupRequest struct {
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Enabled      bool           `json:"enabled"`
	Applications []Applications `json:"applications"`
	// CreationTime   int32          `json:"creationTime,string"`
	// ModifiedBy     int64          `json:"modifiedBy,string"`
	// ModifiedTime   int32          `json:"modifiedTime,string"`
	PolicyMigrated bool `json:"policyMigrated"`
}

type SegmentGroupResponse struct {
	Applications   []Applications `json:"applications,omitempty"`
	ConfigSpace    string         `json:"configSpace"`
	CreationTime   int32          `json:"creationTime,string"`
	Description    string         `json:"description"`
	Enabled        bool           `json:"enabled"`
	ID             int            `json:"id,string"`
	ModifiedBy     int64          `json:"modifiedBy,string"`
	ModifiedTime   int32          `json:"modifiedTime,string"`
	Name           string         `json:"name"`
	PolicyMigrated bool           `json:"policyMigrated"`
}
type AppServerGroup struct {
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
type Applications struct {
	BypassType           string           `json:"bypassType"`
	ConfigSpace          string           `json:"configSpace"`
	CreationTime         int32            `json:"creationTime,string"`
	DefaultIdleTimeout   int32            `json:"defaultIdleTimeout,string"`
	DefaultMaxAge        int32            `json:"defaultMaxAge,string"`
	Description          string           `json:"description"`
	DomainName           string           `json:"domainName"`
	DomainNames          []string         `json:"domainNames"`
	DoubleEncrypt        bool             `json:"doubleEncrypt"`
	Enabled              bool             `json:"enabled"`
	HealthCheckType      string           `json:"healthCheckType"`
	ID                   int64            `json:"id,string"`
	IPAnchored           bool             `json:"ipAnchored"`
	LogFeatures          []string         `json:"logFeatures"`
	ModifiedBy           int64            `json:"modifiedBy,string"`
	ModifiedTime         int32            `json:"modifiedTime,string"`
	Name                 []interface{}    `json:"name"`
	PassiveHealthEnabled bool             `json:"passiveHealthEnabled"`
	AppServerGroup       []AppServerGroup `json:"serverGroups"`
	TCPPortRanges        interface{}      `json:"tcpPortRanges"`
	TCPPortsIn           interface{}      `json:"tcpPortsIn"`
	TCPPortsOut          interface{}      `json:"tcpPortsOut"`
	UDPPortRanges        interface{}      `json:"udpPortRangesg"`
}

func (service *Service) Get(id string) (*SegmentGroupResponse, *http.Response, error) {
	v := new(SegmentGroupResponse)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(segmentGroup SegmentGroupRequest) (*SegmentGroupResponse, *http.Response, error) {
	v := new(SegmentGroupResponse)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, nil, segmentGroup, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, segmentGroupRequest SegmentGroupRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, id)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, segmentGroupRequest, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
