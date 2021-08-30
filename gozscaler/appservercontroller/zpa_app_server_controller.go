package appservercontroller

import (
	"fmt"
	"net/http"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/common"
)

const (
	mgmtConfig                  = "/mgmtconfig/v1/admin/customers/"
	appServerControllerEndpoint = "/server"
)

type ApplicationServersResponse struct {
	TotalPages int32                      `json:"totalPages,string"`
	List       []common.ApplicationServer `json:"list"`
}

func (service *Service) Get(id int64) (*common.ApplicationServer, *http.Response, error) {
	v := new(common.ApplicationServer)
	relativeURL := fmt.Sprintf("%s/%d", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) GetAll() (*ApplicationServersResponse, *http.Response, error) {
	v := new(ApplicationServersResponse)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Create(server common.ApplicationServer) (*common.ApplicationServer, *http.Response, error) {
	v := new(common.ApplicationServer)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, nil, server, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

func (service *Service) Update(id string, appServer common.ApplicationServer) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+appServerControllerEndpoint, id)
	resp, err := service.Client.NewRequestDo("PUT", path, nil, appServer, nil)
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
