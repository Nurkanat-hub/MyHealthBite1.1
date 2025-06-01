package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	statsProto "MyHealthBite/stats-service/proto"
	emailProto "email-service/proto"

	"user-service/internal/client"
	metrics "user-service/internal/handler/metrics"
	"user-service/internal/repository"
	"user-service/proto"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserServer реализует gRPC-сервис UserService
type UserServer struct {
	proto.UnimplementedUserServiceServer
	Repo *repository.UserRepository
}

// Конструктор UserServer
func NewUserServer(repo *repository.UserRepository) *UserServer {
	return &UserServer{Repo: repo}
}

// Register — регистрация нового пользователя
func (s *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	exists, err := s.Repo.EmailExists(ctx, req.Email)
	if err != nil {
		metrics.UserErrorCounter.Inc()
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	if exists {
		metrics.UserErrorCounter.Inc()
		return nil, status.Errorf(codes.AlreadyExists, "email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		metrics.UserErrorCounter.Inc()
		return nil, status.Errorf(codes.Internal, "could not hash password")
	}

	user := repository.User{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.Repo.CreateUser(ctx, user); err != nil {
		metrics.UserErrorCounter.Inc()
		return nil, status.Errorf(codes.Internal, "could not save user")
	}

	metrics.UserRegistrationCounter.Inc()

	_, err = client.StatsClient.InitStats(ctx, &statsProto.InitStatsRequest{
		UserId:         user.ID.Hex(),
		TargetCalories: 2000,
		TargetWaterMl:  2000,
	})
	if err != nil {
		log.Printf("⚠️ Failed to initialize stats for user %s: %v", user.ID.Hex(), err)
	}

	go func() {
		_, err := client.EmailClient.SendEmail(context.Background(), &emailProto.EmailRequest{
			To:      user.Email,
			Subject: "Welcome to MyHealthBite!",
			Body:    fmt.Sprintf("Hello %s,\n\nWelcome to MyHealthBite! We're glad you're here.", user.Name),
		})
		if err != nil {
			log.Printf("⚠️ Failed to send welcome email to %s: %v", user.Email, err)
		}
	}()

	token, err := generateJWT(user.ID.Hex(), user.Name)
	if err != nil {
		metrics.UserErrorCounter.Inc()
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}

	return &proto.AuthResponse{
		UserId: user.ID.Hex(),
		Token:  token,
	}, nil
}

// Login — аутентификация пользователя
func (s *UserServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {
	user, err := s.Repo.FindByEmail(ctx, req.Email)
	if err != nil {
		metrics.UserErrorCounter.Inc()
		log.Println("Пользователь не найден:", err)
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		metrics.UserErrorCounter.Inc()
		log.Println("Неверный пароль:", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	token, err := generateJWT(user.ID.Hex(), user.Name)
	if err != nil {
		metrics.UserErrorCounter.Inc()
		log.Println("Ошибка при создании JWT:", err)
		return nil, status.Errorf(codes.Internal, "could not generate token")
	}

	metrics.UserLoginCounter.Inc()

	return &proto.AuthResponse{
		UserId: user.ID.Hex(),
		Token:  token,
		Name:   user.Name,
	}, nil
}

// generateJWT — создание токена доступа
func generateJWT(userId string, userName string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET not set in environment")
	}

	claims := jwt.MapClaims{
		"user_id":   userId,
		"user_name": userName,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GetProfile — получить профиль пользователя
func (s *UserServer) GetProfile(ctx context.Context, req *proto.UserIdRequest) (*proto.UserResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		log.Println("Неверный ID:", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := s.Repo.FindByID(ctx, objectID)
	if err != nil {
		log.Println("Пользователь не найден:", err)
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &proto.UserResponse{
		UserId:  user.ID.Hex(),
		Name:    user.Name,
		Email:   user.Email,
		Goal:    user.Goal,
		Height:  user.Height,
		Weight:  user.Weight,
		Address: user.Address,
		Phone:   user.Phone,
	}, nil
}

// UpdateProfile — обновить профиль пользователя
func (s *UserServer) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.UserResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	updated, err := s.Repo.UpdateUserProfile(ctx, objectID, repository.User{
		Name:    req.Name,
		Goal:    req.Goal,
		Height:  req.Height,
		Weight:  req.Weight,
		Address: req.Address,
		Phone:   req.Phone,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update profile")
	}

	return &proto.UserResponse{
		UserId:  updated.ID.Hex(),
		Name:    updated.Name,
		Email:   updated.Email,
		Goal:    updated.Goal,
		Height:  updated.Height,
		Weight:  updated.Weight,
		Address: updated.Address,
		Phone:   updated.Phone,
	}, nil
}

// DeleteAccount — удалить пользователя и его статистику
func (s *UserServer) DeleteAccount(ctx context.Context, req *proto.UserIdRequest) (*proto.Empty, error) {
	_, err := client.StatsClient.DeleteStatsByUserId(ctx, &statsProto.UserIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		log.Printf("⚠️ Failed to delete stats for user %s: %v", req.UserId, err)
	}

	objectID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	if err := s.Repo.DeleteByID(ctx, objectID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user")
	}

	return &proto.Empty{}, nil
}
