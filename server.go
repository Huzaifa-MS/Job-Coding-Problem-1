package main

import (
	"context"
	"fmt"
	api "jobproblem1/api"
	"strconv"
	"sync"
)

type ReceiptID struct {
	id string
	api.ReceiptsProcessPostRes
}
type ReceiptPoints struct {
	points int
	api.ReceiptsIDPointsGetRes
}

type ReceiptMetaData struct {
	receipt api.Receipt
	id      string
	points  int
}
type ReceiptService struct {
	receipts map[string]ReceiptMetaData
	mu       sync.Mutex // was in ogen docs. Useful for making sure processing happens independently
}

func (rs *ReceiptService) processReceipt(r *api.Receipt) (ReceiptID, error) {
	// Create id and process points here
	rmd := ReceiptMetaData{
		receipt: *r,
	}
	rmd.id = r.Retailer + r.Total
	rid := ReceiptID{
		id: rmd.id,
	}
	points, err := strconv.Atoi(r.Total)
	if err != nil {
		return ReceiptID{}, err
	}
	rmd.points = points
	rs.receipts[rmd.id] = rmd
	return rid, err
}
func (rs *ReceiptService) getReceiptPoints(params api.ReceiptsIDPointsGetParams) ReceiptPoints {
	rid := params.ID
	println(rid)
	print(rs)
	receipt := rs.receipts[rid]
	fmt.Printf("%+v\n", receipt)
	rpoints := ReceiptPoints{points: receipt.points}
	fmt.Printf("%+v\n", rpoints)
	return rpoints
}

func (rs *ReceiptService) ReceiptsProcessPost(ctx context.Context, r *api.Receipt) (api.ReceiptsProcessPostRes, error) {
	println("ReceiptsProcessPost")
	rid, err := rs.processReceipt(r)
	return rid, err
}

func (rs *ReceiptService) ReceiptsIDPointsGet(ctx context.Context, params api.ReceiptsIDPointsGetParams) (api.ReceiptsIDPointsGetRes, error) {
	println("ReceiptsIDPointsGetParams")
	rpoints := rs.getReceiptPoints(params)
	return rpoints, nil
}
