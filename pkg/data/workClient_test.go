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

type SuiteWorkClient struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
}

var wClientsTest = &WorkClient{
	Id:       1,
	ClientId: 1,
	WorkId:   1,
}

func (s *SuiteWorkClient) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *SuiteWorkClient) SetupSuite() {
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

func TestInitWorkClient(t *testing.T) {
	suite.Run(t, new(SuiteWorkClient))
}

func (s *SuiteWorkClient) TestAdd() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "work_clients"`)).
		WithArgs(wClientsTest.ClientId, wClientsTest.WorkId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()
	id, err := s.repository.AddWorkClient(wClientsTest.ClientId, wClientsTest.WorkId)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 1)
}

func (s *SuiteWorkClient) TestAddErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "work_clients"`)).
		WithArgs(wClientsTest.ClientId, wClientsTest.WorkId).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.repository.AddWorkClient(wClientsTest.ClientId, wClientsTest.WorkId)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *SuiteWorkClient) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "work_clients"`)).
		WithArgs(wClientsTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdWorkClient(wClientsTest.Id)
	require.NoError(s.T(), err)
}
func (s *SuiteWorkClient) TestDeleteErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "work_clients"`)).
		WithArgs(wClientsTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdWorkClient(wClientsTest.Id)
	require.Error(s.T(), err)
}

func (s *SuiteWorkClient) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "work_clients"`)).
		WithArgs(wClientsTest.ClientId, wClientsTest.WorkId, wClientsTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateWorkClient(*wClientsTest)
	require.NoError(s.T(), err)
}

func (s *SuiteWorkClient) TestUpdateErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "work_clients"`)).
		WithArgs(wClientsTest.ClientId, wClientsTest.WorkId, wClientsTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.UpdateWorkClient(*wClientsTest)
	require.Error(s.T(), err)
}
