package appservercontroller

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                  = "/mgmtconfig/v1/admin/customers/"
	appServerControllerEndpoint = "/server"
)

type ApplicationServerRequest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Enabled           bool     `json:"enabled"`
	Address           string   `json:"address"`
	AppServerGroupIds []string `json:"appServerGroupIds,omitempty"`
}
type ApplicationServerResponse struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Enabled           bool     `json:"enabled"`
	Address           string   `json:"address"`
	ID                int64    `json:"id,string"`
	CreationTime      int32    `json:"creationTime,string"`
	ModifiedBy        int64    `json:"modifiedBy,string"`
	ModifiedTime      int32    `json:"modifiedTime,string"`
	AppServerGroupIds []string `json:"appServerGroupIds,omitempty"`
	ConfigSpace       string   `json:"configSpace,omitempty"`
}

func (service *Service) Get(id int64) (*ApplicationServerResponse, *http.Response, error) {
	v := new(ApplicationServerResponse)
	relativeURL := fmt.Sprintf("%s/%d", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*[]ApplicationServerResponse, *http.Response, error) {
	v := new([]ApplicationServerResponse)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(server ApplicationServerRequest) (*ApplicationServerResponse, *http.Response, error) {
	v := new(ApplicationServerResponse)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, nil, server, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, appServerRequest ApplicationServerRequest) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, id)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, appServerRequest, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, id)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
