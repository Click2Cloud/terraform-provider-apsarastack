package apsarastack

import (
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)
func dataSourceApsaraStackPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApsaraStackGetPasswordPolicy,
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

func dataSourceApsaraStackGetPasswordPolicy(d *schema.ResourceData, meta interface{}) error  {
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