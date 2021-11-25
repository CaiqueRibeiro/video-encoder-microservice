package domain_test

import (
	"testing"
	"time"

	"encoder/domain"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "abc"
	video.ResourceID = "aaaa"
	video.FilePath = "/"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "aaaa"
	video.FilePath = "/"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
