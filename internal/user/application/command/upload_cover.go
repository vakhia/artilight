package command

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/common/fileuploader"
	"github.com/vakhia/artilight/internal/user/domain/repository"
	"io"
)

type UploadCoverHandler struct {
	repo    repository.UserRepository
	storage fileuploader.IStorage
}

func NewUploadCoverHandler(repo repository.UserRepository, storage fileuploader.IStorage) UploadCoverHandler {
	return UploadCoverHandler{
		repo:    repo,
		storage: storage,
	}
}

func (h *UploadCoverHandler) Handle(ctx *gin.Context) error {
	contextId, _ := ctx.Get("userId")
	userId, err := uuid.Parse(contextId.(string))
	if err != nil {
		return err
	}

	user, err := h.repo.FindById(userId)
	if err != nil {
		return errors.New("user doesn't exist")
	}

	f, uploadedFile, err := ctx.Request.FormFile("file")
	if err != nil {
		return err
	}
	defer f.Close()

	fileData, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	filePath, err := h.storage.UploadFile(fileData, "covers/"+contextId.(string)+"/"+uploadedFile.Filename)
	if err != nil {
		return err
	}

	user.SetCover(filePath)
	if err := h.repo.Save(user); err != nil {
		return err
	}

	return nil
}
