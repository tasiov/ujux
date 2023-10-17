package utils

import (
	"github.com/syndtr/goleveldb/leveldb/opt"
	tmDb "github.com/tendermint/tm-db"
)

type TxIndexStore struct {
	db tmDb.DB
}

var txIndexStoreName = "tx_index"

func NewTxIndexStore(path string) (*TxIndexStore, error) {
	db, err := tmDb.NewGoLevelDBWithOpts(txIndexStoreName, path, &opt.Options{ReadOnly: true})
	return &TxIndexStore{db}, err
}
