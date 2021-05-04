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

type SuiteRole struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
	person     *Role
}

var rTest = &Role{
	Name:       "Программист",
	DivisionId: 1,
}

func (s *SuiteRole) AfterTest(_, _ string) {
	//require.NoError(s.T(), s.mock.ExpectationsWereMet())
	s.SetupSuite()
}

func (s *SuiteRole) SetupSuite() {
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

func TestInitRole(t *testing.T) {
	suite.Run(t, new(SuiteRole))
}

//func (s *Suite) TestAdd()  {
//	s.mock.ExpectBegin()
//	s.mock.ExpectQuery(regexp.QuoteMeta(
//		`INSERT INTO "roles"`)).
//		WithArgs(rTest.Name, "АСУ").
//		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
//	s.mock.ExpectCommit()
//	id, err := s.repository.Add("testName", "АСУ")
//	require.NoError(s.T(), err)
//	require.Equal(s.T(), id, 1)
//}
func (s *SuiteRole) TestDeleteById() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "roles"`)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdRole(1)
	require.NoError(s.T(), err)
}
func (s *SuiteRole) TestDeleteByIdErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "roles"`)).
		WithArgs(1).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdRole(1)
	require.Error(s.T(), err)
}

//func (s *Suite) TestUpdate() {
//	s.mock.ExpectBegin()
//	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "roles"`)).
//		WithArgs("asd", rTest.Name, rTest.Id).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	s.mock.ExpectCommit()
//	err := s.repository.Update(rTest.Id,"asd", rTest.Name)
//	require.NoError(s.T(), err)
//}
//
//func (s *Suite) TestUpdateServiceErr()  {
//	s.mock.ExpectBegin()
//	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "roles"`)).
//		WithArgs("asd", rTest.Name, rTest.Id).
//		WillReturnError(errors.New("something went wrong"))
//	s.mock.ExpectCommit()
//	err := s.repository.Update(*rTest)
//	require.Error(s.T(), err)
//}
