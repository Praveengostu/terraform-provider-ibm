package ibm

import (
	"fmt"
	"strings"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerWorkerPool() *schema.Resource {

	return &schema.Resource{
		Create:   resourceIBMContainerWorkerPoolCreate,
		Read:     resourceIBMContainerWorkerPoolRead,
		Update:   resourceIBMContainerWorkerPoolUpdate,
		Delete:   resourceIBMContainerWorkerPoolDelete,
		Exists:   resourceIBMContainerWorkerPoolExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"machine_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},

			"worker_pool_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"size_per_zone": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"hardware": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"disk_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"kube_version": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"private_vlan": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"public_vlan": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"worker_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"labels": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
				Elem:             schema.TypeString,
			},
		},
	}
}

func resourceIBMContainerWorkerPoolCreate(d *schema.ResourceData, meta interface{}) error {

	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	workerPoolConfig := v1.WorkerPoolConfig{
		Name:        d.Get("worker_pool_name").(string),
		Size:        d.Get("size_per_zone").(int),
		MachineType: d.Get("machine_type").(string),
	}
	if v, ok := d.GetOk("hardware"); ok {
		workerPoolConfig.Isolation = v.(string)
	}
	if l, ok := d.GetOk("labels"); ok {
		labels := make(map[string]string)
		for k, v := range l.(map[string]interface{}) {
			labels[k] = v.(string)
		}
		workerPoolConfig.Labels = labels
	}
	params := v1.WorkerPoolRequest{
		WorkerPoolConfig: workerPoolConfig,
		DiskEncryption:   d.Get("disk_encryption").(bool),
	}
	cluster := d.Get("cluster").(string)

	workerPoolsAPI := csClient.WorkerPools()

	res, err := workerPoolsAPI.CreateWorkerPool(cluster, params)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", cluster, res.ID))

	return resourceIBMContainerWorkerPoolRead(d, meta)
}

func resourceIBMContainerWorkerPoolRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPoolID := parts[1]

	workerPoolsAPI := csClient.WorkerPools()

	workerPool, err := workerPoolsAPI.GetWorkerPool(cluster, workerPoolID)
	if err != nil {
		return err
	}

	machineType := workerPool.MachineType
	d.Set("worker_pool_name", workerPool.Name)
	d.Set("machine_type", machineType)
	d.Set("size_per_zone", workerPool.Size)
	d.Set("hardware", workerPool.Isolation)
	d.Set("state", workerPool.State)
	d.Set("kube_version", workerPool.WorkerVersion)
	d.Set("labels", workerPool.Labels)
	d.Set("zones", flattenZones(workerPool.Zones))
	d.Set("cluster", cluster)
	if strings.Contains(machineType, "encrypted") {
		d.Set("disk_encryption", true)
	} else {
		d.Set("disk_encryption", false)
	}
	return nil
}

func resourceIBMContainerWorkerPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPoolID := parts[1]
	workerPoolsAPI := csClient.WorkerPools()

	if d.HasChange("size_per_zone") {
		err = workerPoolsAPI.ResizeWorkerPool(cluster, workerPoolID, d.Get("size_per_zone").(int))
		if err != nil {
			return err
		}
	}
	return resourceIBMContainerWorkerPoolRead(d, meta)
}

func resourceIBMContainerWorkerPoolDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	workerPoolID := parts[1]

	workerPoolsAPI := csClient.WorkerPools()

	err = workerPoolsAPI.DeleteWorkerPool(cluster, workerPoolID)
	if err != nil {
		return err
	}

	return nil
}

func resourceIBMContainerWorkerPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
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

	return workerPool.ID == workerPoolID, nil
}
