package apsarastack

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

type Users struct {
	//Code string `json:"code"`
	//Cost int    `json:"cost"`
	Data0 []struct {
		OrganizationID int64  `json:"organizationId"`
		GroupName      string `json:"groupName"`
		AugID          string `json:"augId"`
	} `json:"data"`
}

func dataSourceApsaraStackAscmListUserGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApsaraStackAscmListUserGroupRead,
		Schema: map[string]*schema.Schema{

			"organizations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organizationid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"groupname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"augid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceApsaraStackAscmListUserGroupRead(d *schema.ResourceData, meta interface{}) error {
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
	log.Printf("*************")
	request.Method = "POST"
	request.Product = "ascm"
	request.Version = "2019-05-10"
	var parentId string
	if v, ok := d.GetOk("parent_id"); ok {
		parentId = fmt.Sprint(v.(int))
	} else {
		parentId = client.Department
	}

	request.RegionId = client.RegionId
	request.ApiName = "ListUserGroups"

	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{
		"AccessKeyId":     client.AccessKey,
		"AccessKeySecret": client.SecretKey,
		"Product":         "ascm",
		"RegionId":        client.RegionId,
		"Action":          "ListUserGroups",
		"Version":         "2019-05-10",
		"id":              parentId,
	}

	var response Users
	for {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.ProcessCommonRequest(request)
		})
		log.Printf(" response of raw MeteringWebQuery : %s", raw)

		if err != nil {
			return WrapErrorf(err, DataDefaultErrorMsg, "apsarastack_ascm_list_user_group", request.GetActionName(), ApsaraStackSdkGoERROR)
		}

		bresponse, _ := raw.(*responses.CommonResponse)
		log.Printf("Raw response %v", bresponse)

		err = json.Unmarshal(bresponse.GetHttpContentBytes(), &response)
		if err != nil {
			return WrapError(err)
		}
		log.Printf("Unmarshelled response %v", response)

		if bresponse.IsSuccess() == true || len(response.Data0) > 0 {
			break
		}
	}

	var ids []string
	var s []map[string]interface{}
	for _, rg := range response.Data0 {

		mapping := map[string]interface{}{

			"organizationid": rg.OrganizationID,
			"groupname":      rg.GroupName,
			"augid":          rg.AugID,
		}

		log.Printf("Execute")
		s = append(s, mapping)
		log.Printf("Mapping done")
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("organizations", s); err != nil {
		//return WrapError(err)
		return nil

	}

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}
	return nil
}
