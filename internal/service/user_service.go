package service

import (
	"errors"

	"github.com/minhtridinh/trid-profile-go/internal/middleware"
	"github.com/minhtridinh/trid-profile-go/internal/model"
	"github.com/minhtridinh/trid-profile-go/internal/repository"
)

type UserService interface {
	Register(request model.RegisterRequest) error
	Login(request model.LoginRequest) (*model.AuthResponse, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	ValidateToken(token string) (*middleware.Claims, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service instance
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register creates a new user account
func (s *userService) Register(request model.RegisterRequest) error {
	// Check if username already exists
	existingUser, _ := s.userRepo.FindByUsername(request.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Check if email already exists
	existingEmail, _ := s.userRepo.FindByEmail(request.Email)
	if existingEmail != nil {
		return errors.New("email already exists")
	}

	// Create new user
	user := &model.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		FullName: request.FullName,
		Role:     "user", // Default role
		Active:   true,
	}

	// Hash password before saving
	if err := user.HashPassword(); err != nil {
		return err
	}

	return s.userRepo.Create(user)
}

// Login authenticates a user and returns JWT token
func (s *userService) Login(request model.LoginRequest) (*model.AuthResponse, error) {
	// Find user by username
	user, err := s.userRepo.FindByUsername(request.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Check if user is active
	if !user.Active {
		return nil, errors.New("account is deactivated")
	}

	// Verify password
	if !user.ComparePassword(request.Password) {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, expiresAt, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	// Update last login time
	s.userRepo.UpdateLastLogin(user.ID)

	// Return auth response
	return &model.AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	}, nil
}

// GetUserByID retrieves user information by ID
func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

// UpdateUser updates user information
func (s *userService) UpdateUser(user *model.User) error {
	// Check if the user exists first
	existingUser, err := s.userRepo.FindByID(user.ID)
	if err != nil {
		return err
	}

	// If password is being updated, hash it
	if user.Password != "" && user.Password != existingUser.Password {
		if err := user.HashPassword(); err != nil {
			return err
		}
	}

	return s.userRepo.Update(user)
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

// ValidateToken validates a JWT token
func (s *userService) ValidateToken(token string) (*middleware.Claims, error) {
	return middleware.ValidateToken(token)
}
