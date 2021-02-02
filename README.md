# msg-parser
Tx msg parser for block chain which built with cosmos-sdk

## install

### Requirement

Go version above 1.15

### Use Go Mod

```text
require (
    github.com/kaifei-bianjie/msg-parser latest
)
```

## Usage

### Init Client

The initialization SDK code is as follows:

```go

client := msg_sdk.NewMsgClient()
```

parse Bank Msg of Tx 
```go
bankDoc, ok := client.Bank.HandleTxMsg(&msg)
if ok {
		//db save bank msg
	}
```

use in sync 
```go

var (
	docTx models.Tx
	docTxMsgs []msg_sdk.DocTxMsg
    	)
	authTx := Tx.(signing.Tx)
	......
	msgs := authTx.GetMsgs()
	for _, msg := range msgs {
        if bankDoc, ok := client.Bank.HandleTxMsg(&msg);ok {
            docTxMsgs = append(docTxMsgs, bankDoc.DocTxMsg)
            continue
        }
        .....
    }
 docTx.DocTxMsgs = docTxMsgs

```