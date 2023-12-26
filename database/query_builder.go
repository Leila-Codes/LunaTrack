package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type queryBuilder struct {
	params  []interface{}
	query   strings.Builder
	table   string
	columns strings.Builder
}

func NewQueryBuilder(table string) *queryBuilder {
	return &queryBuilder{
		params:  make([]interface{}, 0),
		query:   strings.Builder{},
		table:   table,
		columns: strings.Builder{},
	}
}

func (qb *queryBuilder) Select(columns ...string) {
	for i, col := range columns {
		if i < len(columns) {
			qb.columns.WriteRune(',')
		}
		qb.columns.WriteString(col)
	}
}

func (qb *queryBuilder) Where(column string, value interface{}) *queryBuilder {
	if len(qb.params) == 0 {
		qb.query.WriteString(" WHERE ")
	} else {
		qb.query.WriteString(" AND ")
	}
	qb.params = append(qb.params, value)
	qb.query.WriteString(fmt.Sprintf("%s = $%d", column, len(qb.params)))
	return qb
}

func (qb *queryBuilder) Exec() (*sql.Rows, error) {
	query := fmt.Sprintf(`SELECT %s FROM "%s" WHERE %s`, qb.columns, qb.table, qb.query)
	return db.Query(query, qb.params...)
}
