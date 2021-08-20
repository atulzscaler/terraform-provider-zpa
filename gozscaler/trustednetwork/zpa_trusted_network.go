package trustednetwork

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig             = "/mgmtconfig/v1/admin/customers/"
	trustedNetworkEndpoint = "/network"
)

// TrustedNetwork ...
type TrustedNetwork struct {
	CreationTime int32  `json:"creationTime,string"`
	Domain       string `json:"domain"`
	ID           string `json:"id"`
	ModifiedBy   int64  `json:"modifiedBy,string"`
	ModifiedTime int32  `json:"modifiedTime,string"`
	Name         string `json:"name"`
	NetworkId    string `json:"networkId"`
	ZscalerCloud string `json:"zscalerCloud"`
}

func (service *Service) Get(networkId string) (*TrustedNetwork, *http.Response, error) {
	v := new(TrustedNetwork)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+trustedNetworkEndpoint, networkId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]TrustedNetwork, *http.Response, error) {
	v := make([]TrustedNetwork, 0)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+trustedNetworkEndpoint, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
