package stix2

// MitreTactic represents a MITRE tactic
type MitreTactic struct {
	attackPatternsLookup map[*AttackPattern]struct{}
	ref                  *Collection
	STIXDomainObject
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Domains     []string `json:"x_mitre_domains,omitempty"`
	ShortName   string   `json:"x_mitre_shortname,omitempty"`
}

// SetReference sets the collection reference and updates collection maps as needed
func (m *MitreTactic) SetReference(collection *Collection) {
	m.ref = collection

	if m.attackPatternsLookup == nil {
		m.attackPatternsLookup = make(map[*AttackPattern]struct{})
	}

	collection.mitreTacticsByShortName[m.ShortName] = m
}

// SetAttackPattern associates an attack pattern
func (m *MitreTactic) SetAttackPattern(attack *AttackPattern) {
	if m.attackPatternsLookup == nil {
		m.attackPatternsLookup = make(map[*AttackPattern]struct{})
	}

	m.attackPatternsLookup[attack] = struct{}{}
}

// AttackPatterns returns all attack patterns associated to this tactic
func (m *MitreTactic) AttackPatterns() []*AttackPattern {
	if m.ref == nil || m.attackPatternsLookup == nil {
		return nil
	}

	results := make([]*AttackPattern, 0, len(m.attackPatternsLookup))
	for ref := range m.attackPatternsLookup {
		results = append(results, ref)
	}

	return results
}

// MitreID returns the external mitre id for this tactic
func (m *MitreTactic) MitreID() string {
	for _, ref := range m.ExternalReferences {
		if ref.IsMitre() {
			return ref.ExternalID
		}
	}

	return ""
}

// MitreURL returns the external mitre url for this attack pattern
func (m *MitreTactic) MitreURL() string {
	for _, ref := range m.ExternalReferences {
		if ref.IsMitre() {
			return ref.URL
		}
	}

	return ""
}
