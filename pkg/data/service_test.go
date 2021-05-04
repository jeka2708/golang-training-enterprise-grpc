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

type SuiteService struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
}

var sTest = &Service{
	Id:   1,
	Name: "Test",
	Cost: 100,
}

func (s *SuiteService) AfterTest(_, _ string) {
	//require.NoError(s.T(), s.mock.ExpectationsWereMet())
	s.SetupSuite()
}

func (s *SuiteService) SetupSuite() {
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

func TestInitService(t *testing.T) {
	suite.Run(t, new(SuiteService))
}

func (s *SuiteService) TestAddService() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "services"`)).
		WithArgs(sTest.Name, sTest.Cost, sTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()
	id, err := s.repository.AddService(*sTest)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 1)
}

func (s *SuiteService) TestAddServiceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "services"`)).
		WithArgs(sTest.Name, sTest.Cost).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.repository.AddService(*sTest)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *SuiteService) TestReadAllServices() {
	rows := sqlmock.NewRows([]string{"id", "name", "cost"}).
		AddRow(sTest.Id, sTest.Name, sTest.Cost)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "services"`)).
		WillReturnRows(rows)
	clients, err := s.repository.ReadAllServices()
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), clients)
	require.Equal(s.T(), clients[0], *sTest)
	require.Len(s.T(), clients, 1)
}

func (s *SuiteService) TestReadAllServicesErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "services"`)).
		WillReturnError(errors.New("something went wrong"))
	clients, err := s.repository.ReadAllServices()
	require.Error(s.T(), err)
	require.Empty(s.T(), clients)
}
func (s *SuiteService) TestDeleteService() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "services"`)).
		WithArgs(sTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdService(sTest.Id)
	require.NoError(s.T(), err)
}
func (s *SuiteService) TestDeleteServiceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "services"`)).
		WithArgs(sTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdService(sTest.Id)
	require.Error(s.T(), err)
}

func (s *SuiteService) TestUpdateService() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "services"`)).
		WithArgs(sTest.Id, sTest.Name, sTest.Cost, sTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateService(*sTest)
	require.NoError(s.T(), err)
}

func (s *SuiteService) TestUpdateServiceErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "services"`)).
		WithArgs(sTest.Id, sTest.Name, sTest.Cost, sTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.UpdateService(*sTest)
	require.Error(s.T(), err)
}
