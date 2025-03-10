package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SoftwareIdentity struct {
	DID         string      `json:"did"`
	Name        string      `json:"name"`
	DidDocument DidDocument `json:"document"`
	Controller  string      `json:"controller"` // 控制者身份（MSPID）
	Status      string      `json:"status"`     // ACTIVE/DOWN
	Version     uint64      `json:"version"`    // 版本号（每次更新递增）
	Timestamp   int64       `json:"timestamp"`  // 最后更新时间戳
}

type DidDocument struct {
	Context            []string           `json:"@context"`
	ID                 string             `json:"id"`
	VerificationMethod VerificationMethod `json:"verificationMethod"`
	Authentication     Authentication     `json:"authentication"`
	Service            []Service          `json:"service"`
	Created            string             `json:"created"`
	Updated            string             `json:"updated"`
}

type VerificationMethod struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Controller string `json:"controller"`
	PublicKey  string `json:"publicKey"`
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
}

type Authentication struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Sign    string `json:"sign"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

// StringArray 表示字符串数组
type StringArray []string

// Scan 实现 Scanner 接口，将数据库中的值反序列化为 StringArray
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, a)
	case string:
		return json.Unmarshal([]byte(v), a)
	default:
		return errors.New("unsupported data type")
	}
}

// Value 实现 Valuer 接口，将 StringArray 序列化为 JSON 存储到数据库
func (a StringArray) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

// GormDataType 返回 GORM 的逻辑数据类型
func (StringArray) GormDataType() string {
	return "json"
}

// GormDBDataType 返回 MySQL 数据库的实际数据类型
func (StringArray) GormDBDataType() string {
	return "JSON"
}
