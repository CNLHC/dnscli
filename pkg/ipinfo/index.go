package ipinfo

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

type IPInfo struct {
	PublicV4 JSONIPResp
	PublicV6 JSONIPResp
}

func GetIP() (res IPInfo) {
	res.PublicV4, _ = JsonIPGet()
	res.PublicV6, _ = JSONIPGetV6()
	return

}

type JSONIPResp struct {
	IP      string `json:"ip"`
	GeoIP   string `json:"geo-ip"`
	APIHelp string `json:"API Help"`
}

func JsonIPGet() (obj JSONIPResp, err error) {
	var res []byte
	code, body, err := fasthttp.Get(res, "https://jsonip.com")
	if code != http.StatusOK || err != nil {
		err = errors.Wrap(err, "jsonip api error")
		return
	}
	err = json.Unmarshal(body, &obj)
	return
}

func JSONIPGetV6() (obj JSONIPResp, err error) {
	cli := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return net.Dial("tcp6", addr)
		},
	}
	var res []byte
	code, body, err := cli.Get(res, "https://jsonip.com:443")
	if code != http.StatusOK || err != nil {
		err = errors.Wrap(err, "jsonip api error")
		return
	}
	err = json.Unmarshal(body, &obj)
	return

}
