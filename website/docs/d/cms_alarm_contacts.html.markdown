---
subcategory: "Cloud Monitor"
layout: "apsarastack"
page_title: "Apsarastack: apsarastack_cms_alarm_contacts"
sidebar_current: "docs-apsarastack-datasource-cms-alarm-contacts"
description: |-
  Provides a list of alarm contact owned by an Apsarastack Cloud account.
---

# apsarastack\_cms\_alarm\_contacts

Provides a list of alarm contact owned by an Apsarastack Cloud account.

## Example Usage

Basic Usage

```
data "apsarastack_cms_alarm_contacts" "example" {

}
output "first-contact" {
  value = data.apsarastack_cms_alarm_contacts.example
}
```

## Argument Reference

The following arguments are supported:

* `ids` - (Optional, ForceNew) A list of alarm contact IDs. 
* `name_regex` - (Optional, ForceNew) A regex string to filter results by alarm contact name. 
* `chanel_type` - (Optional, ForceNew)  The alarm notification method. Alarm notifications can be sent by using `Email` or `DingWebHook`.
* `chanel_value` - (Optional, ForceNew)  The alarm notification target.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`). 


## Attributes Reference

The following attributes are exported:

* `ids` - A list of alarm contact IDs.
* `names` - A list of alarm contact names.
* `contacts` - A list of alarm contacts. Each element contains the following attributes:
    * `id` - The ID of the alarm contact.
    * `alarm_contact_name` - The name of the alarm contact.
