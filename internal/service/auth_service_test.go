package service

import (
	"context"
	"testing"
	"time"

	"github.com/satudata/backend/internal/config"
	"github.com/satudata/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// mockUserRepo adalah mock untuk UserRepositoryInterface.
type mockUserRepo struct {
	users map[string]*domain.User
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{
		users: make(map[string]*domain.User),
	}
}

func (m *mockUserRepo) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, nil
}

func (m *mockUserRepo) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	if u, ok := m.users[id]; ok {
		return u, nil
	}
	return nil, nil
}

func (m *mockUserRepo) CreateUser(ctx context.Context, user *domain.User) error {
	m.users[user.ID.String()] = user
	return nil
}

func (m *mockUserRepo) UpdateUser(ctx context.Context, user *domain.User) error {
	m.users[user.ID.String()] = user
	return nil
}

func TestRegister(t *testing.T) {
	repo := newMockUserRepo()
	jwtCfg := config.JWTConfig{
		Secret:          "test-secret",
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 168 * time.Hour,
		Issuer:          "satudata-test",
	}
	svc := NewAuthService(repo, jwtCfg)

	req := RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
		FullName: "Test User",
	}

	user, err := svc.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if user.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got: %s", user.Email)
	}

	if user.FullName != "Test User" {
		t.Errorf("expected full name 'Test User', got: %s", user.FullName)
	}

	if user.Role != domain.UserRoleEditor {
		t.Errorf("expected role 'editor', got: %s", user.Role)
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	repo := newMockUserRepo()
	jwtCfg := config.JWTConfig{
		Secret:          "test-secret",
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 168 * time.Hour,
		Issuer:          "satudata-test",
	}
	svc := NewAuthService(repo, jwtCfg)

	req := RegisterRequest{
		Email:    "duplicate@example.com",
		Password: "password123",
		FullName: "User One",
	}

	_, err := svc.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error on first register, got: %v", err)
	}

	// Try registering with same email
	_, err = svc.Register(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}
}

func TestLogin_Success(t *testing.T) {
	repo := newMockUserRepo()
	jwtCfg := config.JWTConfig{
		Secret:          "test-secret",
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 168 * time.Hour,
		Issuer:          "satudata-test",
	}
	svc := NewAuthService(repo, jwtCfg)

	// Register first
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &domain.User{
		Email:        "login@example.com",
		PasswordHash: string(hash),
		FullName:     "Login User",
		Role:         domain.UserRoleEditor,
		IsActive:     true,
	}
	_ = repo.CreateUser(context.Background(), user)

	// Login
	req := LoginRequest{
		Email:    "login@example.com",
		Password: "password123",
	}

	tokenResp, err := svc.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if tokenResp.AccessToken == "" {
		t.Error("expected access token, got empty")
	}

	if tokenResp.RefreshToken == "" {
		t.Error("expected refresh token, got empty")
	}

	if tokenResp.TokenType != "Bearer" {
		t.Errorf("expected token type 'Bearer', got: %s", tokenResp.TokenType)
	}

	if tokenResp.User.Email != "login@example.com" {
		t.Errorf("expected email 'login@example.com', got: %s", tokenResp.User.Email)
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	repo := newMockUserRepo()
	jwtCfg := config.JWTConfig{Secret: "test-secret"}
	svc := NewAuthService(repo, jwtCfg)

	hash, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)
	user := &domain.User{
		Email:        "wrong@example.com",
		PasswordHash: string(hash),
		FullName:     "Wrong Password User",
		Role:         domain.UserRoleEditor,
		IsActive:     true,
	}
	_ = repo.CreateUser(context.Background(), user)

	req := LoginRequest{
		Email:    "wrong@example.com",
		Password: "wrongpassword",
	}

	_, err := svc.Login(context.Background(), req)
	if err == nil {
		t.Fatal("expected error for wrong password, got nil")
	}
}

func TestValidateToken(t *testing.T) {
	repo := newMockUserRepo()
	jwtCfg := config.JWTConfig{
		Secret:          "test-secret",
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 168 * time.Hour,
		Issuer:          "satudata-test",
	}
	svc := NewAuthService(repo, jwtCfg)

	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &domain.User{
		Email:        "token@example.com",
		PasswordHash: string(hash),
		FullName:     "Token User",
		Role:         domain.UserRoleAdmin,
		IsActive:     true,
	}
	_ = repo.CreateUser(context.Background(), user)

	req := LoginRequest{
		Email:    "token@example.com",
		Password: "password123",
	}

	tokenResp, err := svc.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to login: %v", err)
	}

	claims, err := svc.ValidateToken(context.Background(), tokenResp.AccessToken)
	if err != nil {
		t.Fatalf("expected no error validating token, got: %v", err)
	}

	if claims.Email != "token@example.com" {
		t.Errorf("expected email 'token@example.com', got: %s", claims.Email)
	}

	if claims.Role != "admin" {
		t.Errorf("expected role 'admin', got: %s", claims.Role)
	}
}
