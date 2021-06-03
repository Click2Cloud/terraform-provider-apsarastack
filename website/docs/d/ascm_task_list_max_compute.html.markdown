---
subcategory: "ASCM"
layout: "apsarastack"
page_title: "ApsaraStack: apsarastack_ascm_task_list_max_compute"
sidebar_current: "docs-apsarastack-datasource-ascm-task_list_max_compute"
description: |-
Provides a list to the user.
---

# apsarastack\_ascm-task_list_max_compute

This data source provides the list of the current Apsara Stack Cloud user.

## Example Usage

```
data "apsarastack_ascm_task_list_max_compute" "default" {

  userid          = "ascm-dw-1615962103985"
  username        = "demo-2"
  organizationid  = 6

}

output "roles" {
  value = data.apsarastack_ascm_task_list_max_compute.default.*
}


```

## Argument Reference

The following arguments are supported:

* `username` - (Optional) name of user.
* `organizationid` - (Optional) ID of the Organization.
* `userid` - (Optional) The ID of user.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `organization` - A list of organizations. Each element contains the following attributes:

    * `id` - ID of the organization.
    * `username` - name of user.
    * `organizationid` - ID of organization.
    * `usertype` - Type of user.
    * `userid` - ID of user.
    * `organizationname` - name of organization.