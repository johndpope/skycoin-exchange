package account

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang/glog"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/util"
)

type AccountManager interface {
	CreateAccountWithPubkey(pk cipher.PubKey) (Accounter, error)
	GetAccount(id AccountID) (Accounter, error)
	Save() error
}

// AccountManager manage all the accounts in the server.
type ExchangeAccountManager struct {
	Accounts        map[AccountID]*ExchangeAccount `json:"accounts"`
	AcntMgrFileName string
	mtx             sync.RWMutex
}

type exchgAcntMgrJson struct {
	Accounts []exchgAcntJson `json:"accounts"`
}

// NewAccountManager
func NewAccountManager(fileName string) AccountManager {
	return &ExchangeAccountManager{
		Accounts:        make(map[AccountID]*ExchangeAccount),
		AcntMgrFileName: fileName,
	}
}

// LoadAccountManager from local disk.
func LoadAccountManager(acntName string) (AccountManager, error) {
	p := filepath.Join(acntDir, acntName)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, err
	}

	a := exchgAcntMgrJson{}
	d, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(d, &a); err != nil {
		return nil, err
	}
	return a.ToExchgAcntMgr(acntName), nil
}

// CreateAccountWithPubkey create an accounter with specific pubkey, this pubkey is generated by client.
func (self *ExchangeAccountManager) CreateAccountWithPubkey(pubkey cipher.PubKey) (Accounter, error) {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	// check duplicate
	if _, ok := self.Accounts[AccountID(pubkey)]; ok {
		return nil, errors.New("duplicate account id")
	}
	at := newExchangeAccount(AccountID(pubkey))
	self.Accounts[AccountID(pubkey)] = &at

	// save the account into disk
	if err := self.save(); err != nil {
		return nil, err
	}

	return &at, nil
}

// GetAccount return the account of specific id.
func (self *ExchangeAccountManager) GetAccount(id AccountID) (Accounter, error) {
	self.mtx.RLock()
	defer self.mtx.RUnlock()
	if account, ok := self.Accounts[id]; ok {
		return account, nil
	} else {
		return nil, errors.New("account does not exist")
	}
}

func (self ExchangeAccountManager) ToMarshalable() exchgAcntMgrJson {
	amj := exchgAcntMgrJson{}

	for _, acnt := range self.Accounts {
		amj.Accounts = append(amj.Accounts, acnt.ToMarshalable())
	}
	return amj
}

func (self *ExchangeAccountManager) Save() error {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	return self.save()
}

// persistance to disc. Save as JSON
func (self *ExchangeAccountManager) save() error {
	glog.Info("save accounts")
	a := self.ToMarshalable()
	// for self.Accounts
	return util.SaveJSON(filepath.Join(acntDir, self.AcntMgrFileName), a, 0600)
}

func (self exchgAcntMgrJson) ToExchgAcntMgr(fileName string) *ExchangeAccountManager {
	acntMap := make(map[AccountID]*ExchangeAccount, len(self.Accounts))
	for _, acnt := range self.Accounts {
		at := acnt.ToExchgAcnt()
		acntMap[at.ID] = at
	}
	return &ExchangeAccountManager{
		Accounts:        acntMap,
		AcntMgrFileName: fileName,
	}
}
