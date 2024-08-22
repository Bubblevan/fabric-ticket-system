package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Order describes basic details of an order
type Order struct {
	UserID     uint      `json:"userID"`
	TicketID   uint      `json:"ticketID"`
	Num        int       `json:"num"`
	TotalPrice float64   `json:"totalPrice"`
	OrderDate  time.Time `json:"orderDate"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Order
}

// InitLedger adds a base set of orders to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fixedTime, _ := time.Parse(time.RFC3339, "2024-07-19T12:00:00Z")
	orders := []Order{
		{UserID: 1, TicketID: 1, Num: 2, TotalPrice: 60.0, OrderDate: fixedTime},
		{UserID: 2, TicketID: 2, Num: 1, TotalPrice: 30.0, OrderDate: fixedTime},
		{UserID: 3, TicketID: 3, Num: 3, TotalPrice: 90.0, OrderDate: fixedTime},
	}

	for i, order := range orders {
		orderAsBytes, _ := json.Marshal(order)
		err := ctx.GetStub().PutState("ORDER"+strconv.Itoa(i), orderAsBytes)

		if err != nil {
			return fmt.Errorf("error: failed to process the request. %s", err.Error())
		}
	}

	return nil
}

// CreateOrder adds a new order to the world state with given details
func (s *SmartContract) CreateOrder(ctx contractapi.TransactionContextInterface, orderID string, userID uint, ticketID uint, num int, totalPrice float64, orderDate string) (*Order, error) {
	parsedOrderDate, err := time.Parse(time.RFC3339, orderDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse orderDate: %s", err.Error())
	}

	order := Order{
		UserID:     userID,
		TicketID:   ticketID,
		Num:        num,
		TotalPrice: totalPrice,
		OrderDate:  parsedOrderDate,
	}

	orderAsBytes, _ := json.Marshal(order)

	err = ctx.GetStub().PutState(orderID, orderAsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to put order state: %s", err.Error())
	}

	return &order, nil
}

// QueryOrder returns the order stored in the world state with given id
func (s *SmartContract) QueryOrder(ctx contractapi.TransactionContextInterface, orderID string) (*Order, error) {
	orderAsBytes, err := ctx.GetStub().GetState(orderID)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if orderAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", orderID)
	}

	order := new(Order)
	_ = json.Unmarshal(orderAsBytes, order)

	return order, nil
}

// QueryAllOrders returns all orders found in world state
func (s *SmartContract) QueryAllOrders(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		order := new(Order)
		_ = json.Unmarshal(queryResponse.Value, order)

		queryResult := QueryResult{Key: queryResponse.Key, Record: order}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeTicketsOwner updates the userID field of order with given orderID in world state
func (s *SmartContract) ChangeTicketsOwner(ctx contractapi.TransactionContextInterface, orderID string, newUserID uint) error {
	order, err := s.QueryOrder(ctx, orderID)

	if err != nil {
		return err
	}

	order.UserID = newUserID

	orderAsBytes, _ := json.Marshal(order)

	return ctx.GetStub().PutState(orderID, orderAsBytes)
}

// QueryOrdersByUserID returns all orders for a given userID
func (s *SmartContract) QueryOrdersByUserID(ctx contractapi.TransactionContextInterface, userID uint) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		order := new(Order)
		_ = json.Unmarshal(queryResponse.Value, order)

		if order.UserID == userID {
			queryResult := QueryResult{Key: queryResponse.Key, Record: order}
			results = append(results, queryResult)
		}
	}

	return results, nil
}

// QueryOrdersByTicketID returns all orders for a given ticketID
func (s *SmartContract) QueryOrdersByTicketID(ctx contractapi.TransactionContextInterface, ticketID uint) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		order := new(Order)
		_ = json.Unmarshal(queryResponse.Value, order)

		if order.TicketID == ticketID {
			queryResult := QueryResult{Key: queryResponse.Key, Record: order}
			results = append(results, queryResult)
		}
	}

	return results, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabasg chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabasg chaincode: %s", err.Error())
	}
}
