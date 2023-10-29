package pipeline_test

import (
	"strings"
	"testing"

	"github.com/fabioelizandro/manitable/modules/pipeline"
	"github.com/stretchr/testify/assert"
)

func Test_pipeline_suite(t *testing.T) {
	t.Run("applies transformation to table", func(t *testing.T) {
		pipe := pipeline.NewPipeline(
			[]pipeline.Transform{
				transformAllInUpperCase,
			},
		)

		newTable := pipe.Run(pipeline.NewTable([]string{"name"}, [][]string{{"fabio"}}))

		assert.Equal(
			t,
			pipeline.NewTable([]string{"name"}, [][]string{{"FABIO"}}),
			newTable,
		)
	})

	t.Run("applies multiple transformations", func(t *testing.T) {
		pipe := pipeline.NewPipeline(
			[]pipeline.Transform{
				transformAllInUpperCase,
				transformAppendCharToAllValues('e'),
			},
		)

		newTable := pipe.Run(pipeline.NewTable([]string{"name"}, [][]string{{"fabio"}}))

		assert.Equal(
			t,
			pipeline.NewTable([]string{"name"}, [][]string{{"FABIOe"}}),
			newTable,
		)
	})

	t.Run("applies transformations in the correct sequence", func(t *testing.T) {
		pipe := pipeline.NewPipeline(
			[]pipeline.Transform{
				transformAppendCharToAllValues('e'),
				transformAllInUpperCase,
			},
		)

		newTable := pipe.Run(pipeline.NewTable([]string{"name"}, [][]string{{"fabio"}}))

		assert.Equal(
			t,
			pipeline.NewTable([]string{"name"}, [][]string{{"FABIOE"}}),
			newTable,
		)
	})
}

func transformAllInUpperCase(table pipeline.Table) pipeline.Table {
	return table.
		Mutate().
		Append(func(_ string, value string) string {
			return strings.ToUpper(value)
		}).
		Run()
}

func transformAppendCharToAllValues(char rune) pipeline.Transform {
	return func(table pipeline.Table) pipeline.Table {
		return table.
			Mutate().
			Append(func(_ string, value string) string {
				return value + string(char)
			}).
			Run()
	}
}
