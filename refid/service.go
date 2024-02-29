package refid

import (
	"strconv"
)

type Service interface {
	RegisterRefId(input RegisterRefIdInput) (RefId, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterRefId(input RegisterRefIdInput) (RefId, error) {
	refid := RefId{}
	refid.ID = strconv.Itoa(input.ID)
	refid.Flgused = 0

	newReffId, err := s.repository.Save(refid)

	if err != nil {
		return refid, nil
	}

	return newReffId, nil
}
