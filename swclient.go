package swclient

type swclient struct {
	user    string
	key     string
	shopurl string
	h       httpclient
	hshr    Hash.Hash
}

func NewSWClient(user string, key string, shopurl string) *swclient {
	return swclient{user: user, key: key, shopurl: shopurl, httpclient{}, md5.New()}
}

func (swc swclient) GetArticleById(id int) (string, error) {
	return swc.h.get(shopurl+"/api/"+strconv.ItoA(id), swc.user, swc.key, swc.hshr)
}
