package oppoHap

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

type ConvParams struct {
	IMEI        string `json:"imei,omitempty" structs:"imei"`
	OAID        string `json:"ouId,omitempty" structs:"ouId"`
	Mac         string `json:"mac,omitempty" structs:"mac"`
	ClientIP    string `json:"clientIp,omitempty" structs:"clientIp"`
	Timestamp   int64  `json:"timestamp,omitempty" structs:"timestamp"`
	Pkg         string `json:"pkg,omitempty" structs:"pkg"`
	DataType    int    `json:"dataType,omitempty" structs:"dataType"`
	CustomType  int    `json:"customType,omitempty" structs:"customType"`
	Channel     int    `json:"channel,omitempty" structs:"channel"`
	Type        int    `json:"type,omitempty" structs:"type"`
	AppType     int    `json:"appType,omitempty" structs:"appType"`
	PayAmount   int64  `json:"payAmount,omitempty" structs:"payAmount"`
	AscribeType int    `json:"ascribeType,omitempty" structs:"ascribeType"`
	ADID        int64  `json:"adId,omitempty" structs:"adId"`
	RequestID   string `json:"requestId,omitempty" structs:"requestId"`
	PayID       string `json:"payId,omitempty" structs:"payId"`
}

type ConvResp struct {
	Ret int    `json:"ret" structs:"ret"`
	Msg string `json:"msg" structs:"msg"`
}

// GetSign 获取签名
func (h *Handler) GetSign(body []byte, timestamp int64, slat string) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d%s", body, timestamp, slat)))

	return strings.ToLower(hex.EncodeToString(hash[:]))
}

// EncryptByAes 设备号AES加密
func (h *Handler) EncryptByAes(deviceByte []byte) (string, error) {
	if len(deviceByte) == 0 {
		return "", nil
	}

	key, err := base64.StdEncoding.DecodeString(aesBase64Key)
	if err != nil {
		return "", fmt.Errorf("base64.StdEncoding.DecodeString %w", err)
	}

	output, err := h.AESECBEncrypt(deviceByte, key)
	if err != nil {
		return "", fmt.Errorf("AESECBEncrypt %w", err)
	}

	return base64.StdEncoding.EncodeToString(output), nil
}

// AESECBEncrypt AES ECB 加密
func (h *Handler) AESECBEncrypt(input []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("aes.NewCipher error: %v", err)
	}

	paddedInput := h.Pkcs5Padding(input, block.BlockSize())

	output := make([]byte, len(paddedInput))
	// ECB 模式需要使用 NewECBEncrypter 方法
	ecb := h.NewECBEncrypter(block)
	ecb.CryptBlocks(output, paddedInput)

	return output, nil
}

// Pkcs5Padding PKCS5 填充
func (h *Handler) Pkcs5Padding(input []byte, blockSize int) []byte {
	padding := blockSize - len(input)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(input, padText...)
}

// 实现 ECB 模式的加密器
type ecbEncrypter struct {
	b         cipher.Block
	blockSize int
}

// NewECBEncrypter ECB 模式
func (h *Handler) NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	blockSize := b.BlockSize()
	return &ecbEncrypter{b, blockSize}
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("src not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
