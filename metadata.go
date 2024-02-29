package librametadata

import (
	"encoding/json"
	"time"
)

// ContributorData contains libra metadata for authors, contributors or advisors
type ContributorData struct {
	ComputeID   string `json:"computeID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Department  string `json:"department"`
	Institution string `json:"institution"`
}

// StudentData contains libra metadata for student authors
type StudentData struct {
	ComputeID   string `json:"computeID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Program     string `json:"program"`
	Institution string `json:"institution"`
}

// ETDWorkFromBytes will create an ETDWork from a byte array
func ETDWorkFromBytes(bytes []byte) (*EDTWork, error) {
	var etdWork EDTWork
	err := json.Unmarshal(bytes, &etdWork)
	if err != nil {
		return nil, err
	}
	return &etdWork, nil
}

// EDTWork contains libra metadata for ETD works
type EDTWork struct {
	Degree      string            `json:"degree"`
	Visibility  string            `json:"visibility"`
	Title       string            `json:"title"`
	Author      StudentData       `json:"author"`
	Advisors    []ContributorData `json:"advisors"`
	Abstract    string            `json:"abstract"`
	License     string            `json:"license"`
	Keywords    []string          `json:"keywords"`
	Language    string            `json:"language"`
	RelatedURLs []string          `json:"relatedURLs"`
	Sponsors    []string          `json:"sponsors"`
	Notes       string            `json:"notes"`
}

// EasyStoreETD is a wrapper around ETD data that returns the metadata in a storage format
type EasyStoreETD struct {
	JSONData   EDTWork
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// MimeType gets the mime type of ETD metadata
func (oa EasyStoreETD) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of ETD metadata
func (oa EasyStoreETD) Payload() ([]byte, error) {
	return json.Marshal(oa.JSONData)
}

// Created gets date when the OpenAccess metadata was created in easystore
func (oa EasyStoreETD) Created() time.Time {
	return oa.CreatedAt
}

// Modified gets last modification date of the ETD metadata
func (oa EasyStoreETD) Modified() time.Time {
	return oa.ModifiedAt
}

// OAWorkFromBytes will create an OAWork from a byte array
func OAWorkFromBytes(bytes []byte) (*OAWork, error) {
	var oaWork OAWork
	err := json.Unmarshal(bytes, &oaWork)
	if err != nil {
		return nil, err
	}
	return &oaWork, nil
}

// OAWork contains libra metadata for openAccess works
type OAWork struct {
	Visibility       string            `json:"visibility"`
	ResourceType     string            `json:"resourceType"`
	Title            string            `json:"title"`
	Authors          []ContributorData `json:"authors"`
	Abstract         string            `json:"abstract"`
	License          string            `json:"license"`
	Languages        []string          `json:"languages"`
	Keywords         []string          `json:"keywords"`
	Contributors     []ContributorData `json:"contributors"`
	Publisher        string            `json:"publisher"`
	Citation         string            `json:"citation"`
	PubllicationData string            `json:"pubDate"`
	RelatedURLs      []string          `json:"relatedURLs"`
	Sponsors         []string          `json:"sponsors"`
	Notes            string            `json:"notes"`
}

// EasyStoreOA is a wrapper around OpenAccess data that returns the metadata in a storage format
type EasyStoreOA struct {
	JSONData   OAWork
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// MimeType gets the mime type of openAccess metadata
func (oa EasyStoreOA) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of OpenAccess metadata
func (oa EasyStoreOA) Payload() ([]byte, error) {
	return json.Marshal(oa.JSONData)
}

// Created gets date when the OpenAccess metadata was created in easystore
func (oa EasyStoreOA) Created() time.Time {
	return oa.CreatedAt
}

// Modified gets last modification date of the OpenAccess metadata
func (oa EasyStoreOA) Modified() time.Time {
	return oa.ModifiedAt
}
