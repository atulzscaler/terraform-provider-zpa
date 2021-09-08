package scimattributeheader

import (
	"fmt"
	"net/http"
)

const (
	userConfig = "/userconfig/v1/customers/"
)

type ScimAttrHeader struct {
	CanonicalValues []string `json:"canonicalValues"`
	CaseSensitive   bool     `json:"caseSensitive"`
	CreationTime    int32    `json:"creationTime,"`
	DataType        int32    `json:"dataType,"`
	Description     string   `json:"description,"`
	ID              int64    `json:"id,"`
	IdpId           int64    `json:"idpId,"`
	ModifiedBy      int64    `json:"modifiedBy,"`
	ModifiedTime    int32    `json:"modifiedTime,"`
	MultiValued     bool     `json:"multivalued,"`
	Mutability      string   `json:"mutability,"`
	Name            string   `json:"name,omitempty"`
	Required        bool     `json:"required,"`
	Returned        string   `json:"returned,"`
	SchemaURI       string   `json:"schemaURI,"`
	Uniqueness      bool     `json:"uniqueness,"`
}

func (service *Service) Get(IdpId, ScimAttrHeaderId int64) (*ScimAttrHeader, *http.Response, error) {
	v := new(ScimAttrHeader)
	relativeURL := fmt.Sprintf(userConfig+service.Client.Config.CustomerID+"/idp/%d/scimattribute/%d", IdpId, ScimAttrHeaderId)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
