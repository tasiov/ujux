package utils

import (
	"fmt"

	tmstore "github.com/cometbft/cometbft/proto/tendermint/store"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/gogo/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/opt"
	tmDb "github.com/tendermint/tm-db"
)

type BlockStore struct {
	db tmDb.DB
}

var blockStoreName = "blockstore"

func NewBlockStore(path string) (*BlockStore, error) {
	db, err := tmDb.NewGoLevelDBWithOpts(blockStoreName, path, &opt.Options{ReadOnly: true})
	return &BlockStore{db}, err
}

var blockStoreKey = []byte("blockStore")

// LoadBlockStoreState returns the BlockStoreState as loaded from disk.
func (bs *BlockStore) LoadBlockStoreState() (blockstoreState *tmstore.BlockStoreState, err error) {
	bz, err := bs.db.Get(blockStoreKey)
	if err != nil {
		return nil, err
	}

	var bss tmstore.BlockStoreState
	err = proto.Unmarshal(bz, &bss)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to BlockStoreState failed: %w", err)
	}

	return &bss, nil
}

func buildBlockMetaKey(height int64) []byte {
	return []byte(fmt.Sprintf("H:%v", height))
}

// LoadBlockMeta returns the BlockMeta for the given height.
func (bs *BlockStore) LoadBlockMeta(height int64) (*tmproto.BlockMeta, error) {
	bz, err := bs.db.Get(buildBlockMetaKey(height))
	if err != nil {
		return nil, err
	}

	var blockMeta tmproto.BlockMeta
	err = proto.Unmarshal(bz, &blockMeta)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to BlockMeta: %w", err)
	}

	return &blockMeta, nil
}

func buildBlockPartKey(height int64, partIndex int) []byte {
	return []byte(fmt.Sprintf("P:%v:%v", height, partIndex))
}

func (bs *BlockStore) LoadBlockPart(height int64, partIndex int) (*tmproto.Part, error) {
	bz, err := bs.db.Get(buildBlockPartKey(height, partIndex))
	if err != nil {
		return nil, err
	}

	var part tmproto.Part
	err = proto.Unmarshal(bz, &part)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to Part failed: %w", err)
	}

	return &part, nil
}

func (bs *BlockStore) LoadBlock(height int64) (*tmproto.Block, error) {
	blockMeta, err := bs.LoadBlockMeta(height)
	if err != nil {
		return nil, err
	}

	numParts := blockMeta.BlockID.PartSetHeader.Total

	buf := []byte{}
	for i := 0; i < int(numParts); i++ {
		part, err := bs.LoadBlockPart(height, i)
		if err != nil {
			return nil, err
		}
		fmt.Println(part)

		buf = append(buf, part.Bytes...)
	}

	var block tmproto.Block
	err = proto.Unmarshal(buf, &block)
	if err != nil {
		return nil, fmt.Errorf("unmarshal to Block failed: %w", err)
	}

	return &block, nil
}
