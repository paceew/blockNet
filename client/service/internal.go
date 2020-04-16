package service

import (
	"fmt"
	"time"

	ch "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

//用eventID注册链码事件
func regitserEvent(client *ch.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

//用eventID接受链码事件结果
func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}
