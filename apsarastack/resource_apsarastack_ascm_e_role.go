package apsarastack

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity/ascm"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"time"
)

func resourceApsaraStackAscmRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceApsaraStackAscmRoleCreate,
		Read:   resourceApsaraStackAscmRoleRead,
		Update: resourceApsaraStackAscmRoleUpdate,
		Delete: resourceApsaraStackAscmRoleDelete,
		Schema: map[string]*schema.Schema{
			"role_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_range": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApsaraStackAscmRoleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	ascmService := AscmService{client}
	rname := d.Get("role_name").(string)
	rrange := d.Get("role_range").(string)
	check, err := ascmService.DescribeAscmRole(rname)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_organization", "ORG alreadyExist", ApsaraStackSdkGoERROR)
	}
	    var requestInfo *ascm.Client
	if len(check.Data) == 0 {
		request := requests.NewCommonRequest()
		request.QueryParams = map[string]string{
			"RegionId":        client.RegionId,
			"AccessKeySecret": client.SecretKey,
			"Product":         "Ascm",
			"Action":          "CreateRole",
			"Version":         "2019-05-10",
			"ProductName":     "ascm",
			"RoleName":        rname,
			"RoleRange":       rrange,
		}
		request.Method = "POST"
		request.Product = "Ascm"
		request.Version = "2019-05-10"
		request.ServiceCode = "ascm"
		request.Scheme = "http"
		request.ApiName = "CreateRole"
		request.RegionId = client.RegionId
		request.Headers = map[string]string{"RegionId": client.RegionId}

		raw, err := client.WithEcsClient(
			func(ecsClient *ecs.Client) (interface{}, error) {
				return ecsClient.ProcessCommonRequest(request)
			}
		)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_role", "CreateRole", raw)
		}
		addDebug("CreateRole", raw, requestInfo, request)

		bresponse, _ := raw.(*responses.CommonResponse)
		if bresponse.GetHttpStatus() != 200 {
			return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_role", "CreateRole", ApsaraStackSdkGoERROR)
		}
		addDebug("CreateRole", raw, requestInfo, bresponse.GetHttpContentString())
	}
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		check, err = ascmService.DescribeAscmRole(rname)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if len(check.Data) != 0 {
			return nil
		}
		return resource.RetryableError(Error("New Role has been created successfully."))
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_role", "Failed to create role",
			ApsaraStackSdkGoERROR)
	}
	return resourceApsaraStackAscmUserUpdate(d, meta)

}

func resourceApsaraStackAscmRoleUpdate(d *schema.ResourceData, meta interface{}) error {

return resourceApsaraStackAscmUserRead(d, meta)

}

func resourceApsaraStackAscmRoleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	ascmService := AscmService{client}
	rname := d.Get("role_name").(string)
	object, err := ascmService.DescribeAscmRole(rname)
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	if len(object.Data) == 0 {
		d.SetId("")
		return nil
	}

	d.Set("role_name", object.Data[0].RoleName)
	d.Set("role_range", object.Data[0].RoleRange)

	return nil
}

func resourceApsaraStackAscmRoleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	var requestInfo *ascm.Client
	rname := d.Get("role_name").(string)

	request := requests.NewCommonRequest()
	request.QueryParams = map[string]string{
		"RegionId":         client.RegionId,
		"AccessKeySecret":  client.SecretKey,
		"Product":          "Ascm",
		"Action":           "RemoveRole",
		"Version":          "2019-05-10",
		"ProductName":      "ascm",
		"RoleName":        rname,
	}
	request.Method = "POST"
	request.Product = "Ascm"
	request.Version = "2019-05-10"
	request.ServiceCode = "ascm"
	request.Scheme = "http"
	request.ApiName = "RemoveRole"
	request.RegionId = client.RegionId
	request.Headers = map[string]string{"RegionId": client.RegionId}

	raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.ProcessCommonRequest(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_role", "RemoveRole", raw)
	}
	addDebug("RemoveRole", raw, requestInfo, request)

	bresponse, _ := raw.(*responses.CommonResponse)
	if bresponse.GetHttpStatus() != 200 {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_role", "RemoveRole", ApsaraStackSdkGoERROR)
	}
	addDebug("RemoveRole", raw, requestInfo, bresponse.GetHttpContentString())
	return nil
}

