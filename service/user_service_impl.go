package service

import (
	"fmt"
	"strconv"
	"time"

	"cek/auth"
	"cek/exception"
	"cek/helper"
	"cek/model/domain"
	"cek/model/web"
	"cek/repository"

	c "cek/configuration"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewUserService(userRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
	db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *UserServiceImpl) Create(request *web.UserCreateRequest) web.UserResponse {
	// Validate Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start Transaction
	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.PanicIfError(err)
	}
	nipUser := strconv.Itoa(request.Nip)
	// Create user
	user := &domain.User{
		// Model:    gorm.Model{ID: uint(nipUser)},
		Role:     request.Role,
		Name:     request.Name,
		Nip:      nipUser,
		JoinDate: request.JoinDate,
		Password: string(hashedPassword),
		Email:    request.Email,
		Phone:    request.Phone,
	}
	user = service.UserRepository.Create(tx, user)

	return user.ToUserResponse()
}

func (service *UserServiceImpl) Registration(request *web.RegistrationRequest) web.UserResponse {
	// Validate Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start Transaction
	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.PanicIfError(err)
	}
	// hd := hashids.NewData()
	// hd.Salt = request.Password
	// h, _ := hashids.NewWithData(hd)
	// id, _ := h.Encode([]int{request.Nip})
	// pass, _ := h.EncodeHex(request.Password)
	nipUser := strconv.Itoa(request.Nip)
	// userID := id + hd.Salt + pass
	// Create user
	user := &domain.User{
		CreatedByID: strconv.Itoa(request.Nip),
		UpdatedByID: strconv.Itoa(request.Nip),
		ID:          uint(request.Nip),
		Role:        request.Role,
		Name:        request.Name,
		Nip:         nipUser,
		JoinDate:    request.JoinDate,
		Password:    string(hashedPassword),
		Email:       request.Email,
		Phone:       request.Phone,
		Image:       "default.png",
	}
	user = service.UserRepository.Create(tx, user)

	return user.ToUserResponse()
}

func (service *UserServiceImpl) Delete(id *int) {
	// Start Transaction
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.UserRepository.Delete(tx, id)
}

func (service *UserServiceImpl) ResetPassword(auth *auth.AccessDetails, id *int) web.UserResponse {
	// Start Transaction
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser := service.UserRepository.FindByID(tx, id)

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(getUser.Nip+".123"), bcrypt.DefaultCost)
	if err != nil {
		helper.PanicIfError(err)
	}

	// Update user
	user := &domain.User{
		Model:    gorm.Model{ID: uint(*id)},
		Password: string(hashedPassword),
	}
	user = service.UserRepository.Update(tx, user)

	return user.ToUserResponse()
}

func (service *UserServiceImpl) ChangePassword(auth *auth.AccessDetails, id *int, request *web.UserChangePasswordRequest) web.UserResponse {
	// Validate Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start Transaction
	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getUser := service.UserRepository.FindByID(tx, id)

	hashedPassword := []byte(request.OldPassword)
	err = bcrypt.CompareHashAndPassword([]byte(getUser.Password), hashedPassword)
	if err != nil {
		err = &exception.ErrorSendToResponse{Err: "Password Old doesn't match "}
		helper.PanicIfError(err)
	}

	if request.NewPassword != request.RetypePassword {
		err = &exception.ErrorSendToResponse{Err: "New Password and Retype Password doesn't match "}
		helper.PanicIfError(err)
	}

	// Hash Password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	// Update user
	user := &domain.User{
		Model:    gorm.Model{ID: uint(*id)},
		Password: string(newHashedPassword),
	}
	user = service.UserRepository.Update(tx, user)

	return user.ToUserResponse()
}

func (service *UserServiceImpl) Update(id *int, request *web.UserUpdateRequest) web.UserResponse {
	// Validate Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start Transaction
	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// profileImage := fmt.Sprintf("%v-%v.png", request.Name, id)
	// Update user
	user := &domain.User{
		Model:      gorm.Model{ID: uint(*id)},
		Role:       request.Role,
		Nip:        request.Nip,
		Name:       request.Name,
		JoinDate:   request.JoinDate,
		ResignDate: request.ResignDate,
		Email:      request.Email,
		Phone:      request.Phone,
		Image:      request.Image.Filename,
	}
	user = service.UserRepository.Update(tx, user)

	return user.ToUserResponse()
}

func (service *UserServiceImpl) FindAll(filters *map[string]string) []web.UserResponse {
	// Start Transaction
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(tx, filters)

	return users.ToUserResponses()
}

func (service *UserServiceImpl) RefreshToken(request *web.RefreshTokenCreateRequest, userAgent, remoteAddress *string) web.TokenResponse {
	// Validate Request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Start Transaction
	tx := service.DB.Begin()
	err = tx.Error
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Parse JWT Refresh Token
	configuration, err := c.LoadConfig()
	helper.PanicIfError(err)

	token, err := jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configuration.RefreshSecret), nil
	})
	if err != nil {
		helper.PanicIfError(exception.ErrRefreshTokenExpired)
	}
	if !token.Valid {
		helper.PanicIfError(exception.ErrUnauthorized)
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		// Check User is Exist
		userID, err := strconv.Atoi(fmt.Sprintf("%.f", claims["user_id"]))
		if err != nil {
			helper.PanicIfError(err)
		}
		user := service.UserRepository.FindByID(tx, &userID)

		// Delete Exist Refresh UUID
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			helper.PanicIfError(err)
		}
		service.SessionRepository.FindByRefreshUUID(tx, &refreshUUID)
		service.SessionRepository.DeleteByRefreshUUID(tx, &refreshUUID)

		// Create New Token
		ts, createErr := auth.CreateToken(&user, user.Nip, userAgent, remoteAddress, func(userID uint, tokenDetails *auth.TokenDetails) {
			sessions := &domain.Session{
				UserID:        userID,
				RefreshUUID:   tokenDetails.RefreshUUID,
				Expired:       time.Unix(tokenDetails.RefreshTokenExpired, 0),
				UserAgent:     tokenDetails.UserAgent,
				RemoteAddress: tokenDetails.RemoteAddress,
			}
			_ = service.SessionRepository.Create(tx, sessions)
		})
		if createErr != nil {
			helper.PanicIfError(err)
		}
		token := web.TokenResponse{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
		return token
	}
	helper.PanicIfError(exception.ErrRefreshTokenExpired)
	return web.TokenResponse{}
}
func (service *UserServiceImpl) Login(nip, password, userAgent, remoteAddress *string) web.TokenResponse {
	tx := service.DB.Begin()
	err := tx.Error
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user := service.UserRepository.FindByNip(tx, nip)

	// hashedPassword := []byte(*password)
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), hashedPassword)
	// if err != nil {
	// 	helper.PanicIfError(exception.ErrUnauthorized)
	// }

	ts, err := auth.CreateToken(&user, user.Nip, userAgent, remoteAddress, func(userID uint, tokenDetails *auth.TokenDetails) {
		sessions := &domain.Session{
			UserID:        userID,
			RefreshUUID:   tokenDetails.RefreshUUID,
			Expired:       time.Unix(tokenDetails.RefreshTokenExpired, 0),
			UserAgent:     tokenDetails.UserAgent,
			RemoteAddress: tokenDetails.RemoteAddress,
		}
		_ = service.SessionRepository.Create(tx, sessions)
	})
	if err != nil {
		helper.PanicIfError(err)
	}

	token := web.TokenResponse{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}

	return token
}
