package impl

import (
	"errors"
	"fmt"
	"time"

	"github.com/ellcrys/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm requires it
)

// ErrChainExist represents an error about an existing chain
var ErrChainExist = errors.New("chain already exists")

// LedgerTableName represents the name of the table where all
// known ledger info are stored.
const LedgerTableName = "ledgers"

// TransactionTableName represents the name of the table where all
// transactions are stored.
const TransactionTableName = "transactions"

// NullHash is the default hash value assigned to columns that require a default hash value
const NullHash = "0000000000000000000000000000000000000000000000000000000000000000"

// Ledger represents a group of linked transactions
type Ledger struct {
	Number         uint   `gorm:"primary_key"`
	Hash           string `json:"hash" gorm:"type:varchar(64);unique_index"`
	PrevLedgerHash string `json:"prev_ledger_hash" gorm:"type:varchar(64);unique_index"`
	NextLedgerHash string `json:"next_ledger_hash" gorm:"type:varchar(64);unique_index"`
	Name           string `json:"name" gorm:"type:varchar(64);unique_index"`
	CocoonCodeID   string `json:"cocoon_code_id"`
	Public         bool   `json:"public"`
	CreatedAt      int64  `json:"created_at"`
}

// Transaction reprents a group of transactions belonging to a ledger.
// All transaction entries must reference the hash of the immediate transaction
// sharing the same ledger name.
type Transaction struct {
	Number     uint   `gorm:"primary_key"`
	Ledger     string `json:"ledger"`
	ID         string `json:"id" gorm:"type:varchar(64);unique_index"`
	Key        string `json:"key" gorm:"type:varchar(64)"`
	Value      string `json:"key" gorm:"type:text"`
	Hash       string `json:"hash" gorm:"type:varchar(64);unique_index"`
	PrevTxHash string `json:"prev_tx_hash" gorm:"type:varchar(64);unique_index"`
	NextTxHash string `json:"next_tx_hash" gorm:"type:varchar(64);unique_index"`
}

// PostgresLedgerChain defines a ledgerchain implementation
// on the postgres database. It implements the LedgerChain interface
type PostgresLedgerChain struct {
	db *gorm.DB
}

// GetBackend returns the database backend this chain
// implementation depends on.
func (ch *PostgresLedgerChain) GetBackend() string {
	return "postgres"
}

// Connect connects to a postgress server and returns a client
// or error if connection failed.
func (ch *PostgresLedgerChain) Connect(dbAddr string) (interface{}, error) {

	var err error
	ch.db, err = gorm.Open("postgres", dbAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ledgerchain backend")
	}

	ch.db.LogMode(false)

	return ch.db, nil
}

// MakeLegderHash takes a ledger and computes a hash
func (ch *PostgresLedgerChain) MakeLegderHash(ledger *Ledger) string {
	return util.Sha256(fmt.Sprintf("%s|%s|%t|%d", ledger.PrevLedgerHash, ledger.Name, ledger.Public, ledger.CreatedAt))
}

// Init initializes the blockchain. Creates the necessary tables such as the
// the table holding records of all ledgers.
func (ch *PostgresLedgerChain) Init() error {

	// create ledger table if not existing
	if !ch.db.HasTable(LedgerTableName) {
		if err := ch.db.CreateTable(&Ledger{}).Error; err != nil {
			return fmt.Errorf("failed to create `%s` table. %s", LedgerTableName, err)
		}
	}

	// create transaction table if not existing
	if !ch.db.HasTable(TransactionTableName) {
		if err := ch.db.CreateTable(&Transaction{}).Error; err != nil {
			return fmt.Errorf("failed to create `%s` table. %s", TransactionTableName, err)
		}
	}

	// Create general ledger if it does not exists
	var c int
	if err := ch.db.Model(&Ledger{}).Count(&c).Error; err != nil {
		return fmt.Errorf("failed to check whether general ledger exists in the ledger list table. %s", err)
	}

	if c == 0 {
		_, err := ch.CreateLedger("general", "", true)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateLedger creates a new ledger.
func (ch *PostgresLedgerChain) CreateLedger(name, cocoonCodeID string, public bool) (interface{}, error) {

	tx := ch.db.Begin()

	err := tx.Exec(`SET TRANSACTION isolation level repeatable read`).Error
	if err != nil {
		return nil, fmt.Errorf("failed to set transaction isolation level. %s", err)
	}

	newLedger := &Ledger{
		Name:           name,
		CocoonCodeID:   cocoonCodeID,
		Public:         public,
		NextLedgerHash: "",
		CreatedAt:      time.Now().Unix(),
	}

	var prevLedger Ledger
	err = tx.Where("next_ledger_hash = ?", "").Last(&prevLedger).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		newLedger.PrevLedgerHash = NullHash
	} else {
		newLedger.PrevLedgerHash = prevLedger.Hash
	}

	newLedger.Hash = ch.MakeLegderHash(newLedger)

	if err = tx.Model(&prevLedger).Update("next_ledger_hash", newLedger.Hash).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(newLedger).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return newLedger, nil
}

// Put creates a new transaction associated to a ledger.
// Returns error if ledger does not exists or nil of successful.
func (ch *PostgresLedgerChain) Put(ledgerName, key, value string) error {

	var ledger Ledger
	err := ch.db.Where("name = ?", ledgerName).First(&ledger).Error
	if err != nil {
		return fmt.Errorf("ledger does not exist")
	}

	return nil
}

// Close releases any resource held
func (ch *PostgresLedgerChain) Close() error {
	if ch.db != nil {
		return ch.Close()
	}
	return nil
}
