package blockchain

import {
    "net/http"
    "net/url"
    "time"
    "fmt"
    "bytes"
    "encoding/binary"
    "encoding/json"
}

type BlockchainInterface interface {

    NewNode(address string) bool // Add a new node to the list of nodes

    VerifyChain(chain Blockchain) bool // Verify whether the blockchain is valid or not

    ResolveConflicts() bool // Consensus Algorithm to resolve conflicts

    CreateNewBlock(proof int64, previousHash string) Block // Add a new block to the blockchain

    RegisterNewTransaction(trans Transaction) int64 // Register a new transaction which goes in the next block

    FinalBlock() Block // Returns the final block in the chain

    ProofOfWork(lastProof int64) int64 // Simplified proof of work algorithm

    ValidProof(lastProof, proof int64) bool // Validates the proof
}

type Block struct {
    Index  int64
    Timestamp int64
    Transactions []Transaction
    Proof int64
    PreviousHash string
}

type Transaction struct {
    Sender string
    Receiver string
    Amount int64
}

type Blockchain struct {
    chain []Block
    transactions []Transaction
    nodes StringSet
}

func (bc *Blockchain) NewNode(address string) bool {
    address, err := url.parse(address)

    if err != nil {
        return false
    }
    return bc.nodes.Add(address.host)
}

func (bc *Blockchain) VerifyChain(chain *[]Block) bool {
    previousBlock := (*chain)[0]

    for currentIndex:=1; currentIndex < len(*chain); currentIndex++ {
        block := (*chain)[currentIndex]

        if block.PreviousHash != GetHash(preiousBlock) {
            return false
        }

        if !bc.ValidProof(previousBlock.Proof, block.Proof) {
            return false
        }

        previousBlock = block
    }

    return true

}

func (bc *Blockchain) ResolveConflicts() bool {

}

func (bc *Blockchain) CreateNewBlock(proof int64, previousHash string) Block {

}

func (bc *Blockchain) RegisterNewTransaction(trans Transaction) int64 {

}

func (bc *Blockchain) FinalBlock() Block {
    return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) ProofOfWork(lastProof int64) int64 {
    var proof int64 = 0
    for !bc.ValidProof(lastproof, proof){
        proof += 1
    }
    return proof
}

func (bc *Blockchain) ValidProof(lastProof, proof int64) bool {
    combined := fmt.Sprintf("%d%d", lastProof, proof)
    computeHash := ComputeHashSha256([]byte(combined))
    return computeHash[:4] == "0000"
}
