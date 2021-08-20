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
						"creationtime": {
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
									"creationtime": {
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
									"issuedcertid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"machinegroupid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machinegroupname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"machinetokenid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"modifiedby": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"modifiedtime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									// Placeholder
									// I am not sure if this is the proper structure to hold this data
									/*"signingcert": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem: map[string]*schema.Schema{
											"additionalProp1": {
												Type:     schema.TypeString,
												Computed: true,
											},
											"additionalProp2": {
												Type:     schema.TypeString,
												Computed: true,
											},
											"additionalProp3": {
												Type:     schema.TypeString,
												Computed: true,
											},
										},
									},*/
								},
							},
						},
						"modifiedby": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"modifiedtime": {
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
			// May not be necessary for Terraform.
			// "totalpages": {
			// 	Type:     schema.TypeInt,
			// 	Optional: true,
			// },
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
			"creationtime": machineGroupItem.CreationTime,
			"description":  machineGroupItem.Description,
			"enabled":      machineGroupItem.Enabled,
			"id":           machineGroupItem.ID,
			"machines":     flattenMachines(machineGroupItem),
			"modifiedby":   machineGroupItem.ModifiedBy,
			"modifiedtime": machineGroupItem.ModifiedTime,
			"name":         machineGroupItem.Name,
		}
	}

	return machineGroupsSlice
}

func flattenMachines(machineGroup machinegroup.MachineGroup) []interface{} {
	machines := make([]interface{}, len(machineGroup.Machines))
	for i, machineItem := range machineGroup.Machines {
		machines[i] = map[string]interface{}{
			"creationtime":     machineItem.CreationTime,
			"description":      machineItem.Description,
			"fingerprint":      machineItem.Fingerprint,
			"id":               machineItem.ID,
			"issuedcertid":     machineItem.IssuedCertID,
			"machinegroupid":   machineItem.MachineGroupID,
			"machinegroupname": machineItem.MachineGroupName,
			"machinetokenid":   machineItem.MachineTokenID,
			"modifiedby":       machineItem.ModifiedBy,
			"modifiedtime":     machineItem.ModifiedTime,
			"name":             machineItem.Name,
			"signingcert":      machineItem.SigningCert,
		}
	}

	return machines
}
