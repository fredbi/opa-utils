package v1alpha1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAttackTrackUnmarshal(t *testing.T) {
	var obj AttackTrack
	file, err := os.ReadFile(filepath.Join("testdata", "attacktrack.json"))
	require.NoError(t, err)

	assert.NoError(t,
		json.Unmarshal(file, &obj),
	)
}
