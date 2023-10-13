package gorm

import (
	"fmt"
	"github.com/loafoe/iamsale/gen/account"
	"github.com/loafoe/iamsale/storage"
	"gorm.io/gorm"
)

var _ storage.Store = (*Storage)(nil)

func New(db *gorm.DB) (storage.Store, error) {
	err := db.AutoMigrate(&account.Account{})
	if err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}
	return &Storage{DB: db}, nil
}

type Storage struct {
	DB *gorm.DB
}

func (s Storage) FindAll() (res []*account.Account, err error) {
	tx := s.DB.Find(&res)
	return res, tx.Error
}

func (s Storage) FindByLoginID(login string) (*account.Account, error) {
	var res account.Account
	tx := s.DB.First(&res, "login = ?", login)
	return &res, tx.Error
}

func (s Storage) Remove(a account.Account) error {
	tx := s.DB.Delete(&account.Account{}, "guid = ?", a.GUID)
	return tx.Error
}

func (s Storage) Add(a account.Account) (*account.Account, error) {
	tx := s.DB.Create(&a)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &a, nil
}

func (s Storage) Update(a account.Account) (*account.Account, error) {
	if a.GUID == nil {
		return nil, fmt.Errorf("update: missing GUID")
	}
	tx := s.DB.Model(a).Updates(map[string]interface{}{"status": a.Status, "guid": a.GUID})
	if tx.Error != nil {
		return nil, tx.Error
	}
	updated, err := s.FindByGUID(*a.GUID)
	return updated, err
}

func (s Storage) FindByGUID(guid string) (*account.Account, error) {
	var res account.Account
	tx := s.DB.First(&res, "guid = ?", guid)
	return &res, tx.Error
}
