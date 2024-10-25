package repository

import (
	"bean-x-echo/dto"
	"errors"
	"fmt"
	"github.com/aivyss/bean"
	"sync"
)

func init() {
	bean.RegisterBeanLazy(NewAccountRepository)
}

type AccountRepository interface {
	Insert(userID string, password string) error
}

type accountRepository struct {
	sequence int
	mutex    sync.Mutex
	memory   map[int]dto.Account
	idSet    map[string]bool
}

func (a *accountRepository) Insert(userID string, password string) error {
	a.mutex.Lock()
	if a.idSet[userID] {
		a.mutex.Unlock()
		return errors.New("duplicated user id")
	}

	account := dto.Account{
		ID:       a.sequence,
		UserID:   userID,
		Password: password,
	}
	a.memory[account.ID] = account
	a.idSet[account.UserID] = true
	a.sequence++
	a.mutex.Unlock()

	return nil
}

func NewAccountRepository() AccountRepository {
	fmt.Println("autowired: AccountRepository")

	return &accountRepository{
		sequence: 0,
		memory:   map[int]dto.Account{},
		idSet:    map[string]bool{},
	}
}
