package api

import (
	"backend/model"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"gopkg.in/yaml.v2"
)

var (
	SDK           *fabsdk.FabricSDK
	channelClient *channel.Client
	channelName   = "mychannel"
	chaincodeName = "fabasg"
	orgName       = "Org1"
	orgAdmin      = "Admin"
	org1Peer0     = "peer0.org1.example.com"
	org2Peer0     = "peer0.org2.example.com"
)

type Ticket = model.Ticket
type Order = model.Order

func ChannelExecute(funcName string, args [][]byte) (channel.Response, error) {
	var err error
	configPath := "./config/config.yaml"
	configProvider := config.FromFile(configPath)
	SDK, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("Failed to create new SDK: %s", err)
	}
	ctx := SDK.ChannelContext(channelName, fabsdk.WithOrg(orgName), fabsdk.WithUser(orgAdmin))
	channelClient, err = channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}

	response, err := channelClient.Execute(channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        args,
	})
	if err != nil {
		return response, err
	}
	SDK.Close()
	return response, nil
}

func ChannelQuery(funcName string, args [][]byte) (channel.Response, error) {
	configPath := "./config/config.yaml"
	configProvider := config.FromFile(configPath)
	SDK, err := fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("Failed to create new SDK: %s", err)
	}
	defer SDK.Close()

	ctx := SDK.ChannelContext(channelName, fabsdk.WithOrg(orgName), fabsdk.WithUser(orgAdmin))
	channelClient, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}

	response, err := channelClient.Query(channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        args,
	})
	if err != nil {
		return response, err
	}

	return response, nil
}

func BlockchainRoutes(r *gin.Engine, configPath string) {
	r.GET("/stats", func(c *gin.Context) {
		blockCount, transactionCount, nodeCount, chaincodeCount, err := getBlockchainStats(configPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"blockCount":       blockCount,
			"transactionCount": transactionCount,
			"nodeCount":        nodeCount,
			"chaincodeCount":   chaincodeCount,
		})
	})
}

type Config struct {
	Channels map[string]ChannelConfig `yaml:"channels"`
	Peers    map[string]PeerConfig    `yaml:"peers"`
}

type ChannelConfig struct {
	Peers map[string]PeerInfo `yaml:"peers"`
}

type PeerInfo struct {
	EndorsingPeer  bool `yaml:"endorsingPeer"`
	ChaincodeQuery bool `yaml:"chaincodeQuery"`
	LedgerQuery    bool `yaml:"ledgerQuery"`
	EventSource    bool `yaml:"eventSource"`
}

type PeerConfig struct {
	URL string `yaml:"url"`
}

func getBlockchainStats(configPath string) (blockCount uint64, transactionCount int64, nodeCount int, chaincodeCount int, err error) {
	// 加载配置文件
	configProvider := config.FromFile(configPath)

	// 读取配置文件内容
	configData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Failed to read config file: %s\n", err)
		return
	}

	// 解析配置文件
	var cfg Config
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		fmt.Printf("Failed to unmarshal config file: %s\n", err)
		return
	}

	// 计算 nodeCount
	nodeCount = 0
	for _, peer := range cfg.Peers {
		if peer.URL != "" {
			nodeCount++
		}
	}

	// 创建SDK实例
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		return
	}
	defer sdk.Close()

	// 创建通道上下文
	channelName := "mychannel"
	orgName := "Org1"
	user := "Admin"
	ctxProvider := sdk.ChannelContext(channelName, fabsdk.WithUser(user), fabsdk.WithOrg(orgName))

	// 创建账本客户端
	ledgerClient, err := ledger.New(ctxProvider)
	if err != nil {
		fmt.Printf("Failed to create new ledger client: %s\n", err)
		return
	}

	// 创建资源管理客户端
	resMgmtClient, err := resmgmt.New(sdk.Context(fabsdk.WithUser(user), fabsdk.WithOrg(orgName)))
	if err != nil {
		fmt.Printf("Failed to create new resource management client: %s\n", err)
		return
	}

	// 获取 blockCount
	info, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("Failed to query channel info: %s\n", err)
		return
	}
	blockCount = info.BCI.Height

	// 获取 transactionCount
	block, err := ledgerClient.QueryBlock(blockCount - 1)
	if err != nil {
		fmt.Printf("Failed to query block: %s\n", err)
		return
	}
	for _, tx := range block.GetData().GetData() {
		if tx != nil {
			transactionCount++
		}
	}

	// 获取 chaincodeCount
	chaincodes, err := resMgmtClient.QueryInstalledChaincodes()
	if err != nil {
		fmt.Printf("Failed to query installed chaincodes: %s\n", err)
		return
	}
	chaincodeCount = len(chaincodes.Chaincodes)

	return
}
