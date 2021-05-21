package apsarastack

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

type Organization1 struct {

	Code string `json:"code"`
	Cost int    `json:"cost"`
	Data            []struct {
		OrganizationID   int    `json:"organizationId"`
		MuserID          string `json:"muserId"`
		OrganizationName string `json:"organizationName"`
		Ctime            int64  `json:"ctime"`
		AasPk            string `json:"aasPk"`
		ID               int    `json:"id"`
		UserType         string `json:"userType"`
		Mtime            int64  `json:"mtime"`
		UserName         string `json:"userName"`
		CuserID          string `json:"cuserId"`
		UserID           string `json:"userId"`
	} `json:"data"`
	Message      string `json:"message"`
	PureListData bool   `json:"pureListData"`
	Redirect     bool   `json:"redirect"`
	RequestID    string `json:"requestId"`
	Success      bool   `json:"success"`
}


func dataSourceApsaraStackAscmTaskListMaxCompute() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApsaraStackAscmaskListMaxCompute,
		Schema: map[string]*schema.Schema{
			"organizations": {
				Type: schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"organizationid": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"muserid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"organizationname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"aaspk": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ctime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"usertype": {
							Type:     schema.TypeString,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"mtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cuserid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"userid": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceApsaraStackAscmaskListMaxCompute(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	request := requests.NewCommonRequest()
	if strings.ToLower(client.Config.Protocol) == "https" {
		request.Scheme = "https"
	} else {
		request.Scheme = "http"
	}
	if client.Config.Insecure {
		request.SetHTTPSInsecure(client.Config.Insecure)
	}
	request.Method = "POST"
	request.Product = "ascm"
	request.Version = "2019-05-10"

	request.RegionId = client.RegionId
	request.ApiName = "GetOdpsUserList"
	department := "43"
	resource_group := "52"
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{
		"AccessKeyId":     client.AccessKey,
		"AccessKeySecret": client.SecretKey,
		"Department": 		department,
		"ResourceGroup": 	resource_group,
		"Product":         "ascm",
		"RegionId":         client.RegionId,
		"Action":          "GetOdpsUserList",
		"Version":         "2019-05-10",
		"SignatureVersion": "1.0",
	}

	log.Printf("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

	var response Organization1
	for {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.ProcessCommonRequest(request)
		})
		log.Printf(" response of raw MeteringWebQuery : %s", raw)

		if err != nil {
			return WrapErrorf(err, DataDefaultErrorMsg, "apsarastack_ascm_task_list_max_compute", request.GetActionName(), ApsaraStackSdkGoERROR)
		}
		log.Printf("YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYy")
		bresponse, _ := raw.(*responses.CommonResponse)

		headers := bresponse.GetHttpHeaders()
		if headers["X-Acs-Response-Success"][0] == "false" {
			if len(headers["X-Acs-Response-Errorhint"]) > 0 {
				return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm", "API Action", headers["X-Acs-Response-Errorhint"][0])
			} else {
				return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm", "API Action", bresponse.GetHttpContentString())
			}
		}
		log.Printf("Raw response %v",bresponse)

		err = json.Unmarshal(bresponse.GetHttpContentBytes(), &response)
		if err != nil {
			return WrapError(err)
		}
		log.Printf("Unmarshelled response %v",response)
		if bresponse.IsSuccess() == true || len(response.Data) > 0 {
			break
		}

	}

	var ids []string
	var s []map[string]interface{}
	for _, rg := range response.Data {

		mapping := map[string]interface{}{

			"organizationid": 	rg.OrganizationID ,
			"muserid": 			rg.MuserID,
			"organizationname": rg.OrganizationName ,
			"ctime": 			rg.Ctime ,
			"aaspk":			rg.AasPk ,
			"id": 				rg.ID ,
			"usertype": 		rg.UserType ,
			"mtime": 			rg.Mtime ,
			"username": 		rg.UserName ,
			"cuserid": 			rg.CuserID ,
			"userid": 			rg.UserID,

		}

		log.Printf("Execute")
		s = append(s, mapping)
		log.Printf("Mapping done")
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("organizations", s); err != nil {
		return nil
	}

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}
	return nil
}