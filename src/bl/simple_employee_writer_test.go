package bl

import (
	"aiplus_demo/src/da"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockValidator *MockEmployeeValidator
var mockRepo *da.MockEmployeeRepository
var mockDb *da.MockDb
var mockTx *da.MockTx

func TestWriteOk(t *testing.T) {
	initMocks()
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockRepo.On("IsEmployeeWithNumberExist", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	mockRepo.On("CreateEmployee", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockTx.On("Commit").Return(nil)
	mockTx.On("Rollback").Return(nil)
	mockDb.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, nil)
	writer := SimpleEmployeeWriter{
		validator: mockValidator,
		repo:      mockRepo,
		db:        mockDb,
	}

	isSuccess, err := writer.Write(context.TODO(), da.Employee{})
	assert.Nil(t, err)
	assert.True(t, isSuccess)
}

func TestWriteValidationFail(t *testing.T) {
	initMocks()
	mockValidator.On("Validate", mock.Anything).Return(NewValidationError("Some.Target", "some reason"))
	mockRepo.On("IsEmployeeWithNumberExist", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	mockRepo.On("CreateEmployee", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockTx.On("Commit").Return(nil)
	mockTx.On("Rollback").Return(nil)
	mockDb.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, nil)
	writer := SimpleEmployeeWriter{
		validator: mockValidator,
		repo:      mockRepo,
		db:        mockDb,
	}

	isSuccess, err := writer.Write(context.TODO(), da.Employee{})
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, err.ErrType, BlValidationError)
	}
	assert.False(t, isSuccess)
}

func TestWriteBeginTxFail(t *testing.T) {
	initMocks()
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockRepo.On("IsEmployeeWithNumberExist", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	mockRepo.On("CreateEmployee", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockTx.On("Commit").Return(nil)
	mockTx.On("Rollback").Return(nil)
	mockDb.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, errors.New("some reason"))
	writer := SimpleEmployeeWriter{
		validator: mockValidator,
		repo:      mockRepo,
		db:        mockDb,
	}

	isSuccess, err := writer.Write(context.TODO(), da.Employee{})
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, err.ErrType, BlDefaultError)
	}
	assert.False(t, isSuccess)
}

func TestWriteBusinessErr(t *testing.T) {
	initMocks()
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockRepo.On("IsEmployeeWithNumberExist", mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
	mockRepo.On("CreateEmployee", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockTx.On("Commit").Return(nil)
	mockTx.On("Rollback").Return(nil)
	mockDb.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, nil)
	writer := SimpleEmployeeWriter{
		validator: mockValidator,
		repo:      mockRepo,
		db:        mockDb,
	}

	isSuccess, err := writer.Write(context.TODO(), da.Employee{})
	assert.Nil(t, err)
	assert.False(t, isSuccess)
}

func TestWriteCreateErr(t *testing.T) {
	initMocks()
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockRepo.On("IsEmployeeWithNumberExist", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	mockRepo.On("CreateEmployee", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some reason"))
	mockTx.On("Commit").Return(nil)
	mockTx.On("Rollback").Return(nil)
	mockDb.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, nil)
	writer := SimpleEmployeeWriter{
		validator: mockValidator,
		repo:      mockRepo,
		db:        mockDb,
	}

	isSuccess, err := writer.Write(context.TODO(), da.Employee{})
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, err.ErrType, BlDefaultError)
	}
	assert.False(t, isSuccess)
}

func TestWriteCommitFail(t *testing.T) {
	initMocks()
	mockValidator.On("Validate", mock.Anything).Return(nil)
	mockRepo.On("IsEmployeeWithNumberExist", mock.Anything, mock.Anything, mock.Anything).Return(false, nil)
	mockRepo.On("CreateEmployee", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mockTx.On("Commit").Return(errors.New("some reason"))
	mockTx.On("Rollback").Return(nil)
	mockDb.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, nil)
	writer := SimpleEmployeeWriter{
		validator: mockValidator,
		repo:      mockRepo,
		db:        mockDb,
	}

	isSuccess, err := writer.Write(context.TODO(), da.Employee{})
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, err.ErrType, BlDefaultError)
	}
	assert.False(t, isSuccess)
}

func initMocks() {
	mockValidator = new(MockEmployeeValidator)
	mockRepo = new(da.MockEmployeeRepository)
	mockDb = new(da.MockDb)
	mockTx = new(da.MockTx)
}
