package manitable_test

import (
	"testing"

	"github.com/fabioelizandro/manitable"
	"github.com/stretchr/testify/assert"
)

func Test_manitable_suite(t *testing.T) {
	t.Run("it renames columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.RenameColumn("c-a", "c-b")
		assert.Equal(t, "c-b\nv-1\n", table.String())
	})

	t.Run("it renames multiple columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a", "c-b"}, [][]string{{"v-1", "v-2"}}))
		table = table.RenameColumn("c-a", "c-b").RenameColumn("c-b", "c-c")
		assert.Equal(t, "c-b,c-c\nv-1,v-2\n", table.String())
	})

	t.Run("it keeps the last rename for the given column", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.RenameColumn("c-a", "c-b").RenameColumn("c-a", "c-c")
		assert.Equal(t, "c-b\nv-1\n", table.String())
	})

	t.Run("it does not chain multiple renames for the same column", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.RenameColumn("c-a", "c-b").RenameColumn("c-b", "c-c")
		assert.Equal(t, "c-b\nv-1\n", table.String())
	})

	t.Run("it ignores renames for non existing columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.RenameColumn("c-b", "c-c")
		assert.Equal(t, "c-a\nv-1\n", table.String())
	})

	t.Run("it deletes columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a", "c-b"}, [][]string{{"v-1", "v-2"}}))
		table = table.DeleteColumn("c-a")
		assert.Equal(t, "c-b\nv-2\n", table.String())
	})

	t.Run("it deletes multiple columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a", "c-b", "c-c"}, [][]string{{"v-1", "v-2", "v-3"}}))
		table = table.DeleteColumn("c-a").DeleteColumn("c-c")
		assert.Equal(t, "c-b\nv-2\n", table.String())
	})

	t.Run("it ignores deletes for non existing columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.DeleteColumn("c-b")
		assert.Equal(t, "c-a\nv-1\n", table.String())
	})

	t.Run("it adds column", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.AddColumn("c-b", manitable.NewInlineFTableCell(func(row manitable.TableRow) string {
			return "v-2"
		}))
		assert.Equal(t, "c-a,c-b\nv-1,v-2\n", table.String())
	})

	t.Run("it adds multiple columns", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))
		table = table.
			AddColumn("c-b", manitable.NewInlineFTableCell(func(row manitable.TableRow) string {
				return "v-2"
			})).
			AddColumn("c-c", manitable.NewInlineFTableCell(func(row manitable.TableRow) string {
				return "v-3"
			}))

		assert.Equal(t, "c-a,c-b,c-c\nv-1,v-2,v-3\n", table.String())
	})

	t.Run("it transforms values for a given column", func(t *testing.T) {
		table := manitable.New(manitable.NewTableSource([]string{"c-a"}, [][]string{{"v-1"}}))

		table = table.TransformColumn("c-a", manitable.NewInlineFTableCell(func(row manitable.TableRow) string {
			return "v-2"
		}))

		assert.Equal(t, "c-a\nv-2\n", table.String())
	})
}
