package stix2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMitreTactic_AttackPatterns(t *testing.T) {
	t.Run("should return nil if ref is nil", func(t *testing.T) {
		tac := &MitreTactic{}
		require.Nil(t, tac.AttackPatterns())
	})

	t.Run("should return nil if attack lookup is nil", func(t *testing.T) {
		tac := &MitreTactic{
			ref: &Collection{},
		}
		require.Nil(t, tac.AttackPatterns())
	})

	t.Run("should set attack lookup when nil", func(t *testing.T) {
		tac := &MitreTactic{
			ref: &Collection{},
		}
		tac.SetAttackPattern(nil)
		require.NotNil(t, tac.attackPatternsLookup)
	})

	t.Run("should return empty mitre url and id", func(t *testing.T) {
		tac := &MitreTactic{}
		require.Empty(t, tac.MitreID())
		require.Empty(t, tac.MitreURL())
	})
}
