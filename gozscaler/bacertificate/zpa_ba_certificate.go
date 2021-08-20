package bacertificate

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	baCertificateEndpoint = "/clientlessCertificate"
	baCertificateAll      = "/issued"
)

type BaCertificate struct {
	CName               string   `json:"cName"`
	CertChain           string   `json:"certChain"`
	Certificate         string   `json:"certificate"`
	CreationTime        int32    `json:"creationTime,string"`
	Description         string   `json:"description"`
	ID                  int64    `json:"id,string"`
	IssuedBy            string   `json:"issuedBy"`
	IssuedTo            string   `json:"issuedTo"`
	ModifiedBy          int64    `json:"modifiedBy,string"`
	ModifiedTime        int32    `json:"modifiedTime,string"`
	Name                string   `json:"name"`
	PublicKey           string   `json:"publicKey"`
	San                 []string `json:"san"`
	SerialNo            string   `json:"serialNo"`
	Status              string   `json:"status"`
	ValidFromInEpochSec int64    `json:"validFromInEpochSec,string"`
	ValidToInEpochSec   int64    `json:"validToInEpochSec,string"`
}

func (service *Service) Get(baCertificateId string) (*BaCertificate, *http.Response, error) {
	v := new(BaCertificate)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+baCertificateEndpoint, baCertificateId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetAll() ([]BaCertificate, *http.Response, error) {
	v := make([]BaCertificate, 0)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+baCertificateEndpoint+baCertificateAll, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
