package binance

import (
	"context"
)

// TransferToSubAccountService transfer to subaccount
type TransferToSubAccountService struct {
	c       *Client
	toEmail string
	asset   string
	amount  string
}

// ToEmail set toEmail
func (s *TransferToSubAccountService) ToEmail(toEmail string) *TransferToSubAccountService {
	s.toEmail = toEmail
	return s
}

// Asset set asset
func (s *TransferToSubAccountService) Asset(asset string) *TransferToSubAccountService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *TransferToSubAccountService) Amount(amount string) *TransferToSubAccountService {
	s.amount = amount
	return s
}

func (s *TransferToSubAccountService) transferToSubaccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"toEmail": s.toEmail,
		"asset":   s.asset,
		"amount":  s.amount,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *TransferToSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *TransferToSubAccountResponse, err error) {
	data, err := s.transferToSubaccount(ctx, "/sapi/v1/sub-account/transfer/subToSub", opts...)
	if err != nil {
		return nil, err
	}
	res = &TransferToSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransferToSubAccountResponse define transfer to subaccount response
type TransferToSubAccountResponse struct {
	TxnID int64 `json:"txnId"`
}

type SubaccountDepositAddressService struct {
	c       *Client
	email   string
	coin    string
	network string
}

// Email set email
func (s *SubaccountDepositAddressService) Email(email string) *SubaccountDepositAddressService {
	s.email = email
	return s
}

// Coin set coin
func (s *SubaccountDepositAddressService) Coin(coin string) *SubaccountDepositAddressService {
	s.coin = coin
	return s
}

// Network set network
func (s *SubaccountDepositAddressService) Network(network string) *SubaccountDepositAddressService {
	s.network = network
	return s
}

func (s *SubaccountDepositAddressService) subaccountDepositAddress(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email":   s.email,
		"coin":    s.coin,
		"network": s.network,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountDepositAddressResponse, err error) {
	data, err := s.subaccountDepositAddress(ctx, "/sapi/v1/capital/deposit/subAddress", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountDepositAddressResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountDepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

type SubaccountAssetsService struct {
	c     *Client
	email string
}

// Email set email
func (s *SubaccountAssetsService) Email(email string) *SubaccountAssetsService {
	s.email = email
	return s
}

func (s *SubaccountAssetsService) subaccountAssets(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email": s.email,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountAssetsResponse, err error) {
	data, err := s.subaccountAssets(ctx, "/sapi/v3/sub-account/assets", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountAssetsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountAssetsResponse Query Sub-account Assets response
type SubaccountAssetsResponse struct {
	Balances []AssetBalance `json:"balances"`
}

type AssetBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

type SubaccountSpotSummaryService struct {
	c     *Client
	email *string
	page  *int32
	size  *int32
}

// Email set email
func (s *SubaccountSpotSummaryService) Email(email string) *SubaccountSpotSummaryService {
	s.email = &email
	return s
}

func (s *SubaccountSpotSummaryService) Page(page int32) *SubaccountSpotSummaryService {
	s.page = &page
	return s
}

func (s *SubaccountSpotSummaryService) Size(size int32) *SubaccountSpotSummaryService {
	s.size = &size
	return s
}

func (s *SubaccountSpotSummaryService) subaccountSpotSummary(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.size != nil {
		r.setParam("size", *s.size)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountSpotSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountSpotSummaryResponse, err error) {
	data, err := s.subaccountSpotSummary(ctx, "/sapi/v1/sub-account/spotSummary", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountSpotSummaryResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountSpotSummaryResponse Query Sub-account Spot Assets Summary response
type SubaccountSpotSummaryResponse struct {
	TotalCount                int64                       `json:"totalCount"`
	MasterAccountTotalAsset   string                      `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []SpotSubUserAssetBtcVoList `json:"spotSubUserAssetBtcVoList"`
}

type SpotSubUserAssetBtcVoList struct {
	Email      string `json:"email"`
	TotalAsset string `json:"totalAsset"`
}

// SubAccountListService Query Sub-account List (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#query-sub-account-list-for-master-account
type SubAccountListService struct {
	c           *Client
	email       *string
	isFreeze    bool
	page, limit int
}

func (s *SubAccountListService) Email(v string) *SubAccountListService {
	s.email = &v
	return s
}

func (s *SubAccountListService) IsFreeze(v bool) *SubAccountListService {
	s.isFreeze = v
	return s
}

func (s *SubAccountListService) Page(v int) *SubAccountListService {
	s.page = v
	return s
}

func (s *SubAccountListService) Limit(v int) *SubAccountListService {
	s.limit = v
	return s
}

func (s *SubAccountListService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountList, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/list",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.isFreeze {
		r.setParam("isFreeze", "true")
	} else {
		r.setParam("isFreeze", "false")
	}
	if s.page > 0 {
		r.setParam("page", s.page)
	}
	if s.limit > 200 {
		r.setParam("limit", 200)
	} else if s.limit <= 0 {
		r.setParam("limit", 10)
	} else {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountList)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountList struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  uint64 `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}

// 新增为子账户API Key开启/关闭IP白名单
type SubaccountIpRestrictionService struct {
	c                *Client
	email            string
	subAccountApiKey string
	ipRestrict       bool
	method           string
}

// Email set email
func (s *SubaccountIpRestrictionService) Email(email string) *SubaccountIpRestrictionService {
	s.email = email
	return s
}

// SubAccountApiKey set SubAccountApiKey
func (s *SubaccountIpRestrictionService) SubAccountApiKey(subAccountApiKey string) *SubaccountIpRestrictionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

// ipRestrict set ipRestrict
func (s *SubaccountIpRestrictionService) IpRestrict(ipRestrict bool) *SubaccountIpRestrictionService {
	s.ipRestrict = ipRestrict
	return s
}

// IpAddress set ipAddress
func (s *SubaccountIpRestrictionService) IpList() *SubaccountIpRestrictionService {
	s.method = "GET"
	return s
}

func (s *SubaccountIpRestrictionService) subaccountIpRestrictions(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	if s.method == "" {
		s.method = "POST"
	}
	r := &request{
		method:   s.method,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email":            s.email,
		"subAccountApiKey": s.subAccountApiKey,
		"ipRestrict":       s.ipRestrict,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountIpRestrictionService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountIpRestrictionResponse, err error) {
	data, err := s.subaccountIpRestrictions(ctx, "/sapi/v1/sub-account/subAccountApi/ipRestriction", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountIpRestrictionResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountIpRestrictionResponse struct {
	IPRestrict string   `json:"ipRestrict"`
	IPList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	APIKey     string   `json:"apiKey"`
}

// 新增为子账户API Key添加IP白名单功能
type SubaccountIpRestrictionAddService struct {
	c                *Client
	email            string
	subAccountApiKey string
	ipAddress        string
}

// Email set email
func (s *SubaccountIpRestrictionAddService) Email(email string) *SubaccountIpRestrictionAddService {
	s.email = email
	return s
}

// SubAccountApiKey set subAccountApiKey
func (s *SubaccountIpRestrictionAddService) SubAccountApiKey(subAccountApiKey string) *SubaccountIpRestrictionAddService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

// IpAddress set ipAddress
func (s *SubaccountIpRestrictionAddService) IpAddress(ipAddress string) *SubaccountIpRestrictionAddService {
	s.ipAddress = ipAddress
	return s
}

func (s *SubaccountIpRestrictionAddService) subaccountIpRestrictionAdds(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email":            s.email,
		"subAccountApiKey": s.subAccountApiKey,
		"ipAddress":        s.ipAddress,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountIpRestrictionAddService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountIpRestrictionAddResponse, err error) {
	data, err := s.subaccountIpRestrictionAdds(ctx, "/sapi/v1/sub-account/subAccountApi/ipRestriction/ipList", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountIpRestrictionAddResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountIpRestrictionAddResponse struct {
	IP         string `json:"ip"`
	UpdateTime int64  `json:"updateTime"`
	APIKey     string `json:"apiKey"`
}
