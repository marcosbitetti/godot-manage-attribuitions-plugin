package infra

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/domain"
	"github.com/pkg/errors"

	_ "github.com/mattn/go-sqlite3"
)

const AtribuitionModelPath = "ATRIBUITION_HANDLER_PATH"

type StorageInterface interface {
	CloseDatabase()
	AddType(name string) error
	UpdateType(id int64, name string) error
	DeleteType(id int64) error
	ListTypes() ([]domain.Type, error)
	AddLicence(name string, link string) error
	UpdateLicence(id int64, name string, link string) error
	DeleteLicence(id int64) error
	ListLicences() ([]domain.Licence, error)
	AddAttribuition(name string, fileame string, author string, link string, ctype string, licence string) error
	FindAttribuitions(ascDesc string, search string) ([]domain.Attribuition, error)
	UpdateAttribuition(id int64, name string, fileame string, author string, link string, ctype string, licence string) error
	DeleteAttribuition(id int64) error
}

type Storage struct {
	db     *sql.DB
	locker sync.Mutex
}

var _ StorageInterface = &Storage{db: nil}

func NewStorage(path string) (*Storage, error) {
	var err error
	var needToInit = false
	_, err = os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			handler, err := os.Create(path)
			if err != nil {
				return nil, errors.Wrap(err, "error to create database file")
			}
			if err := handler.Close(); err != nil {
				return nil, errors.Wrap(err, "error prepare db file")
			}
			needToInit = true
		} else {
			return nil, errors.Wrap(err, "error to acess/create database file")
		}
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}

	storage := &Storage{
		db: db,
	}

	if needToInit {
		initDatabase(storage)
	}
	return storage, nil
}

func (s *Storage) CloseDatabase() {
	s.locker.Lock()
	defer s.locker.Unlock()

	if s.db == nil {
		return
	}
	if err := s.db.Close(); err != nil {
		panic(errors.Wrap(err, "error closing database").Error())
	}

}

func ParseDatabasePath(args []string) (string, error) {
	file := args[1]
	return file, nil
}

func initDatabase(storage *Storage) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createBaseTable(ctx, storage)
	dumpFirstTypes(storage)
	dumpFirstLicences(storage)
}

func (s *Storage) AddType(name string) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`
		INSERT INTO types(name) VALUES(?);
	`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to add type")
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to add type").Error())
		}
	}()

	_, err = stmt.Exec(name)
	if err != nil {
		return errors.Wrap(err, "cant exec to add type")
	}
	return nil
}

func (s *Storage) UpdateType(id int64, name string) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`
		UPDATE types SET name=? WHERE _id=?
	`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to update type")
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to add type").Error())
		}
	}()

	_, err = stmt.Exec(name, id)
	if err != nil {
		return errors.Wrap(err, "cant exec to update type")
	}
	return nil
}

func (s *Storage) DeleteType(id int64) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`DELETE FROM types WHERE _id = ?`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to delete type")
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to delete type").Error())
		}
	}()
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "cant exec to delete type")
	}
	return nil
}

func (s *Storage) ListTypes() ([]domain.Type, error) {
	s.locker.Lock()
	defer s.locker.Unlock()

	list := make([]domain.Type, 0)
	rows, err := s.db.Query(`
		SELECT _id, name FROM types ORDER BY name COLLATE NOCASE ASC
	`)
	if err != nil {
		return nil, errors.Wrap(err, "cant read rows from types")
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(errors.Wrap(err, "cant close rows from types").Error())
		}
	}()
	for rows.Next() {
		data := domain.Type{}
		if err := rows.Scan(&data.Id, &data.Name); err != nil {
			return nil, errors.Wrap(err, "cant read row from types")
		}
		list = append(list, data)
	}
	return list, nil
}

func (s *Storage) AddLicence(name string, link string) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`
		INSERT INTO licences(name, link) VALUES(?, ?);
	`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to add licence")
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to add licence").Error())
		}
	}()

	_, err = stmt.Exec(name, link)
	if err != nil {
		return errors.Wrap(err, "cant exec to add Licence")
	}
	return nil
}

func (s *Storage) UpdateLicence(id int64, name string, link string) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`
		UPDATE licences SET name=?, link=? WHERE _id=?
	`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to update licence")
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to add licence").Error())
		}
	}()

	_, err = stmt.Exec(name, link, id)
	if err != nil {
		return errors.Wrap(err, "cant exec to update licence")
	}
	return nil
}

func (s *Storage) DeleteLicence(id int64) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`DELETE FROM licences WHERE _id = ?`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to delete licence")
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to delete licence").Error())
		}
	}()
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "cant exec to delete licence")
	}
	return nil
}

func (s *Storage) ListLicences() ([]domain.Licence, error) {
	s.locker.Lock()
	defer s.locker.Unlock()

	list := make([]domain.Licence, 0)
	rows, err := s.db.Query(`
		SELECT _id, name, link FROM licences ORDER BY name COLLATE NOCASE ASC
	`)
	if err != nil {
		return nil, errors.Wrap(err, "cant read rows from licences")
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(errors.Wrap(err, "cant close rows from licences").Error())
		}
	}()
	for rows.Next() {
		data := domain.Licence{}
		if err := rows.Scan(&data.Id, &data.Name, &data.Link); err != nil {
			return nil, errors.Wrap(err, "cant read row from licences")
		}
		list = append(list, data)
	}
	return list, nil
}

func (s *Storage) AddAttribuition(name string, fileame string, author string, link string, ctype string, licence string) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`
		INSERT InTO credits
		(name, filename, author, link, type_id, licence_id)
		VALUES
		(?, ?, ?, ?,
			(SELECT _id FROM types WHERE name=?),
			(SELECT _id FROM licences WHERE name=?)
		)
	`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to add attribuition")
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to add attribuition").Error())
		}
	}()
	_, err = stmt.Exec(name, fileame, author, link, ctype, licence)
	if err != nil {
		return errors.Wrap(err, "cant exec to add attribuition")
	}
	return nil
}

func (s *Storage) FindAttribuitions(ascDesc string, search string) ([]domain.Attribuition, error) {
	s.locker.Lock()
	defer s.locker.Unlock()

	list := make([]domain.Attribuition, 0)
	whereClause, args := mountQueryWhere(search)
	query := fmt.Sprintf(`
		SELECT c._id, c.name, filename, author, c.link,
			t.name as type,
			l.name as licence,
			l.link as licence_link
		FROM credits c
		LEFT JOIN types t ON t._id = c.type_id
		LEFT JOIN licences l ON l._id = c.licence_id
		%s
		ORDER BY c.name COLLATE NOCASE %s
	`, whereClause, ascDesc)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "cant read rows from attribuitions")
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(errors.Wrap(err, "cant close rows from attribuitions").Error())
		}
	}()
	for rows.Next() {
		data := domain.Attribuition{}
		if err := rows.Scan(&data.Id, &data.Name, &data.FileName, &data.Author, &data.Link, &data.Type, &data.Licence, &data.LicenceUrl); err != nil {
			return nil, errors.Wrap(err, "cant read row from attribuitions")
		}
		list = append(list, data)
	}
	return list, nil
}

func mountQueryWhere(q string) (string, []interface{}) {
	if q == "" {
		return "", nil
	}
	tokens := strings.Fields(q)
	joined := "%" + strings.Join(tokens, "%") + "%"
	return "WHERE c.name LIKE ? OR c.author LIKE ?", []interface{}{joined, joined}
}

func (s *Storage) UpdateAttribuition(id int64, name string, fileame string, author string, link string, ctype string, licence string) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`
		UPDATE credits SET
			name=?,
			filename=?,
			author=?,
			link=?,
			type_id=(SELECT _id FROM types WHERE name=?),
			licence_id=(SELECT _id FROM licences WHERE name=?)
		WHERE _id = ?
	`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to add attribuition")
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to add attribuition").Error())
		}
	}()
	_, err = stmt.Exec(name, fileame, author, link, ctype, licence, id)
	if err != nil {
		return errors.Wrap(err, "cant exec to add attribuition")
	}
	return nil
}

func (s *Storage) DeleteAttribuition(id int64) error {
	s.locker.Lock()
	defer s.locker.Unlock()

	stmt, err := s.db.Prepare(`DELETE FROM credits WHERE _id = ?`)
	if err != nil {
		return errors.Wrap(err, "cant prepare to delete attribuition")
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			panic(errors.Wrap(err, "cant close prepare to delete attribuition").Error())
		}
	}()
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "cant exec to delete attribuition")
	}
	return nil
}
