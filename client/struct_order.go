package client

// 订单数据
type Order struct {
	OrderId string `json:"orderId"` // 必须,订单编号
	TradeOrderId string `json:"tradeOrderId"` // 可选,平台订单号
	PayId string `json:"payId"` // 可选,支付单号
	PayTime string `json:"payTime"` // 必须,支付时间
	Status int `json:"status"` // 必须,订单状态,40=待发货,50=已发货
	ReceiverProvince string `json:"receiverProvince"` // 必须,收货人省份
	ReceiverCity string `json:"receiverCity"` // 必须,收货人城市
	ReceiverArea string `json:"receiverArea"` // 必须,收货人地区
	ReceiverAddress string `json:"receiverAddress"` // 必填,收货人地址
	ReceiverMobile string `json:"receiverMobile"` // 可选,收货人手机（电话和手机必存在一个）
	ReceiverPhone string `json:"receiverPhone"` // 可选,收货人电话（电话和手机必存在一个）
	ReceiverName string `json:"receiverName"` // 必填,收货人名字
	ReceiverZipCode string `json:"receiverZipCode"` // 可选,收货人邮编
	BuyerComment string `json:"buyerComment"` // 可选,买家留言
	IdCardType int `json:"idCardType"` // 可选,证件类型:1=身份证
	CardNumber string `json:"cardNumber"` // 可选,证件号码
	TotalPrice float64 `json:"totalPrice"` // 可选,订单总金额(含邮费和税费)
	ItemPrice float64 `json:"itemPrice"` // 可选,商品总金额
	TaxPrice float64 `json:"taxPrice"` 	// 可选,税金
	LogisticsPrice float64 `json:"logisticsPrice"` // 可选,邮费
	PayEntName string  `json:"payEntName"` // 可选,企业支付名称
	PayEntNo string `json:"payEntNo"` 	// 可选,企业支付编码
	HasInvoice int `json:"hasInvoice"`	// 可选,发票:0=不需要,1=需要
	InvoiceHead	string `json:"invoiceHead"` // 可选,发票抬头
	InvoiceContent string `json:"invoiceContent"`	// 可选,发票内容
	CreateTime string `json:"createTime"` // 必填,订单创建时间
	ModifyTime string `json:"modifyTime"`	// 必填,订单修改时间
	PayMoney float64 `json:"payMoney"` // 可选,实际支付金额（国内订单暂不开放，默认传“0”）
	PayerName string `json:"payerName"`	// 可选,订购人名字
	PayerCardId string `json:"payerCardId"` // 可选,订购人身份证
	Substract float64 `json:"substract"` // 可选,立减金额
	MeetSubstract 	float64 `json:"meetSubstract"` // 可选,满减
	MergePrice 	float64 `json:"mergePrice"` // 可选,任选优惠
	Coupon 	float64 `json:"coupon"` // 可选,优惠券抵扣
	InvoiceAbilityAmount float64 `json:"invoiceAbilityAmount"` // 可选,可开票金额(需申请后开放)
	OrderDetail []OrderDetail `json:"orderDetail"` // 必填,订单详情，参考OrderDetail描述
}
