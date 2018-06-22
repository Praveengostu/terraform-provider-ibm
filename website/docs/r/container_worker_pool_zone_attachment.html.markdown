---
layout: "ibm"
page_title: "IBM: container_worker_pool_zone_attachment"
sidebar_current: "docs-ibm-resource-container-worker-pool-zone-attachment"
description: |-
  Manages IBM container worker pool zone attachment.
---

# ibm\_container_cluster

Create, update, or delete a zone. This resource creates the zone and attaches it to the specified worker pool.

## Example Usage

In the following example, you can create a zone:

```hcl
resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster      = "my_cluster"
  worker_pool  = "my_cluster_pool"
  zone         = "dal12"
  private_vlan = "2320267"
  public_vlan  = "2320265"
}

```

## Argument Reference

The following arguments are supported:

* `zone` - (Required, string) The name of the zone.
* `cluster` - (Required, string) The name or id of the cluster.
* `worker_pool` - (Required, string) The name or id of the worker pool.
* `private_vlan_id` - (Required, string) The private VLAN of the worker node. You can retrieve the value by running the `bx cs vlans <data-center>` command in the IBM Cloud CLI.
* `public_vlan_id` - (Optional, string) The public VLAN of the worker node. You can retrieve the value by running the `bx cs vlans <data-center>` command in the IBM Cloud CLI..


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker pool zone attachment resource. The id is composed of \<cluster_name_id\>/\< worker_pool_name_id\>/\<zone/>
* `worker_count` - Number of workers attached to this zone.

## Import

ibm_container_worker_pool_zone_attachment can be imported using cluster_name_id, worker_pool_name_id and zone, eg

```
$ terraform import ibm_container_worker_pool_zone_attachment.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35/dal10