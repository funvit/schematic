package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSerialize(t *testing.T) {
	var source *User

	source = UserWithDefaults()
	source.Binary = nil
	source.CreatedAt = time.Unix(333, 0)

	b, err := json.Marshal(source)
	require.NoError(t, err, "json marshall")

	var dest User
	err = json.Unmarshal(b, &dest)
	require.NoError(t, err, "json unmarshal")

	require.EqualValues(t, *source, dest)
}
