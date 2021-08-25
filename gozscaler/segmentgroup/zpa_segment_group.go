package segmentgroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig           = "/mgmtconfig/v1/admin/customers/"
	segmentGroupEndpoint = "/segmentGroup"
)

type SegmentGroup struct {
	Applications        []Application `json:"applications,omitempty"`
	ConfigSpace         string        `json:"configSpace,omitempty"`
	CreationTime        int32         `json:"creationTime,string,omitempty"`
	Description         string        `json:"description,omitempty"`
	Enabled             bool          `json:"enabled,omitempty"`
	ID                  int64         `json:"id,string,omitempty"`
	ModifiedBy          int64         `json:"modifiedBy,string,omitempty"`
	ModifiedTime        int32         `json:"modifiedTime,string,omitempty"`
	Name                string        `json:"name"`
	PolicyMigrated      bool          `json:"policyMigrated,omitempty"`
	TcpKeepAliveEnabled int           `json:"tcpKeepAliveEnabled,string,omitempty"`
}

/*
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
*/
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
type Application struct {
	BypassType           string           `json:"bypassType,omitempty"`
	ConfigSpace          string           `json:"configSpace,omitempty"`
	CreationTime         int32            `json:"creationTime,string,omitempty"`
	DefaultIdleTimeout   int32            `json:"defaultIdleTimeout,string,omitempty"`
	DefaultMaxAge        int32            `json:"defaultMaxAge,string,omitempty"`
	Description          string           `json:"description,omitempty"`
	DomainName           string           `json:"domainName,omitempty"`
	DomainNames          []string         `json:"domainNames,omitempty"`
	DoubleEncrypt        bool             `json:"doubleEncrypt,omitempty"`
	Enabled              bool             `json:"enabled,omitempty"`
	HealthCheckType      string           `json:"healthCheckType,omitempty"`
	ID                   int              `json:"id,string,omitempty"`
	IPAnchored           bool             `json:"ipAnchored,omitempty"`
	LogFeatures          []string         `json:"logFeatures,omitempty"`
	ModifiedBy           int64            `json:"modifiedBy,string,omitempty"`
	ModifiedTime         int32            `json:"modifiedTime,string,omitempty"`
	Name                 string           `json:"name"`
	PassiveHealthEnabled bool             `json:"passiveHealthEnabled,omitempty"`
	ServerGroup          []AppServerGroup `json:"serverGroups,omitempty"`
	TCPPortRanges        interface{}      `json:"tcpPortRanges,omitempty"`
	TCPPortsIn           interface{}      `json:"tcpPortsIn,omitempty"`
	TCPPortsOut          interface{}      `json:"tcpPortsOut,omitempty"`
	UDPPortRanges        interface{}      `json:"udpPortRangesg,omitempty"`
}

func (service *Service) Get(segmentGroupId int64) (*SegmentGroup, *http.Response, error) {
	v := new(SegmentGroup)
	relativeURL := fmt.Sprintf("%v/%v", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, segmentGroupId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(segmentGroup SegmentGroup) (*SegmentGroup, *http.Response, error) {
	v := new(SegmentGroup)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, nil, segmentGroup, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(segmentGroupId int64, segmentGroupRequest SegmentGroup) (*http.Response, error) {
	path := fmt.Sprintf("%v/%v", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, segmentGroupId)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, segmentGroupRequest, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(segmentGroupId int64) (*http.Response, error) {
	path := fmt.Sprintf("%v/%v", mgmtConfig+service.Client.Config.CustomerID+segmentGroupEndpoint, segmentGroupId)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
