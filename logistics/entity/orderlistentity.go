package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ShippingDocumentParameterRequestOrderListEntity
type ShippingDocumentParameterRequestOrderListEntity struct {
	OrderSn       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
}

//String
func (s ShippingDocumentParameterRequestOrderListEntity) String() string {
	return lib.ObjectToString(s)
}

//CreateShippingDocumentRequestOrderListEntity
type CreateShippingDocumentRequestOrderListEntity struct {
	OrderSn              string `json:"order_sn"`
	PackageNumber        string `json:"package_number"`
	TrackingNumber       string `json:"tracking_number"`
	ShippingDocumentType string `json:"shipping_document_type"`
}

//String
func (c CreateShippingDocumentRequestOrderListEntity) String() string {
	return lib.ObjectToString(c)
}

//GetShippingDocumentResultRequestOrderListEntity
type GetShippingDocumentResultRequestOrderListEntity struct {
	OrderSn              string `json:"order_sn"`
	PackageNumber        string `json:"package_number"`
	ShippingDocumentType string `json:"shipping_document_type"`
}

//String
func (c GetShippingDocumentResultRequestOrderListEntity) String() string {
	return lib.ObjectToString(c)
}

//DownloadShippingDocumentRequestOrderListEntity
type DownloadShippingDocumentRequestOrderListEntity struct {
	OrderList            []ShippingDocumentParameterRequestOrderListEntity `json:"order_list"`
	ShippingDocumentType string                                            `json:"shipping_document_type"`
}

//String
func (d DownloadShippingDocumentRequestOrderListEntity) String() string {
	return lib.ObjectToString(d)
}

//BatchShipOrderRequestOrderListEntity
type BatchShipOrderRequestOrderListEntity struct {
	OrderSn string `json:"order_sn"`
	//PackageNumber string `json:"package_number"`
}

//String
func (b BatchShipOrderRequestOrderListEntity) String() string {
	return lib.ObjectToString(b)
}
