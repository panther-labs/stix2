package stix2

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttackPattern(t *testing.T) {
	assert := assert.New(t)

	name := "spearphishing"

	t.Run("missing_property", func(t *testing.T) {
		r, err := NewAttackPattern("")
		assert.Nil(r)
		assert.Equal(ErrPropertyMissing, err)
	})

	t.Run("no_optional", func(t *testing.T) {
		r, err := NewAttackPattern(name, nil)
		assert.NotNil(r)
		assert.NoError(err)
	})

	t.Run("with_options", func(t *testing.T) {
		conf := 50
		desc := "My description"
		ts := &Timestamp{time.Now()}
		createdBy := NewIdentifier(TypeIdentity)
		ref := &ExternalReference{}
		marking := make([]*GranularMarking, 0)
		labels := []string{"tag1", "tag2"}
		lang := "en"
		objmark := []Identifier{Identifier("id")}
		specVer := "2.0"

		aliases := []string{"name1", "name2"}
		killChain := []*KillChainPhase{{}}

		opts := []STIXOption{
			OptionConfidence(conf),
			OptionDescription(desc),
			OptionCreated(ts),
			OptionModified(ts),
			OptionCreatedBy(createdBy),
			OptionExternalReferences([]*ExternalReference{ref}),
			OptionGranularMarking(marking),
			OptionLabels(labels),
			OptionLang(lang),
			OptionObjectMarking(objmark),
			OptionRevoked(true),
			OptionSpecVersion(specVer),
			//
			OptionAliases(aliases),
			OptionKillChainPhase(killChain),
		}
		a, err := NewAttackPattern(name, opts...)
		assert.NotNil(a)
		assert.NoError(err)
		assert.Equal(conf, a.Confidence)
		assert.Equal(desc, a.Description)
		assert.Equal(ts, a.Created)
		assert.Equal(ts, a.Modified)
		assert.Equal(createdBy, a.CreatedBy)
		assert.Contains(a.ExternalReferences, ref)
		assert.Equal(marking, a.GranularMarking)
		assert.Equal(labels, a.Labels)
		assert.Equal(lang, a.Lang)
		assert.Equal(objmark, a.ObjectMarking)
		assert.True(a.Revoked)
		assert.Equal(specVer, a.SpecVersion)

		assert.Equal(aliases, a.Aliases)
		assert.Equal(killChain, a.KillChainPhase)
	})

	t.Run("parse_json", func(t *testing.T) {
		data := []byte(`{
  "type": "attack-pattern",
  "spec_version": "2.1",
  "id": "attack-pattern--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061",
  "created": "2016-05-12T08:17:27.000Z",
  "modified": "2016-05-12T08:17:27.000Z",
  "name": "Spear Phishing",
  "description": "...",
  "external_references": [
    {
      "source_name": "capec",
      "external_id": "CAPEC-163"
    }
  ]
}`)
		ts, err := time.Parse(time.RFC3339Nano, "2016-05-12T08:17:27.000Z")
		assert.NoError(err)
		var obj *AttackPattern
		err = json.Unmarshal(data, &obj)
		assert.NoError(err)
		assert.Equal(Identifier("attack-pattern--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061"), obj.ID)
		assert.Equal("2.1", obj.SpecVersion)
		assert.Equal(TypeAttackPattern, obj.Type)
		assert.Equal("Spear Phishing", obj.Name)
		assert.Equal("...", obj.Description)
		assert.Len(obj.ExternalReferences, 1)
		assert.Equal("capec", obj.ExternalReferences[0].Name)
		assert.Equal("CAPEC-163", obj.ExternalReferences[0].ExternalID)
		assert.Equal(ts, obj.Created.Time)
		assert.Equal(ts, obj.Modified.Time)
	})

	t.Run("should return empty mitre url and id", func(t *testing.T) {
		attack := &AttackPattern{}
		require.Empty(t, attack.MitreID())
		require.Empty(t, attack.MitreURL())
	})
}

func TestAttackPatternAddDelivers(t *testing.T) {
	assert := assert.New(t)

	t.Run("valid", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeMalware)
		rel, err := a.AddDelivers(id)
		assert.NoError(err)
		assert.Equal(id, rel.Target)
		assert.Equal(a.ID, rel.Source)
	})

	t.Run("invalid_type", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeIPv4Addr)
		rel, err := a.AddDelivers(id)
		assert.Equal(err, ErrInvalidParameter)
		assert.Nil(rel)
	})
}

func TestAttackPatternAddTargets(t *testing.T) {
	assert := assert.New(t)

	t.Run("identity", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeIdentity)
		rel, err := a.AddTargets(id)
		assert.NoError(err)
		assert.Equal(id, rel.Target)
		assert.Equal(a.ID, rel.Source)
	})

	t.Run("location", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeLocation)
		rel, err := a.AddTargets(id)
		assert.NoError(err)
		assert.Equal(id, rel.Target)
		assert.Equal(a.ID, rel.Source)
	})

	t.Run("identity", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeVulnerability)
		rel, err := a.AddTargets(id)
		assert.NoError(err)
		assert.Equal(id, rel.Target)
		assert.Equal(a.ID, rel.Source)
	})

	t.Run("invalid_type", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeIPv4Addr)
		rel, err := a.AddTargets(id)
		assert.Equal(err, ErrInvalidParameter)
		assert.Nil(rel)
	})
}

func TestAttackPatternAddUses(t *testing.T) {
	assert := assert.New(t)

	t.Run("malware", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeMalware)
		rel, err := a.AddUses(id)
		assert.NoError(err)
		assert.Equal(id, rel.Target)
		assert.Equal(a.ID, rel.Source)
	})

	t.Run("tool", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeTool)
		rel, err := a.AddUses(id)
		assert.NoError(err)
		assert.Equal(id, rel.Target)
		assert.Equal(a.ID, rel.Source)
	})

	t.Run("invalid_type", func(t *testing.T) {
		a, err := NewAttackPattern("name")
		assert.NoError(err)
		id := NewIdentifier(TypeIPv4Addr)
		rel, err := a.AddUses(id)
		assert.Equal(err, ErrInvalidParameter)
		assert.Nil(rel)
	})
}
