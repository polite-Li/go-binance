### go-binance

A Golang SDK for [binance](https://www.binance.com) API.

基于https://github.com/adshao/go-binance.git 新增接口

### Documentation

### REST API

#### SubAccount

```golang
var accountClient = binance.NewClient(apiKey, secretKey)

// 为子账户API Key开启/关闭IP白名单
list, err := accountClient.NewSubaccountIpRestrictionService().
		Email("").SubAccountApiKey("").IpRestrict(true).
		Do(context.Background())

// 查询子账户API Key IP白名单
list, err := accountClient.NewSubaccountIpRestrictionService().
		Email("").SubAccountApiKey("").IpList().
		Do(context.Background())

// 为子账户API Key添加IP白名单
list, err := accountClient.NewSubaccountIpRestrictionAddService().
		Email("").SubAccountApiKey("").IpAddress("").
		Do(context.Background())