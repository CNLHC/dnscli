package shim

import "errors"

type DNSType string

const (
	RecordA     = "A"
	RecordAAAA  = "AAAA"
	RecordTXT   = "TXT"
	RecordCNAME = "CNAME"
	RecordMX    = "MX"
	RecordSRV   = "SRV"
	RecordCAA   = "CAA"
)

type DNSRecord struct {
	DomainName string
	Value      string
	Host       string
	Type       DNSType
}

type DNSProvider interface {
	CreateRecord(r DNSRecord) error
	UpdateRecord(r DNSRecord) error
	DeleteRecord(r DNSRecord) error
	ListRecord(domain string) ([]DNSRecord, error)
}

var (
	ErrNoSuchRecord = errors.New("No Such Record")
)
