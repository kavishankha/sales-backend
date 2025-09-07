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

func TestCountryLevelRevenue_Success(t *testing.T) {
	// Create sqlmock DB
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	// Fake JSON result
	fakeData := []models.CountryReport{
		{
			Country:      "Sri Lanka",
			Products:     []string{"Tea", "Cinnamon"},
			TotalRevenue: 10000,
			Transactions: 5,
		},
		{
			Country:      "India",
			Products:     []string{"Rice"},
			TotalRevenue: 8000,
			Transactions: 3,
		},
	}
	rawJSON, _ := json.Marshal(map[string]interface{}{
		"data": fakeData,
	})

	// Mock expected query + result
	rows := sqlmock.NewRows([]string{"result"}).AddRow(rawJSON)
	mock.ExpectQuery("SELECT jsonb_build_object").
		WillReturnRows(rows)

	// Call repository
	ctx := context.Background()
	result, err := dashboardrepo.CountryLevelRevenue(ctx)

	assert.NoError(t, err)
	assert.Equal(t, fakeData, result)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryLevelRevenue_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	// Force DB error
	mock.ExpectQuery("SELECT jsonb_build_object").
		WillReturnError(sql.ErrConnDone)

	ctx := context.Background()
	_, err = dashboardrepo.CountryLevelRevenue(ctx)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCountryLevelRevenue_JSONError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.DB = db

	// Return invalid JSON
	rows := sqlmock.NewRows([]string{"result"}).AddRow([]byte("invalid-json"))
	mock.ExpectQuery("SELECT jsonb_build_object").
		WillReturnRows(rows)

	ctx := context.Background()
	_, err = dashboardrepo.CountryLevelRevenue(ctx)

	assert.Error(t, err) // should fail JSON unmarshal
	assert.NoError(t, mock.ExpectationsWereMet())
}
