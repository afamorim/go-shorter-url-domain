package service

import (
	"testing"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UlRepositoryMock struct {
	mock.Mock
}

func (r *UlRepositoryMock) Save(url *model.Url) error {
	url.Id = 1
	return nil
}

func (r *UlRepositoryMock) FindByShorter(shoterUrl string) (model.Url, error) {
	url := model.Url{
		Id:          1,
		OriginalUrl: "http://www.teste.com.br",
		CompressUrl: "http://yougotit.com",
	}

	return url, nil
}

func TestSave(t *testing.T) {
	urlRepository := UlRepositoryMock{}
	urlService := NewUrlService(&urlRepository)
	url := model.Url{}
	urlService.Save(url)
	assert.Equal(t, 1, url.Id)
	t.Error("ERROR")
}
