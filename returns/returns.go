package returns

import (
	"github.com/wjp-letgo/letgo/lib"
	shopeeConfig "github.com/wjp-letgo/shopeego/config"
	returnsentity "github.com/wjp-letgo/shopeego/returns/entity"
)

//Returns
type Returns struct {
	Config *shopeeConfig.Config
}

//GetReturnDetail
//@Title Use this api to get detail information of a return by return id.
//@Description https://open.shopee.com/documents?module=99&type=1&id=607&version=2
func (m *Returns) GetReturnDetail(params returnsentity.GetReturnDetailRequest) returnsentity.GetReturnDetailResult {
	method := "returns/get_return_detail"
	result := returnsentity.GetReturnDetailResult{}
	err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}

//GetReturnList
//@Title Use this api to get detail information of many return by shop id.
//@Description https://open.shopee.com/documents?module=99&type=1&id=608&version=2
func (m *Returns) GetReturnList(params returnsentity.GetReturnListRequest) returnsentity.GetReturnListResult {
	method := "returns/get_return_list"
	result := returnsentity.GetReturnListResult{}
	err := m.Config.HttpGet(method, lib.JSONToMap(params.String()), &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}

//Confirm
//@Title Confirm return
//@Description https://open.shopee.com/documents?module=99&type=1&id=609&version=2
func (m *Returns) Confirm(params returnsentity.ConfirmRequest) returnsentity.ConfirmResult {
	method := "returns/confirm"
	result := returnsentity.ConfirmResult{}
	err := m.Config.HttpPost(method, params, &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}

//Dispute
//@Title Dispute return
//@Description https://open.shopee.com/documents?module=99&type=1&id=610&version=2
func (m *Returns) Dispute(params returnsentity.DisputeRequest) returnsentity.DisputeResult {
	method := "returns/dispute"
	result := returnsentity.DisputeResult{}
	err := m.Config.HttpPost(method, params, &result)
	if err != nil {
		result.Error = err.Error()
	}
	return result
}
