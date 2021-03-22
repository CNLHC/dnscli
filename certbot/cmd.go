package certbot

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/CNLHC/dnscli/config"
	"github.com/CNLHC/dnscli/providers"
	"github.com/CNLHC/dnscli/shim"
	"github.com/spf13/cobra"
)

var logger = config.GetLogger()

func GetCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certbot",
		Short: "certbot hooks",
	}
	auth_cmd := &cobra.Command{
		Use: "auth",
		Run: func(c *cobra.Command, args []string) {
			Auth()
		},
	}
	clean_cmd := &cobra.Command{
		Use: "clean",
	}
	cmd.AddCommand(auth_cmd)
	cmd.AddCommand(clean_cmd)
	return cmd
}

func Auth() {
	provider := providers.GetDefaultProvider()
	acme_root := os.Getenv("CERTBOT_DOMAIN")
	token := os.Getenv("CERTBOT_VALIDATION")

	if acme_root == "" || token == "" {
		fmt.Printf("Invalid Certbot Credentials")
		os.Exit(1)
	}

	r, err := BuildRecord(acme_root)
	if err != nil {
		os.Exit(1)
	}
	r.Value = token
	logger.Info().Msgf("ACME Auth for %+v", r)
	provider.DeleteRecord(r)
	err = provider.CreateRecord(r)
	time.Sleep(time.Second * 3)

	if err != nil {
		panic(err)
	} else {
		os.Exit(0)
	}
}

func CleanUp() {
	os.Exit(0)
}

func BuildRecord(acme_raw_domain string) (r shim.DNSRecord, err error) {
	ingredients := strings.Split(acme_raw_domain, ".")
	r.Type = shim.RecordTXT
	hosts := []string{"_acme-challenge"}

	if len(ingredients) > 2 {
		r.DomainName = strings.Join(ingredients[len(ingredients)-2:], ".")
		hosts = append(hosts, ingredients[0:len(ingredients)-2]...)

	} else if len(ingredients) == 2 {
		r.DomainName = acme_raw_domain
	} else {
		err = errors.New("Wrong Domain")
	}
	r.Host = strings.Join(hosts, ".")

	return
}
