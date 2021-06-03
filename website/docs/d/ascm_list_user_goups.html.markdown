---
subcategory: "ASCM"
layout: "apsarastack"
page_title: "ApsaraStack: apsarastack_ascm_list_user_groups"
sidebar_current: "docs-apsarastack-datasource-ascm-list-user-groups"
description: |-
Provides a list of user groups to the user.
---

# apsarastack\_ascm_list_user_groups

This data source provides the list of the current Apsara Stack Cloud user.

## Example Usage

```
data "apsarastack_ascm_roles" "default" {

  augid           = "aug-10122706fd040a1e0d19000p"
  groupname       = "uitester_usergroup"
  organizationid  = 6

}

output "roles" {
  value = data.apsarastack_ascm_roles.default.*
}


```

## Argument Reference

The following arguments are supported:

* `groupname` - (Optional) User group name.
* `organizationid` - (Optional) ID of the Organization.
* `augid` - (Optional) The ID.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `organization` - A list of organizations. Each element contains the following attributes:
  
  * `id` - ID of the organization.
  * `name` - organization name.
  * `uuid` - uuid of organization.
  * `parent_d` - parentid of organization.
  

* `roles` - A list of roles. Each element contains the following attributes:

    * `id` - ID of the role.
    * `role_name` - role name.
    * `description` - Description about the role.
    * `role_level` - role level.
    * `role_type` - types of role.
    * `ram_role` - ram authorized role.
    * `role_range` - specific range for a role.
    * `default` - default role.
    * `active` - Role status.
    * `organization_id` - ID of the organization where role belongs.
    * `code` - role code.
  
  
* `resourceSets` - A list of resource sets. Each element contains the following attributes:
    * `id` - ID of resource set.
    * `organization_id` - ID of the organization where role belongs.
    * `resourcegroupname` - resource group name.
    * `resourcegrouptype` - resource group type.
    * `rsid` - ID of rs.
  

* `users` - A list of users. Each element contains the following attributes:
    * `accounttype` - Account type of user.
    * `active` - User status.
    * `backendaccounttype` - backend account type of user.
    * `cellphonenum` - phone number of user.
    * `default` - default user.
    * `defaultroleid` - default role id of user.
    * `email` - email of user.
    * `enableemail` - status of email. 
    * `id` - ID of user.
    * `loginname` - login name of user. 
    * `logintime` - login time of user.
    * `mobilenationcode`- mobile nation code of user.
    * `organizationid` - organization id of user.
    * `ramuser` - status of ram user.
    * `userloginctrlid` - login ctrol id of user.
    * `username` - name of the user.






