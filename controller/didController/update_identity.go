package didController

import (
	"DIDTrustCore/model"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"log"
)

func UpdateDocument(contract *client.Contract, didID string, newDidDoc model.DidDocument) error {
	identityJson, jsonErr := json.Marshal(newDidDoc)
	if jsonErr != nil {
		log.Printf("序列化失败：%v", jsonErr)
		return fmt.Errorf("序列化失败：%v", jsonErr)
	}
	fmt.Printf("\n--> Evaluate Transaction: UpdateDocument, \n")
	params := []string{didID, string(identityJson)}
	_, err := contract.SubmitTransaction("UpdateDocument", params...)
	if err != nil {
		fmt.Printf("failed to evaluate transaction: %w", err)
		return fmt.Errorf("failed to evaluate transaction: %w", err)
	}
	fmt.Printf("*** Result:更新成功")
	return nil
}
