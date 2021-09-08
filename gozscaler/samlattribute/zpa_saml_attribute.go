package samlattribute

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	samlAttributeEndpoint = "/samlAttribute"
)

type SamlAttribute struct {
	CreationTime  string `json:"creationTime,omitempty"`
	ID            string `json:"id,omitempty"`
	IdpId         string `json:"idpId,omitempty"`
	IdpName       string `json:"idpName,omitempty"`
	ModifiedBy    string `json:"modifiedBy,omitempty"`
	ModifiedTime  string `json:"modifiedTime,omitempty"`
	Name          string `json:"name,omitempty"`
	SamlName      string `json:"samlName,omitempty"`
	UserAttribute bool   `json:"userAttribute,omitempty"`
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

func (service *Service) GetByName(name string) (*SamlAttribute, *http.Response, error) {
	var v []SamlAttribute
	relativeURL := fmt.Sprintf(mgmtConfig + service.Client.Config.CustomerID + samlAttributeEndpoint)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, struct{ pagesize int }{
		pagesize: 500,
	}, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	for _, samlAttribute := range v {
		if strings.EqualFold(samlAttribute.Name, name) {
			return &samlAttribute, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("no saml attribute named '%s' was found", name)
}
