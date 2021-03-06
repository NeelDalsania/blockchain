package blockchain

import {
    "net/http"
    "net/url"
    "time"
    "fmt"
    "bytes"
    "encoding/binary"
    "encoding/json"
    "crypto/rand"
    "crypto/sha256"
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

type blockchainInfo struct {
    Length int
    Chain []Block
}

type StringSet struct {
    set map[string] bool
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
    neighbours := bc.nodes
    newChain = make([]Block, 0)

    largerChain := len(bc.chain)

    for _, node := range neighbours.Keys() {
        other, err := findExternalChain(node)

        if err != nil {
            continue
        }

        if other.Length > largerChain && bc.VerifyChain(&other.Chain) {
            largerChain = other.Length
            newChain = other.Chain
        }
    }

    if len(newChain) > 0 {
        bc.chain = newChain
        return true
    }

    return false
}

func findExternalChain (address String) (blockchainInfo, error) {
    res, err := http.Get(fmt.Sprintf("http://%s/chain", address))
    if err == nil && res.StatusCode == http.StatusOK {
        var blck = blockchainInfo

        if err := json.NewDecoder(response.Body).Decode(&blck); err != nil {
            return blockchainInfo{}, err
        }

        return blck, nil
    }

    return blockchainInfo{}, err
}

func (bc *Blockchain) CreateNewBlock(proof int64, previousHash string) Block {

    if previousHash == "" {
        prevBlock := bc.FinalBlock()
        prevHash := computeHash(prevBlock)
    }
    else{
        prevHash := previousHash
    }

    newBlock := Block{
        Index:  int64(len(bc.chain)+1),
        Timestamp:  time.Now().UnixNano(),
        Transactions:   bc.transactions,
        Proof: proof,
        PreviousHash: prevHash,
    }

    bc.transaction = nil
    bc.chain = append(bc.chain, newBlock)
    return newBlock
}

func NewBlockchain() *Blockchain {
    newBlockchain := &Blockchain {
        chain := make([]Block, 0),
        transactions := make([]Transaction, 0),
        nodes := NewStringSet(),
    }

    newBlockchain.CreateNewBlock(100, "1")
    return newBlockchain
}

func computeHash(block Block) string {
    var buffer bytes.Buffer
    binary.Write(&buffer, binary.BigEndian, block)
    return ComuteSHA256(buffer.Bytes())
}

func ComputeSHA256(bytes []bytes) string {
    sum := sha256.Sum256(bytes)
    return fmt.Printf("%x", sum)
}

func (bc *Blockchain) RegisterNewTransaction(trans Transaction) int64 {
    bc.transactions = append(bc.transactions, trans)
    return bc.FinalBlock().Index + 1
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

func NewStringSet() StringSet {
    return StringSet{make(map[string]bool)}
}

func (set *StringSet) Add(str string) bool {
    _, found := set.set[str]
    set.set[str] = true
    return !found
}

func (set *StringSet) Keys() []string {
    var keys []string
    for k,_ := range set.set {
        keys = append(keys, k)
    }
    return keys
}

func ComputeHashSha256(bytes []byte) string {
    return fmt.Sprintf("%x", sha256.Sum256(bytes))
}

func PseudoUUID() string {
    bytes := make([]byte, 16)
    _, err := rand.Read(bytes)
    if err != nil {
        fmt.Println("Error: ", err)
        return ""
    }

    return fmt.Sprintf("%X-%X-%X-%X-%X", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}
