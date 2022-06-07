package service

import (
	model "HW/app/models"
	"HW/app/repo"
	"errors"
	"reflect"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	type fields struct {
		repo repo.UserRepositoryInterface
	}
	type args struct {
		email    string
		password string
	}
	mockRepo := &mockUserRepository{
		items: []model.User{
			{ID: 2,
				Name:     "Ahmet",
				Email:    "ahmetbayrak@kgstore.com",
				Password: "456789",
				Phone:    "05355816760",
				Roles: []*model.Role{
					{
						Name: "customer",
					},
				}},
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.User
	}{
		{
			name: "userGet_Success",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				email:    "ahmetbayrak@kgstore.com",
				password: "456789",
			},
			want: &model.User{
				ID:       2,
				Name:     "Ahmet",
				Email:    "ahmetbayrak@kgstore.com",
				Password: "456789",
				Phone:    "05355816760",
				Roles: []*model.Role{
					{
						Name: "customer",
					},
				},
			},
		},
		{
			name: "userGet_Fail",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				email:    "kadirgönen@kgstore.com",
				password: "456786",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.repo.GetUserEmailPassword(tt.args.email, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserHasRole(t *testing.T) {
	type fields struct {
		repo repo.UserRepositoryInterface
	}
	type args struct {
		id   int
		role string
	}
	mockRepo := &mockUserRepository{
		items: []model.User{
			{ID: 3,
				Name:     "Ahmet",
				Email:    "ahmetbayrak@kgstore.com",
				Password: "456789",
				Phone:    "05355816760",
				Roles: []*model.Role{
					{
						Name: "customer",
					},
				},
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
		result  bool
	}{
		{
			name: "userHasRole_Success",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				id:   3,
				role: "customer",
			},
			want: &model.User{
				ID:       3,
				Name:     "Ahmet",
				Email:    "ahmetbayrak@kgstore.com",
				Password: "456789",
				Phone:    "05355816760",
				Roles: []*model.Role{
					{
						Name: "customer",
					},
				},
			},
			wantErr: true,
			result:  true,
		},
		{
			name: "userHasRole_Fail",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				id:   2,
				role: "customer",
			},
			want: &model.User{
				ID:       2,
				Name:     "Ahmet",
				Email:    "ahmetbayrak@kgstore.com",
				Password: "456789",
				Phone:    "05355816760",
				Roles: []*model.Role{
					{
						Name: "customer",
					},
				},
			},
			wantErr: false,
			result:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.repo.GetUserWithRoles(tt.args.id) //got != tt.wantErr {
			if got == nil {
				//t.Errorf("UserService.UserHasRole() = %v, want %v", got, tt.want)
				tt.wantErr = false
			} else {
				for i := range got.Roles {
					if got.Roles[i].Name == tt.args.role {
						tt.wantErr = true
					}
				}
			}
			if tt.wantErr != tt.result {
				t.Errorf("UserService.UserHasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		repo repo.UserRepositoryInterface
	}
	type args struct {
		user *model.User
	}
	mockRepo := &mockUserRepository{
		items: []model.User{
			{ID: 3,
				Name:     "Ahmet",
				Email:    "ahmetbayrak@kgstore.com",
				Password: "456789",
				Phone:    "05355816760",
				Roles: []*model.Role{
					{
						Name: "customer",
					},
				},
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "userCreateUser_Fail",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				user: &model.User{
					ID:       3,
					Name:     "Ahmet",
					Email:    "ahmetbayrak@kgstore.com",
					Password: "456789",
					Phone:    "05355816760",
					Roles: []*model.Role{
						{
							Name: "customer",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "userCreateUser_Success",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				user: &model.User{
					ID:       3,
					Name:     "Ahmet",
					Email:    "kadirgonen@kgstore.com",
					Password: "456789",
					Phone:    "05355816760",
					Roles: []*model.Role{
						{
							Name: "customer",
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.repo.GetUserEmail(tt.args.user.Email)
			if got == nil {
				if err := tt.fields.repo.CreateUser(tt.args.user); err != nil {
					t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if tt.wantErr != false {
					t.Errorf("UserService.CreateUser() error = %v, wantErr %v", got, tt.wantErr)
				}
			}
		})
	}
}

// Mock UserRepository
var (
	errCRUD      = errors.New("mock: error crud operation")
	userNotFound = errors.New("user not found")
)

type mockUserRepository struct {
	items []model.User
}

// GetAllUsers get user from mock repository
func (u *mockUserRepository) GetAllUsers() ([]model.User, int) {

	// if id == 0 {
	// 	return nil, errCRUD
	// }

	// for _, item := range u.items {
	// 	if item.ID == id {
	// 		return &item, nil
	// 	}
	// }
	// return nil, userNotFound
	var users []model.User
	var count int64
	return users, int(count)
}

// GetUser get user from mock repository
func (u *mockUserRepository) GetUser(id int) *model.User {

	if id == 0 {
		return nil
	}

	for _, item := range u.items {
		if item.ID == id {
			return nil
		}
	}
	return nil
}
func (u *mockUserRepository) GetUserWithRoles(id int) *model.User {
	//var user model.User
	//var role model.User
	if id == 0 {
		return nil
	}
	for _, item := range u.items {
		if item.ID == id {
			return &item
		}
	}

	return nil
}

// GetUserEmail get user by email from mock repository
func (u *mockUserRepository) GetUserEmail(email string) *model.User {
	//var user model.User

	if len(email) == 0 {
		return nil
	}
	for _, item := range u.items {
		if item.Email == email {
			return &item
		}
	}
	return nil
}

// GetUserEmailPassword get user by email from mock repository
func (u *mockUserRepository) GetUserEmailPassword(email string, password string) *model.User {
	//var user model.User

	if len(email) == 0 {
		return nil
	}
	if len(password) == 0 {
		return nil
	}
	for _, item := range u.items {
		if item.Email == email && item.Password == password {
			return &item
		}
	}
	return nil
}

// CreateUser get user from mock repository
func (u *mockUserRepository) CreateUser(User *model.User) error {

	// if id == 0 {
	// 	return nil
	// }

	// for _, item := range u.items {
	// 	if item.ID == id {
	// 		return &item, nil
	// 	}
	// }
	return nil
}

// UpdateUser get user from mock repository
func (u *mockUserRepository) UpdateUser(User *model.User) error {

	// if id == 0 {
	// 	return nil
	// }

	// for _, item := range u.items {
	// 	if item.ID == id {
	// 		return &item, nil
	// 	}
	// }
	return nil
}

// DeleteUser get user from mock repository
func (u *mockUserRepository) DeleteUser(id int) error {

	if id == 0 {
		return nil
	}

	for _, item := range u.items {
		if item.ID == id {
			return nil
		}
	}
	return nil
}
