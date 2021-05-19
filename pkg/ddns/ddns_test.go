package ddns

import (
	"testing"
	"time"

	"github.com/CNLHC/dnscli/config"
	"github.com/CNLHC/dnscli/providers"
	"github.com/spf13/viper"
)

func TestDDNS(t *testing.T) {
	cfg := []DDNSCfg{}
	config.GetKey("")
	viper.UnmarshalKey("ddns", &cfg)
	o := DDNS{
		Cfg:         cfg,
		Interval:    5 * time.Second,
		DNSProvider: providers.GetDefaultProvider(),
	}
	o.Run()

}
