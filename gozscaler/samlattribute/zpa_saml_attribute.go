package samlattribute

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	samlAttributeEndpoint = "/samlAttribute"
)

// SamlAttribute ...
type SamlAttribute struct {
	CreationTime  int32  `json:"creationTime,string"`
	ID            int64  `json:"id,string"`
	IdpId         int64  `json:"idpId,string"`
	IdpName       string `json:"idpName"`
	ModifiedBy    int64  `json:"modifiedBy,string"`
	ModifiedTime  int32  `json:"modifiedTime,string"`
	Name          string `json:"name"`
	SamlName      string `json:"samlName"`
	UserAttribute bool   `json:"userAttribute"`
}

func (service *Service) Get(samlattributeId string) (*SamlAttribute, *http.Response, error) {
	v := new(SamlAttribute)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+samlAttributeEndpoint, samlattributeId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]SamlAttribute, *http.Response, error) {
	v := make([]SamlAttribute, 0)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+samlAttributeEndpoint, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
