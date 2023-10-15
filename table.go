package manitable

type Table struct {
	columns []string
	rows    [][]string
}

func NewTable(columns []string, rows [][]string) Table {
	return Table{columns: columns, rows: rows}
}

func (t Table) Mutate() TableMutation {
	return newTableMutation(t, []func(column string, value string) string{})
}

// TableMutation

type TableMutation struct {
	table     Table
	mutations []func(column string, value string) string
}

func newTableMutation(table Table, mutations []func(column string, value string) string) TableMutation {
	return TableMutation{table: table, mutations: mutations}
}

func (m TableMutation) Append(f func(column string, value string) string) TableMutation {
	return newTableMutation(m.table, append(m.mutations, f))
}

func (m TableMutation) Run() Table {
	rows := m.table.rows
	newRows := make([][]string, len(rows))
	for i, row := range rows {
		newRows[i] = make([]string, len(row))
		for j, value := range row {
			for _, mutation := range m.mutations {
				value = mutation(m.table.columns[j], value)
			}

			newRows[i][j] = value
		}
	}

	return NewTable(m.table.columns, newRows)
}
