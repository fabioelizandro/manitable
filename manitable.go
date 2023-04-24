package manitable

import (
	"bytes"
	"encoding/csv"

	"github.com/fabioelizandro/manitable/assert"
)

type ManiTable struct {
	source    *TableSource
	renames   map[string]string
	deletes   map[string]bool
	additions map[string]TableCell
}

func New(source *TableSource) ManiTable {
	return ManiTable{
		source:    source,
		renames:   map[string]string{},
		deletes:   map[string]bool{},
		additions: map[string]TableCell{},
	}
}

func (t ManiTable) RenameColumn(origName string, newName string) ManiTable {
	t.renames[origName] = newName

	return ManiTable{
		source:    t.source,
		renames:   t.renames,
		deletes:   t.deletes,
		additions: t.additions,
	}
}

func (t ManiTable) DeleteColumn(s string) ManiTable {
	t.deletes[s] = true

	return ManiTable{
		source:    t.source,
		renames:   t.renames,
		deletes:   t.deletes,
		additions: t.additions,
	}
}

func (t ManiTable) AddColumn(columnName string, cell TableCell) ManiTable {
	t.additions[columnName] = cell

	return ManiTable{
		source:    t.source,
		renames:   t.renames,
		deletes:   t.deletes,
		additions: t.additions,
	}
}

func (t ManiTable) String() string {
	buffer := bytes.NewBufferString("")

	columns := []string{}
	toBeDeleted := map[int]bool{}
	for index, column := range t.source.columns {
		if t.deletes[column] {
			toBeDeleted[index] = true
			continue
		}

		if newName, ok := t.renames[column]; ok {
			column = newName
		}

		columns = append(columns, column)
	}

	cells := []TableCell{}
	for columnName, cell := range t.additions {
		columns = append(columns, columnName)
		cells = append(cells, cell)
	}

	rows := [][]string{}
	for _, row := range t.source.rows {
		newRow := []string{}
		for index, cell := range row {
			if toBeDeleted[index] {
				continue
			}

			newRow = append(newRow, cell)
		}

		for _, cell := range cells {
			newRow = append(newRow, cell.Value(row))
		}

		rows = append(rows, newRow)
	}

	output := [][]string{columns}
	output = append(output, rows...)

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

// -- Table Cell

type TableCell interface {
	Value(row TableRow) string
}

type InlineFTableCell struct {
	f func(row TableRow) string
}

func NewInlineFTableCell(f func(row TableRow) string) InlineFTableCell {
	return InlineFTableCell{f: f}
}

func (c InlineFTableCell) Value(row TableRow) string {
	return c.f(row)
}

// -- Table Row

type TableRow interface {
}
