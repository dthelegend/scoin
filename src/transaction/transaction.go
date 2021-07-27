package transaction

import (
	"crypto/sha256"
	"math/big"
)

type Transaction struct {
	previous *Transaction
}

func (t Transaction) calculateHash() []byte {
	h := sha256.New()
	return h.Sum(nil)
}

type Block struct {
	previous *Block
	previousBlockHash []byte
	nonce big.Int
	transactions []Transaction
}

func (b Block) calculateBlockHash() []byte {
	h := sha256.New()
	h.Write(b.previousBlockHash)
	h.Write(b.nonce.Bytes())
	h.Write(b.calculateRootHash())
	return h.Sum(nil)
}

func calculateRootHash(hashes [][]byte) []byte {
	hashes_length := len(hashes)

	if(hashes_length == 0) {
		return nil
	}

	if(hashes_length == 1) {
		return hashes[0]
	}
	
	h := sha256.New()
	h.Write(calculateRootHash(hashes[:(len(hashes) / 2)]))
	h.Write(calculateRootHash(hashes[(len(hashes) / 2):]))
	return h.Sum(nil)
}

func (b Block) calculateRootHash() []byte {
	var hashes [][]byte
	for _, transaction := range b.transactions {
		hashes = append(hashes, transaction.calculateHash())
	}
	return calculateRootHash(hashes)
}
