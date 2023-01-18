package shopeego

import (
	"github.com/wjp-letgo/shopeego/accounthealth"
	accounthealthentity "github.com/wjp-letgo/shopeego/accounthealth/entity"
	"github.com/wjp-letgo/shopeego/addondeal"
	addondealentity "github.com/wjp-letgo/shopeego/addondeal/entity"
	"github.com/wjp-letgo/shopeego/auth"
	authEntity "github.com/wjp-letgo/shopeego/auth/entity"
	"github.com/wjp-letgo/shopeego/bundledeal"
	bundledealentity "github.com/wjp-letgo/shopeego/bundledeal/entity"
	"github.com/wjp-letgo/shopeego/commonentity"
	shopeeConfig "github.com/wjp-letgo/shopeego/config"
	"github.com/wjp-letgo/shopeego/discount"
	discountentity "github.com/wjp-letgo/shopeego/discount/entity"
	"github.com/wjp-letgo/shopeego/firstmile"
	firstmileentity "github.com/wjp-letgo/shopeego/firstmile/entity"
	"github.com/wjp-letgo/shopeego/followprize"
	followprizeentity "github.com/wjp-letgo/shopeego/followprize/entity"
	"github.com/wjp-letgo/shopeego/globalproduct"
	globalProductEntity "github.com/wjp-letgo/shopeego/globalproduct/entity"
	"github.com/wjp-letgo/shopeego/logistics"
	logisticsEntity "github.com/wjp-letgo/shopeego/logistics/entity"
	"github.com/wjp-letgo/shopeego/mediaspace"
	mediaspaceEntity "github.com/wjp-letgo/shopeego/mediaspace/entity"
	"github.com/wjp-letgo/shopeego/merchant"
	merchantEntity "github.com/wjp-letgo/shopeego/merchant/entity"
	"github.com/wjp-letgo/shopeego/order"
	orderEntity "github.com/wjp-letgo/shopeego/order/entity"
	"github.com/wjp-letgo/shopeego/payment"
	paymententity "github.com/wjp-letgo/shopeego/payment/entity"
	"github.com/wjp-letgo/shopeego/product"
	productEntity "github.com/wjp-letgo/shopeego/product/entity"
	"github.com/wjp-letgo/shopeego/public"
	publicentity "github.com/wjp-letgo/shopeego/public/entity"
	"github.com/wjp-letgo/shopeego/push"
	pushentity "github.com/wjp-letgo/shopeego/push/entity"
	"github.com/wjp-letgo/shopeego/returns"
	returnsentity "github.com/wjp-letgo/shopeego/returns/entity"
	"github.com/wjp-letgo/shopeego/sellerchat"
	sellerchatentity "github.com/wjp-letgo/shopeego/sellerchat/entity"
	"github.com/wjp-letgo/shopeego/shop"
	shopEntity "github.com/wjp-letgo/shopeego/shop/entity"
	"github.com/wjp-letgo/shopeego/shopcategory"
	shopcategoryentity "github.com/wjp-letgo/shopeego/shopcategory/entity"
	"github.com/wjp-letgo/shopeego/toppicks"
	toppicksentity "github.com/wjp-letgo/shopeego/toppicks/entity"
	"github.com/wjp-letgo/shopeego/voucher"
	voucherentity "github.com/wjp-letgo/shopeego/voucher/entity"
)

//Shopeer
type Shopeer interface {
	GetPartnerID() int64
	//auth
	AuthorizationURL() string
	GetAccesstoken(code string, shopID int64) authEntity.GetAccessTokenResult
	RefreshAccessToken(shop commonentity.ShopInfo) authEntity.RefreshAccessTokenResult
	//order
	GetOrderList(
		timeRangeField order.TimeRangeField,
		timeFrom, timeTo, pageSize int,
		cursor int,
		orderStatus order.OrderStatus,
		responseOptionalFields string) orderEntity.GetOrderListResult
	GetShipmentList(cursor string, pageSize int) orderEntity.GetShipmentListResult
	GetOrderDetail(orderSnList []string, responseOptionalFields ...string) orderEntity.GetOrderDetailResult
	SplitOrder(orderSn string, packageList []orderEntity.PackageListRequestEntity) orderEntity.SplitOrderResult
	UnSplitOrder(orderSn string) orderEntity.UnSplitOrderResult
	CancelOrder(orderSn string, cancelReason order.CancelReason, itemList []orderEntity.CancelOrderRequestEntity) orderEntity.CancelOrderResult
	HandleBuyerCancellation(orderSn string, operation order.Operation) orderEntity.HandleBuyerCancellationResult
	SetNote(orderSn, note string) orderEntity.SetNoteResult
	AddInvoiceData(orderSn string, invoiceData orderEntity.InvoiceDataEntity) orderEntity.AddInvoiceDataResult
	//logistics
	GetShippingParameter(orderSn string) logisticsEntity.GetShippingParameterResult
	GetTrackingNumber(orderSn, packageNumber string, responseOptionalFields ...string) logisticsEntity.GetTrackingNumberResult
	ShipOrder(orderSn, packageNumber string, pickup *logisticsEntity.ShipOrderRequestPickupEntity, dropoff *logisticsEntity.ShipOrderRequestDropoffEntity, nonIntegrated *logisticsEntity.ShipOrderRequestNonIntegratedEntity) logisticsEntity.ShipOrderResult
	UpdateShippingOrder(orderSn, packageNumber string, pickup *logisticsEntity.UpdateShippingOrderRequestPickupEntity) logisticsEntity.UpdateShippingOrderResult
	GetShippingDocumentParameter(orderList []logisticsEntity.ShippingDocumentParameterRequestOrderListEntity) logisticsEntity.GetShippingDocumentParameterResult
	CreateShippingDocument(orderList []logisticsEntity.CreateShippingDocumentRequestOrderListEntity) logisticsEntity.CreateShippingDocumentResult
	GetShippingDocumentResult(orderList []logisticsEntity.GetShippingDocumentResultRequestOrderListEntity) logisticsEntity.GetShippingDocumentResult
	DownloadShippingDocument(orderList logisticsEntity.DownloadShippingDocumentRequestOrderListEntity) logisticsEntity.DownloadShippingDocumentResult
	GetShippingDocumentInfo(orderSn, packageNumber string) logisticsEntity.GetShippingDocumentInfoResult
	GetTrackingInfo(orderSn, packageNumber string) logisticsEntity.GetTrackingInfoResult
	GetAddressList() logisticsEntity.GetAddressListResult
	SetAddressConfig(showPickupAddress bool, AddressTypeConfig logisticsEntity.AddressTypeConfigEntity) logisticsEntity.SetAddressConfigResult
	DeleteAddress(addressID int64) logisticsEntity.DeleteAddressResult
	GetChannelList() logisticsEntity.GetChannelListResult
	UpdateChannel(logisticsChannelID int64, enableds ...bool) logisticsEntity.UpdateChannelResult
	BatchShipOrder(orderList []logisticsEntity.BatchShipOrderRequestOrderListEntity, pickup *logisticsEntity.BatchShipOrderRequestPickupEntity, dropoff *logisticsEntity.BatchShipOrderRequestDropoffEntity, nonIntegrated *logisticsEntity.BatchShipOrderRequestNonIntegratedEntity) logisticsEntity.BatchShipOrderResult
	//product
	GetComment(itemID, commentID int64, cursor int, pageSize int) productEntity.GetCommentResult
	ReplyComment(commentList []productEntity.ReplyCommentRequestCommentEntity) productEntity.ReplyCommentResult
	GetItemBaseInfo(itemIdList []int64) productEntity.GetItemBaseInfoResult
	GetItemExtraInfo(itemIdList []int64) productEntity.GetItemExtraInfoResult
	GetItemList(offset, pageSize, updateTimeFrom, updateTimeTo int, itemStatus product.ItemStatus) productEntity.GetItemListResult
	DeleteItem(itemID int64) productEntity.DeleteItemResult
	AddItem(item productEntity.AddItemRequestItemEntity) productEntity.AddItemResult
	UpdateItem(item productEntity.UpdateItemRequestItemEntity) productEntity.UpdateItemResult
	UpdateItemSku(item productEntity.UpdateItemRequestItemV2Entity) productEntity.UpdateItemResult
	GetModelList(itemID int64) productEntity.GetModelListResult
	UpdateSizeChart(itemID int64, sizeChart string) productEntity.UpdateSizeChartResult
	UnlistItem(itemList []productEntity.UnlistItemItemListEntity) productEntity.UnlistItemResult
	BoostItem(itemIdList []int64) productEntity.BoostItemResult
	GetBoostedList() productEntity.GetBoostedListResult
	GetDtsLimit(categoryID int64) productEntity.GetDtsLimitResult
	GetItemLimit(categoryID int64) productEntity.GetItemLimitResult
	SupportSizeChart(category int64) productEntity.SupportSizeChartResult
	InitTierVariation(itemID int64, tierVariation []productEntity.TierVariationEntity, model []productEntity.InitTierVariationModelEntity) productEntity.InitTierVariationResult
	UpdateTierVariation(itemID int64, tierVariation []productEntity.TierVariationEntity) productEntity.UpdateTierVariationResult
	UpdateModel(itemID int64, model []productEntity.UpdateModelEntity) productEntity.UpdateModelResult
	AddModel(itemID int64, modelList []productEntity.InitTierVariationModelEntity) productEntity.AddModelResult
	DeleteModel(itemID, modelID int64) productEntity.DeleteModelResult
	UpdatePrice(itemID int64, priceList []productEntity.UpdatePricePriceInfoEntity) productEntity.UpdatePriceResult
	UpdateStock(itemID int64, stockList []productEntity.UpdateStockStockInfoEntity) productEntity.UpdateStockResult
	GetCategory(language string) productEntity.GetCategoryResult
	GetAttributes(language string, categoryID int64) productEntity.GetAttributesResult
	GetBrandList(offset, pageSize int, categoryID int64, status product.BrandStatus) productEntity.GetBrandListResult
	CategoryRecommend(itemName string) productEntity.CategoryRecommendResult
	GetItemPromotion(itemIdList []int64) productEntity.GetItemPromotionResult
	UpdateSipItemPrice(itemID int64, sipItemPrice []productEntity.SipItemPriceEntity) productEntity.UpdateSipItemPriceResult
	SearchItem(offset string, pageSize int, itemName string, attributeStatus product.AttributeStatus, itemSku string) productEntity.SearchItemResult
	
	//global_product
	GetGlobalCategory(language string) globalProductEntity.GetCategoryResult
	GetGlobalAttributes(language string, categoryID int64) globalProductEntity.GetAttributesResult
	GetGlobalBrandList(offset, pageSize int, categoryID int64, status globalproduct.BrandStatus) globalProductEntity.GetBrandListResult
	GetGlobalItemLimit(categoryID int64) globalProductEntity.GetGlobalItemLimitResult
	GetGlobalDtsLimit(categoryID int64) globalProductEntity.GetDtsLimitResult
	GetGlobalItemList(offset string, pageSize int) globalProductEntity.GetGlobalItemListResult
	GetGlobalItemInfo(globalItemIdList []int64) globalProductEntity.GetGlobalItemInfoResult
	AddGlobalItem(item globalProductEntity.AddItemRequestItemEntity) globalProductEntity.AddGlobalItemResult
	UpdateGlobalItem(item globalProductEntity.UpdateItemRequestItemEntity) globalProductEntity.AddGlobalItemResult
	DeleteGlobalItem(globalItemID int64) globalProductEntity.DeleteGlobalItemResult
	InitGlobalTierVariation(globalItemID int64, tierVariation globalProductEntity.TierVariationEntity, model globalProductEntity.InitTierVariationModelEntity) globalProductEntity.InitTierVariationResult
	UpdateGlobalTierVariation(globalItemID int64, tierVariation []globalProductEntity.TierVariationEntity) globalProductEntity.UpdateTierVariationResult
	AddGlobalModel(globalItemID int64, modelList []globalProductEntity.InitTierVariationModelEntity) globalProductEntity.AddGlobalModelResult
	UpdateGlobalModel(globalItemID int64, model []globalProductEntity.UpdateModelEntity) globalProductEntity.UpdateGlobalModelResult
	DeleteGlobalModel(globalItemID, globalModelID int64) globalProductEntity.DeleteGlobalModelResult
	GetGlobalModelList(globalItemID int64) globalProductEntity.GetModelListResult
	SupportGlobalSizeChart(categoryID int64) globalProductEntity.SupportGlobalSizeChartResult
	UpdateGlobalSizeChart(globalItemID int64, sizeChart string) globalProductEntity.UpdateGlobalSizeChartResult
	CreatePublishTask(globalItemID, shopID int64, shopRegion string, item globalProductEntity.CreatePublishTaskItemEntity) globalProductEntity.CreatePublishTaskResult
	GetPublishableShop(globalItemID int64) globalProductEntity.GetPublishableShopResult
	GetPublishTaskResult(publishTaskID int64) globalProductEntity.GetPublishTaskResult
	GetPublishedList(globalItemID int64) globalProductEntity.GetPublishedListResult
	UpdateGlobalPrice(globalItemID int64, priceList []globalProductEntity.PriceEntity) globalProductEntity.UpdateGlobalPriceResult
	UpdateGlobalStock(globalItemID int64, stockList []globalProductEntity.UpdateStockStockInfoEntity) globalProductEntity.UpdateGlobalStockResult
	SetSyncField(shopSyncList []globalProductEntity.ShopSyncEntity) globalProductEntity.SetSyncFieldResult
	GetGlobalItemID(shopID int64, itemIdList []int64) globalProductEntity.GetGlobalItemIDResult
	//media_space
	InitVideoUpload(fileMd5 string, fileSize int) mediaspaceEntity.InitVideoUploadResult
	UploadVideoPart(videoUploadID string, partSeq int, contentMd5 string, partContentPath string) mediaspaceEntity.UploadVideoPartResult
	CompleteVideoUpload(videoUploadID string, partSeqList []int, reportData mediaspaceEntity.ReportDataEntity) mediaspaceEntity.CompleteVideoUploadResult
	GetVideoUploadResult(videoUploadID string) mediaspaceEntity.GetVideoUploadResult
	CancelVideoUpload(videoUploadID string) mediaspaceEntity.CancelVideoUploadResult
	UploadImage(image string) mediaspaceEntity.UploadImageResult
	//shop
	GetShopInfo() shopEntity.GetShopInfoResult
	GetProfile() shopEntity.GetProfileResult
	UpdateProfile(shopName, shopLogo, description string) shopEntity.UpdateProfileResult
	//merchant
	GetMerchantInfo() merchantEntity.GetMerchantInfoResult
	GetShopListByMerchant(pageNo, pageSize int) merchantEntity.GetShopListByMerchantResult
	//firstmile
	GetUnbindOrderList(params firstmileentity.GetUnbindOrderListRequest) firstmileentity.GetUnbindOrderListResult
	GetDetail(params firstmileentity.GetDetailRequest) firstmileentity.GetDetailResult
	GenerateFirstMileTrackingNumber(params firstmileentity.GenerateFirstMileTrackingNumberRequest) firstmileentity.GenerateFirstMileTrackingNumberResult
	BindFirstMileTrackingNumber(params firstmileentity.BindFirstMileTrackingNumberRequest) firstmileentity.BindFirstMileTrackingNumberResult
	UnbindFirstMileTrackingNumber(params firstmileentity.UnbindFirstMileTrackingNumberRequest) firstmileentity.UnbindFirstMileTrackingNumberResult
	GetTrackingNumberList(params firstmileentity.GetTrackingNumberListRequest) firstmileentity.GetTrackingNumberListResult
	GetWaybill(params firstmileentity.GetWaybillRequest) firstmileentity.GetWaybillResult
	FGetChannelList(params firstmileentity.GetChannelListRequest) firstmileentity.GetChannelListResult
	//payment
	GetEscrowDetail(params paymententity.GetEscrowDetailRequest) paymententity.GetEscrowDetailResult
	SetShopInstallmentStatus(params paymententity.SetShopInstallmentStatusRequest) paymententity.SetShopInstallmentStatusResult
	GetShopInstallmentStatus() paymententity.GetShopInstallmentStatusResult
	GetPayoutDetail(params paymententity.GetPayoutDetailRequest) paymententity.GetPayoutDetailResult
	SetItemInstallmentStatus(params paymententity.SetItemInstallmentStatusRequest) paymententity.SetItemInstallmentStatusResult
	GetItemInstallmentStatus(params paymententity.GetItemInstallmentStatusRequest) paymententity.GetItemInstallmentStatusResult
	GetPaymentMethodList(params paymententity.GetPaymentMethodListRequest) paymententity.GetPaymentMethodListResult
	GetWalletTransactionList(params paymententity.GetWalletTransactionListRequest) paymententity.GetWalletTransactionListResult
	GetEscrowList(params paymententity.GetEscrowListRequest) paymententity.GetEscrowListResult
	//discount
	AddDiscount(params discountentity.AddDiscountRequest) discountentity.AddDiscountResult
	AddDiscountItem(params discountentity.AddDiscountItemRequest) discountentity.AddDiscountItemResult
	DeleteDiscount(params discountentity.DeleteDiscountRequest) discountentity.DeleteDiscountResult
	DeleteDiscountItem(params discountentity.DeleteDiscountItemRequest) discountentity.DeleteDiscountItemResult
	GetDiscount(params discountentity.GetDiscountRequest) discountentity.GetDiscountResult
	GetDiscountList(params discountentity.GetDiscountListRequest) discountentity.GetDiscountListResult
	UpdateDiscount(params discountentity.UpdateDiscountRequest) discountentity.UpdateDiscountResult
	UpdateDiscountItem(params discountentity.UpdateDiscountItemRequest) discountentity.UpdateDiscountItemResult
	EndDiscount(params discountentity.EndDiscountRequest) discountentity.EndDiscountResult
	//bundledeal
	AddBundleDeal(params bundledealentity.AddBundleDealRequest) bundledealentity.AddBundleDealResult
	AddBundleDealItem(params bundledealentity.AddBundleDealItemRequest) bundledealentity.AddBundleDealItemResult
	DeleteBundleDeal(params bundledealentity.DeleteBundleDealRequest) bundledealentity.DeleteBundleDealResult
	DeleteBundleDealItem(params bundledealentity.DeleteBundleDealItemRequest) bundledealentity.DeleteBundleDealItemResult
	EndBundleDeal(params bundledealentity.EndBundleDealRequest) bundledealentity.EndBundleDealResult
	GetBundleDealList(params bundledealentity.GetBundleDealListRequest) bundledealentity.GetBundleDealListResult
	GetBundleDeal(params bundledealentity.GetBundleDealRequest) bundledealentity.GetBundleDealResult
	GetBundleDealItem(params bundledealentity.GetBundleDealItemRequest) bundledealentity.GetBundleDealItemResult
	UpdateBundleDeal(params bundledealentity.UpdateBundleDealRequest) bundledealentity.UpdateBundleDealResult
	UpdateBundleDealItem(params bundledealentity.UpdateBundleDealItemRequest) bundledealentity.UpdateBundleDealItemResult
	//addondeal
	AddAddOnDealMainItem(params addondealentity.AddAddOnDealMainItemRequest) addondealentity.AddAddOnDealMainItemResult
	AddAddOnDealSubItem(params addondealentity.AddAddOnDealSubItemRequest) addondealentity.AddAddOnDealSubItemResult
	AddAddOnDeal(params addondealentity.AddAddOnDealRequest) addondealentity.AddAddOnDealResult
	DeleteAddOnDealMainItem(params addondealentity.DeleteAddOnDealMainItemRequest) addondealentity.DeleteAddOnDealMainItemResult
	DeleteAddOnDealSubItem(params addondealentity.DeleteAddOnDealSubItemRequest) addondealentity.DeleteAddOnDealSubItemResult
	DeleteAddOnDeal(params addondealentity.DeleteAddOnDealRequest) addondealentity.DeleteAddOnDealResult
	EndAddOnDeal(params addondealentity.EndAddOnDealRequest) addondealentity.EndAddOnDealResult
	GetAddOnDealList(params addondealentity.GetAddOnDealListRequest) addondealentity.GetAddOnDealListResult
	GetAddOnDealMainItem(params addondealentity.GetAddOnDealMainItemRequest) addondealentity.GetAddOnDealMainItemResult
	GetAddOnDealSubItem(params addondealentity.GetAddOnDealSubItemRequest) addondealentity.GetAddOnDealSubItemResult
	GetAddOnDeal(params addondealentity.GetAddOnDealRequest) addondealentity.GetAddOnDealResult
	UpdateAddOnDealMainItem(params addondealentity.UpdateAddOnDealMainItemRequest) addondealentity.UpdateAddOnDealMainItemResult
	UpdateAddOnDealSubItem(params addondealentity.UpdateAddOnDealSubItemRequest) addondealentity.UpdateAddOnDealSubItemResult
	UpdateAddOnDeal(params addondealentity.UpdateAddOnDealRequest) addondealentity.UpdateAddOnDealResult
	//voucher
	AddVoucher(params voucherentity.AddVoucherRequest) voucherentity.AddVoucherResult
	DeleteVoucher(params voucherentity.DeleteVoucherRequest) voucherentity.DeleteVoucherResult
	EndVoucher(params voucherentity.EndVoucherRequest) voucherentity.EndVoucherResult
	UpdateVoucher(params voucherentity.UpdateVoucherRequest) voucherentity.UpdateVoucherResult
	GetVoucherDetail(params voucherentity.GetVoucherDetailRequest) voucherentity.GetVoucherDetailResult
	GetVoucherList(params voucherentity.GetVoucherListRequest) voucherentity.GetVoucherListResult
	//followprize
	AddFollowPrize(params followprizeentity.AddFollowPrizeRequest) followprizeentity.AddFollowPrizeResult
	DeleteFollowPrize(params followprizeentity.DeleteFollowPrizeRequest) followprizeentity.DeleteFollowPrizeResult
	EndFollowPrize(params followprizeentity.EndFollowPrizeRequest) followprizeentity.EndFollowPrizeResult
	UpdateFollowPrize(params followprizeentity.UpdateFollowPrizeRequest) followprizeentity.UpdateFollowPrizeResult
	GetFollowPrizeDetail(params followprizeentity.GetFollowPrizeDetailRequest) followprizeentity.GetFollowPrizeDetailResult
	GetFollowPrizeList(params followprizeentity.GetFollowPrizeListRequest) followprizeentity.GetFollowPrizeListResult
	//toppicks
	GetTopPicksList(params toppicksentity.GetTopPicksListRequest) toppicksentity.GetTopPicksListResult
	AddTopPicks(params toppicksentity.AddTopPicksRequest) toppicksentity.AddTopPicksResult
	UpdateTopPicks(params toppicksentity.UpdateTopPicksRequest) toppicksentity.UpdateTopPicksResult
	DeleteTopPicks(params toppicksentity.DeleteTopPicksRequest) toppicksentity.DeleteTopPicksResult
	//shopcategory
	AddShopCategory(params shopcategoryentity.AddShopCategoryRequest) shopcategoryentity.AddShopCategoryResult
	GetShopCategoryList(params shopcategoryentity.GetShopCategoryListRequest) shopcategoryentity.GetShopCategoryListResult
	DeleteShopCategory(params shopcategoryentity.DeleteShopCategoryRequest) shopcategoryentity.DeleteShopCategoryResult
	UpdateShopCategory(params shopcategoryentity.UpdateShopCategoryRequest) shopcategoryentity.UpdateShopCategoryResult
	AddItemList(params shopcategoryentity.AddItemListRequest) shopcategoryentity.AddItemListResult
	CGetItemList(params shopcategoryentity.GetItemListRequest) shopcategoryentity.GetItemListResult
	DeleteItemList(params shopcategoryentity.DeleteItemListRequest) shopcategoryentity.DeleteItemListResult
	//returns
	GetReturnDetail(params returnsentity.GetReturnDetailRequest) returnsentity.GetReturnDetailResult
	GetReturnList(params returnsentity.GetReturnListRequest) returnsentity.GetReturnListResult
	Confirm(params returnsentity.ConfirmRequest) returnsentity.ConfirmResult
	Dispute(params returnsentity.DisputeRequest) returnsentity.DisputeResult
	//accounthealth
	ShopPerformance() accounthealthentity.ShopPerformanceResult
	ShopPenalty() accounthealthentity.ShopPenaltyResult
	//sellerchat
	GetMessage(params sellerchatentity.GetMessageRequest) sellerchatentity.GetMessageResult
	SendMessage(params sellerchatentity.SendMessageRequest) sellerchatentity.SendMessageResult
	GetConversationList(params sellerchatentity.GetConversationListRequest) sellerchatentity.GetConversationListResult
	GetOneConversation(params sellerchatentity.GetOneConversationRequest) sellerchatentity.GetOneConversationResult
	DeleteConversation(params sellerchatentity.DeleteConversationRequest) sellerchatentity.DeleteConversationResult
	PinConversation(params sellerchatentity.PinConversationRequest) sellerchatentity.PinConversationResult
	UnpinConversation(params sellerchatentity.UnpinConversationRequest) sellerchatentity.UnpinConversationResult
	ReadConversation(params sellerchatentity.ReadConversationRequest) sellerchatentity.ReadConversationResult
	UnreadConversation(params sellerchatentity.UnreadConversationRequest) sellerchatentity.UnreadConversationResult
	GetOfferToggleStatus(params sellerchatentity.GetOfferToggleStatusRequest) sellerchatentity.GetOfferToggleStatusResult
	SetOfferToggleStatus(params sellerchatentity.SetOfferToggleStatusRequest) sellerchatentity.SetOfferToggleStatusResult
	SUploadImage(file string) sellerchatentity.UploadImageResult
	//public
	GetShopsByPartner(params publicentity.GetShopsByPartnerRequest) publicentity.GetShopsByPartnerResult
	GetMerchantsByPartner(params publicentity.GetMerchantsByPartnerRequest) publicentity.GetMerchantsByPartnerResult
	GetRefreshTokenByUpgradeCode(upgradeCode string, shopIdList []int64) publicentity.GetRefreshTokenByUpgradeCodeResult
	//push
	GetPushConfig(params pushentity.GetPushConfigRequest) pushentity.GetPushConfigResult
	SetPushConfig(params pushentity.SetPushConfigRequest) pushentity.SetPushConfigResult
}

//Shopee
type Shopee struct {
	auth.Auth
	order.Order
	logistics.Logistics
	product.Product
	globalproduct.GlobalProduct
	mediaspace.MediaSpace
	shop.Shop
	merchant.Merchant
	firstmile.FirstMile
	payment.Payment
	discount.Discount
	bundledeal.BundleDeal
	addondeal.AddOnDeal
	voucher.Voucher
	followprize.FollowPrize
	toppicks.TopPicks
	shopcategory.ShopCategory
	returns.Returns
	accounthealth.AccountHealth
	sellerchat.Sellerchat
	public.Public
	push.Push
}

//shopeeList 接口列表
var shopeeList map[string]Shopeer

//Register
func Register(name string, cfg *shopeeConfig.Config) {
	shopeeList[name] = &Shopee{
		auth.Auth{Config: cfg},
		order.Order{Config: cfg},
		logistics.Logistics{Config: cfg},
		product.Product{Config: cfg},
		globalproduct.GlobalProduct{Config: cfg},
		mediaspace.MediaSpace{Config: cfg},
		shop.Shop{Config: cfg},
		merchant.Merchant{Config: cfg},
		firstmile.FirstMile{Config: cfg},
		payment.Payment{Config: cfg},
		discount.Discount{Config: cfg},
		bundledeal.BundleDeal{Config: cfg},
		addondeal.AddOnDeal{Config: cfg},
		voucher.Voucher{Config: cfg},
		followprize.FollowPrize{Config: cfg},
		toppicks.TopPicks{Config: cfg},
		shopcategory.ShopCategory{Config: cfg},
		returns.Returns{Config: cfg},
		accounthealth.AccountHealth{Config: cfg},
		sellerchat.Sellerchat{Config: cfg},
		public.Public{Config: cfg},
		push.Push{Config: cfg},
	}
}

//GetApi
func GetApi(name string) Shopeer {
	return shopeeList[name]
}

//NewApi
func NewApi(cfg *shopeeConfig.Config) Shopeer {
	return &Shopee{
		auth.Auth{Config: cfg},
		order.Order{Config: cfg},
		logistics.Logistics{Config: cfg},
		product.Product{Config: cfg},
		globalproduct.GlobalProduct{Config: cfg},
		mediaspace.MediaSpace{Config: cfg},
		shop.Shop{Config: cfg},
		merchant.Merchant{Config: cfg},
		firstmile.FirstMile{Config: cfg},
		payment.Payment{Config: cfg},
		discount.Discount{Config: cfg},
		bundledeal.BundleDeal{Config: cfg},
		addondeal.AddOnDeal{Config: cfg},
		voucher.Voucher{Config: cfg},
		followprize.FollowPrize{Config: cfg},
		toppicks.TopPicks{Config: cfg},
		shopcategory.ShopCategory{Config: cfg},
		returns.Returns{Config: cfg},
		accounthealth.AccountHealth{Config: cfg},
		sellerchat.Sellerchat{Config: cfg},
		public.Public{Config: cfg},
		push.Push{Config: cfg},
	}
}

//init
func init() {
	shopeeList = make(map[string]Shopeer)
}
