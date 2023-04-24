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

	t.Run("it renames multiple columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a", "c-b"}, [][]string{{"v-1", "v-2"}}))
		table = table.Rename("c-a", "c-b").Rename("c-b", "c-c")
		assert.Equal(t, "c-b,c-c\nv-1,v-2\n", table.String())
	})

	t.Run("it ignores non existing columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.Rename("c-b", "c-c")
		assert.Equal(t, "c-a\nv-1\n", table.String())
	})
}
