---
subcategory: "ASCM"
layout: "apsarastack"
page_title: "Apsarastack: apsarastack_ascm_user_group"
sidebar_current: "docs-apsarastack-resource-ascm-user_group"
description: |-
Provides Ascm user group.
---

# apsarastack\_ascm_user_group

Provides Ascm user group.

## Example Usage

```
resource "apsarastack_ascm_user_group" "group"{
  group_name= "new-demo"
  organization_id="43"
}
output "User_group" {
  value = apsarastack_ascm_user_group.group.*
}
```
## Argument Reference

The following arguments are supported:

* `group_name` - (Required) User group name.
* `organization_id` - (Required) ID of the Organization.


## Attributes Reference

The following attributes are exported:

* `id` - User Group ID of the user.
* `group_name` - User Group Name of the user.
* `organization_id ` - The ID of the organization.