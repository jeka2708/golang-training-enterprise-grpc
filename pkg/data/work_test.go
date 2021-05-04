package data

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type SuiteWork struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
}

var wTest = &Work{
	Id:        1,
	WorkerId:  1,
	ServiceId: 1,
}

func (s *SuiteWork) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *SuiteWork) SetupSuite() {
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

func TestInitWork(t *testing.T) {
	suite.Run(t, new(SuiteWork))
}

func (s *SuiteWork) TestAddWork() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "works"`)).
		WithArgs(wTest.WorkerId, wTest.ServiceId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()
	id, err := s.repository.AddWork(wTest.WorkerId, wTest.ServiceId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 1)
}

func (s *SuiteWork) TestAddWorkErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "works"`)).
		WithArgs(wTest.WorkerId, wTest.ServiceId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.repository.AddWork(wTest.WorkerId, wTest.ServiceId)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *SuiteWork) TestDeleteWork() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "works"`)).
		WithArgs(wTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdWork(wTest.Id)
	require.NoError(s.T(), err)
}
func (s *SuiteWork) TestDeleteWorkErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "works"`)).
		WithArgs(wTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdWork(wTest.Id)
	require.Error(s.T(), err)
}

func (s *SuiteWork) TestUpdateWork() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "works"`)).
		WithArgs(wTest.Id, wTest.WorkerId, wTest.ServiceId, wTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateWork(*wTest)
	require.NoError(s.T(), err)
}

func (s *SuiteWork) TestUpdateWorkErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "works"`)).
		WithArgs(wTest.Id, wTest.WorkerId, wTest.ServiceId, wTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.UpdateWork(*wTest)
	require.Error(s.T(), err)
}

//func (s *SuiteWork) TestReadAll() {
//	rows := sqlmock.NewRows([]string{"worker_id", "services_id"}).
//		AddRow(wTest.WorkerId, wTest.ServiceId)
//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`SELECT * FROM "works"`)).
//		WillReturnRows(rows)
//	works, err := s.repository.ReadAll()
//	require.NoError(s.T(),err)
//	require.NotEmpty(s.T(),works)
//	require.Equal(s.T(), works[0], *wTest)
//	require.Len(s.T(), works, 1)
//}

//func (s Suite) TestReadAllErr()  {
//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`SELECT * FROM "services"`)).
//		WillReturnError(errors.New("something went wrong"))
//	clients, err := s.repository.ReadAll()
//	require.Error(s.T(),err)
//	require.Empty(s.T(),clients)
//}
