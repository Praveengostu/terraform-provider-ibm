package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
)

func dataSourceIBMLbaas() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMLbaasRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datacenter": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_instances_up": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"server_instances_down": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"active_connections": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocols": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"frontend_protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"frontend_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"backend_protocol": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backend_port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"load_balancing_method": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"session_stickiness": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_conn": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tls_certificate_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"protocol_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"server_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"member_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMLbaasRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetNetworkLBaaSLoadBalancerService(sess)
	lbs, err := service.Mask("datacenter,members,listeners.defaultPool,listeners.defaultPool.sessionAffinity").Filter(filter.Build(
		filter.Path("name").Eq(name))).GetAllObjects()
	if err != nil {
		return err
	}
	if len(lbs) != 1 {
		return fmt.Errorf("No load balancer with name: %s", name)
	}
	result := lbs[0]

	//Get statistics
	lbStat, err := service.GetLoadBalancerStatistics(result.Uuid)
	if err != nil {
		return fmt.Errorf("Error retrieving load balancer statistics: %s", err)
	}
	//Get members health
	lbMembersHealth, err := service.GetLoadBalancerMemberHealth(result.Uuid)
	if err != nil {
		return fmt.Errorf("Error retrieving load balancer members: %s", err)
	}
	members := flattenServerInstances(result.Members)

	for _, lbHealth := range lbMembersHealth {
		for _, lbMemHealth := range lbHealth.MembersHealth {
			for _, member := range members {
				if member["member_id"] == *lbMemHealth.Uuid {
					member["status"] = *lbMemHealth.Status
				}
			}
		}
	}

	var lbType string
	if *result.IsPublic == 1 {
		lbType = "PUBLIC"
	} else {
		lbType = "PRIVATE"
	}

	d.SetId(*result.Uuid)

	d.Set("name", result.Name)
	d.Set("description", result.Description)
	d.Set("server_instances_up", lbStat.NumberOfMembersUp)
	d.Set("server_instances_down", lbStat.NumberOfMembersDown)
	d.Set("active_connections", lbStat.TotalConnections)
	if result.Datacenter != nil {
		d.Set("datacenter", result.Datacenter.Name)
	}
	d.Set("type", lbType)
	d.Set("status", result.OperatingStatus)
	d.Set("vip", result.Address)
	d.Set("protocols", flattenProtocols(result.Listeners))
	d.Set("server_instances", members)
	return nil
}
