package pipeline

type Table struct {
	columns []string
	rows    [][]string
}

func NewTable(columns []string, rows [][]string) Table {
	return Table{columns: columns, rows: rows}
}

func (t Table) Mutate() TableMutation {
	return newTableMutation(t, map[int]bool{}, []func(column string, value string) string{})
}

// TableMutation

type TableMutation struct {
	table         Table
	mutations     []func(column string, value string) string
	columnsToDrop map[int]bool
}

func newTableMutation(table Table, columnsToDrop map[int]bool, mutations []func(column string, value string) string) TableMutation {
	return TableMutation{table: table, columnsToDrop: columnsToDrop, mutations: mutations}
}

func (m TableMutation) DropColumn(columnName string) TableMutation {
	for i, column := range m.table.columns {
		if column == columnName {
			columnsToDrop := m.columnsToDrop
			columnsToDrop[i] = true

			return newTableMutation(m.table, columnsToDrop, m.mutations)
		}
	}

	return m
}

func (m TableMutation) Append(f func(column string, value string) string) TableMutation {
	return newTableMutation(m.table, m.columnsToDrop, append(m.mutations, f))
}

func (m TableMutation) Run() Table {
	rows := m.table.rows
	newRows := make([][]string, len(rows))
	for i, row := range rows {
		newRows[i] = make([]string, len(row))
		for j, value := range row {
			if m.columnsToDrop[j] {
				break
			}

			for _, mutation := range m.mutations {
				value = mutation(m.table.columns[j], value)
			}

			newRows[i][j] = value
		}
	}

	return NewTable(m.table.columns, newRows)
}
