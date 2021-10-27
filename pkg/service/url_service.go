package service

import (
	"errors"
	"time"

	"github.com/afamorim/go-shorter-url-domain/pkg/model"
	"github.com/afamorim/go-shorter-url-domain/pkg/repository"
	hashids "github.com/speps/go-hashids"
)

const (
	base_url = "http://localhost:8008/"
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
		return url, errors.New("original url is mandatory")
	}
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	id, _ := h.Encode([]int{int(now.Unix())})
	url.Id = id
	url.CompressUrl = base_url + id
	err := s.urlRepository.Save(url)
	return url, err
}

func (s service) FindByShorter(shoterUrl string) (model.Url, error) {
	url, err := s.urlRepository.FindByShorter(shoterUrl)
	return url, err
}
