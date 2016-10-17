package events

import (
	"database/sql"
	"log"
	"time"

	"github.com/mkappus/checkin/src/users"

	_ "github.com/mattn/go-sqlite3"
)

type EventStore struct {
	*sql.DB
}

// Queries for checkin events
const (
	//ALL_CHECKINS_QUERY = "select s.perm_id, s.grd, s.student_name, c.id, c.loc, c.start, c.end, c.is_active from students as s inner join checkins as c using(perm_id) group by perm_id"
	ALL_CHECKINS_QUERY = "select distinct s.perm_id, s.grd, s.student_name, c.id, c.loc, c.start, c.end, c.is_active from students as s inner join checkins as c using(perm_id)"
	CHECKOUT_QUERY     = "update checkins set end=?, is_active=? where id=?"
	ADD_CHECKIN_QUERY  = "insert into checkins(perm_id, loc, start, end, is_active) values(?,?,?,?,?)"
)

// CREATE TABLE checkins (id integer primary key autoincrement, student_id text, loc text, start datetime, end datetime, is_active integer, foreign key(student_id) references students(perm_id));

func (es *EventStore) AllCheckins() (events []*Checkin, err error) {
	rows, err := es.Query(ALL_CHECKINS_QUERY)
	if err != nil {
		return events, err
	}

	for rows.Next() {
		var c Checkin
		var s users.Student
		err = rows.Scan(&s.ID, &s.Grade, &s.StudentName, &c.ID, &c.Loc, &c.Start, &c.End, &c.IsActive)
		if err != nil {
			continue
		}
		c.User = &s
		events = append(events, &c)

	}
	log.Println("length of events", len(events))
	return events, nil
}

// AddCheckin inserts checkin to databases
func (es *EventStore) AddCheckin(s *users.Student, loc string) (*Checkin, error) {
	stmt, err := es.Prepare(ADD_CHECKIN_QUERY)
	if err != nil {
		return nil, err
	}
	c := &Checkin{
		User:     s,
		Loc:      loc,
		Start:    time.Now(),
		End:      time.Now(),
		IsActive: true,
	}
	res, err := stmt.Exec(s.ID, c.Loc, c.Start, c.End, c.IsActive)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	c.ID = int(id)
	return c, nil

}

func (es *EventStore) Checkout(cID int) error {
	stmt, err := es.Prepare(CHECKOUT_QUERY)
	if err != nil {
		return err
	}
	//end=?, is_active=? where id=?
	_, err = stmt.Exec(time.Now(), false, cID)
	return err
}

