package manitable

import (
	"bytes"
	"encoding/csv"

	"github.com/fabioelizandro/manitable/assert"
)

type ManiTable struct {
	source  *TableSource
	renames map[string]string
}

func New(source *TableSource) ManiTable {
	return ManiTable{
		source:  source,
		renames: map[string]string{},
	}
}

func (t ManiTable) Rename(origName string, newName string) ManiTable {
	t.renames[origName] = newName

	return ManiTable{
		source:  t.source,
		renames: t.renames,
	}
}

func (t ManiTable) String() string {
	buffer := bytes.NewBufferString("")

	columns := []string{}
	for _, column := range t.source.columns {
		if newName, ok := t.renames[column]; ok {
			column = newName
		}

		columns = append(columns, column)
	}

	output := [][]string{columns}
	output = append(output, t.source.rows...)

	w := csv.NewWriter(buffer)
	assert.NoErr(w.WriteAll(output))
	assert.NoErr(w.Error())

	return buffer.String()
}

// -- Source

type TableSource struct {
	columns []string
	rows    [][]string
}

func NewTableSource(columns []string, rows [][]string) *TableSource {
	return &TableSource{
		columns: columns,
		rows:    rows,
	}
}
