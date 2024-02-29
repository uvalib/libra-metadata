package librametadata

import (
	"encoding/json"
	"time"
)

// AuthorData contains libra metadata for authors and contributors
type AuthorData struct {
	ComputeID   string `json:"computeID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Department  string `json:"department"`
	Institution string `json:"institution"`
}

// OADepositData contsins libra metadata for openAccess submissions
type OADepositData struct {
	Visibility       string       `json:"visibility"`
	ResourceType     string       `json:"resourceType"`
	Title            string       `json:"title"`
	Authors          []AuthorData `json:"authors"`
	Abstract         string       `json:"abstract"`
	License          string       `json:"license"`
	Languages        []string     `json:"languages"`
	Keywords         []string     `json:"keywords"`
	Contributors     []AuthorData `json:"contributors"`
	Publisher        string       `json:"publisher"`
	Citation         string       `json:"citation"`
	PubllicationData string       `json:"pubDate"`
	RelatedURLs      []string     `json:"relatedURLs"`
	Sponsors         []string     `json:"sponsors"`
	Notes            string       `json:"notes"`
}

// EasyStoreOAWrapper is a wrapper around OpenAccess data that returns the metadata in a storage format
type EasyStoreOAWrapper struct {
	JSONData   OADepositData
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// MimeType gets the mime type of openAccess metadata
func (oa EasyStoreOAWrapper) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of OpenAccess metadata
func (oa EasyStoreOAWrapper) Payload() ([]byte, error) {
	return json.Marshal(oa.JSONData)
}

// Created gets date when the OpenAccess metadata was created in easystore
func (oa EasyStoreOAWrapper) Created() time.Time {
	return oa.CreatedAt
}

// Modified gets last modification date of the OpenAccess metadata
func (oa EasyStoreOAWrapper) Modified() time.Time {
	return oa.ModifiedAt
}
