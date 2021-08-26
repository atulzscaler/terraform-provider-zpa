package machinegroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                  = "/mgmtconfig/v1/admin/customers/"
	machineGroupEndpoint string = "/machineGroup"
)

// MachineGroupsResponse Response structure from API endpoint
type MachineGroupsResponse struct {
	TotalPages int32          `json:"totalPages,string"`
	List       []MachineGroup `json:"list"`
}

type MachineGroup struct {
	CreationTime int32      `json:"creationTime,string,omitempty"`
	Description  string     `json:"description,omitempty"`
	Enabled      bool       `json:"enabled,omitempty"`
	ID           int64      `json:"id,string,omitempty"`
	Machines     []Machines `json:"machines,omitempty"`
	ModifiedBy   int64      `json:"modifiedBy,string,omitempty"`
	ModifiedTime int32      `json:"modifiedTime,string,omitempty"`
	Name         string     `json:"name"`
}

type Machines struct {
	CreationTime     int32                  `json:"creationTime,string,omitempty"`
	Description      string                 `json:"description,omitempty"`
	Fingerprint      string                 `json:"fingerprint,omitempty"`
	ID               int64                  `json:"id,string,omitempty"`
	IssuedCertID     int64                  `json:"issuedCertId,string,omitempty"`
	MachineGroupID   string                 `json:"machineGroupId,omitempty"`
	MachineGroupName string                 `json:"machineGroupName,omitempty"`
	MachineTokenID   int64                  `json:"machineTokenId,string,omitempty"`
	ModifiedBy       int64                  `json:"modifiedBy,string,omitempty"`
	ModifiedTime     int32                  `json:"modifiedTime,string,omitempty"`
	Name             string                 `json:"name"`
	SigningCert      map[string]interface{} `json:"signingCert,omitempty"`
	//SigningCert      map[string]string `json:"signingCert,omitempty"`
}

// Get ...
func (service *Service) Get(machineGroupId string) (*MachineGroup, *http.Response, error) {
	v := new(MachineGroup)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+machineGroupEndpoint, machineGroupId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() (*MachineGroupsResponse, *http.Response, error) {
	v := new(MachineGroupsResponse)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+machineGroupEndpoint, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
