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

func (r *UrlRepositoryMock) Save(url model.Url) (model.Url, error) {
	args := r.Called(url)

	return args.Get(0).(model.Url), args.Error(1)
}

func (r *UrlRepositoryMock) FindByShorter(shorterUrl string) (model.Url, error) {
	url := model.Url{
		Id:          1,
		OriginalUrl: "http://www.teste.com.br",
		CompressUrl: "http://yougotit.com",
	}

	return url, nil
}

func TestSaveSucess(t *testing.T) {
	url := model.Url{
		OriginalUrl: "http://www.teste.com",
	}

	urlRepositoryMock := UrlRepositoryMock{}
	urlRepositoryMock.On("Save", url).Return(
		model.Url{
			Id: 2,
		},
		nil)

	urlService := NewUrlService(&urlRepositoryMock)

	newUrl, err := urlService.Save(url)

	assert.Equal(t, 2, newUrl.Id)
	assert.NoError(t, err)

	urlRepositoryMock.AssertExpectations(t)
}

/*
func TestSaveEmptyUrlError(t *testing.T) {
	url := model.Url{
		OriginalUrl: "",
	}

	urlRepositoryMock := UrlRepositoryMock{}
	urlRepositoryMock.On("Save", url).Return(
		model.Url{},
		errors.New("original url is mandatory"))

	urlService := NewUrlService(&urlRepositoryMock)

	newUrl, err := urlService.Save(url)
	//assert.True(t, err != nil, err.Error())
	//fmt.Print(err)
	assert.EqualError(t, err, "original url is mandatory")
	assert.Equal(t, 0, newUrl.Id)

	urlRepositoryMock.AssertExpectations(t)
}
*/
