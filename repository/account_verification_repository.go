package repository

import (
	"fmt"
	"github.com/aivyss/bean"
	"github.com/google/uuid"
	"sync"
)

func init() {
	bean.RegisterBeanLazy(NewAccountVerificationRepository)
}

type AccountVerificationRepository interface {
	Insert(userID string, verificationCode uuid.UUID)
	IsOK(userID string, verificationCode string) bool
}

type accountVerificationRepository struct {
	memory map[struct {
		userID string
		code   string
	}]bool
	OkUsers map[string]bool
	mutex   sync.Mutex
}

func (a *accountVerificationRepository) Insert(userID string, verificationCode uuid.UUID) {
	a.mutex.Lock()
	a.memory[struct {
		userID string
		code   string
	}{userID: userID, code: verificationCode.String()}] = true
	a.mutex.Unlock()
}

func (a *accountVerificationRepository) IsOK(userID string, verificationCode string) bool {
	ok := a.memory[struct {
		userID string
		code   string
	}{userID: userID, code: verificationCode}]
	if ok {
		a.mutex.Lock()
		a.OkUsers[userID] = true
		a.mutex.Unlock()
	}

	return ok
}

func NewAccountVerificationRepository() AccountVerificationRepository {
	fmt.Println("autowired: AccountVerificationRepository")

	return &accountVerificationRepository{
		memory: map[struct {
			userID string
			code   string
		}]bool{},
	}
}
