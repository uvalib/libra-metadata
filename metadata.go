package librametadata

import (
	"encoding/json"
	"fmt"
	"time"
)

// update this when an incompatible schema change is made.
// note adding and removing fields maintains compatibility, renaming and retyping does not.
var schemaVersion = "1"

// ErrSchemaVersion is the error that is thrown when there is a schema mismatch
var ErrSchemaVersion = fmt.Errorf("incompatible schema versions, some data may be lost")

// SchemaVersion mechanism to manage schema versioning
type SchemaVersion struct {
	Version string `json:"version"`
}

// ContributorData contains libra metadata for authors (student or otherwise), contributors or advisors
type ContributorData struct {
	ComputeID   string `json:"computeID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Department  string `json:"department,omitempty"` // This will be blank for student ETD author
	Institution string `json:"institution"`
	ORCID       string `json:"orcid"`
}

// FileData describes a file submitted to libra
type FileData struct {
	MimeType  string    `json:"mimeType"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	// TODO more fields... URL ? Stream? Payload?
}

// ETDWorkFromBytes will create an ETDWork from a byte array
func ETDWorkFromBytes(bytes []byte) (*ETDWork, error) {
	var schema SchemaVersion
	err := json.Unmarshal(bytes, &schema)
	if err != nil {
		return nil, err
	}
	if schema.Version != schemaVersion {
		return nil, ErrSchemaVersion
	}

	var etdWork ETDWork
	err = json.Unmarshal(bytes, &etdWork)
	if err != nil {
		return nil, err
	}

	if etdWork.Keywords == nil {
		etdWork.Keywords = make([]string, 0)
	}
	if etdWork.RelatedURLs == nil {
		etdWork.RelatedURLs = make([]string, 0)
	}
	if etdWork.Sponsors == nil {
		etdWork.Sponsors = make([]string, 0)
	}
	if etdWork.Advisors == nil {
		etdWork.Advisors = make([]ContributorData, 0)
	}

	return &etdWork, nil
}

// ETDWork contains libra metadata for ETD works
type ETDWork struct {
	SchemaVersion
	Program     string            `json:"program"`
	Degree      string            `json:"degree"`
	Title       string            `json:"title"`
	Author      ContributorData   `json:"author"`
	Advisors    []ContributorData `json:"advisors"`
	Abstract    string            `json:"abstract"`
	License     string            `json:"license"`
	LicenseURL  string            `json:"licenseURL"`
	Keywords    []string          `json:"keywords"`
	Language    string            `json:"language"`
	RelatedURLs []string          `json:"relatedURLs"`
	Sponsors    []string          `json:"sponsors"`
	Notes       string            `json:"notes"`
	AdminNotes  string            `json:"adminNotes"`
}

// IsAuthor checks if the passed computeID is a work author
func (etd ETDWork) IsAuthor(computeID string) bool {
	return etd.Author.ComputeID == computeID
}

// MimeType gets the mime type of ETD metadata
func (etd ETDWork) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of ETD metadata
func (etd ETDWork) Payload() ([]byte, error) {
	etd.SchemaVersion.Version = schemaVersion
	return json.Marshal(etd)
}

// Created gets date when the OpenAccess metadata was created in easystore.
func (etd ETDWork) Created() time.Time {
	return time.Time{}
}

// Modified gets last modification date of the ETD metadata
func (etd ETDWork) Modified() time.Time {
	return time.Time{}
}

// OAWorkFromBytes will create an OAWork from a byte array
func OAWorkFromBytes(bytes []byte) (*OAWork, error) {
	var schema SchemaVersion
	err := json.Unmarshal(bytes, &schema)
	if err != nil {
		return nil, err
	}
	if schema.Version != schemaVersion {
		return nil, ErrSchemaVersion
	}

	var oaWork OAWork
	err = json.Unmarshal(bytes, &oaWork)
	if err != nil {
		return nil, err
	}
	if oaWork.Keywords == nil {
		oaWork.Keywords = make([]string, 0)
	}
	if oaWork.Languages == nil {
		oaWork.Languages = make([]string, 0)
	}
	if oaWork.Sponsors == nil {
		oaWork.Sponsors = make([]string, 0)
	}
	if oaWork.RelatedURLs == nil {
		oaWork.RelatedURLs = make([]string, 0)
	}
	if oaWork.Authors == nil {
		oaWork.Authors = make([]ContributorData, 0)
	}
	if oaWork.Contributors == nil {
		oaWork.Contributors = make([]ContributorData, 0)
	}

	return &oaWork, nil
}

// OAWork contains libra metadata for openAccess works
type OAWork struct {
	SchemaVersion
	ResourceType    string            `json:"resourceType"`
	Title           string            `json:"title"`
	Authors         []ContributorData `json:"authors"`
	Abstract        string            `json:"abstract"`
	License         string            `json:"license"`
	LicenseURL      string            `json:"licenseURL"`
	Languages       []string          `json:"languages"`
	Keywords        []string          `json:"keywords"`
	Contributors    []ContributorData `json:"contributors"`
	Publisher       string            `json:"publisher"`
	Citation        string            `json:"citation"`
	PublicationDate string            `json:"pubDate"`
	RelatedURLs     []string          `json:"relatedURLs"`
	Sponsors        []string          `json:"sponsors"`
	Notes           string            `json:"notes"`
	AdminNotes      string            `json:"adminNotes"`
}

// IsAuthor checks if the passed computeID is a work author
func (oa OAWork) IsAuthor(computeID string) bool {
	isAuthor := false
	for _, author := range oa.Authors {
		if author.ComputeID == computeID {
			isAuthor = true
			break
		}
	}
	return isAuthor
}

// MimeType gets the mime type of openAccess metadata
func (oa OAWork) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of OpenAccess metadata
func (oa OAWork) Payload() ([]byte, error) {
	oa.SchemaVersion.Version = schemaVersion
	return json.Marshal(oa)
}

// Created gets date when the OpenAccess metadata was created in easystore
func (oa OAWork) Created() time.Time {
	return time.Time{}
}

// Modified gets last modification date of the OpenAccess metadata
func (oa OAWork) Modified() time.Time {
	return time.Time{}
}

// Audit metadata
type Audit struct {
	Who       string    `json:"who"`
	Oid       string    `json:"oid"`
	Namespace string    `json:"namespace"`
	FieldName string    `json:"fieldName"`
	Before    string    `json:"before"`
	After     string    `json:"after"`
	EventTime time.Time `json:"eventTime"`
}

// AuditsFromBytes will create an Audit array from a byte array
func AuditsFromBytes(bytes []byte) (*[]Audit, error) {

	var audit []Audit
	err := json.Unmarshal(bytes, &audit)
	if err != nil {
		return nil, err
	}

	return &audit, nil
}
