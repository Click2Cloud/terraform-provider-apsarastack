package apsarastack

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApsaraStackRamRoleAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApsaraStackInstanceRoleAttachmentCreate,
		Read:   resourceApsaraStackInstanceRoleAttachmentRead,
		Delete: resourceApsaraStackInstanceRoleAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"role_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_ids": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApsaraStackInstanceRoleAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	ramService := RamService{client}

	instanceIds := convertListToJsonString(d.Get("instance_ids").(*schema.Set).List())

	request := ecs.CreateAttachInstanceRamRoleRequest()
	request.RegionId = client.RegionId
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": client.SecretKey, "Product": "ecs"}
	request.InstanceIds = instanceIds
	request.RamRoleName = d.Get("role_name").(string)

	err := ramService.JudgeRolePolicyPrincipal(request.RamRoleName)
	if err != nil {
		return WrapError(err)
	}

	return resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.AttachInstanceRamRole(request)
		})
		if err != nil {
			if IsExpectedErrors(err, []string{"unexpected end of JSON input"}) {
				return resource.RetryableError(WrapError(Error("Please trying again.")))
			}
			return resource.NonRetryableError(WrapErrorf(err, DefaultErrorMsg, "ram_role_attachment", request.GetActionName(), ApsaraStackSdkGoERROR))
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		d.SetId(d.Get("role_name").(string) + COLON_SEPARATED + instanceIds)
		return resource.NonRetryableError(WrapError(resourceApsaraStackInstanceRoleAttachmentRead(d, meta)))
	})
}

func resourceApsaraStackInstanceRoleAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}
	roleName := parts[0]
	client := meta.(*connectivity.ApsaraStackClient)
	ramService := RamService{client}
	object, err := ramService.DescribeRamRoleAttachment(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	instRoleSets := object.InstanceRamRoleSets.InstanceRamRoleSet
	var instIds []string
	for _, item := range instRoleSets {
		if item.RamRoleName == roleName {
			instIds = append(instIds, item.InstanceId)
		}
	}
	d.Set("role_name", object.InstanceRamRoleSets.InstanceRamRoleSet[0].RamRoleName)
	d.Set("instance_ids", instIds)
	return nil

}

func resourceApsaraStackInstanceRoleAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.ApsaraStackClient)
	ramService := RamService{client}
	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}
	roleName := parts[0]
	instanceIds := parts[1]

	request := ecs.CreateDetachInstanceRamRoleRequest()
	request.RegionId = client.RegionId
	request.Headers = map[string]string{"RegionId": client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": client.SecretKey, "Product": "ecs"}
	request.RamRoleName = roleName
	request.InstanceIds = instanceIds

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DetachInstanceRamRole(request)
		})
		if err != nil {
			if IsExpectedErrors(err, []string{"unexpected end of JSON input"}) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(WrapErrorf(err, DefaultTimeoutMsg, d.Id(), request.GetActionName(), ApsaraStackSdkGoERROR))
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), ApsaraStackSdkGoERROR)
	}
	return WrapError(ramService.WaitForRamRoleAttachment(d.Id(), Deleted, DefaultTimeout))
}
