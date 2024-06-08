package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/entity"
)

type UserService interface {
	FindOwnerById(userID uuid.UUID) (entity.Owner, error)
}

type FileUploader interface {
	UploadFile(ctx context.Context, fileContent []byte, fileName string) (string, error)
}
