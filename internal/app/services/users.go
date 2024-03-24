package services

import "github.com/Pk05999/rest_api_gin/internal/app/repository"

type userService struct {
	repository interface.IRepositoryUsers
}