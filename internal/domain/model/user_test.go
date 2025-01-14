package model

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args struct {
			id        int
			companyID int
			name      string
			email     string
			password  string
		}
		wantErr bool
	}{
		{
			name: "valid user creation",
			args: struct {
				id        int
				companyID int
				name      string
				email     string
				password  string
			}{
				id:        1,
				companyID: 123,
				name:      "John Doe",
				email:     "john@example.com",
				password:  "securepassword",
			},
			wantErr: false,
		},
		{
			name: "empty password",
			args: struct {
				id        int
				companyID int
				name      string
				email     string
				password  string
			}{
				id:        2,
				companyID: 456,
				name:      "Jane Doe",
				email:     "jane@example.com",
				password:  "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewUser(tt.args.id, tt.args.companyID, tt.args.name, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_hashPassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args struct {
			password string
		}
		wantErr bool
	}{
		{
			name: "valid password",
			args: struct {
				password string
			}{
				password: "validpassword",
			},
			wantErr: false,
		},
		{
			name: "empty password",
			args: struct {
				password string
			}{
				password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := hashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_VerifyPassword(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		fields struct {
			ID        int
			CompanyID int
			Name      string
			Email     string
			Password  string
			CreatedAt time.Time
			UpdatedAt time.Time
		}
		args struct {
			password string
		}
		wantErr bool
	}{
		{
			name: "correct password",
			fields: struct {
				ID        int
				CompanyID int
				Name      string
				Email     string
				Password  string
				CreatedAt time.Time
				UpdatedAt time.Time
			}{
				ID:        1,
				CompanyID: 123,
				Name:      "John Doe",
				Email:     "john@example.com",
				Password:  "hashedpassword",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			args: struct {
				password string
			}{
				password: "hashedpassword",
			},
			wantErr: false,
		},
		{
			name: "incorrect password",
			fields: struct {
				ID        int
				CompanyID int
				Name      string
				Email     string
				Password  string
				CreatedAt time.Time
				UpdatedAt time.Time
			}{
				ID:        2,
				CompanyID: 456,
				Name:      "Jane Doe",
				Email:     "jane@example.com",
				Password:  "hashedpassword",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			args: struct {
				password string
			}{
				password: "wrongpassword",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &User{
				ID:        tt.fields.ID,
				CompanyID: tt.fields.CompanyID,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := u.VerifyPassword(tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("User.VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
