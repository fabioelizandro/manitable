package manitable

import (
	"bytes"
	"encoding/csv"

	"github.com/fabioelizandro/manitable/assert"
)

type ManiTable struct {
	source       *TableSource
	renames      map[string]string
	deletes      map[string]bool
	additions    map[string]TableCell
	transformers map[string]TableCell
}

func New(source *TableSource) ManiTable {
	return ManiTable{
		source:       source,
		renames:      map[string]string{},
		deletes:      map[string]bool{},
		additions:    map[string]TableCell{},
		transformers: map[string]TableCell{},
	}
}

func (t ManiTable) RenameColumn(origName string, newName string) ManiTable {
	t.renames[origName] = newName

	return ManiTable{
		source:       t.source,
		renames:      t.renames,
		deletes:      t.deletes,
		additions:    t.additions,
		transformers: t.transformers,
	}
}

func (t ManiTable) DeleteColumn(s string) ManiTable {
	t.deletes[s] = true

	return ManiTable{
		source:       t.source,
		renames:      t.renames,
		deletes:      t.deletes,
		additions:    t.additions,
		transformers: t.transformers,
	}
}

func (t ManiTable) AddColumn(columnName string, cell TableCell) ManiTable {
	t.additions[columnName] = cell

	return ManiTable{
		source:       t.source,
		renames:      t.renames,
		deletes:      t.deletes,
		additions:    t.additions,
		transformers: t.transformers,
	}
}

func (t ManiTable) TransformColumn(columnName string, cell TableCell) ManiTable {
	t.transformers[columnName] = cell

	return ManiTable{
		source:       t.source,
		renames:      t.renames,
		deletes:      t.deletes,
		additions:    t.additions,
		transformers: t.transformers,
	}
}

func (t ManiTable) String() string {
	buffer := bytes.NewBufferString("")

	columns := []string{}
	toBeDeleted := map[int]bool{}
	toBeTransformed := map[int]TableCell{}
	for index, column := range t.source.columns {
		if t.deletes[column] {
			toBeDeleted[index] = true
			continue
		}

		if cell, ok := t.transformers[column]; ok {
			toBeTransformed[index] = cell
		}

		if newName, ok := t.renames[column]; ok {
			column = newName
		}

		columns = append(columns, column)
	}

	addedCells := []TableCell{}
	for columnName, cell := range t.additions {
		columns = append(columns, columnName)
		addedCells = append(addedCells, cell)
	}

	rows := [][]string{}
	for _, row := range t.source.rows {
		newRow := []string{}
		for index, cell := range row {
			if toBeDeleted[index] {
				continue
			}

			if transformer, ok := toBeTransformed[index]; ok {
				cell = transformer.Value(row)
			}

			newRow = append(newRow, cell)
		}

		for _, cell := range addedCells {
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
