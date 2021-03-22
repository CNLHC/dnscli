package ali

import (
	"fmt"

	"github.com/CNLHC/dnscli/config"
	"github.com/CNLHC/dnscli/shim"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type aliProvider struct {
	_client *alidns.Client
}

func NewAliProvider() (prov shim.DNSProvider, err error) {
	_client, err := alidns.NewClientWithAccessKey("cn-beijing",
		config.GetKey("ALI_AID"),
		config.GetKey("ALI_AKEY"))
	prov = &aliProvider{_client: _client}
	return

}

func (c *aliProvider) CreateRecord(r shim.DNSRecord) (err error) {
	req := alidns.CreateAddDomainRecordRequest()
	req.DomainName = r.DomainName
	req.RR = r.Host
	req.Type = string(r.Type)
	req.Value = r.Value
	_, err = c._client.AddDomainRecord(req)
	return
}

func (c *aliProvider) UpdateRecord(r shim.DNSRecord) (err error) {

	req := alidns.CreateUpdateDomainRecordRequest()
	rid, err := c.getIDByRecord(r)
	if err != nil {
		return
	}
	req.RecordId = rid
	req.RR = r.Host
	req.Type = string(r.Type)
	req.Value = r.Value
	_, err = c._client.UpdateDomainRecord(req)
	return
}

func (c *aliProvider) DeleteRecord(r shim.DNSRecord) (err error) {
	cli := c._client
	rid, err := c.getIDByRecord(r)
	if err != nil {
		return
	}
	req := alidns.CreateDeleteDomainRecordRequest()
	req.RecordId = rid
	_, err = cli.DeleteDomainRecord(req)
	return
}

func (c *aliProvider) ListRecord(domain string) (res []shim.DNSRecord, err error) {
	return
}

func (c *aliProvider) getIDByRecord(r shim.DNSRecord) (id string, err error) {
	fullname := fmt.Sprintf("%s.%s", r.Host, r.DomainName)
	resp, err := c.describeRecord(fullname)
	if err != nil || resp.TotalCount <= 0 {
		err = shim.ErrNoSuchRecord
		return
	}
	return resp.DomainRecords.Record[0].RecordId, nil
}

func (c *aliProvider) describeRecord(fulldomain string) (resp *alidns.DescribeSubDomainRecordsResponse, err error) {
	req := alidns.CreateDescribeSubDomainRecordsRequest()
	req.SubDomain = fulldomain
	resp, err = c._client.DescribeSubDomainRecords(req)
	return
}
