// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package block

import (
    fct "github.com/FactomProject/factoid"
    "github.com/FactomProject/factoid/wallet"
)

var _ = fct.Prt

var genesisBlock IFBlock

// Provide the initial Exchange Rate, and the number of addresses to 
// fund, and how much to put at each address.
//
// GetGenesisBlock(1000000,10,200000000000)
//
// This is close to .1 a penny per Entry Credit with a Factoid price of
// 15 cents, and 2000 Factoids distributed to 10 addresses.
//
func GetGenesisBlock(ftime uint64, ExRate uint64, addressCnt int, Factoids uint64 ) IFBlock {
    if genesisBlock != nil { return genesisBlock }
    genesisBlock = NewFBlock(1000000, uint32(0))  
    
    w := new(wallet.SCWallet)        
    t := w.CreateTransaction(ftime)
    
    for i:=0; i<10; i++ {
        h,_ := w.GenerateFctAddress([]byte("test "+string(i)),1,1)
        w.AddOutput(t,h,200000000000)   // 2000 factoids per address
    }
    
	flg, err := genesisBlock.AddCoinbase(t)
	if !flg || err != nil { 
        fct.Prtln("Flag: ",flg," Error: ",err)
    }
	return genesisBlock
}
