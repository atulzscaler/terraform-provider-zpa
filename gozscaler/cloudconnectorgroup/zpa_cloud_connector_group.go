package cloudconnectorgroup

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                  = "/mgmtconfig/v1/admin/customers/"
	cloudConnectorGroupEndpoint = "/cloudConnectorGroup"
)

type CloudConnectorGroup struct {
	CreationTime    int32             `json:"creationTime,string"`
	Description     string            `json:"description"`
	CloudConnectors []CloudConnectors `json:"cloudConnectors"`
	Enabled         bool              `json:"enabled"`
	GeolocationId   int64             `json:"geoLocationId,string"`
	ID              int               `json:"id,string"`
	ModifiedBy      string            `json:"modifiedBy,string"`
	ModifiedTime    int32             `json:"modifiedTime,string"`
	Name            string            `json:"name"`
	ZiaCloud        string            `json:"ziaCloud"`
	ZiaOrgid        int64             `json:"ziaOrgId,string"`
}
type CloudConnectors struct {
	CreationTime int32    `json:"creationTime,string"`
	Description  string   `json:"description"`
	Enabled      bool     `json:"enabled"`
	Fingerprint  string   `json:"fingerprint"`
	ID           int      `json:"id,string"`
	IpAcl        []string `json:"ipAcl"`
	IssuedCertId int64    `json:"issuedCertId,string"`
	ModifiedBy   int64    `json:"modifiedBy,string"`
	ModifiedTime int32    `json:"modifiedTime,string"`
	Name         string   `json:"name"`
}

func (service *Service) Get() (*CloudConnectorGroup, *http.Response, error) {
	v := new(CloudConnectorGroup)
	relativeURL := fmt.Sprintf(mgmtConfig + service.Client.Config.CustomerID + cloudConnectorGroupEndpoint)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
