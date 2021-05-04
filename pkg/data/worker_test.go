package data

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SuiteWorker struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
	person     *Worker
}

var wcTest = &Worker{
	Id:          1,
	FirstName:   "Дим",
	LastName:    "Иванович",
	MiddleName:  "Сергеевич",
	PhoneNumber: "375251264567",
	RoleId:      1,
}

func (s *SuiteWorker) AfterTest(_, _ string) {
	//require.NoError(s.T(), s.mock.ExpectationsWereMet())
	s.SetupSuite()
}

func (s *SuiteWorker) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = NewDataEnterprise(s.DB)
}

func TestInitWorker(t *testing.T) {
	suite.Run(t, new(SuiteWorker))
}

func (s *SuiteWorker) TestAdd() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "workers"`)).
		WithArgs(wcTest.FirstName, wcTest.LastName, wcTest.MiddleName, wcTest.PhoneNumber, wcTest.RoleId, wcTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()
	id, err := s.repository.AddWorker(*wcTest)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 1)
}
func (s *SuiteWorker) TestAddErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "workers"`)).
		WithArgs(wcTest.FirstName, wcTest.LastName, wcTest.MiddleName, wcTest.PhoneNumber, wcTest.RoleId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.repository.AddWorker(*wcTest)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

//func (s *SuiteWorker) TestReadAll() {
//	rows := sqlmock.NewRows([]string{"first_name", "last_name", "middle_name", "phone_number", "role_id"}).
//		AddRow(wcTest.FirstName, wcTest.LastName, wcTest.MiddleName, wcTest.PhoneNumber, wcTest.RoleId)
//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`SELECT * FROM "workers"`)).
//		WillReturnRows(rows)
//	clients, err := s.repository.ReadAll()
//	require.NoError(s.T(), err)
//	require.NotEmpty(s.T(), clients)
//	require.Equal(s.T(), clients[0], *wcTest)
//	require.Len(s.T(), clients, 1)
//}
//
//func (s *SuiteWorker) TestReadAllErr() {
//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`SELECT * FROM "workers"`)).
//		WillReturnError(errors.New("something went wrong"))
//	clients, err := s.repository.ReadAll()
//	require.Error(s.T(), err)
//	require.Empty(s.T(), clients)
//}

func (s *SuiteWorker) TestDeleteById() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "workers"`)).
		WithArgs(wcTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdWorker(wcTest.Id)
	require.NoError(s.T(), err)
}

func (s *SuiteWorker) TestDeleteByIdErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "workers"`)).
		WithArgs(wcTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdWorker(wcTest.Id)
	require.Error(s.T(), err)
}

func (s *SuiteWorker) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "workers"`)).
		WithArgs(wcTest.FirstName, wcTest.LastName, wcTest.MiddleName, wcTest.PhoneNumber, wcTest.RoleId, wcTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateWorker(*wcTest)
	require.NoError(s.T(), err)
}

func (s *SuiteWorker) TestUpdateErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "workers"`)).
		WithArgs(wcTest.FirstName, wcTest.LastName, wcTest.MiddleName, wcTest.PhoneNumber, wcTest.RoleId, wcTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.UpdateWorker(*wcTest)
	require.Error(s.T(), err)
}
