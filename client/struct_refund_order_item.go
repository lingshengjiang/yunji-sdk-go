package client

// 售后订单
type RefundOrderItem struct {
	RefundId string `json:"refundId"` // 必须,售后订单编号
	OrderId string `json:"orderId"` // 必须,订单编号
	SkuNo string `json:"skuNo"` // 必须,商家编码
	BarCode string `json:"barCode"` // 可选,商品条码
	Name string `json:"name"` // 可选,商品名称
	ReturnStatus int `json:returnStatus` // 必填,订单退款状态,1=申请退款,3=已退款,4=已关闭
	Price float64 `json:"price"` // 可选,商品价格
	ReturnQty int `json:"returnQty"` // 可选,退货数量
	ReturnMoney float64 `json:"returnMoney"` // 可选,退款金额
	LogisticsCompanyName string `json:"logisticsCompanyName"` // 可选,买家寄回物流公司
	LogisticsNumber string `json:"logisticsNumber"` // 可选,买家寄回物流单号
	BuerDesc string `json:"buerDesc"` // 可选,买家售后备注
	CreateTime string `json:"createTime"` // 必须,售后创建时间
	ReturnType int `json:"returnType"` // 必须,订单退款类型,0=无,1=仅退款,2=退款不退货,3=退货退款
	ReturnReason string `json:"returnReason"` // 可选,售后原因
	ModifyTime string `json:"modifyTime"` // 必须,售后修改时间
}
