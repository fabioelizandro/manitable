package manitable_test

import (
	"strings"
	"testing"

	"github.com/fabioelizandro/manitable"
	"github.com/stretchr/testify/assert"
)

func Test_pipeline_suite(t *testing.T) {
	t.Run("applies transformation to table", func(t *testing.T) {
		pipeline := manitable.NewPipeline(
			[]manitable.Transform{
				transformAllInUpperCase,
			},
		)

		newTable := pipeline.Run(manitable.NewTable([]string{"name"}, [][]string{{"fabio"}}))

		assert.Equal(
			t,
			manitable.NewTable([]string{"name"}, [][]string{{"FABIO"}}),
			newTable,
		)
	})

	t.Run("applies multiple transformations", func(t *testing.T) {
		pipeline := manitable.NewPipeline(
			[]manitable.Transform{
				transformAllInUpperCase,
				transformAppendCharToAllValues('e'),
			},
		)

		newTable := pipeline.Run(manitable.NewTable([]string{"name"}, [][]string{{"fabio"}}))

		assert.Equal(
			t,
			manitable.NewTable([]string{"name"}, [][]string{{"FABIOe"}}),
			newTable,
		)
	})

	t.Run("applies transformations in the correct sequence", func(t *testing.T) {
		pipeline := manitable.NewPipeline(
			[]manitable.Transform{
				transformAppendCharToAllValues('e'),
				transformAllInUpperCase,
			},
		)

		newTable := pipeline.Run(manitable.NewTable([]string{"name"}, [][]string{{"fabio"}}))

		assert.Equal(
			t,
			manitable.NewTable([]string{"name"}, [][]string{{"FABIOE"}}),
			newTable,
		)
	})
}

func transformAllInUpperCase(table manitable.Table) manitable.Table {
	rows := table.Rows()
	newRows := make([][]string, len(rows))
	for i, row := range rows {
		newRows[i] = make([]string, len(row))
		for j, value := range row {
			newRows[i][j] = strings.ToUpper(value)
		}
	}

	return manitable.NewTable(table.Columns(), newRows)
}

func transformAppendCharToAllValues(char rune) manitable.Transform {
	return func(table manitable.Table) manitable.Table {
		rows := table.Rows()
		newRows := make([][]string, len(rows))
		for i, row := range rows {
			newRows[i] = make([]string, len(row))
			for j, value := range row {
				newRows[i][j] = value + string(char)
			}
		}

		return manitable.NewTable(table.Columns(), newRows)
	}
}
