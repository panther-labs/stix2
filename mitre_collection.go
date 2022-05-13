package stix2

// MitreCollection represents a MITRE collection
type MitreCollection struct {
	STIXDomainObject
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"x_mitre_version"`
}
