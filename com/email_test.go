package com

import (
	"github.com/bwmarrin/snowflake"
	"github.com/golang/glog"
	"github.com/shopspring/decimal"
	"github.com/wumansgy/goEncrypt/rsa"
	"time"

	//"crypto/rsa"
	"fmt"
	sn "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"sync"
	"testing"
)

type Msg struct {
	Index int
	Vale  int
}

func TestSign1(*testing.T) {
	sig, _ := hexutil.Decode("0x6274e7a92311176180ced575be04b1c60bc239ed6fe9a212f5e8548e7207a7222d9aa4e3300874c56ac368ac940c0cba62c8d66250a0cdab8e700ac702ad98a61b")
	b := ValidSignerV1(sig, "0x3581930D218107F5C28D68A160F79C8a94225DDc", common.HexToAddress("0x3581930D218107F5C28D68A160F79C8a94225DDc"))
	fmt.Println(b)
}

func TestSign(*testing.T) {
	//	fmt.Println(common.IsHexAddress("0xe8e3a028903e5435bb2cd7da30119d37eef66999"))

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//}

	//	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	b := ValidSigner(signature, "hello", common.HexToAddress("0x96216849c49358B10257cb55b28eA603c874b05E"))
	fmt.Println(b)
	//sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pub, _ := crypto.UnmarshalPubkey(sigPublicKey)
	//addr := crypto.PubkeyToAddress(*pub)
	//fmt.Println(addr)
	//matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	//fmt.Println(matches) // true
	//
	//sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	//matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	//fmt.Println(matches) // true
	//
	//signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	//verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	//fmt.Println(verified) // true
}

func TestJs1(t *testing.T) {

	fmt.Println(Encode(uint64(50)))
	//
	//
	a := common.HexToAddress("0xE8e3a028903e5435bB2CD7DA30119d37eEf66999")
	b := common.HexToAddress("0xe8e3a028903e5435bb2cd7da30119d37eef66999")
	if a == b {
		fmt.Println("hello ==")
		return
	}

	fmt.Println("hello ==")

	ch := make(chan int, 2)
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 3, 3, 3, 3, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5}

	go func() {
		for v := range data {
			ch <- v
		}
		close(ch)
	}()
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			for l := range ch {
				fmt.Println(l)
			}
			w.Done()
		}(wg)
	}
	wg.Wait()
	fmt.Println("hi finished!")
}
func TestJs(t *testing.T) {
	jsfile := "./helper.js"
	bytes, err := ioutil.ReadFile(jsfile)
	if err != nil {
		fmt.Println(err)
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	_, err = vm.Call("loadData", nil)
	if err != nil {
		fmt.Println(err)
	}
	value, err := vm.Call("getRoot", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value.String())

}

//[{"address":"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266","amount":0,"index":0},{"address":"0x07cB5cD417d8EB9373849e8F482Cb2031d9F1e43","amount":1,"index":1},{"address":"0xE8e3a028903e5435bB2CD7DA30119d37eEf66999","amount":2,"index":2}];

func TestEmail(t *testing.T) {
	msg := "床前明月光，疑是地上霜，举头望明月，低头思故乡"
	rsaBase64Key, err := rsa.GenerateRsaKeyBase64(1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	base64Sign, err := rsa.RsaSignBase64([]byte(msg), rsaBase64Key.PrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("rsa签名的base64为:\n%s\n", base64Sign)
	res := rsa.RsaVerifySignBase64([]byte(msg), base64Sign, rsaBase64Key.PublicKey)
	fmt.Printf("rsa验签结果为:\n%v\n", res)

	fmt.Println("priv:", rsaBase64Key.PrivateKey)
	fmt.Println("public:", rsaBase64Key.PublicKey)
}

func TestSnow(t *testing.T) {
	for i := 0; i < 1000; i++ {
		node, err := snowflake.NewNode(1)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Generate a snowflake ID.
		id := node.Generate()

		// Print out the ID in a few different ways.
		fmt.Printf("Int64  ID: %d\n", id)
	}

}
func TestBwmarrinSnowflakeLoad(t *testing.T) {
	var wg sync.WaitGroup

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	var check sync.Map
	t1 := time.Now()
	for i := 0; i < 200000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//node, err = snowflake.NewNode(1)
			//if err != nil {
			//	fmt.Println(err)
			//}
			val := node.Generate().Int64()

			strId := fmt.Sprintf("%v", val)

			strId = strId[len(strId)-15:]
			dec, _ := decimal.NewFromString(strId)
			val = dec.IntPart()
			if _, ok := check.Load(val); ok {
				// id冲突检查
				log.Println(fmt.Errorf("error#unique: val:%v", val))
				return
			}
			fmt.Println(val)
			check.Store(val, 0)
			if val == 0 {
				log.Println(fmt.Errorf("error"))
				return
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(t1)
	println(int64(elapsed))
}

func TestSnowflakeLoad(t *testing.T) {
	var wg sync.WaitGroup
	//node, err := snowflake.NewNode(1)
	s, err := sn.NewSnowflake(int64(0), int64(0))
	if err != nil {
		glog.Error(err)
		return
	}
	var check sync.Map
	t1 := time.Now()
	for i := 0; i < 200000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := s.NextVal()
			if _, ok := check.Load(val); ok {
				// id冲突检查
				fmt.Println(fmt.Errorf("error#unique: val:%v", val))
				return
			}
			fmt.Println(val)
			check.Store(val, 0)
			if val == 0 {
				glog.Error(fmt.Errorf("error"))
				return
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(t1)
	println(int64(elapsed))
}
