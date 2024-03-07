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

// FileData describes a file submitted to libra
type FileData struct {
	ID       string `json:"id"`
	MimeType string `json:"mimeType"`
	Name     string `json:"name"`
	// TODO more fields... URL ? Stream? Payload?
}

// ETDWorkFromBytes will create an ETDWork from a byte array
func ETDWorkFromBytes(bytes []byte) (*ETDWork, error) {
	var etdWork ETDWork
	err := json.Unmarshal(bytes, &etdWork)
	if err != nil {
		return nil, err
	}
	return &etdWork, nil
}

// ETDWork contains libra metadata for ETD works
type ETDWork struct {
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

// MimeType gets the mime type of ETD metadata
func (etd ETDWork) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of ETD metadata
func (etd ETDWork) Payload() ([]byte, error) {
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

// MimeType gets the mime type of openAccess metadata
func (oa OAWork) MimeType() string {
	return "application/json"
}

// Payload gets the encoded binary representation of OpenAccess metadata
func (oa OAWork) Payload() ([]byte, error) {
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
