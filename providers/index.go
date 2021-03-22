package providers

import (
	"github.com/CNLHC/dnscli/providers/ali"
	"github.com/CNLHC/dnscli/shim"
)

func GetDefaultProvider() shim.DNSProvider {
	prov, err := ali.NewAliProvider()
	if err != nil {
		panic(err)
	}
	return prov

}
