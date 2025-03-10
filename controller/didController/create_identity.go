package didController

import (
	"DIDTrustCore/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

// CreateSoftwareIdentity Submit a transaction synchronously, blocking until it has been committed to the ledger.
func CreateSoftwareIdentity(contract *client.Contract, softwareIdentity model.SoftwareIdentity) error {
	didDocJson, jsonErr := json.Marshal(softwareIdentity)
	if jsonErr != nil {
		log.Printf("序列化失败：%v", jsonErr)
		return fmt.Errorf("序列化失败：%v", jsonErr)
	}
	fmt.Printf("\n--> Submit Transaction: CreateSoftwareDID\n")
	params := []string{softwareIdentity.DID, string(didDocJson)}
	_, err := contract.SubmitTransaction("CreateSoftwareDID", params...)
	if err != nil {
		log.Printf("failed to submit transaction: %v", err)
		return fmt.Errorf("failed to submit transaction: %v", err)
	}
	fmt.Printf("*** Transaction committed successfully\n")
	return nil
}
