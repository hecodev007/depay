package com

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	//	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func ToValidateAddress(address string) string {
	addrLowerStr := strings.ToLower(address)
	if strings.HasPrefix(addrLowerStr, "0x") {
		addrLowerStr = addrLowerStr[2:]
		address = address[2:]
	}
	var binaryStr string
	addrBytes := []byte(addrLowerStr)
	hash256 := crypto.Keccak256Hash([]byte(addrLowerStr)) //注意，这里是直接对字符串转换成byte切片然后哈希

	for i, e := range addrLowerStr {
		//如果是数字则跳过
		if e >= '0' && e <= '9' {
			continue
		} else {
			binaryStr = fmt.Sprintf("%08b", hash256[i/2]) //注意，这里一定要填充0
			if binaryStr[4*(i%2)] == '1' {
				addrBytes[i] -= 32
			}
		}
	}

	return "0x" + string(addrBytes)
}

func IsEthAddress(address string) bool {
	return ToValidateAddress(address) == address
}

func ValidSigner(signature []byte, message string, address common.Address) bool {
	data := []byte(message)
	hash := crypto.Keccak256Hash(data)
	//fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	//fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return false
	}
	pub, _ := crypto.UnmarshalPubkey(sigPublicKey)
	addr := crypto.PubkeyToAddress(*pub)

	fmt.Println("addr:", addr)
	return addr == address
}

func ValidSignerV1(signature []byte, message string, address common.Address) bool {
	data := []byte(message)
	//hash := crypto.Keccak256Hash(data)
	//fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
	if len(signature) != crypto.SignatureLength {
		fmt.Println(fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength))
	}
	if signature[crypto.RecoveryIDOffset] != 27 && signature[crypto.RecoveryIDOffset] != 28 {
		fmt.Println(fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)"))
	}
	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	//fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
	pub, err := crypto.SigToPub(accounts.TextHash(data), signature)
	if err != nil {
		fmt.Println(err)
	}
	//sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	//if err != nil {
	//	return false
	//}
	//pub, _ := crypto.UnmarshalPubkey(sigPublicKey)

	addr := crypto.PubkeyToAddress(*pub)

	fmt.Println("addr:", addr)
	return addr == address
}
