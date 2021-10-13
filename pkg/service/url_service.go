package service

import (
	"errors"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/afamorim/go-shorter-url-domain/pkg/repository"
)

type UrlService interface {
	Save(url model.Url) (model.Url, error)
	FindByShorter(shoterUrl string) (model.Url, error)
}

type service struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(urlRepImpl repository.UrlRepository) UrlService {
	//urlRepository = ur
	return service{
		urlRepository: urlRepImpl,
	}
}

func (s service) Save(url model.Url) (model.Url, error) {
	if url.OriginalUrl == "" {
		return url, errors.New("Original URL is mandatory")
	}
	newUrl, err := s.urlRepository.Save(url)
	return newUrl, err
}

func (s service) FindByShorter(shoterUrl string) (model.Url, error) {
	url, err := s.urlRepository.FindByShorter(shoterUrl)
	return url, err
}
