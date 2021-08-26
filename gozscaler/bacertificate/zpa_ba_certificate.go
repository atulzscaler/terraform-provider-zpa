package bacertificate

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig            = "/mgmtconfig/v1/admin/customers/"
	baCertificateEndpoint = "/clientlessCertificate"
	//baCertificateAll      = "/issued"
)

type BaCertificate struct {
	CName               string   `json:"cName,omitempty"`
	CertChain           string   `json:"certChain,omitempty"`
	Certificate         string   `json:"certificate,omitempty"`
	CreationTime        int32    `json:"creationTime,string,omitempty"`
	Description         string   `json:"description,omitempty"`
	ID                  int64    `json:"id,string,omitempty"`
	IssuedBy            string   `json:"issuedBy,omitempty"`
	IssuedTo            string   `json:"issuedTo,omitempty"`
	ModifiedBy          int64    `json:"modifiedBy,string,omitempty"`
	ModifiedTime        int32    `json:"modifiedTime,string,omitempty"`
	Name                string   `json:"name"`
	PublicKey           string   `json:"publicKey,omitempty"`
	San                 []string `json:"san,omitempty"`
	SerialNo            string   `json:"serialNo,omitempty"`
	Status              string   `json:"status,omitempty"`
	ValidFromInEpochSec int64    `json:"validFromInEpochSec,string,omitempty"`
	ValidToInEpochSec   int64    `json:"validToInEpochSec,string,omitempty"`
}

func (service *Service) Get(baCertificateId int64) (*BaCertificate, *http.Response, error) {
	v := new(BaCertificate)
	relativeURL := fmt.Sprintf("%v/%v", mgmtConfig+service.Client.Config.CustomerID+baCertificateEndpoint, baCertificateId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*
func (service *Service) GetAll() ([]BaCertificate, *http.Response, error) {
	v := make([]BaCertificate, 0)
	resp, err := service.Client.NewRequestDo("GET", mgmtConfig+service.Client.Config.CustomerID+baCertificateEndpoint+baCertificateAll, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}
*/
