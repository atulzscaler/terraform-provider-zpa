package zscaler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/machinegroup"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMachineGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMachineGroupRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creation_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"machines": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"creation_time": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fingerprint": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"issued_cert_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"machine_group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machine_group_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machine_token_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"modifiedby": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"modified_time": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"signing_cert": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"modifiedby": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modified_time": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceMachineGroupRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	machineGroupID := ""
	if machineGroupIDInterface, ok := d.GetOk("id"); ok {
		machineGroupID = machineGroupIDInterface.(string)
	}

	if len(machineGroupID) != 0 {
		log.Printf("[INFO] Getting data for machine group %s\n", machineGroupID)
		// Getting specific machine group ID
		resp, _, err := zClient.machinegroup.Get(machineGroupID)
		if err != nil {
			return err
		}

		// Add the one machine group ID we received
		d.SetId(fmt.Sprintf("%d", resp.ID))
		// Now we make a slice of machine groups, so we can add the one machine group to the resource list after flattening
		machineGroups := make([]machinegroup.MachineGroup, 0)
		machineGroups = append(machineGroups, *resp)
		d.Set("list", flattenMachineGroups(machineGroups))
	} else {
		log.Printf("[INFO] Getting data for all machine groups \n")
		// Getting all machine groups
		resp, _, err := zClient.machinegroup.GetAll()
		if err != nil {
			return err
		}

		// In case of all machine groups returned, I don't now which ID to set as ID here, so I add time from documentation
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		d.Set("list", flattenMachineGroups(resp.List))
	}

	return nil
}

func flattenMachineGroups(machineGroups []machinegroup.MachineGroup) []interface{} {
	machineGroupsSlice := make([]interface{}, len(machineGroups))
	for i, machineGroupItem := range machineGroups {
		machineGroupsSlice[i] = map[string]interface{}{
			"creation_time": machineGroupItem.CreationTime,
			"description":   machineGroupItem.Description,
			"enabled":       machineGroupItem.Enabled,
			"id":            machineGroupItem.ID,
			"modifiedby":    machineGroupItem.ModifiedBy,
			"modified_time": machineGroupItem.ModifiedTime,
			"name":          machineGroupItem.Name,
			"machines":      flattenMachines(machineGroupItem),
		}
	}

	return machineGroupsSlice
}

func flattenMachines(machineGroup machinegroup.MachineGroup) []interface{} {
	machines := make([]interface{}, len(machineGroup.Machines))
	for i, machineItem := range machineGroup.Machines {
		machines[i] = map[string]interface{}{
			"creation_time":      machineItem.CreationTime,
			"description":        machineItem.Description,
			"fingerprint":        machineItem.Fingerprint,
			"id":                 machineItem.ID,
			"issued_cert_id":     machineItem.IssuedCertID,
			"machine_group_id":   machineItem.MachineGroupID,
			"machine_group_name": machineItem.MachineGroupName,
			"machine_token_id":   machineItem.MachineTokenID,
			"modifiedby":         machineItem.ModifiedBy,
			"modified_time":      machineItem.ModifiedTime,
			"name":               machineItem.Name,
			"signing_cert":       machineItem.SigningCert,
		}
	}

	return machines
}
