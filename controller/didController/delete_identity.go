package didController

import (
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func DeleteSoftwareByID(contract *client.Contract, didID string) error {
	_, err := contract.SubmitTransaction("RevokeSoftwareIdentity", didID)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate RevokeSoftwareIdentity: %w", err))
	}
	fmt.Printf("*** Result:注销成功")
	return nil
}
