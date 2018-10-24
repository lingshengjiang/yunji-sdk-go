package client

// 订单详情
type OrderDetail struct {
	Qty int `json:"qty"` // 必须,数量
	Id string `json:"id"` // 必须,子订单编号
	TradeOrderDetailId string `json:"tradeOrderDetailId"` // TU订单编号
	SkuNo string `json:"skuNo"` // 必须,商家编码
	BarCode string `json:"barCode"` // 必须,商品条码
	ReturnStatus int `json:"returnStatus"` // 必须,退款状态,0=无退款,1=申请退款,3=已退款,4=已关闭
	Name string `json:"name"` // 必须,商品名称 （预售标拼接：【预售 *日期发货】+ 商品名称）
	ItemPrice float64 `json:"itemPrice"` // 可选,商品单价
	TotalPrice float64 `json:"totalPrice"` // 可选,商品总金额
	Description string `json:"description"` // 可选,描述
	Unit string `json:"unit"` // 可选,商品单位
	Substract float64 `json:"substract"` // 可选,立减金额
	MeetSubstract float64 `json:"meetSubstract"` // 可选,满减
	MergePrice float64 `json:"mergePrice"` // 可选,任选优惠
	Coupon float64 `json:"coupon"` // 可选,优惠券抵扣
	TaxPrice float64 `json:"taxPrice"` // 可选,税金
	LogisticsPrice float64 `json:"logisticsPrice"` // 可选,邮费
	OrderLineNo float64 `json:"orderLineNo"` // 可选,商品编号
}
