package scimgroup

import (
	"fmt"
	"net/http"
)

const (
	userConfig        = "/userconfig/v1/customers/"
	scimGroupEndpoint = "/scimgroup/idpId"
)

type ScimGroup struct {
	CreationTime int32  `json:"creationTime,string"`
	ID           int64  `json:"id,string"`
	IdpGroupId   string `json:"idpGroupId"`
	IdpId        int64  `json:"idpId,string"`
	ModifiedTime int64  `json:"modifiedTime,string"`
	Name         string `json:"name"`
}

func (service *Service) Get(scimGroupId string) (*ScimGroup, *http.Response, error) {
	v := new(ScimGroup)
	relativeURL := fmt.Sprintf("%s/%s", userConfig+service.Client.Config.CustomerID+scimGroupEndpoint, scimGroupId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]ScimGroup, *http.Response, error) {
	v := make([]ScimGroup, 0)
	resp, err := service.Client.NewRequestDo("GET", userConfig+service.Client.Config.CustomerID+scimGroupEndpoint, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
