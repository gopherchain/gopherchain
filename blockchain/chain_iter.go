package blockchain

import "github.com/dgraph-io/badger/v2"

type BlockChainIterartor struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (chain *BlockChain) Iterator() *BlockChainIterartor {
	iter := &BlockChainIterartor{chain.LastHash, chain.Database}

	return iter
}

func (iter *BlockChainIterartor) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		encodedBlock, err := item.ValueCopy(nil)
		block = Deserialize(encodedBlock)

		return err
	})

	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}
