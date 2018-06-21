package ibm

import (
	"fmt"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerWorkerPoolZoneAttachment() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerWorkerPoolZoneAttachmentCreate,
		Read:     resourceIBMContainerWorkerPoolZoneAttachmentRead,
		Update:   resourceIBMContainerWorkerPoolZoneAttachmentUpdate,
		Delete:   resourceIBMContainerWorkerPoolZoneAttachmentDelete,
		Exists:   resourceIBMContainerWorkerPoolZoneAttachmentExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cluster": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"worker_pool": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"private_vlan_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"public_vlan_id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"worker_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceIBMContainerWorkerPoolZoneAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	zone := d.Get("zone").(string)
	workerPoolZoneNetwork := v1.WorkerPoolZoneNetwork{
		PrivateVLAN: d.Get("private_vlan_id").(string),
	}

	if v, ok := d.GetOk("public_vlan_id"); ok {
		workerPoolZoneNetwork.PublicVLAN = v.(string)
	}

	workerPoolZone := v1.WorkerPoolZone{
		ID: zone,
		WorkerPoolZoneNetwork: workerPoolZoneNetwork,
	}

	cluster := d.Get("cluster").(string)
	workerPool := d.Get("worker_pool").(string)

	workerPoolsAPI := csClient.WorkerPools()

	err = workerPoolsAPI.AddZone(cluster, workerPool, workerPoolZone)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", cluster, workerPool, zone))

	return resourceIBMContainerWorkerPoolZoneAttachmentRead(d, meta)

}

func resourceIBMContainerWorkerPoolZoneAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPool := parts[1]
	zoneName := parts[2]

	workerPoolsAPI := csClient.WorkerPools()

	workerPoolRes, err := workerPoolsAPI.GetWorkerPool(cluster, workerPool)
	if err != nil {
		return err
	}
	zones := workerPoolRes.Zones

	for _, zone := range zones {
		if zone.ID == zoneName {
			d.Set("public_vlan_id", zone.PublicVLAN)
			d.Set("private_vlan_id", zone.PrivateVLAN)
			d.Set("worker_count", zone.WorkerCount)
			d.Set("zone", zone.ID)
			d.Set("cluster", cluster)
			d.Set("worker_pool", workerPool)
			break
		}
	}

	return nil
}

func resourceIBMContainerWorkerPoolZoneAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	workerPoolsAPI := csClient.WorkerPools()

	if d.HasChange("private_vlan_id") || d.HasChange("public_vlan_id") {
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		cluster := parts[0]
		workerPool := parts[1]
		zone := parts[2]
		err = workerPoolsAPI.UpdateZoneNetwork(cluster, zone, workerPool, d.Get("private_vlan_id").(string), d.Get("public_vlan_id").(string))
		if err != nil {
			return err
		}
	}

	return resourceIBMContainerWorkerPoolZoneAttachmentRead(d, meta)
}

func resourceIBMContainerWorkerPoolZoneAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPool := parts[1]
	zone := parts[2]

	workerPoolsAPI := csClient.WorkerPools()
	err = workerPoolsAPI.RemoveZone(cluster, zone, workerPool)
	if err != nil {
		return err
	}

	return nil
}

func resourceIBMContainerWorkerPoolZoneAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	cluster := parts[0]
	workerPoolID := parts[1]
	zoneID := parts[2]

	workerPoolsAPI := csClient.WorkerPools()

	workerPool, err := workerPoolsAPI.GetWorkerPool(cluster, workerPoolID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	zones := workerPool.Zones
	var zone v1.WorkerPoolZoneResponse
	for _, z := range zones {
		if z.ID == zoneID {
			zone = z
		}
	}
	return zone.ID == zoneID, nil
}
