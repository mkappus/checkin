package users

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type StudentStore struct {
	*sql.DB
}

func (ss *StudentStore) All() (students []*Student, err error) {
	rows, err := ss.Query("SELECT DISTINCT  perm_id, student_name, grd  FROM students")
	if err != nil {
		return
	}

	for rows.Next() {
		var s Student
		err = rows.Scan(&s.ID, &s.StudentName, &s.Grade)
		if err != nil {
			continue
		}
		students = append(students, &s)
	}
	return students, err
}
