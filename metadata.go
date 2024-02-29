package libra

import (
	"encoding/json"
	"time"
)

type authorData struct {
	ComputeID   string `json:"computeID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Department  string `json:"department"`
	Institution string `json:"institution"`
}

type oaDepositData struct {
	Visibility       string       `json:"visibility"`
	ResourceType     string       `json:"resourceType"`
	Title            string       `json:"title"`
	Authors          []authorData `json:"authors"`
	Abstract         string       `json:"abstract"`
	License          string       `json:"license"`
	Languages        []string     `json:"languages"`
	Keywords         []string     `json:"keywords"`
	Contributors     []authorData `json:"contributors"`
	Publisher        string       `json:"publisher"`
	Citation         string       `json:"citation"`
	PubllicationData string       `json:"pubDate"`
	RelatedURLs      []string     `json:"relatedURLs"`
	Sponsors         []string     `json:"sponsors"`
	Notes            string       `json:"notes"`
}

type easystoreOAWrapper struct {
	JSONData   oaDepositData
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (oa easystoreOAWrapper) MimeType() string {
	return "application/json"
}

func (oa easystoreOAWrapper) Payload() []byte {
	out, _ := json.Marshal(oa.JSONData)
	return out
}

func (oa easystoreOAWrapper) Created() time.Time {
	return oa.CreatedAt
}

func (oa easystoreOAWrapper) Modified() time.Time {
	return oa.ModifiedAt
}
