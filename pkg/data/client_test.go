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

type SuiteClient struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
}

var cTest = &Client{
	Id:           1,
	FirstNameC:   "Дим",
	LastNameC:    "Иванович",
	MiddleNameC:  "Сергеевич",
	PhoneNumberC: "375251264567",
}

func (s *SuiteClient) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *SuiteClient) SetupSuite() {
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

func TestInitClient(t *testing.T) {
	suite.Run(t, new(SuiteClient))
}

func (s *SuiteClient) TestAddClient() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "clients"`)).
		WithArgs(cTest.FirstNameC, cTest.LastNameC, cTest.MiddleNameC, cTest.PhoneNumberC, cTest.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()
	id, err := s.repository.AddClient(*cTest)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 1)
}
func (s *SuiteClient) TestAddClientErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "clients"`)).
		WithArgs(cTest.FirstNameC, cTest.LastNameC, cTest.MiddleNameC, cTest.PhoneNumberC, cTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.repository.AddClient(*cTest)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *SuiteClient) TestReadAllClients() {
	rows := sqlmock.NewRows([]string{"id", "first_name_c", "last_name_c", "middle_name_c", "phone_number_c"}).
		AddRow(cTest.Id, cTest.FirstNameC, cTest.LastNameC, cTest.MiddleNameC, cTest.PhoneNumberC)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "clients"`)).
		WillReturnRows(rows)
	clients, err := s.repository.ReadAllClients()
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), clients)
	require.Equal(s.T(), clients[0], *cTest)
	require.Len(s.T(), clients, 1)
}

func (s *SuiteClient) TestReadAllClientsErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "clients"`)).
		WillReturnError(errors.New("something went wrong"))
	clients, err := s.repository.ReadAllClients()
	require.Error(s.T(), err)
	require.Empty(s.T(), clients)
}

func (s *SuiteClient) TestDeleteById() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "clients"`)).
		WithArgs(cTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdClient(cTest.Id)
	require.NoError(s.T(), err)
}

func (s *SuiteClient) TestDeleteByIdClientErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "clients"`)).
		WithArgs(cTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdClient(cTest.Id)
	require.Error(s.T(), err)
}

func (s *SuiteClient) TestUpdateClient() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "clients"`)).
		WithArgs(cTest.FirstNameC, cTest.LastNameC, cTest.MiddleNameC, cTest.PhoneNumberC, cTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateClient(*cTest)
	require.NoError(s.T(), err)
}

func (s *SuiteClient) TestUpdateClientErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "clients"`)).
		WithArgs(cTest.FirstNameC, cTest.LastNameC, cTest.MiddleNameC, cTest.PhoneNumberC, cTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.UpdateClient(*cTest)
	require.Error(s.T(), err)
}
