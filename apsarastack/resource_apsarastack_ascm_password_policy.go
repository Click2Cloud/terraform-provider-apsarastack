package apsarastack

import (
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity/ascm"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApsaraStackAscmPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceApsaraStackAscmPasswordPolicyCreate,
		Read:   resourceApsaraStackAscmPasswordPolicyRead,
		Delete: resourceApsaraStackAscmPasswordPolicyDelete,
        Update: resourceApsaraStackAscmPasswordPolicyUpdate,
		Schema: map[string]*schema.Schema{
			//"Action": {
			//	Type:     schema.TypeString,
			//	Required: true,
			//	ForceNew: true,
			//},
			//"Id": {
			//	Type:     schema.TypeInt,
			//	Required: true,
			//	ForceNew: true,
			//},
			"HardExpiry" :{
				Type:     schema.TypeBool,
				Required: false,
			},
			"RequireLowercaseCharacters" :{
				Type:     schema.TypeBool,
				Required: false,
			},
			"RequireNumbers" :{
				Type:     schema.TypeBool,
				Required: false,
			},
			"RequireSymbols" :{
				Type:     schema.TypeBool,
				Required: false,
			},
			"RequireUppercaseCharacters" :{
				Type:     schema.TypeBool,
				Required: false,
			},
			"MaxLoginAttemps": {
				Type:     schema.TypeInt,
				Required: false,
			},
			"MaxPasswordAge": {
				Type:     schema.TypeInt,
				Required: false,
			},
			"MinimumPasswordLength": {
				Type:     schema.TypeInt,
				Required: false,
			},
			"PasswordErrorLockPeriod": {
				Type:     schema.TypeInt,
				Required: false,
			},
			"PasswordErrorTolerancePeriod": {
				Type:     schema.TypeInt,
				Required: false,
			},
			"PasswordReusePrevention": {
				Type:     schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceApsaraStackAscmPasswordPolicyCreate(d *schema.ResourceData, meta interface{}) error{
	client := meta.(*connectivity.ApsaraStackClient)
	request , err := buildPasswordPolicyCreateRequest(d)
	request.Method = "POST"
	request.Scheme = "http"
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{"AccessKeyId": client.AccessKey,
		"AccessKeySecret": client.SecretKey, "Product": "ascm", "RegionId": client.RegionId,
		"Action": "SetPasswordPolicy", "Version": "2019-05-10", "Department": client.Department,
		"ResourceGroup": client.ResourceGroup ,"Id": d.Id()}
	request.RegionId = client.RegionId
	if err != nil {
		return WrapError(err)
	}
	raw, err := client.WithAscmClient(func(ascmClient *ascm.Client) (interface{}, error) {
		return ascmClient.SetPasswordPolicy(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	response, _ := raw.(*ascm.SetPasswordPolicyResponse)
	addDebug(request.GetActionName(), raw, request.RpcRequest, request, response)
	return  nil
}

func resourceApsaraStackAscmPasswordPolicyRead(d *schema.ResourceData, meta interface{}) error{
	client := meta.(*connectivity.ApsaraStackClient)
	ascmService := AscmService{client}
	response, err := ascmService.GetPasswordPolicy()
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	object := response.Data
	d.Set("HardExpiry", object.HardExpiry)
	d.Set("RequireLowercaseCharacters", object.RequireLowercaseCharacters)
	d.Set("RequireUppercaseCharacters", object.RequireUppercaseCharacters)
	d.Set("RequireNumbers", object.RequireNumbers)
	d.Set("RequireSymbols", object.RequireSymbols)
	d.Set("MaxLoginAttemps", object.MaxLoginAttemps)
	d.Set("MaxPasswordAge", object.MaxPasswordAge)
	d.Set("MinimumPasswordLength", object.MinimumPasswordLength)
	d.Set("PasswordErrorLockPeriod", object.PasswordErrorLockPeriod)
	d.Set("PasswordErrorTolerancePeriod", object.PasswordErrorTolerancePeriod)
	d.Set("PasswordReusePrevention", object.PasswordReusePrevention)
	return  nil
}

func resourceApsaraStackAscmPasswordPolicyUpdate(d *schema.ResourceData, meta interface{}) error{
	client := meta.(*connectivity.ApsaraStackClient)
	ascmService := AscmService{client}
	response, err := ascmService.GetPasswordPolicy()
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	object := response.Data
	d.Set("HardExpiry", object.HardExpiry)
	d.Set("RequireLowercaseCharacters", object.RequireLowercaseCharacters)
	d.Set("RequireUppercaseCharacters", object.RequireUppercaseCharacters)
	d.Set("RequireNumbers", object.RequireNumbers)
	d.Set("RequireSymbols", object.RequireSymbols)
	d.Set("MaxLoginAttemps", object.MaxLoginAttemps)
	d.Set("MaxPasswordAge", object.MaxPasswordAge)
	d.Set("MinimumPasswordLength", object.MinimumPasswordLength)
	d.Set("PasswordErrorLockPeriod", object.PasswordErrorLockPeriod)
	d.Set("PasswordErrorTolerancePeriod", object.PasswordErrorTolerancePeriod)
	d.Set("PasswordReusePrevention", object.PasswordReusePrevention)
	return  nil
}

func resourceApsaraStackAscmPasswordPolicyDelete(d *schema.ResourceData, meta interface{}) error{
	client := meta.(*connectivity.ApsaraStackClient)
	request := ascm.CreateRemovePasswordPolicyByIDRequest()
	request.Method = "POST"
	request.Scheme = "http"
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{"AccessKeyId": client.AccessKey,
		"AccessKeySecret": client.SecretKey, "Product": "ascm", "RegionId": client.RegionId,
		"Action": "RemoveLoginPolicyById", "Version": "2019-05-10", "Department": client.Department,
		"ResourceGroup": client.ResourceGroup ,"Id": d.Id()}
	request.RegionId = client.RegionId

	raw, err := client.WithAscmClient(func(ascmClient *ascm.Client) (interface{}, error) {
		return ascmClient.RemovePasswordPolicyByID(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	response, _ := raw.(*ascm.RemovePasswordPolicyByIDResponse)
	addDebug(request.GetActionName(), raw, request.RpcRequest, request, response)
	return  nil
}

func buildPasswordPolicyCreateRequest(d *schema.ResourceData) (*ascm.SetPasswordPolicyRequest, error) {
	request := ascm.CreateSetPasswordPolicyRequest()
	request.HardExpiry = d.Get("HardExpiry").(bool)
	request.RequireLowercaseCharacters = d.Get("RequireLowercaseCharacters").(bool)
	request.RequireUppercaseCharacters = d.Get("RequireUppercaseCharacters").(bool)
	request.RequireNumbers =d.Get("RequireNumbers").(bool)
	request.RequireSymbols = d.Get("RequireSymbols").(bool)
	request.MaxLoginAttemps = d.Get("MaxLoginAttemps").(int)
	request.MaxPasswordAge = d.Get("MaxPasswordAge").(int)
	request.MinimumPasswordLength = d.Get("MinimumPasswordLength").(int)
	request.PasswordErrorLockPeriod = d.Get("PasswordErrorLockPeriod").(int)
	request.PasswordErrorTolerancePeriod = d.Get("PasswordErrorTolerancePeriod").(int)
	request.PasswordReusePrevention = d.Get("PasswordReusePrevention").(int)
	return request, nil
}