package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var global int
var hashit string
var lastblockhashwithoutchange string
var lastblockhashwithchange string

type Block struct {
	transaction string
	nounce      int
	prevhash    string
	currenthash string
}

type Blockchain struct {
	list []*Block
}

func NewBlock(transaction string, nonce int) *Block {
	s := new(Block)
	s.transaction = transaction
	s.nounce = nonce
	s.currenthash = CalculateHash(strconv.Itoa(nonce) + transaction + hashit)
	lastblockhashwithoutchange = s.currenthash
	if global == 0 {
		hashit = s.currenthash
		global++
	} else if global >= 1 {
		s.prevhash = hashit
		hashit = s.currenthash
	}
	return s

}
func (obj *Blockchain) createblock(transaction string, nounce int) *Block {

	as1 := NewBlock(transaction, nounce)
	obj.list = append(obj.list, as1)
	return as1
}

func printblock(obj *Blockchain, size int) {

	for i := 0; i < size; i++ {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))

		fmt.Println("Transaction: ", obj.list[i].transaction)
		fmt.Println("Nounce: ", obj.list[i].nounce)
		fmt.Println("Prev Hash: ", obj.list[i].prevhash)

	}

}

func CalculateHash(hashh string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func changeblock(block_num int, obj *Blockchain, new_transaction string, size int) {

	for i := 0; i < size; i++ {

		if i == block_num {

			obj.list[i].transaction = new_transaction
			obj.list[i].currenthash = CalculateHash(strconv.Itoa(obj.list[i].nounce) + obj.list[i].transaction + obj.list[i].prevhash)
			hashit = obj.list[i].currenthash
		}
		if i > block_num {
			obj.list[i].prevhash = hashit
			obj.list[i].currenthash = CalculateHash(strconv.Itoa(obj.list[i].nounce) + obj.list[i].transaction + obj.list[i].prevhash)
			hashit = obj.list[i].currenthash
			lastblockhashwithchange = obj.list[i].currenthash
		}

	}

}

func verify(lastblockhashwithchange string, lastblockhashwithoutchange string, obj *Blockchain, size int) {
	for i := 0; i < size; i++ {
		if i == size-1 {
			if CalculateHash(strconv.Itoa(obj.list[size-1].nounce)+obj.list[size-1].transaction+obj.list[size-1].prevhash) != lastblockhashwithoutchange {
				fmt.Println("Woops!! the blocks are tempered :(")
			} else {
				fmt.Println("Rest Assured!! The blocks are not tempered :)")
			}

		}

	}
}

func main() {
	fmt.Printf("\n")
	fmt.Println("                       ---Weolcome---                   ")
	fmt.Println("               Lets beign by creating blocks!            ")
	fmt.Printf("\n")
	s1 := new(Blockchain)
	s1.createblock("Alice to Bob", 2)
	s1.createblock("Bob to Charlie", 4)
	s1.createblock("Charlie to jack", 6)
	s1.createblock("Jack to Alice", 7)
	fmt.Print("Blockchain before data is tempered")
	fmt.Printf("\n")
	fmt.Printf("\n")
	printblock(s1, 4)
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Print("Blockchain after data is tempered")
	fmt.Printf("\n")
	fmt.Printf("\n")
	changeblock(2, s1, "Alice to Alice", 4)
	printblock(s1, 4)
	fmt.Printf("\n")
	fmt.Print("**********Now verifying***********")
	fmt.Printf("\n")

	fmt.Printf("\n")
	verify(lastblockhashwithchange, lastblockhashwithoutchange, s1, 4)

}
