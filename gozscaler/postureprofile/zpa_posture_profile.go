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
	CreationTime      string `json:"creationTime,omitempty"`
	Domain            string `json:"domain,omitempty"`
	ID                string `json:"id,omitempty"`
	ModifiedBy        string `json:"modifiedBy,omitempty"`
	ModifiedTime      string `json:"modifiedTime,omitempty"`
	Name              string `json:"name,omitempty"`
	PostureudId       string `json:"postureUdid,omitempty"`
	ZscalerCloud      string `json:"zscalerCloud,omitempty"`
	ZscalerCustomerId string `json:"zscalerCustomerId,omitempty"`
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
