package stix2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMitreMatrix_Tactics(t *testing.T) {
	t.Run("should return nil if no reference set", func(t *testing.T) {
		m := &MitreMatrix{}

		require.Nil(t, m.Tactics())
	})
}
