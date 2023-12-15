package contract

import (
	"context"
	//	"depay/contract"
	"depay/model"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"time"

	//"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"math/big"

	"strings"
)

func FilOne(host string, conAddr []string) {
	//9601359
	//12508458
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Println(err)
		return
	}
	height, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	blockHeight := &model.BlockHeight{}
	err = model.DB.Model(model.BlockHeight{}).First(blockHeight).Error
	if err != nil {
		log.Println(err)
		return
	}
	if blockHeight.Height < int64(height) {
		for i := blockHeight.Height; i <= int64(height); i++ {
			FilLog(conAddr, uint64(i), client)
		}
	}

}
func toEthDbAmount(amount decimal.Decimal) decimal.Decimal {
	return amount.Div(decimal.New(1, 18))
}
func toUsdtDbAmount(amount decimal.Decimal) decimal.Decimal {
	return amount.Div(decimal.New(1, 6))
}

func FilLog(conAddr []string, height uint64, client *ethclient.Client) {

	addresses := make([]common.Address, 0)
	for _, addr := range conAddr {
		contractAddress := common.HexToAddress(addr)
		addresses = append(addresses, contractAddress)
	}
	PayOrderEventSignature := []byte("PayOrderEvent(uint256,address,address,address,uint256,uint256,uint256,uint256)")
	payOrderEvent := crypto.Keccak256Hash(PayOrderEventSignature)

	SubScribeEventSignature := []byte("SubScribe(address,address)")
	subScribeEvent := crypto.Keccak256Hash(SubScribeEventSignature)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(height)),
		ToBlock:   big.NewInt(int64(height)),
		Addresses: addresses,
		Topics:    [][]common.Hash{{payOrderEvent, subScribeEvent}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
		return
	}

	payAbi, err := abi.JSON(strings.NewReader(PayABI))
	if err != nil {
		log.Println(err)
		return
	}

	tx := model.DB.Begin()
	for _, vLog := range logs {
		fmt.Println("height:", vLog.BlockNumber)

		for _, topic := range vLog.Topics {

			if topic == payOrderEvent {
				event := PayPayOrderEvent{}
				err := payAbi.UnpackIntoInterface(&event, "PayOrderEvent", vLog.Data)
				if err != nil {
					log.Println("UnpackIntoInterface PayOrderEvent", err)
					return
				}

				order := &model.PayOrder{}
				if err := model.DB.Where("order_id=?", event.OrderId.Int64()).First(order).Error; err != nil {
					log.Printf("[PayOrderEvent] find order err,id %v,err is %v", event.OrderId.Int64(), err)
					continue
				}
				log.Info("get payoder event")
				if common.HexToAddress(order.MerchantAddress) != event.Merchant {
					log.Info("[PayOrderEvent],merchant address not equal")
					continue
				}
				order.TokenAmount = toEthDbAmount(decimal.NewFromBigInt(event.TokenAmount, 0))
				order.TokenAddress = event.PayToken.String()
				order.UserAddress = event.User.String()
				order.SwapAmount = toUsdtDbAmount(decimal.NewFromBigInt(event.SwapAmount, 0))
				order.PayedUsdt = toUsdtDbAmount(order.PayedUsdt.Add(decimal.NewFromBigInt(event.PayAmount, 0)))
				order.UpdateTime = time.Now()
				if order.PayedUsdt.Cmp(order.UsdtAmount) > 0 {
					order.Status = model.PAYED
				}
				if order.PayedUsdt.GreaterThan(decimal.Zero) && order.UsdtAmount.GreaterThan(order.PayedUsdt) {
					order.Status = model.PART
				}

				if err := tx.Where("id=?", order.Id).Save(order).Error; err != nil {
					log.Println("[PayOrderEvent] 修改支付订单错误：", order.OrderId)
					continue
				}
				//ChangeHeight = vLog.BlockNumber
			}
			if topic == subScribeEvent {
				event := PaySubScribe{}
				err := payAbi.UnpackIntoInterface(&event, "SubScribe", vLog.Data)
				if err != nil {
					log.Println(err)
					return
				}
				sub := &model.Subscribes{
					MerchantAddress: event.Merchant.String(),
					UserAddress:     event.User.String(),
					CreateTime:      time.Now(),
				}
				if err := tx.Create(sub).Error; err != nil {
					log.Println("add sub err:", event.Merchant, event.User)
					continue
				}

				//ChangeHeight = vLog.BlockNumber
			}

		}

		//fmt.Println(topics[0]) //7
	}
	blHeight := &model.BlockHeight{
		Height: int64(height),
	}
	if err := tx.Save(blHeight); err != nil {
		log.Error(err)
		tx.Rollback()

	}
	tx.Commit()
}
