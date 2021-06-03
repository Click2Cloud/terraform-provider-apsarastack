package apsarastack

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
	"time"

	_ "context"

	_ "io/ioutil"

	_ "net/http"
	_ "os"
	_ "os/user"
	_ "strconv"

	_ "time"
)

type ListUser struct {
	Redirect        bool   `json:"redirect"`
	EagleEyeTraceID string `json:"eagleEyeTraceId"`
	AsapiSuccess    bool   `json:"asapiSuccess"`
	Code            string `json:"code"`
	Cost            int    `json:"cost"`
	Data            []struct {
		OrganizationID int    `json:"organizationId"`
		GroupName      string `json:"groupName"`
		AugID          string `json:"augId"`
		Roles          []struct {
			RoleRange              string `json:"roleRange"`
			ArID                   string `json:"arId"`
			MuserID                string `json:"muserId"`
			RAMRole                bool   `json:"rAMRole"`
			Code                   string `json:"code"`
			Active                 bool   `json:"active"`
			Description            string `json:"description"`
			RoleType               string `json:"roleType"`
			Default                bool   `json:"default"`
			OwnerOrganizationID    int    `json:"ownerOrganizationId"`
			Enable                 bool   `json:"enable"`
			RoleName               string `json:"roleName"`
			ID                     int    `json:"id"`
			RoleLevel              int64  `json:"roleLevel"`
			CuserID                string `json:"cuserId"`
			OrganizationVisibility string `json:"organizationVisibility"`
		} `json:"roles"`
		CreateTimeStamp int64 `json:"createTimeStamp"`
		ResourceSets    []struct {
			OrganizationID    int    `json:"organizationID"`
			Creator           string `json:"creator"`
			GmtModified       int64  `json:"gmtModified"`
			ResourceGroupName string `json:"resourceGroupName"`
			GmtCreated        int64  `json:"gmtCreated"`
			ResourceGroupType int    `json:"resourceGroupType"`
			RsID              string `json:"rsId"`
			ID                int    `json:"id"`
		} `json:"resourceSets"`
		ID    int `json:"id"`
		Users []struct {
			CellphoneNum       string `json:"cellphoneNum"`
			MuserID            string `json:"muserId,omitempty"`
			AliyunUser         bool   `json:"aliyunUser"`
			BackendAccountType string `json:"backendAccountType"`
			EnableDingTalk     bool   `json:"enableDingTalk"`
			AccountType        int    `json:"accountType"`
			Active             bool   `json:"active"`
			Mtime              int64  `json:"mtime,omitempty"`
			OrganizationID     int    `json:"organizationId"`
			Default            bool   `json:"default"`
			Deleted            bool   `json:"deleted"`
			LoginTime          int64  `json:"loginTime,omitempty"`
			UserLoginCtrlID    int    `json:"userLoginCtrlId"`
			LoginName          string `json:"loginName"`
			Ctime              int64  `json:"ctime"`
			ID                 int    `json:"id"`
			MobileNationCode   string `json:"mobileNationCode"`
			CuserID            string `json:"cuserId"`
			EnableEmail        bool   `json:"enableEmail"`
			RAMUser            bool   `json:"ramUser"`
			DefaultRoleID      int    `json:"defaultRoleId"`
			Email              string `json:"email"`
			EnableShortMessage bool   `json:"enableShortMessage"`
			Username           string `json:"username"`
		} `json:"users"`
		Organization struct {
			MuserID           string        `json:"muserId"`
			Internal          bool          `json:"internal"`
			MultiCloudStatus  string        `json:"multiCloudStatus"`
			SupportRegionList []interface{} `json:"supportRegionList"`
			Level             string        `json:"level"`
			Name              string        `json:"name"`
			Alias             string        `json:"alias"`
			ID                int           `json:"id"`
			CuserID           string        `json:"cuserId"`
			UUID              string        `json:"uuid"`
			ParentID          int           `json:"parentId"`
		} `json:"organization,omitempty"`
	} `json:"data"`
	PageInfo struct {
		Total       int `json:"total"`
		TotalPage   int `json:"totalPage"`
		PageSize    int `json:"pageSize"`
		CurrentPage int `json:"currentPage"`
	} `json:"pageInfo"`
	Message        string `json:"message"`
	ServerRole     string `json:"serverRole"`
	AsapiRequestID string `json:"asapiRequestId"`
	Success        bool   `json:"success"`
	Domain         string `json:"domain"`
	PureListData   bool   `json:"pureListData"`
	API            string `json:"api"`
	AsapiErrorCode string `json:"asapiErrorCode"`
}

func resourceApsaraStackAscmUserGroup() *schema.Resource {
	return &schema.Resource{

		Read:   resourceApsaraStackAscmUserGroupRead,
		Create: resourceApsaraStackAscmUserGroupCreate,
		Update: resourceApsaraStackAscmUserGroupUpdate,
		Delete: resourceApsaraStackAscmUserGroupDelete,
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApsaraStackAscmUserGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	var requestInfo *ecs.Client
	ascmService := AscmService{client}
	groupname := d.Get("group_name").(string)
	organizationid := d.Get("organization_id").(string)
	check, err := ascmService.DescribeUserGroup(groupname)
	did := strings.Split(d.Id(), COLON_SEPARATED)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_user_group", "IsUserGroupExist", ApsaraStackSdkGoERROR)
	}
	if len(check.Data) == 0 {
		request := requests.NewCommonRequest()
		if client.Config.Insecure {
			request.SetHTTPSInsecure(client.Config.Insecure)
		}
		//department := "43"
		//resource_group := "52"
		request.QueryParams = map[string]string{
			"RegionId":        client.RegionId,
			"AccessKeySecret": client.SecretKey,
			"Department":      client.Department,
			"ResourceGroup":   client.ResourceGroup,
			"Product":         "ascm",
			"Action":          "CreateUserGroup",
			"Version":         "2019-05-10",
			"group_name":      groupname,
			"organization_id": organizationid,
			"oid":             did[1],
		}
		request.Method = "POST"
		request.Product = "ascm"
		request.Version = "2019-05-10"
		request.ServiceCode = "ascm"
		request.Domain = client.Domain
		if strings.ToLower(client.Config.Protocol) == "https" {
			request.Scheme = "https"
		} else {
			request.Scheme = "http"
		}
		request.ApiName = "CreateUserGroup"
		request.RegionId = client.RegionId
		request.Headers = map[string]string{"RegionId": client.RegionId}

		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.ProcessCommonRequest(request)
		})

		log.Printf(" response of raw CreateUserGroup : %s", raw)

		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_user_group", "CreateUserGroup", raw)
		}
		addDebug("CreateUserGroup", raw, requestInfo, request)

		bresponse, _ := raw.(*responses.CommonResponse)
		headers := bresponse.GetHttpHeaders()
		if headers["X-Acs-Response-Success"][0] == "false" {
			if len(headers["X-Acs-Response-Errorhint"]) > 0 {
				return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_user_group", "API Action", headers["X-Acs-Response-Errorhint"][0])
			} else {
				return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_user_group", "API Action", bresponse.GetHttpContentString())
			}
		}
		if bresponse.GetHttpStatus() != 200 {
			return WrapErrorf(err, DefaultErrorMsg, "apsarastack_ascm_user_group", "CreateUserGroups", ApsaraStackSdkGoERROR)
		}
		addDebug("CreateUserGroups", raw, requestInfo, bresponse.GetHttpContentString())
	}
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		check, err = ascmService.DescribeUserGroup(groupname)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return resource.RetryableError(err)
	})
	d.SetId(groupname + COLON_SEPARATED + fmt.Sprint(check.Data[0].OrganizationID))
	return resourceApsaraStackAscmUserGroupUpdate(d, meta)

}
func resourceApsaraStackAscmUserGroupUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceApsaraStackAscmUserGroupRead(d, meta)

}

func resourceApsaraStackAscmUserGroupRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.ApsaraStackClient)
	ascmService := AscmService{client}
	object, err := ascmService.DescribeUserGroup(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("group_name", object.Data[0].GroupName)
	d.Set("organization_id", object.Data[0].OrganizationID)
	return nil
}

func resourceApsaraStackAscmUserGroupDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}
