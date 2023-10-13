package storage

import (
	"github.com/loafoe/iamsale/gen/account"
)

type Store interface {
	Remove(a account.Account) error
	Add(a account.Account) (*account.Account, error)
	Update(a account.Account) (*account.Account, error)
	FindByGUID(id string) (*account.Account, error)
	FindByLoginID(login string) (*account.Account, error)
	FindAll() ([]*account.Account, error)
}
