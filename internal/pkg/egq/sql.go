package egq

import "gorm.io/gorm"

type SQLBuilder struct {
	db        *gorm.DB
	condition []string
	args      []interface{}
}

// NewSQLBuilder ...
func NewSQLBuilder(db *gorm.DB) *SQLBuilder {
	return &SQLBuilder{
		db:        db,
		condition: make([]string, 0),
		args:      make([]interface{}, 0),
	}
}

func (b *SQLBuilder) Where(condition string, args ...interface{}) *SQLBuilder {
	return nil
}
