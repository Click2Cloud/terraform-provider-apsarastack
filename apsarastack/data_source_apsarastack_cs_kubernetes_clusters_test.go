package apsarastack

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
)

func TestAccApsarastackCSKubernetesClustersDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000000, 9999999)
	resourceId := "data.apsarastack_cs_kubernetes_clusters.default"

	testAccConfig := dataSourceTestAccConfigFunc(resourceId,
		fmt.Sprintf("tf-testacckubernetes-%d", rand),
		dataSourceCSKubernetesClustersConfigDependence)

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"enable_details": "true",
			"ids":            []string{"${apsarastack_cs_kubernetes.default.id}"},
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"enable_details": "true",
			"ids":            []string{"${apsarastack_cs_kubernetes.default.id}-fake"},
		}),
	}

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"enable_details": "true",
			"name_regex":     "${apsarastack_cs_kubernetes.default.name}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"enable_details": "true",
			"name_regex":     "${apsarastack_cs_kubernetes.default.name}-fake",
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"enable_details": "true",
			"ids":            []string{"${apsarastack_cs_kubernetes.default.id}"},
			"name_regex":     "${apsarastack_cs_kubernetes.default.name}",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"enable_details": "true",
			"ids":            []string{"${apsarastack_cs_kubernetes.default.id}"},
			"name_regex":     "${apsarastack_cs_kubernetes.default.name}-fake",
		}),
	}
	var existCSKubernetesClustersMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                                      "1",
			"ids.0":                                      CHECKSET,
			"names.#":                                    "1",
			"names.0":                                    REGEXMATCH + fmt.Sprintf("tf-testacckubernetes-%d", rand),
			"clusters.#":                                 "1",
			"clusters.0.id":                              CHECKSET,
			"clusters.0.name":                            REGEXMATCH + fmt.Sprintf("tf-testacckubernetes-%d", rand),
			"clusters.0.availability_zone":               CHECKSET,
			"clusters.0.security_group_id":               CHECKSET,
			"clusters.0.nat_gateway_id":                  CHECKSET,
			"clusters.0.vpc_id":                          CHECKSET,
			"clusters.0.worker_numbers.#":                "1",
			"clusters.0.worker_numbers.0":                "1",
			"clusters.0.master_nodes.#":                  "3",
			"clusters.0.worker_disk_category":            "cloud_ssd",
			"clusters.0.master_disk_size":                "50",
			"clusters.0.master_disk_category":            "cloud_efficiency",
			"clusters.0.worker_disk_size":                "40",
			"clusters.0.connections.%":                   "4",
			"clusters.0.connections.master_public_ip":    CHECKSET,
			"clusters.0.connections.api_server_internet": CHECKSET,
			"clusters.0.connections.api_server_intranet": CHECKSET,
			"clusters.0.connections.service_domain":      CHECKSET,
		}
	}

	var fakeCSKubernetesClustersMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":      "0",
			"names.#":    "0",
			"clusters.#": "0",
		}
	}

	var csKubernetesClustersCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existCSKubernetesClustersMapFunc,
		fakeMapFunc:  fakeCSKubernetesClustersMapFunc,
	}
	preCheck := func() {
		testAccPreCheckWithRegions(t, true, connectivity.KubernetesSupportedRegions)
	}
	csKubernetesClustersCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, preCheck, idsConf, nameRegexConf, allConf)
}

func dataSourceCSKubernetesClustersConfigDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}
data "apsarastack_zones" default {
  available_resource_creation = "VSwitch"
}

data "apsarastack_instance_types" "default_m" {
	availability_zone = "${data.apsarastack_zones.default.zones.0.id}"
	cpu_core_count = 2
	memory_size = 4
	kubernetes_node_role = "Master"
}

data "apsarastack_instance_types" "default_w" {
	availability_zone = "${data.apsarastack_zones.default.zones.0.id}"
	cpu_core_count = 2
	memory_size = 4
	kubernetes_node_role = "Worker"
}

resource "apsarastack_vpc" "default" {
  name = "${var.name}"
  cidr_block = "10.1.0.0/21"
}

resource "apsarastack_vswitch" "default" {
  name = "${var.name}"
  vpc_id = "${apsarastack_vpc.default.id}"
  cidr_block = "10.1.1.0/24"
  availability_zone = "${data.apsarastack_zones.default.zones.0.id}"
}

resource "apsarastack_cs_kubernetes" "default" {
  name = "${var.name}"
  master_vswitch_ids = ["${apsarastack_vswitch.default.id}","${apsarastack_vswitch.default.id}","${apsarastack_vswitch.default.id}"]
  worker_vswitch_ids = ["${apsarastack_vswitch.default.id}"]
  new_nat_gateway = true
  master_instance_types = ["${data.apsarastack_instance_types.default_m.instance_types.0.id}","${data.apsarastack_instance_types.default_m.instance_types.0.id}","${data.apsarastack_instance_types.default_m.instance_types.0.id}"]
  worker_instance_types = ["${data.apsarastack_instance_types.default_w.instance_types.0.id}"]
  worker_number = "1"
  password = "Yourpassword1234"
  pod_cidr = "192.168.1.0/24"
  service_cidr = "192.168.2.0/24"
  enable_ssh = true
  install_cloud_monitor = true
  worker_disk_category  = "cloud_ssd"
  worker_data_disk_category = "cloud_ssd"
  worker_data_disk_size =  200
  master_disk_size = 50
}
`, name)
}
