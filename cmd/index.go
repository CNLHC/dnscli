package cmd

import (
	"github.com/CNLHC/dnscli/config"
	"github.com/CNLHC/dnscli/providers"
	"github.com/CNLHC/dnscli/shim"
	"github.com/spf13/cobra"
)

var r shim.DNSRecord
var logger = config.GetLogger()
var dnstype string
var dry bool

func DecorateRootCmd(c *cobra.Command) {
	add := &cobra.Command{
		Use: "create",
		Run: func(c *cobra.Command, args []string) {
			r.Type = shim.DNSType(dnstype)

			logger.Info().Msgf("%+v", r)
			if !dry {
				p := providers.GetDefaultProvider()
				err := p.CreateRecord(r)
				if err != nil {
					logger.Error().Msgf("err %+v", err)
				} else {
					logger.Info().Msg("success")
				}
			}
		},
	}

	del := &cobra.Command{
		Use: "remove",
		Run: func(c *cobra.Command, args []string) {
			p := providers.GetDefaultProvider()
			err := p.DeleteRecord(r)
			if err != nil {
				logger.Error().Msgf("err %+v", err)
			} else {
				logger.Info().Msg("success")
			}
		},
	}

	add_flags := add.Flags()

	add_flags.StringVarP(&r.DomainName, "domain", "d", "", "base domain name")
	add_flags.StringVarP(&dnstype, "type", "t", "", "(A,AAAA,CNAME,TXT)")
	add_flags.StringVarP(&r.Host, "rr", "r", "", "host name")
	add_flags.StringVarP(&r.Value, "value", "v", "", "value")
	add_flags.BoolVar(&dry, "dry", false, "dry")

	del_flags := del.Flags()
	del_flags.StringVarP(&r.DomainName, "domain", "d", "", "base domain name")
	del_flags.StringVarP(&r.Host, "rr", "r", "", "host name")
	del_flags.BoolVar(&dry, "dry", false, "dry")

	//TODO check
	add.MarkFlagRequired("domain")
	add.MarkFlagRequired("type")
	add.MarkFlagRequired("host")
	add.MarkFlagRequired("value")

	del.MarkFlagRequired("domain")
	del.MarkFlagRequired("rr")

	c.AddCommand(add)
	c.AddCommand(del)

}
