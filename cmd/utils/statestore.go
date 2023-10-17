package utils

import (
	"fmt"

	tmstate "github.com/cometbft/cometbft/proto/tendermint/state"
	"github.com/gogo/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/opt"
	tmDb "github.com/tendermint/tm-db"
)

type StateStore struct {
	db tmDb.DB
}

var stateStoreName = "state"

func NewStateStore(path string) (*StateStore, error) {
	db, err := tmDb.NewGoLevelDBWithOpts(stateStoreName, path, &opt.Options{ReadOnly: true})
	return &StateStore{db}, err
}

func buildABCIResponsesKey(height int64) []byte {
	return []byte(fmt.Sprintf("abciResponsesKey:%v", height))
}

// LoadABCIResponses returns the ABCIResponses for the given height.
func (ss *StateStore) LoadABCIResponses(height int64) (*tmstate.ABCIResponses, error) {
	bz, err := ss.db.Get(buildABCIResponsesKey(height))
	if err != nil {
		return nil, err
	}

	var abciResponses tmstate.ABCIResponses
	err = proto.Unmarshal(bz, &abciResponses)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to tmstate.ABCIResponses: %w", err)
	}

	return &abciResponses, nil
}
