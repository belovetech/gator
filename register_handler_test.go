package main

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/belovetech/gator.git/internal/database"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mocking the database.Queries
// type mockDB struct {
// 	mock.Mock
// }

// func (m *mockDB) GetUser(ctx context.Context, name string) (database.User, error) {
// 	args := m.Called(ctx, name)
// 	return args.Get(0).(database.User), args.Error(1)
// }

// func (m *mockDB) CreateUser(ctx context.Context, params database.CreateUserParams) (database.User, error) {
// 	args := m.Called(ctx, params)
// 	return args.Get(0).(database.User), args.Error(1)
// }

// // Mocking the config.Config
// type mockConfig struct {
// 	mock.Mock
// }

// func (m *mockConfig) SetUser(name string) {
// 	m.Called(name)
// }

// func TestHandleRegister_UserAlreadyExists(t *testing.T) {
// 	mockDB := new(mockDB)
// 	mockConfig := new(mockConfig)
// 	state := &state{
// 		db:     mockDB, // mockDB now satisfies the database.Queries interface
// 		config: mockConfig,
// 	}

// 	cmd := command{
// 		name: "register",
// 		args: []string{"existing_user"},
// 	}

// 	// Simulate that the user already exists
// 	existingUser := database.User{
// 		ID: uuid.New(),
// 	}

// 	mockDB.On("GetUser", mock.Anything, "existing_user").Return(existingUser, nil)

// 	err := handleRegister(state, cmd)

// 	assert.NotNil(t, err)
// 	assert.EqualError(t, err, "user already exists: existing_user")

// 	// Ensure the user is not created
// 	mockDB.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
// }

// func TestHandleRegister_UserDoesNotExist_Success(t *testing.T) {
// 	mockDB := new(mockDB)
// 	mockConfig := new(mockConfig)
// 	state := &state{
// 		db:     mockDB,
// 		config: mockConfig,
// 	}

// 	cmd := command{
// 		name: "register",
// 		args: []string{"new_user"},
// 	}

// 	// Simulate that the user does not exist
// 	mockDB.On("GetUser", mock.Anything, "new_user").Return(database.User{}, errors.New("sql: no rows in result set"))

// 	// Simulate successful user creation
// 	createdUser := database.User{
// 		ID:        uuid.New(),
// 		Name:      "new_user",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	mockDB.On("CreateUser", mock.Anything, mock.Anything).Return(createdUser, nil)
// 	mockConfig.On("SetUser", "new_user").Return()

// 	err := handleRegister(state, cmd)

// 	assert.Nil(t, err)
// 	mockDB.AssertCalled(t, "CreateUser", mock.Anything, mock.Anything)
// 	mockConfig.AssertCalled(t, "SetUser", "new_user")
// }

// func TestHandleRegister_EmptyUsername(t *testing.T) {
// 	state := &state{}
// 	cmd := command{
// 		name: "register",
// 		args: []string{""},
// 	}

// 	err := handleRegister(state, cmd)

// 	assert.NotNil(t, err)
// 	assert.EqualError(t, err, "user cannot be empty")
// }

// func TestHandleRegister_ErrorCheckingUserExistence(t *testing.T) {
// 	mockDB := new(mockDB)
// 	state := &state{
// 		db: mockDB,
// 	}

// 	cmd := command{
// 		name: "register",
// 		args: []string{"new_user"},
// 	}

// 	// Simulate error in checking user existence
// 	mockDB.On("GetUser", mock.Anything, "new_user").Return(database.User{}, errors.New("some db error"))

// 	err := handleRegister(state, cmd)

// 	assert.NotNil(t, err)
// 	assert.EqualError(t, err, "error checking user existence: some db error")

// 	// Ensure the user is not created
// 	mockDB.AssertNotCalled(t, "CreateUser", mock.Anything, mock.Anything)
// }
