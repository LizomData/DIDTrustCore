package didController

import (
	"DIDTrustCore/model"
	"DIDTrustCore/util"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func QuerySoftwareIdentityByID(contract *client.Contract, didID string) (*model.SoftwareIdentity, error) {
	fmt.Printf("\n--> Evaluate Transaction: QuerySoftwareIdentity, function returns asset attributes\n")
	evaluateResult, err := contract.EvaluateTransaction("QuerySoftwareIdentity", didID)
	if err != nil {
		fmt.Printf("failed to evaluate transaction: %v", err)
		return nil, fmt.Errorf("failed to evaluate transaction: %w", err)
	}
	result := util.FormatJSON(evaluateResult)
	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("*** Result: %s\n", string(data))
	return result, nil
}

func QueryAllSoftwareIdentity(contract *client.Contract) ([]model.SoftwareIdentity, error) {
	fmt.Printf("\n--> Evaluate Transaction: QueryAllSoftwareIdentity, function returns identity attributes\n")
	evaluateResult, err := contract.EvaluateTransaction("QueryAllSoftwareIdentity")
	if err != nil {
		fmt.Printf("failed to evaluate transaction: %v", err)
		return nil, fmt.Errorf("failed to evaluate transaction: %w", err)
	}
	var identities []model.SoftwareIdentity
	// 使用 json.Unmarshal 解析 JSON 数据到结构体数组
	if err := json.Unmarshal(evaluateResult, &identities); err != nil {
		fmt.Println("JSON 解析失败:", err.Error())
		return nil, err
	}
	fmt.Printf("*** Result: 执行成功\n")
	return identities, nil
}
