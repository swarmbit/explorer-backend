package model

import (
    "go.mongodb.org/mongo-driver/bson"
    pb "github.com/spacemeshos/api/release/go/spacemesh/v1"
    "github.com/spacemeshos/explorer-backend/utils"
)

type Transaction struct {
    Id		string

    Layer	uint64
    Block	string
    Index	uint32	// the index of the tx in the ordered list of txs to be executed by stf in the layer
    Result	int

    GasProvided	uint64
    GasPrice	uint64
    GasUsed	uint64	// gas units used by the transaction (gas price in tx)
    Fee		uint64	// transaction fee charged for the transaction

    Amount	uint64	// amount of coin transfered in this tx by sender
    Counter	uint64	// tx counter aka nonce

    Type	int
    Scheme	int	// the signature's scheme
    Signature	string	// the signature itself
    PublicKey	string	// included in schemes which require signer to provide a public key

    Sender	string	// tx originator, should match signer inside Signature
    Receiver	string
    SvmData	string	// svm binary data. Decode with svm-codec
}

type TransactionReceipt struct {
    Id		string
    Layer	uint64
    Index	uint32	// the index of the tx in the ordered list of txs to be executed by stf in the layer
    Result	int
    GasUsed	uint64	// gas units used by the transaction (gas price in tx)
    Fee		uint64	// transaction fee charged for the transaction
    SvmData	string	// svm binary data. Decode with svm-codec
}

type TransactionService interface {
    GetTransaction(ctx context.Context, query *bson.D) (*Transaction, error)
    GetTransactions(ctx context.Context, query *bson.D) ([]*Transaction, error)
    SaveTransaction(ctx context.Context, in *Transaction) error
}

func NewTransactionReceipt(txReceipt *pb.TransactionReceipt) *TransactionReceipt {
    return &TransactionReceipt{
        Id: utils.BytesToHex(in.GetId().GetId()),
        Result: txReceipt.GetResult(),
        GasUsed: txReceipt.GetGasUsed(),
        Fee: txReceipt.GetFee().GetValue(),
        Layer: txReceipt.GetLayerNumber(),
        Index: txReceipt.GetIndex(),
        SvmData: txReceipt.GetSvmData(),
    }
}

func NewTransaction(in *pb.Transaction, layer uint64, blockId string) *Transaction {
    gas := in.GetGasOffered()
    sig := in.GetSignature()

    tx := &Transaction{
        Id: utils.BytesToHex(in.GetId().GetId()),
        Sender: utils.BytesToAddressString(in.GetSender().GetAddress()),
        Amount: uint64(in.GetAmount().GetValue()),
        Counter: in.GetCounter(),
        Layer: layer,
        Block: blockId,
        GasProvided: gas.GetGasProvided(),
        GasPrice: gas.GetGasPrice(),
        Scheme: sig.GetScheme(),
        Signature: utils.BytesToHex(sig.GetSignature()),
        PublicKey: utils.BytesToHex(sig.GetPublicKey()),
    }

    if data := t.GetCoinTransfer(); data != nil {
        tx.Receiver = utils.BytesToAddressString(data.GetReceiver().GetAddress()),
    } else if data := t.GetSmartContract(); data != nil {
        tx.Type = data.GetType(),
        tx.SvmData = data.GetData(),
        tx.Receiver = utils.BytesToAddressString(data.GetAccountId().GetAddress()),
    }

    return tx
}
