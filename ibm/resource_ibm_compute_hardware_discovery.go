package ibm

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMComputeHardwareDiscovery() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMComputeHardwareDiscoveryCreate,
		Read:   resourceIBMComputeHardwareDiscoveryRead,
		Delete: resourceIBMComputeHardwareDiscoveryDelete,

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

			"hardwares": {
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

func resourceIBMComputeHardwareDiscoveryCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMComputeHardwareDiscoveryRead(d, meta)
}

func resourceIBMComputeHardwareDiscoveryRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	username := d.Get("user_name")
	var bms []datatypes.Hardware
	var err error

	if username != "" {
		bms, err = service.
			Mask(
				"id,hostname,provisionDate,fullyQualifiedDomainName,billingItem.orderItem.order.userRecord.username",
			).Filter(
			filter.Build(
				filter.Path("hardware.billingItem.orderItem.order.userRecord.username").Eq(username),
			),
		).GetHardware()
		if err != nil {
			return fmt.Errorf("Error retrieving bare metal details: %s", err)
		}
	} else {
		bms, err = service.
			Mask(
				"id,hostname,provisionDate,fullyQualifiedDomainName,billingItem.orderItem.order.userRecord.username",
			).GetHardware()
		if err != nil {
			return fmt.Errorf("Error retrieving virtual guest details for host : %s", err)
		}
	}

	if len(bms) == 0 {
		return fmt.Errorf("No bare metals found ")
	}

	d.Set("total_bms", len(bms))
	result := make([]map[string]interface{}, 0, len(bms))
	ids := ""
	var age int
	if v, ok := d.GetOk("min_age"); ok {
		age = v.(int)
	}
	for _, i := range bms {
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
		ids = "bmdiscovery"
	}
	d.SetId(ids)
	d.Set("hardwares", result)

	return nil
}

func resourceIBMComputeHardwareDiscoveryDelete(d *schema.ResourceData, meta interface{}) error {

	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetHardwareService(sess)
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	for _, i := range parts {

		id, err := strconv.Atoi(i)
		if err != nil {
			return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
		}

		_, err = waitForNoBareMetalActiveTransactions(id, meta)
		if err != nil {
			return fmt.Errorf("Error deleting bare metal server while waiting for zero active transactions: %s", err)
		}

		billingItem, err := service.Id(id).GetBillingItem()
		if err != nil {
			return fmt.Errorf("Error getting billing item for bare metal server: %s", err)
		}

		// Monthly bare metal servers only support an anniversary date cancellation option.
		billingItemService := services.GetBillingItemService(sess)
		_, err = billingItemService.Id(*billingItem.Id).CancelItem(
			sl.Bool(true), sl.Bool(true), sl.String("No longer required"), sl.String("Please cancel this server"),
		)
		if err != nil {
			return fmt.Errorf("Error canceling the bare metal server (%d): %s", id, err)
		}
	}

	return nil

}
