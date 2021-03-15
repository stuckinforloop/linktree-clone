package shortener

import (
	"fmt"

	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepo RedirectRepository) *redirectService {
	return &redirectService{
		redirectRepo: redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return fmt.Errorf("invalid redirect provided: %v", err)
	}

	var err error
	redirect.Code, err = shortid.Generate()
	if err != nil {
		return fmt.Errorf("error generating unique id: %v", err)
	}
	return r.redirectRepo.Store(redirect)
}
