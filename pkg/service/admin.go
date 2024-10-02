package service

import (
	"github.com/gofrs/uuid"
	"stroycity/pkg/model"
	"stroycity/pkg/repository"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}
func (s *AdminService) AdminSignUp(admin model.Admin) error {
	admin.ID = uuid.Must(uuid.NewV4()).String()
	admin.Password = GeneratePasswordHash(admin.Password)
	return s.repo.AdminSignUp(admin)
}

func (s *AdminService) AdminSignIn(login, password string) (string, error) {
	hashedPassword := GeneratePasswordHash(password)

	admin, err := s.repo.AdminSignIn(login, hashedPassword)
	if err != nil {
		return "", err
	}

	token := CreateToken(admin.ID, "admin")

	return token, err
}
