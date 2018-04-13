package constants

type redisItem struct {
	key key
	exp int
}

type key string

// redis keys
var (
	WXAccessToken key = "access_token"
	WXJsTicket    key = "js_ticket"
)

// Keys is the enum collection
var Keys map[key]redisItem

func init() {
	Keys = make(map[key]redisItem)

	Keys[WXAccessToken] = redisItem{
		key: WXAccessToken,
	}
	Keys[WXJsTicket] = redisItem{
		key: WXJsTicket,
	}
}
