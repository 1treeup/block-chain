package main
import "fmt"
type Block struct {
	Timestamp     int64  //创建时的时间戳
    Data          []byte //相关交易信息
    PrevBlockHash []byte //上一个块的hash
    Hash          []byte //本次交易块的hash
}

func (this *Block) SetHash() {
    Timestamp := []byte(strconv.FormatInt(this.Timestamp, 10)) //将交易的时间转换为int数据类型
    headers := bytes.Join([][]byte{this.PrevBlockHash, this.Data, Timestamp}, []byte{})//headers 是将前一个的hash、自身的交易数据以及时间拼接起来存放到了byte切片里面
    hash := sha256.Sum256(headers) 
    this.Hash = hash[:] 
}

func NewBlock(data string, PrevBlockHash []byte) *Block {
    block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}} //用传入的数据将block实例化
    block.SetHash()//调用sethash，将block里面的数据传入到前一个sethash（）方法中，得到档次的hash值
    return block
}

type Blockchain struct {
    blocks []*Block//这里blocks 是一串地址的数组
}

func (this *Blockchain) AddBlock(data string) {
    prevBlock := this.blocks[len(this.blocks)-1] //上一个块
    newBlock := NewBlock(data, prevBlock.Hash) //将本次的交易数据以及上一个块的hash传入，调用创建新区块的方法
    this.blocks = append(this.blocks, newBlock)
}

func NewGenesisBlock() *Block {
    return NewBlock("我是初始模块", []byte{})
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
    bc := NewBlockchain() 

    bc.AddBlock("send 1 btc to ivan")
    bc.AddBlock("send 2 more btc to ivan")

    for _, block := range bc.blocks {
        fmt.Printf("Prev.hash: %x \n", block.PrevBlockHash)
        fmt.Printf("Data:%s \n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        fmt.Println()
    }
}