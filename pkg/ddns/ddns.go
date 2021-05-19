package ddns

import (
	"errors"
	"time"

	"github.com/CNLHC/dnscli/config"
	"github.com/CNLHC/dnscli/pkg/ipinfo"
	"github.com/CNLHC/dnscli/providers"
	"github.com/CNLHC/dnscli/shim"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DDNSCfg struct {
	Type   string `json:"type"`
	Domain string `json:"domain"`
	RR     string `json:"rr"`
}

type DDNS struct {
	Cfg         []DDNSCfg
	Interval    time.Duration
	DNSProvider shim.DNSProvider
	cached_v4   string
	cached_v6   string
}

func (c *DDNS) Run() {
	ticker := time.NewTicker(c.Interval)
	logger := config.GetLogger()
	for ; true; <-ticker.C {
		ip := ipinfo.GetIP()
		logrus.WithFields(logrus.Fields{
			"Action": "Acquire IPInfo",
		}).Info(ip)

		for _, cfg := range c.Cfg {
			r, err := c.cfgToRecord(cfg, ip)
			if err != nil {
				logger.Error().Err(err).Msg("")
				continue
			}
			logger.Info().Msgf("Update Record %v cached %s", r, c.cached_v4)
			err = shim.UpdateOrCreate(c.DNSProvider, r)
			if err != nil {
				logger.Err(err).Msg("")
			} else {
				if cfg.Type == "A" {
					c.cached_v4 = ip.PublicV4.IP
				} else {
					c.cached_v6 = ip.PublicV6.IP
				}

			}

		}

	}
}

func (c *DDNS) cfgToRecord(cfg DDNSCfg, ip ipinfo.IPInfo) (record shim.DNSRecord, err error) {

	switch cfg.Type {
	case "A":
		if ip.PublicV4.IP == "" {
			err = errors.New("No V4 Ip")
			return
		}

		if ip.PublicV4.IP == c.cached_v4 {
			err = errors.New("ipv4 not change")
			return
		}

		record = shim.DNSRecord{
			DomainName: cfg.Domain,
			Value:      ip.PublicV4.IP,
			Host:       cfg.RR,
			Type:       shim.DNSType(cfg.Type),
		}
		return
	case "AAAA":
		if ip.PublicV6.IP == "" {
			err = errors.New("No V6 Ip")
			return
		}

		if ip.PublicV6.IP == c.cached_v6 {
			err = errors.New("ipv6 not change")
			return
		}
		record = shim.DNSRecord{
			DomainName: cfg.Domain,
			Value:      ip.PublicV6.IP,
			Host:       cfg.RR,
			Type:       shim.DNSType(cfg.Type),
		}
		return
	default:
		err = errors.New("invalid config")
		return

	}
}

func RegisterCMD(c *cobra.Command) {

	del := &cobra.Command{
		Use: "ddns",
		Run: func(c *cobra.Command, args []string) {
			cfg := []DDNSCfg{}
			config.GetKey("")
			viper.UnmarshalKey("ddns", &cfg)
			o := DDNS{
				Cfg:         cfg,
				Interval:    5 * time.Second,
				DNSProvider: providers.GetDefaultProvider(),
			}
			o.Run()

		},
	}

	c.AddCommand(del)
}
