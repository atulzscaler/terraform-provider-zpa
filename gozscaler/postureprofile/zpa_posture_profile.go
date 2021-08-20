package postureprofile

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig             = "/mgmtconfig/v1/admin/customers/"
	postureProfileEndpoint = "/posture"
)

// PostureProfile ...
type PostureProfile struct {
	CreationTime      int32  `json:"creationTime,string"`
	Domain            string `json:"domain,omitempty"`
	ID                int64  `json:"id,string"`
	ModifiedBy        int64  `json:"modifiedBy,string"`
	ModifiedTime      int32  `json:"modifiedTime,string"`
	Name              string `json:"name,omitempty"`
	PostureudId       string `json:"postureUdid"`
	ZscalerCloud      string `json:"zscalerCloud"`
	ZscalerCustomerId int64  `json:"zscalerCustomerId,string"`
}

func (service *Service) Get(id string) (*PostureProfile, *http.Response, error) {
	v := new(PostureProfile)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+postureProfileEndpoint, id)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]PostureProfile, *http.Response, error) {
	v := make([]PostureProfile, 0)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+postureProfileEndpoint, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
