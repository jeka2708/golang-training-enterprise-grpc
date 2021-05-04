package data

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"regexp"
	"testing"
)

type SuiteDivision struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository *DataEnterprise
}

var dvTest = &Division{
	Id:           1,
	DivisionName: "Кадры",
}

func (s *SuiteDivision) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *SuiteDivision) SetupSuite() {
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

func TestInitDivision(t *testing.T) {
	suite.Run(t, new(SuiteDivision))
}

func (s *SuiteDivision) TestAdd() {
	var name = "tesname"
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "divisions" ("division_name") VALUES ($1) RETURNING "id"`)).
		WithArgs(name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(6))
	s.mock.ExpectCommit()
	id, err := s.repository.AddDivision(name)
	require.NoError(s.T(), err)
	require.Equal(s.T(), id, 6)
}
func (s *SuiteDivision) TestAddErr() {
	var name = "tesname"
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "divisions"`)).
		WithArgs(name).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	id, err := s.repository.AddDivision(name)
	require.Error(s.T(), err)
	require.Equal(s.T(), id, -1)
}

func (s *SuiteDivision) TestReadAll() {
	rows := sqlmock.NewRows([]string{"id", "division_name"}).AddRow(dvTest.Id, dvTest.DivisionName)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "divisions"`)).
		WillReturnRows(rows)
	dv, err := s.repository.ReadAllDivision()
	log.Println(dv)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), dv)
	require.Equal(s.T(), dv[0], *dvTest)
	require.Len(s.T(), dv, 1)
}

func (s *SuiteDivision) TestReadAllErr() {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "divisions"`)).
		WillReturnError(errors.New("something went wrong"))
	dv, err := s.repository.ReadAllDivision()
	require.Error(s.T(), err)
	require.Empty(s.T(), dv)
}
func (s *SuiteDivision) TestDeleteById() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "divisions"`)).
		WithArgs(dvTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdDivision(dvTest.Id)
	require.NoError(s.T(), err)
}
func (s *SuiteDivision) TestDeleteByIdErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "divisions"`)).
		WithArgs(dvTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.DeleteByIdDivision(dvTest.Id)
	require.Error(s.T(), err)
}

func (s *SuiteDivision) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "divisions"`)).
		WithArgs(dvTest.DivisionName, dvTest.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.repository.UpdateDivision(*dvTest)
	require.NoError(s.T(), err)
}

func (s *SuiteDivision) TestUpdateErr() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "divisions"`)).
		WithArgs(dvTest.DivisionName, dvTest.Id).
		WillReturnError(errors.New("something went wrong"))
	s.mock.ExpectCommit()
	err := s.repository.UpdateDivision(*dvTest)
	require.Error(s.T(), err)
}
