package service

import (
	"golang-app-todolist/internal/entity"
	"golang-app-todolist/internal/model"
	"golang-app-todolist/internal/repository"
	"golang-app-todolist/pkg/exception"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request model.RegisterUserRequest) error
	Login(request model.LoginUserRequest) (*model.TokenResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
	validation     *validator.Validate
	log            *logrus.Logger
	viper          *viper.Viper
}

func UserServiceImpl(
	userRepository repository.UserRepository,
	validation *validator.Validate,
	log *logrus.Logger,
	viper *viper.Viper,
) UserService {
	return &userService{
		userRepository: userRepository,
		validation:     validation,
		log:            log,
		viper:          viper,
	}
}

func (s *userService) Register(request model.RegisterUserRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.log.WithError(err).Warn("validation error")
		return err
	}
	countUsername, err := s.userRepository.CountByUsername(request.Username)
	if err != nil {
		s.log.WithError(err).Error("failed count user to database")
	}
	if countUsername > 0 {
		s.log.Warn("username already exists")
		return exception.NewError(400, "username already exists")
	}
	email := strings.ToLower(request.Email)
	countEmail, err := s.userRepository.CountByEmail(email)
	if err != nil {
		s.log.WithError(err).Error("failed count user to database")
	}
	if countEmail > 0 {
		s.log.Warn("email already exists")
		return exception.NewError(400, "email already exists")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed generate hash password")
		return err
	}
	user := &entity.User{
		Email:    email,
		Username: request.Username,
		Password: string(hashPassword),
	}
	if err := s.userRepository.Save(*user); err != nil {
		s.log.WithError(err).Error("failed save user to dayabase")
		return err
	}
	return nil
}
func (s *userService) Login(request model.LoginUserRequest) (*model.TokenResponse, error) {
	if err := s.validation.Struct(&request); err != nil {
		s.log.WithError(err).Warn("validation error")
		return nil, err
	}
	user, err := s.userRepository.FindByUsername(request.Username)
	if err != nil {
		s.log.WithError(err).Warn("find by username failed")
		return nil, exception.NewError(401, "username or password wrong")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		s.log.WithError(err).Warn("failed compare password hash")
		return nil, exception.NewError(401, "username or password wrong")
	}

	token, err := JwtGenerateAccesToken(user.Username, []byte(s.viper.GetString("jwt.key")))
	if err != nil {
		s.log.WithError(err).Error("failed generate access token")
		return nil, err
	}
	return &model.TokenResponse{Token: token}, nil
}
