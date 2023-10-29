package must_test

import (
	"errors"
	"testing"

	"github.com/fabioelizandro/manitable/modules/must"
	"github.com/stretchr/testify/assert"
)

func TestAssertSuite(t *testing.T) {
	t.Run("no err", func(t *testing.T) {
		must.NoErr(nil)

		f := func() {
			must.NoErr(errors.New("foo"))
		}

		assert.Panics(t, f)
	})

	t.Run("true", func(t *testing.T) {
		must.True(true, "s")

		f := func() {
			must.True(false, "t")
		}

		assert.Panics(t, f)
	})

	t.Run("false", func(t *testing.T) {
		must.False(false, "s")

		f := func() {
			must.False(true, "t")
		}

		assert.Panics(t, f)
	})

	t.Run("must", func(t *testing.T) {
		assert.True(t, must.Return(true, nil))
		assert.Equal(t, 1, must.Return(1, nil))

		assert.PanicsWithError(t, "some error", func() {
			must.Return(false, errors.New("some error"))
		})
	})
}
