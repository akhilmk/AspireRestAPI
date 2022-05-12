package service

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/akhilmk/GoRESTAPI/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository IUserRepo
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		Conn:       db,
		DriverName: "postgres",
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	//s.DB.lo(true)

	s.repository = GetUserRepo(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestAddUser() {

	user := model.User{UserID: 1, UserName: "test", Email: "test@mail.com", Gender: "male", Age: 30}
	row := sqlmock.NewRows([]string{"user_id"}).AddRow(int(1))

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs("test", "test@mail.com", "male", 30, 1).
		WillReturnRows(row)
	s.mock.ExpectCommit()
	err := s.repository.AddUser(user)

	require.NoError(s.T(), err)
}
