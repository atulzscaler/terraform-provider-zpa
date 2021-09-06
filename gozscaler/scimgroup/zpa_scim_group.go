package scimgroup

import (
	"fmt"
	"net/http"
)

const (
	userConfig        = "/userconfig/v1/customers/"
	scimGroupEndpoint = "/scimgroup"
)

type ScimGroup struct {
	CreationTime string `json:"creationTime"`
	ID           string `json:"id"`
	IdpGroupId   string `json:"idpGroupId"`
	IdpId        string `json:"idpId"`
	ModifiedTime string `json:"modifiedTime"`
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
