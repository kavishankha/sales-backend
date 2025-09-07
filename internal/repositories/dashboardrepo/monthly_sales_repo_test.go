package dashboardrepo_test

import (
	"context"
	"database/sql"
	"echo-app/internal/models"
	"echo-app/internal/repositories/dashboardrepo"
	"echo-app/pkg/database"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMonthlySales_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	fakeData := []models.MonthlySale{
		{Month: "Jan", UnitsSold: 10},
		{Month: "Feb", UnitsSold: 20},
	}
	raw, _ := json.Marshal(map[string]interface{}{"monthlySales": fakeData})

	rows := sqlmock.NewRows([]string{"result"}).AddRow(raw)
	mock.ExpectQuery("SELECT jsonb_build_object").WillReturnRows(rows)

	result, err := dashboardrepo.MonthlySales(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, fakeData, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMonthlySales_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	mock.ExpectQuery("SELECT jsonb_build_object").WillReturnError(sql.ErrConnDone)

	_, err = dashboardrepo.MonthlySales(context.Background())
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMonthlySales_JSONError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	rows := sqlmock.NewRows([]string{"result"}).AddRow([]byte("invalid-json"))
	mock.ExpectQuery("SELECT jsonb_build_object").WillReturnRows(rows)

	_, err = dashboardrepo.MonthlySales(context.Background())
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
