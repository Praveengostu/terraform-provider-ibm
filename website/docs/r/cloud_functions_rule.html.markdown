---
layout: "ibm"
page_title: "IBM : cloud_functions_rule"
sidebar_current: "docs-ibm-resource-cloud-functions-rule"
description: |-
  Manages IBM Cloud Functions rule.
---

# ibm\_cloud_functions_rule

Create, update, or delete for [IBM Cloud Functions trigger](https://console.bluemix.net/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers).

## Example Usage

```hcl
resource "ibm_cloud_functions_action" "action" {
  name = "hello"

  exec = {
    kind = "nodejs:6"
    code = "${file("test-fixtures/hellonode.js")}"
  }
}

resource "ibm_cloud_functions_trigger" "trigger" {
  name = "alarmtrigger"

  feed = [
    {
      name = "/whisk.system/alarms/alarm"

      parameters = <<EOF
					[
						{
							"key":"cron",
							"value":"0 */2 * * *"
						}
					]
                EOF
    },
  ]
}

resource "ibm_cloud_functions_rule" "rule" {
  name         = "alarmrule"
  trigger_name = "${ibm_cloud_functions_trigger.trigger.name}"
  action_name  = "${ibm_cloud_functions_action.action.name}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of rule.
* `trigger_name` - (Required, string) Name of trigger.
* `action_name` - (Required, string) Name of action.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new rule.
* `publish` - Rule visbility.
* `version` - Semantic version of the item.
* `status` - Status of the rule.

## Import

ibm_cloud_functions_rule can be imported using their id, e.g.

```
$ terraform import ibm_cloud_functions_rule.sampleRule alarmrule

```
