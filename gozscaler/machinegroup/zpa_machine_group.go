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
	CreationTime int32      `json:"creationTime,string"`
	Description  string     `json:"description"`
	Enabled      bool       `json:"enabled"`
	ID           int64      `json:"id,string"`
	Machines     []Machines `json:"machines"`
	ModifiedBy   int64      `json:"modifiedBy,string"`
	ModifiedTime int32      `json:"modifiedTime,string"`
	Name         string     `json:"name"`
}

type Machines struct {
	CreationTime     int32       `json:"creationTime,string"`
	Description      string      `json:"description"`
	Fingerprint      string      `json:"fingerprint"`
	ID               int64       `json:"id,string"`
	IssuedCertID     int64       `json:"issuedCertId,string"`
	MachineGroupID   string      `json:"machineGroupId"`
	MachineGroupName string      `json:"machineGroupName"`
	MachineTokenID   int64       `json:"machineTokenId,string"`
	ModifiedBy       int64       `json:"modifiedBy,string"`
	ModifiedTime     int32       `json:"modifiedTime,string"`
	Name             string      `json:"name"`
	SigningCert      SigningCert `json:"signingCert"`
}

type SigningCert struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
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
