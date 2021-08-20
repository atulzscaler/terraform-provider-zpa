package idpcontroller

import (
	"fmt"
	"net/http"
)

const (
	mgmtConfig                 = "/mgmtconfig/v1/admin/customers/"
	idpControllerGroupEndpoint = "/idp"
)

type IdpControllerRequest struct {
	// AutoProvision               int32    `json:"autoProvision,string"`
	// CreationTime                int32    `json:"creationTime,string"`
	Description            string   `json:"description"`
	DisableSamlBasedPolicy bool     `json:"disableSamlBasedPolicy"`
	Domainlist             []string `json:"domainList"`
	EnableScimBasedPolicy  bool     `json:"enableScimBasedPolicy"`
	Enabled                bool     `json:"enabled"`
	// ID                     int64    `json:"id,string"`
	IdpEntityId        string `json:"idpEntityId"`
	LoginNameAttribute string `json:"loginNameAttribute"`
	LoginUrl           string `json:"loginUrl"`
	// ModifiedBy                  int64    `json:"modifiedBy,string"`
	// ModifiedTime                int32    `json:"modifiedTime,string"`
	Name                        string `json:"name"`
	ReauthOnUserUpdate          bool   `json:"reauthOnUserUpdate"`
	RedirectBinding             bool   `json:"redirectBinding"`
	ScimEnabled                 bool   `json:"scimEnabled"`
	ScimServiceProviderEndpoint string `json:"scimServiceProviderEndpoint"`
	ScimSharedSecret            string `json:"scimSharedSecret"`
	ScimSharedSecretExists      bool   `json:"scimSharedSecretExists"`
	// SignSamlRequest             int32    `json:"signSamlRequest,string"`
	SsoType             []string `json:"ssoType"`
	UseCustomSpMetadata bool     `json:"useCustomSPMetadata"`
	//AdminMetadata               *AdminMetadata `json:"adminMetadata,omitempty"`
	UserMetadata UserMetadata   `json:"userMetadata"`
	Certificates []Certificates `json:"certificates"`
}

type IdpControllerResponse struct {
	AutoProvision               int32    `json:"autoProvision,string"`
	CreationTime                int32    `json:"creationTime,string"`
	Description                 string   `json:"description"`
	DisableSamlBasedPolicy      bool     `json:"disableSamlBasedPolicy"`
	Domainlist                  []string `json:"domainList"`
	EnableScimBasedPolicy       bool     `json:"enableScimBasedPolicy"`
	Enabled                     bool     `json:"enabled"`
	ID                          int64    `json:"id,string"`
	IdpEntityId                 string   `json:"idpEntityId"`
	LoginNameAttribute          string   `json:"loginNameAttribute"`
	LoginUrl                    string   `json:"loginUrl"`
	ModifiedBy                  int64    `json:"modifiedBy,string"`
	ModifiedTime                int32    `json:"modifiedTime,string"`
	Name                        string   `json:"name"`
	ReauthOnUserUpdate          bool     `json:"reauthOnUserUpdate"`
	RedirectBinding             bool     `json:"redirectBinding"`
	ScimEnabled                 bool     `json:"scimEnabled"`
	ScimServiceProviderEndpoint string   `json:"scimServiceProviderEndpoint"`
	ScimSharedSecret            string   `json:"scimSharedSecret"`
	ScimSharedSecretExists      bool     `json:"scimSharedSecretExists"`
	SignSamlRequest             int32    `json:"signSamlRequest,string"`
	SsoType                     []string `json:"ssoType"`
	UseCustomSpMetadata         bool     `json:"useCustomSPMetadata"`
	//AdminMetadata               *AdminMetadata `json:"adminMetadata,omitempty"`
	UserMetadata UserMetadata   `json:"userMetadata"`
	Certificates []Certificates `json:"certificates"`
}

// type AdminMetadata struct {
// 	CertificateUrl string `json:"certificateUrl"`
// 	SpentityId     string `json:"spEntityId"`
// 	Spmetadataurl  string `json:"spMetadataUrl"`
// 	SpPostUrl      string `json:"spPostUrl"`
// }
type Certificates struct {
	Cname          string `json:"cName"`
	Certificate    string `json:"certificate"`
	Serialno       string `json:"serialNo"`
	ValidFrominSec int64  `json:"validFromInSec,string"`
	ValidToinSec   int64  `json:"validToInSec,string"`
}
type UserMetadata struct {
	CertificateUrl string `json:"certificateUrl"`
	SpEntityId     string `json:"spEntityId"`
	SpMetadataUrl  string `json:"spMetadataUrl"`
	SpPostUrl      string `json:"spPostUrl"`
}

func (service *Service) Get(IdpID string) (*IdpControllerResponse, *http.Response, error) {
	v := new(IdpControllerResponse)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+idpControllerGroupEndpoint, IdpID)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
