package abstraction

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/central-ves/model/internal/database"
)

// the database which used by others
type UserBase interface {
	// insert accounts maps from guid to account
	InsertAccount(userName string, acc uiptypes.Account) error

	// DeleteAccount(userName string, Account) error

	// DeleteAccountByName(userName string) error

	// DeleteAccountByID(user_id) error

	// find accounts which guid is corresponding to user
	FindUser(userName string) (*database.User, error)

	// find accounts which guid is corresponding to user
	FindAccounts(userName string, chainID uint64) ([]uiptypes.Account, error)

	// return true if user has this account
	HasAccount(userName string, acc uiptypes.Account) (has bool, err error)

	// return the user which has this account
	InvertFind(uiptypes.Account) (*database.User, error)
}