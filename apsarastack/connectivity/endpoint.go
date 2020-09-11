package connectivity

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ServiceCode Load endpoints from endpoints.xml or environment variables to meet specified application scenario, like private cloud.
type ServiceCode string

const (
	ECSCode             = ServiceCode("ECS")
	VPCCode             = ServiceCode("VPC")
	KMSCode             = ServiceCode("KMS")
	BSSOPENAPICode      = ServiceCode("BSSOPENAPI")
	RAMCode             = ServiceCode("RAM")
	CDNCode             = ServiceCode("CDN")
	SLBCode             = ServiceCode("SLB")
	OSSCode             = ServiceCode("OSS")
	LOCATIONCode        = ServiceCode("LOCATION")
	CONTAINCode         = ServiceCode("CS")
	ESSCode             = ServiceCode("ESS")
	DcdnCode            = ServiceCode("DCDN")
	MseCode             = ServiceCode("MSE")
	ActiontrailCode     = ServiceCode("ACTIONTRAIL")
	OosCode             = ServiceCode("OOS")
	EcsCode             = ServiceCode("ECS")
	NasCode             = ServiceCode("NAS")
	EciCode             = ServiceCode("ECI")
	DdoscooCode         = ServiceCode("DDOSCOO")
	BssopenapiCode      = ServiceCode("BSSOPENAPI")
	AlidnsCode          = ServiceCode("ALIDNS")
	ResourcemanagerCode = ServiceCode("RESOURCEMANAGER")
	WafOpenapiCode      = ServiceCode("WAFOPENAPI")
	DmsEnterpriseCode   = ServiceCode("DMSENTERPRISE")
	DnsCode             = ServiceCode("DNS")
	KmsCode             = ServiceCode("KMS")
	CbnCode             = ServiceCode("CBN")

	RDSCode = ServiceCode("RDS")

	ONSCode           = ServiceCode("ONS")
	ALIKAFKACode      = ServiceCode("ALIKAFKA")
	CRCode            = ServiceCode("CR")
	CMSCode           = ServiceCode("CMS")
	OTSCode           = ServiceCode("OTS")
	DNSCode           = ServiceCode("DNS")
	PVTZCode          = ServiceCode("PVTZ")
	LOGCode           = ServiceCode("LOG")
	FCCode            = ServiceCode("FC")
	DDSCode           = ServiceCode("DDS")
	GPDBCode          = ServiceCode("GPDB")
	STSCode           = ServiceCode("STS")
	CENCode           = ServiceCode("CEN")
	KVSTORECode       = ServiceCode("KVSTORE")
	POLARDBCode       = ServiceCode("POLARDB")
	DATAHUBCode       = ServiceCode("DATAHUB")
	MNSCode           = ServiceCode("MNS")
	CLOUDAPICode      = ServiceCode("APIGATEWAY")
	DRDSCode          = ServiceCode("DRDS")
	ELASTICSEARCHCode = ServiceCode("ELASTICSEARCH")
	DDOSCOOCode       = ServiceCode("DDOSCOO")
	DDOSBGPCode       = ServiceCode("DDOSBGP")
	SAGCode           = ServiceCode("SAG")
	EMRCode           = ServiceCode("EMR")
	CasCode           = ServiceCode("CAS")
	YUNDUNDBAUDITCode = ServiceCode("YUNDUNDBAUDIT")
	MARKETCode        = ServiceCode("MARKET")
	HBASECode         = ServiceCode("HBASE")
	ADBCode           = ServiceCode("ADB")
	MAXCOMPUTECode    = ServiceCode("MAXCOMPUTE")
	EDASCode          = ServiceCode("EDAS")
	CassandraCode     = ServiceCode("CASSANDRA")
)

type Endpoints struct {
	Endpoint []Endpoint `xml:"Endpoint"`
}

type Endpoint struct {
	Name      string    `xml:"name,attr"`
	RegionIds RegionIds `xml:"RegionIds"`
	Products  Products  `xml:"Products"`
}

type RegionIds struct {
	RegionId string `xml:"RegionId"`
}

type Products struct {
	Product []Product `xml:"Product"`
}

type Product struct {
	ProductName string `xml:"ProductName"`
	DomainName  string `xml:"DomainName"`
}

var localEndpointPath = "./endpoints.xml"
var localEndpointPathEnv = "TF_ENDPOINT_PATH"
var loadLocalEndpoint = false

func hasLocalEndpoint() bool {
	data, err := ioutil.ReadFile(localEndpointPath)
	if err != nil || len(data) <= 0 {
		d, e := ioutil.ReadFile(os.Getenv(localEndpointPathEnv))
		if e != nil {
			return false
		}
		data = d
	}
	return len(data) > 0
}

func loadEndpoint(region string, serviceCode ServiceCode) string {
	endpoint := strings.TrimSpace(os.Getenv(fmt.Sprintf("%s_ENDPOINT", string(serviceCode))))
	if endpoint != "" {
		return endpoint
	}

	// Load current path endpoint file endpoints.xml, if failed, it will load from environment variables TF_ENDPOINT_PATH
	if !loadLocalEndpoint {
		return ""
	}
	data, err := ioutil.ReadFile(localEndpointPath)
	if err != nil || len(data) <= 0 {
		d, e := ioutil.ReadFile(os.Getenv(localEndpointPathEnv))
		if e != nil {
			return ""
		}
		data = d
	}
	var endpoints Endpoints
	err = xml.Unmarshal(data, &endpoints)
	if err != nil {
		return ""
	}
	for _, endpoint := range endpoints.Endpoint {
		if endpoint.RegionIds.RegionId == string(region) {
			for _, product := range endpoint.Products.Product {
				if strings.ToLower(product.ProductName) == strings.ToLower(string(serviceCode)) {
					return strings.TrimSpace(product.DomainName)
				}
			}
		}
	}

	return ""
}
