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
	CreationTime int64  `json:"creationTime,omitempty"`
	ID           int64  `json:"id,omitempty"`
	IdpGroupId   string `json:"idpGroupId,omitempty"`
	IdpId        int64  `json:"idpId,omitempty"`
	ModifiedTime int64  `json:"modifiedTime,omitempty"`
	Name         string `json:"name,omitempty"`
}

func (service *Service) Get(scimGroupId int64) (*ScimGroup, *http.Response, error) {
	v := new(ScimGroup)
	relativeURL := fmt.Sprintf("%s/%d", userConfig+service.Client.Config.CustomerID+scimGroupEndpoint, scimGroupId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
