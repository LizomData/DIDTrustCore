package didController

import (
	"DIDTrustCore/model"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/big"
	"strings"
	"time"
)

func generateDID() string {
	u, err := uuid.NewRandom() // 生成随机 UUID
	if err != nil {
		log.Fatal("生成 UUID 失败: ", err)
	}
	// 移除 UUID 中的 "-"，确保符合w3c did规范
	uuidStr := strings.ReplaceAll(u.String(), "-", "")
	// 取前 21 个字符
	return "did:fabric:" + uuidStr[:21]
}

func generateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("生成密钥失败：%v", err)
	}
	return privateKey
}

func CreateDIDDocument(name string) model.SoftwareIdentity {
	did := generateDID()
	pk := generateKey()
	publicKeyBytes := elliptic.Marshal(pk.PublicKey.Curve, pk.PublicKey.X, pk.PublicKey.Y)
	var services []model.Service
	service := model.Service{
		ID:              did + "#service",
		Type:            "web_page",
		ServiceEndpoint: "https://did.jzhangluo.com/" + did,
	}
	services = append(services, service)
	var didDoc = model.SoftwareIdentity{
		DID:       did,
		Name:      name,
		Status:    "ACTIVE",
		Version:   1,
		Timestamp: time.Now().UnixNano(),
		DidDocument: model.DidDocument{
			Context: []string{"https://www.w3.org/ns/did/v1"},
			ID:      did,
			VerificationMethod: model.VerificationMethod{
				ID:         did + "#keys-1",
				Type:       "ECDSA",
				Controller: did,
				PublicKey:  hex.EncodeToString(publicKeyBytes),
				Created:    time.Now().UnixNano(),
				Updated:    time.Now().UnixNano(),
			},
			Service: services,
			Created: time.Now().Format(time.RFC3339),
			Updated: time.Now().Format(time.RFC3339),
		},
	}
	fmt.Printf("%+v\n", didDoc)
	err := signDocument(&didDoc, pk)
	if err != nil {
		log.Fatalf("文档签名错误：%v", err)
	}
	return didDoc
}

func signDocument(doc *model.SoftwareIdentity, pk *ecdsa.PrivateKey) error {
	// 创建签名用文档副本
	signDoc := doc.DidDocument
	signDoc.Authentication.Sign = ""

	// 序列化文档
	docBytes, err := json.Marshal(signDoc)
	if err != nil {
		return err
	}

	// 计算哈希
	hash := sha256.Sum256(docBytes)

	// 生成签名
	r, s, err := ecdsa.Sign(rand.Reader, pk, hash[:])
	if err != nil {
		return err
	}

	// 编码签名
	signature := append(padBytes(r, 32), padBytes(s, 32)...)
	authentication := model.Authentication{
		ID:      doc.DID + "#auth",
		Type:    "ECDSA",
		Sign:    hex.EncodeToString(signature),
		Created: time.Now().UnixNano(),
		Updated: time.Now().UnixNano(),
	}
	doc.DidDocument.Authentication = authentication
	return nil
}

func padBytes(b *big.Int, size int) []byte {
	bytes := b.Bytes()
	if len(bytes) >= size {
		return bytes
	}
	padded := make([]byte, size)
	copy(padded[size-len(bytes):], bytes)
	return padded
}
