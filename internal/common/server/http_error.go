package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhia/artilight/internal/common/errors"
	"net/http"
)

func RespondWithError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case *errors.ValidationError:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	case *errors.DatabaseError:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	case *errors.NotFoundError:
		ctx.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
	case *errors.UnauthorizedError:
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	}
}
