package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type MockRepo struct {
	CreateFunc  func(any) error
	ReadAllFunc func(any) error
	UpdateFunc  func(any) error
	DeleteFunc  func(any) error
}

func (m *MockRepo) Migrate(any) error {
	return nil
}

func (m *MockRepo) Create(user any) error {
	return m.CreateFunc(user)
}

func (m *MockRepo) ReadAll(users any) error {
	return m.ReadAllFunc(users)
}

func (m *MockRepo) Update(user any, _ uint) error {
	return m.UpdateFunc(user)
}

func (m *MockRepo) Delete(user any, _ uint) error {
	return m.DeleteFunc(user)
}

func TestUserHandler_Create(t *testing.T) {
	type fields struct {
		repo   Repo
		router *http.ServeMux
	}
	tests := []struct {
		name           string
		fields         fields
		input          any
		expectedStatus int
	}{
		{
			name: "Successful creation",
			fields: fields{
				repo: &MockRepo{
					CreateFunc: func(user any) error {
						return nil
					},
				},
				router: http.NewServeMux(),
			},
			input:          User{Username: "Username"},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Failed creation",
			fields: fields{
				repo: &MockRepo{
					CreateFunc: func(user any) error {
						return errors.New("creation failed")
					},
				},
				router: http.NewServeMux(),
			},
			input:          User{Username: ""}, // поскольку мы не делаем валидацию, то принудительно возвращаем ошибку БД
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				repo:   tt.fields.repo,
				router: tt.fields.router,
			}

			body, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			u.create()(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestUserHandler_ReadAll(t *testing.T) {
	type fields struct {
		repo   Repo
		router *http.ServeMux
	}
	tests := []struct {
		name           string
		fields         fields
		expectedUsers  []User
		expectedStatus int
		expectedError  error
	}{
		{
			name: "Successful read all users",
			fields: fields{
				repo: &MockRepo{
					ReadAllFunc: func(users any) error {
						userList := users.(*[]User)
						*userList = []User{
							{Username: "User1"},
							{Username: "User2"},
						}
						return nil
					},
				},
				router: http.NewServeMux(),
			},
			expectedUsers:  []User{{Username: "User1"}, {Username: "User2"}},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Failed read all users",
			fields: fields{
				repo: &MockRepo{
					ReadAllFunc: func(users any) error {
						return errors.New("read failed")
					},
				},
				router: http.NewServeMux(),
			},
			expectedUsers:  nil,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				repo:   tt.fields.repo,
				router: tt.fields.router,
			}

			req := httptest.NewRequest("GET", "/user", nil)
			rr := httptest.NewRecorder()

			u.readAll()(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusOK {
				var users []User
				if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if !reflect.DeepEqual(users, tt.expectedUsers) {
					t.Errorf("handler returned unexpected body: got %v want %v", users, tt.expectedUsers)
				}
			}
		})
	}
}

func TestUserHandler_Update(t *testing.T) {
	type fields struct {
		repo   Repo
		router *http.ServeMux
	}
	tests := []struct {
		name           string
		fields         fields
		input          User
		expectedStatus int
		expectedError  error
	}{
		{
			name: "Successful update user",
			fields: fields{
				repo: &MockRepo{
					UpdateFunc: func(user any) error {
						return nil
					},
				},
				router: http.NewServeMux(),
			},
			input:          User{Username: "User1"},
			expectedStatus: http.StatusNoContent,
		},
		{
			name: "Failed update user",
			fields: fields{
				repo: &MockRepo{
					UpdateFunc: func(user any) error {
						return errors.New("update failed")
					},
				},
				router: http.NewServeMux(),
			},
			input:          User{Username: ""}, //пустое имя будем считать недопустимым для БД, принудительная ошибка.
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				repo:   tt.fields.repo,
				router: tt.fields.router,
			}

			body, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			req := httptest.NewRequest("PUT", "/user/", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			u.update()(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestUserHandler_Delete(t *testing.T) {
	type fields struct {
		repo   Repo
		router *http.ServeMux
	}
	tests := []struct {
		name           string
		fields         fields
		input          User
		expectedStatus int
		expectedError  error
	}{
		{
			name: "Successful delete user",
			fields: fields{
				repo: &MockRepo{
					DeleteFunc: func(user any) error {
						return nil
					},
				},
				router: http.NewServeMux(),
			},
			input:          User{Username: "User1"},
			expectedStatus: http.StatusNoContent,
		},
		{
			name: "Failed delete user",
			fields: fields{
				repo: &MockRepo{
					DeleteFunc: func(user any) error {
						return errors.New("delete failed")
					},
				},
				router: http.NewServeMux(),
			},
			input:          User{Username: ""},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserHandler{
				repo:   tt.fields.repo,
				router: tt.fields.router,
			}

			body, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal input: %v", err)
			}

			req := httptest.NewRequest("DELETE", "/user", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			u.delete()(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
