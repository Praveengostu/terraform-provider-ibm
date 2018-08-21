package ibm

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
)

func resourceIBMComputeVmInstancesDiscovery() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMVmInstancesDiscoveryCreate,
		Read:   resourceIBMVmInstancesDiscoveryRead,
		Delete: resourceIBMVmInstancesDiscoveryDelete,

		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"min_age": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			"vms": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"provision_date": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMVmInstancesDiscoveryCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMVmInstancesDiscoveryRead(d, meta)
}

func resourceIBMVmInstancesDiscoveryRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	username := d.Get("user_name")
	var vgs []datatypes.Virtual_Guest
	var err error

	if username != "" {
		vgs, err = service.
			Mask(
				"id,hostname,provisionDate,fullyQualifiedDomainName,billingItem.orderItem.order.userRecord.username",
			).Filter(
			filter.Build(
				filter.Path("virtualGuests.billingItem.orderItem.order.userRecord.username").Eq(username),
			),
		).GetVirtualGuests()
		if err != nil {
			return fmt.Errorf("Error retrieving virtual guest details for host : %s", err)
		}
	} else {
		vgs, err = service.
			Mask(
				"id,hostname,provisionDate,fullyQualifiedDomainName,billingItem.orderItem.order.userRecord.username",
			).GetVirtualGuests()
		if err != nil {
			return fmt.Errorf("Error retrieving virtual guest details for host : %s", err)
		}
	}

	if len(vgs) == 0 {
		return fmt.Errorf("No virtual guest found ")
	}

	result := make([]map[string]interface{}, 0, len(vgs))
	ids := ""
	var age int
	if v, ok := d.GetOk("min_age"); ok {
		age = v.(int)
	}

	for _, i := range vgs {
		days := 0
		if i.ProvisionDate != nil {
			y1, m1, d1 := i.ProvisionDate.Date()
			y2, m2, d2 := time.Now().Date()
			t1 := Date(y1, m1, d1)
			t2 := Date(y2, m2, d2)
			days = int(t2.Sub(t1).Hours() / 24)
		}
		if days >= age {
			if ids != "" {
				ids = ids + "/"
			}
			id := strconv.Itoa(*i.Id)
			ids = ids + id
			l := map[string]interface{}{
				"hostname":  *i.Hostname,
				"user_name": *i.BillingItem.Billing_Item.OrderItem.Order.UserRecord.Username,
			}
			if i.ProvisionDate != nil {
				l["provision_date"] = i.ProvisionDate.String()
			}
			result = append(result, l)
		}
	}
	if ids == "" {
		ids = "vgdiscovery"
	}
	d.SetId(ids)
	d.Set("vms", result)

	return nil
}

//Date ...
func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func resourceIBMVmInstancesDiscoveryDelete(d *schema.ResourceData, meta interface{}) error {

	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetVirtualGuestService(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	for _, i := range parts {
		id, err := strconv.Atoi(i)
		if err != nil {
			return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
		}

		_, err = WaitForNoActiveTransactions(d, meta)

		if err != nil {
			return fmt.Errorf("Error deleting virtual guest, couldn't wait for zero active transactions: %s", err)
		}
		err = detachSecurityGroupNetworkComponentBindings(d, meta, id)
		if err != nil {
			return err
		}
		ok, err := service.Id(id).DeleteObject()
		if err != nil {
			return fmt.Errorf("Error deleting virtual guest: %s", err)
		}

		if !ok {
			return fmt.Errorf(
				"API reported it was unsuccessful in removing the virtual guest '%d'", id)
		}
	}

	return nil

}
