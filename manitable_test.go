package manitable_test

import (
	"testing"

	"github.com/fabioelizandro/manitable"
	"github.com/stretchr/testify/assert"
)

func Test_manitable_suite(t *testing.T) {
	t.Run("it renames columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.Rename("c-a", "c-b")
		assert.Equal(t, "c-b\nv-1\n", table.String())
	})
}
