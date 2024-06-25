// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
}

type CreateResponse struct {
	Id      int    `json:"id"`
	PlayUrl string `json:"payUrl"`
}

type GoodsInfo struct {
	GoodsId int64 `json:"goodsId"`
	Num     int64 `json:"num"`
	Checked bool  `json:"checked"`
}

type OrderDetailRequst struct {
	Id int `path:"id"`
}

type OrderInfoResponse struct {
	Data interface{} `json:"data"`
}

type OrderListRequest struct {
}

type Request struct {
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}