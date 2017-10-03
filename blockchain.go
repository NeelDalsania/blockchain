package blockchain

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

}

type Transaction struct {

}

type Blockchain struct {

}

func (bc *Blockchain) NewNode(address string) bool {

}

func (bc *Blockchain) VerifyChain(chain Blockchain) bool {

}

func (bc *Blockchain) ResolveConflicts() bool {

}

func (bc *Blockchain) CreateNewBlock(proof int64, previousHash string) Block {

}

func (bc *Blockchain) RegisterNewTransaction(trans Transaction) int64 {

}

func (bc *Blockchain) FinalBlock() Block {

}

func (bc *Blockchain) ProofOfWork(lastProof int64) int64 {

}
