package manitable

type Table struct {
	columns []string
	rows    [][]string
}

func NewTable(columns []string, rows [][]string) Table {
	return Table{columns: columns, rows: rows}
}

func (t Table) Columns() []string {
	return t.columns
}

func (t Table) Rows() [][]string {
	return t.rows
}
