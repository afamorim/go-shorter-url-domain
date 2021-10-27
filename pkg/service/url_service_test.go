package service

import (
	"testing"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UrlRepositoryMock struct {
	mock.Mock
}

func (r *UrlRepositoryMock) Save(url model.Url) error {
	args := r.Called(url)

	//return args.Get(0).(model.Url), args.Error(1)
	return args.Error(1)
}

func (r *UrlRepositoryMock) FindByShorter(shorterUrl string) (model.Url, error) {
	url := model.Url{
		Id:          "SHORTER_ID",
		OriginalUrl: "http://www.teste.com.br",
		CompressUrl: shorterUrl,
	}

	return url, nil
}

func TestSaveSucess(t *testing.T) {
	url := model.Url{
		OriginalUrl: "http://www.teste.com",
	}

	urlRepositoryMock := UrlRepositoryMock{}
	urlRepositoryMock.On("Save", url).Return(nil)

	urlService := NewUrlService(&urlRepositoryMock)

	newUrl, err := urlService.Save(url)

	assert.Equal(t, 2, newUrl.Id)
	assert.NoError(t, err)

	urlRepositoryMock.AssertExpectations(t)
}

func TestSaveEmptyUrlError(t *testing.T) {
	url := model.Url{
		OriginalUrl: "",
	}

	urlRepositoryMock := UrlRepositoryMock{}

	urlService := NewUrlService(&urlRepositoryMock)

	newUrl, err := urlService.Save(url)
	//assert.True(t, err != nil, err.Error())
	//fmt.Print(err)
	assert.EqualError(t, err, "original url is mandatory")
	assert.Equal(t, "", newUrl.Id)

	urlRepositoryMock.AssertExpectations(t)
}

func TestFindByShorter(t *testing.T) {
	shorter := "http://localhost:8080/SHORTER_ID"
	urlRepositoryMock := UrlRepositoryMock{}

	urlService := NewUrlService(&urlRepositoryMock)

	newUrl, err := urlService.FindByShorter(shorter)

	assert.Equal(t, "SHORTER_ID", newUrl.Id)
	assert.Nil(t, err)
}
