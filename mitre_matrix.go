package stix2

// MitreMatrix represents a MITRE matrix
type MitreMatrix struct {
	tactics []*MitreTactic
	collectionReference
	STIXDomainObject
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	TacticRefs  []string `json:"tactic_refs,omitempty"`
	Domains     []string `json:"x_mitre_domains,omitempty"`
}

// Tactics returns all tactics associated to this matrix in the expected order
func (m *MitreMatrix) Tactics() []*MitreTactic {
	if m.ref == nil {
		return nil
	}
	if len(m.tactics) > 0 {
		return m.tactics
	}

	tactics := make([]*MitreTactic, 0, len(m.TacticRefs))
	for _, id := range m.TacticRefs {
		info := m.ref.Get(Identifier(id))
		if tactic, ok := info.(*MitreTactic); ok {
			tactics = append(tactics, tactic)
		}
	}

	m.SetTactics(tactics)

	return tactics
}

func (m *MitreMatrix) SetTactics(tactics []*MitreTactic) {
	m.tactics = tactics
}

// MitreID returns the external mitre id for this matrix
func (m *MitreMatrix) MitreID() string {
	for _, ref := range m.ExternalReferences {
		if ref.IsMitre() {
			return ref.ExternalID
		}
	}

	return ""
}

// MitreURL returns the external mitre url for this matrix
func (m *MitreMatrix) MitreURL() string {
	for _, ref := range m.ExternalReferences {
		if ref.IsMitre() {
			return ref.URL
		}
	}

	return ""
}
