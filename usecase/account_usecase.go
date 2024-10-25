package usecase

import (
	"bean-x-echo/repository"
	"context"
	"fmt"
	"github.com/aivyss/go-bean"
	"github.com/google/uuid"
)

func init() {
	bean.RegisterBeanLazy(NewAccountUsecase)
}

type AccountUsecase interface {
	Signup(ctx context.Context, userID string, password string) error
}

type accountUsecase struct {
	accountRepository             repository.AccountRepository
	accountVerificationRepository repository.AccountVerificationRepository
}

func (a *accountUsecase) Signup(_ context.Context, userID string, password string) error {
	if err := a.accountRepository.Insert(userID, password); err != nil {
		return err
	}

	a.accountVerificationRepository.Insert(userID, uuid.New()) // 例のため、transactionなど考慮しない。

	return nil
}

func NewAccountUsecase(
	accountRepository repository.AccountRepository,
	accountVerificationRepository repository.AccountVerificationRepository,
) AccountUsecase {
	fmt.Println("autowired: AccountUsecase")

	return &accountUsecase{
		accountRepository:             accountRepository,
		accountVerificationRepository: accountVerificationRepository,
	}
}
