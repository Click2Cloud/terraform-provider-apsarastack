package apsarastack

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccApsaraStackEssScalinggroupsDataSource(t *testing.T) {
	rand := acctest.RandIntRange(5, 1000)
	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand, map[string]string{
			"name_regex": `"${apsarastack_ess_scaling_group.default.scaling_group_name}"`,
		}),
		fakeConfig: testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand, map[string]string{
			"name_regex": `"${apsarastack_ess_scaling_group.default.scaling_group_name}_fake"`,
		}),
	}

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand, map[string]string{
			"ids": `["${apsarastack_ess_scaling_group.default.id}"]`,
		}),
		fakeConfig: testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand, map[string]string{
			"ids": `["${apsarastack_ess_scaling_group.default.id}_fake"]`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand, map[string]string{
			"ids":        `["${apsarastack_ess_scaling_group.default.id}"]`,
			"name_regex": `"${apsarastack_ess_scaling_group.default.scaling_group_name}"`,
		}),
		fakeConfig: testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand, map[string]string{
			"ids":        `["${apsarastack_ess_scaling_group.default.id}_fake"]`,
			"name_regex": `"${apsarastack_ess_scaling_group.default.scaling_group_name}"`,
		}),
	}

	var existEssScalingGroupsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"groups.#":                     "1",
			"ids.#":                        "1",
			"names.#":                      "1",
			"groups.0.id":                  CHECKSET,
			"groups.0.name":                fmt.Sprintf("tf-test-%d", rand),
			"groups.0.region_id":           CHECKSET,
			"groups.0.min_size":            "0",
			"groups.0.max_size":            "2",
			"groups.0.cooldown_time":       "20",
			"groups.0.removal_policies.#":  "2",
			"groups.0.removal_policies.0":  "OldestInstance",
			"groups.0.removal_policies.1":  "NewestInstance",
			"groups.0.load_balancer_ids.#": "0",
			"groups.0.db_instance_ids.#":   "0",
			"groups.0.vswitch_ids.#":       "1",
			"groups.0.total_capacity":      CHECKSET,
			"groups.0.active_capacity":     CHECKSET,
			"groups.0.pending_capacity":    CHECKSET,
			"groups.0.removing_capacity":   CHECKSET,
			"groups.0.creation_time":       CHECKSET,
			"groups.0.lifecycle_state":     CHECKSET,
		}
	}

	var fakeEssScalingGroupsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"groups.#": "0",
			"ids.#":    "0",
			"names.#":  "0",
		}
	}

	var essScalingGroupsCheckInfo = dataSourceAttr{
		resourceId:   "data.apsarastack_ess_scaling_groups.default",
		existMapFunc: existEssScalingGroupsMapFunc,
		fakeMapFunc:  fakeEssScalingGroupsMapFunc,
	}

	essScalingGroupsCheckInfo.dataSourceTestCheck(t, rand, nameRegexConf, idsConf, allConf)
}

func testAccCheckApsaraStackEssScalinggroupsDataSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}

	config := fmt.Sprintf(`
%s

variable "name" {
	default = "tf-test-%d"
}

resource "apsarastack_ess_scaling_group" "default" {
	min_size = 0
	max_size = 2
	scaling_group_name = "${var.name}"
	default_cooldown = 20
	removal_policies = ["OldestInstance", "NewestInstance"]
	vswitch_ids = ["${apsarastack_vswitch.default.id}"]
}

data "apsarastack_ess_scaling_groups" "default" {
  %s
}
`, EcsInstanceCommonTestCase, rand, strings.Join(pairs, "\n  "))
	return config
}
