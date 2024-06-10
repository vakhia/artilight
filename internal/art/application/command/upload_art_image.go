package command

import (
	"github.com/google/uuid"
	"github.com/vakhia/artilight/internal/art/domain/entity"
	"github.com/vakhia/artilight/internal/art/domain/repository"
	"github.com/vakhia/artilight/internal/common/fileuploader"
)

type UploadArtImageHandler struct {
	artRepository repository.ArtRepository
	storage       fileuploader.IStorage
}

func NewUploadArtImageHandler(artRepository repository.ArtRepository, storage fileuploader.IStorage) UploadArtImageHandler {
	return UploadArtImageHandler{
		artRepository: artRepository,
		storage:       storage,
	}
}

func (h *UploadArtImageHandler) Handle(artId uuid.UUID, fileData []byte, filename string) error {
	art, err := h.artRepository.FindArtById(artId)
	if err != nil {
		return err
	}

	filePath, err := h.storage.UploadFile(fileData, "arts/"+artId.String()+"/"+filename)
	if err != nil {
		return err
	}

	art.AddImage(entity.NewArtImage(filePath))

	err = h.artRepository.Save(art)
	return nil
}
