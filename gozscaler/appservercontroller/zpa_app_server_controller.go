package appservercontroller

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                  = "/mgmtconfig/v1/admin/customers/"
	appServerControllerEndpoint = "/server"
)

type ApplicationServer struct {
	Name              string   `json:"name"`
	Description       string   `json:"description,omitempty"`
	Enabled           bool     `json:"enabled,omitempty"`
	Address           string   `json:"address,omitempty"`
	ID                int64    `json:"id,string,omitempty"`
	CreationTime      int32    `json:"creationTime,string,omitempty"`
	ModifiedBy        int64    `json:"modifiedBy,string,omitempty"`
	ModifiedTime      int32    `json:"modifiedTime,string,omitempty"`
	AppServerGroupIds []string `json:"appServerGroupIds,omitempty"`
	ConfigSpace       string   `json:"configSpace,omitempty"`
}

func (service *Service) Get(appServerID int64) (*ApplicationServer, *http.Response, error) {
	v := new(ApplicationServer)
	relativeURL := fmt.Sprintf("%v/%v", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, appServerID)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(server ApplicationServer) (*ApplicationServer, *http.Response, error) {
	v := new(ApplicationServer)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, nil, server, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(appServerID string, appServerRequest ApplicationServer) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, appServerID)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, appServerRequest, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (service *Service) Delete(appServerID string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, appServerID)
	resp, err := service.Client.NewRequestDo("DELETE", path, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, err
}
