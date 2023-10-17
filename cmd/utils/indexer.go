package utils

import (
	"fmt"

	tmstore "github.com/cometbft/cometbft/proto/tendermint/store"
	"github.com/gogo/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/opt"
	tmDb "github.com/tendermint/tm-db"
)

type levelDbs struct {
	blockstore tmDb.DB
	state      tmDb.DB
}

var blockStoreKey = []byte("blockStore")

func newLevelDbs(path string) (*levelDbs, error) {
	blockstore, err := tmDb.NewGoLevelDBWithOpts("blockstore", path, &opt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}

	state, err := tmDb.NewGoLevelDBWithOpts("state", path, &opt.Options{ReadOnly: true})
	if err != nil {
		return nil, err
	}

	return &levelDbs{
		blockstore: blockstore,
		state:      state,
	}, nil
}

// LoadBlockStoreState returns the BlockStoreState as loaded from disk.
// If no BlockStoreState was previously persisted, it returns the zero value.
func (ldbs *levelDbs) loadBlockStoreState() (base int64, height int64) {
	bytes, err := ldbs.blockstore.Get(blockStoreKey)

	if err != nil {
		panic(err)
	}

	if len(bytes) == 0 {
		return 0, 0
	}

	var bss tmstore.BlockStoreState
	err = proto.Unmarshal(bytes, &bss)
	if err != nil {
		panic(fmt.Sprintf("Could not unmarshal bytes: %X", bytes))
	}

	return bss.Base, bss.Height
}

func GetBaseHeight(path string) (int64, int64, error) {
	levelDbs, err := newLevelDbs(path)
	if err != nil {
		return 0, 0, err
	}

	base, height := levelDbs.loadBlockStoreState()

	return base, height, nil
}
