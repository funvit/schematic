package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMarshalJSON(t *testing.T) {

	t.Run("marshal", func(t *testing.T) {
		m := UserWithDefaults()
		m.Id = 17
		m.Name = "Bob"
		m.Binary = []byte("some bytes here")
		m.CreatedAt = time.Unix(400, 0)

		b, err := json.Marshal(m)
		require.NoError(t, err)
		require.NotEmpty(t, b)

		t.Run("unmarshal", func(t *testing.T) {

			var m2 User
			err := json.Unmarshal(b, &m2)
			require.NoError(t, err)
			require.EqualValues(t, *m, m2)
		})
	})
}
