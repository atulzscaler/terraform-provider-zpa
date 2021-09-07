package bacertificate

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	baCertificateEndpoint = "/clientlessCertificate"
)

type BaCertificate struct {
	CName               string   `json:"cName,omitempty"`
	CertChain           string   `json:"certChain,omitempty"`
	Certificate         string   `json:"certificate,omitempty"`
	CreationTime        string   `json:"creationTime,omitempty"`
	Description         string   `json:"description,omitempty"`
	ID                  string   `json:"id"`
	IssuedBy            string   `json:"issuedBy,omitempty"`
	IssuedTo            string   `json:"issuedTo,omitempty"`
	ModifiedBy          string   `json:"modifiedBy,omitempty"`
	ModifiedTime        string   `json:"modifiedTime,omitempty"`
	Name                string   `json:"name"`
	PublicKey           string   `json:"publicKey,omitempty"`
	San                 []string `json:"san,omitempty"`
	SerialNo            string   `json:"serialNo,omitempty"`
	Status              string   `json:"status,omitempty"`
	ValidFromInEpochSec string   `json:"validFromInEpochSec,omitempty"`
	ValidToInEpochSec   string   `json:"validToInEpochSec,omitempty"`
}

func (service *Service) Get(baCertificateId string) (*BaCertificate, *http.Response, error) {
	v := new(BaCertificate)
	relativeURL := fmt.Sprintf("%v/%v", mgmtConfig+service.Client.Config.CustomerID+baCertificateEndpoint, baCertificateId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetByName(baCertificateName string) (*BaCertificate, *http.Response, error) {
	var v struct {
		List []BaCertificate `json:"list"`
	}

	relativeURL := mgmtConfig + service.Client.Config.CustomerID + baCertificateEndpoint
	resp, err := service.Client.NewRequestDo("GET", relativeURL, struct{ pagesize int }{
		pagesize: 500,
	}, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	for _, app := range v.List {
		if strings.EqualFold(app.Name, baCertificateName) {
			return &app, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("no application named '%s' was found", baCertificateName)
}
