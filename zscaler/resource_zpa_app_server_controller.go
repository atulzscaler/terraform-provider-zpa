package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/common"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This field defines the name of the server.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field defines the description of the server.",
			},
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This field defines the domain or IP address of the server.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This field defines the status of the server.",
			},
			// App Server Group ID can only be attached if Dynamic Server Discovery in Server Group is False
			"appservergroupids": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "This field defines the list of server groups IDs.",
				Optional:    true,
			},
			"configspace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: resourceApplicationServerCreate,
		Read:   resourceApplicationServerRead,
		Update: resourceApplicationServerUpdate,
		Delete: resourceApplicationServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceApplicationServerCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreateAppServerRequest(d)
	log.Printf("[INFO] Creating zpa application server with request\n%+v\n", req)

	appservercontroller, _, err := zClient.appservercontroller.Create(req)
	if err != nil {
		return err
	}
	//d.SetId(strconv.Itoa(appservercontroller.ID))
	d.SetId(strconv.FormatInt(int64(appservercontroller.ID), 10))
	return resourceApplicationServerRead(d, m)
}

func resourceApplicationServerRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.appservercontroller.Get(id)
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing application server %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	log.Printf("[INFO] Getting application server:\n%+v\n", resp)
	_ = d.Set("address", resp.Address)
	_ = d.Set("appservergroupids", resp.AppServerGroupIds)
	_ = d.Set("configspace", resp.ConfigSpace)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("name", resp.Name)
	return nil

}

func resourceApplicationServerUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Println("An updated occurred")

	if d.HasChange("appservergroupids") || d.HasChange("name") || d.HasChange("address") {
		log.Println("The AppServerGroupID, name or address has been changed")

		if _, err := zClient.appservercontroller.Update(d.Id(), common.ApplicationServer{
			AppServerGroupIds: resourceTypeSetToStringSlice(d.Get("appservergroupids").(*schema.Set)),
			Name:              d.Get("name").(string),
			Address:           d.Get("address").(string),
			Enabled:           d.Get("enabled").(bool),
		}); err != nil {
			return err
		}
	}

	return nil
}

func resourceApplicationServerDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	log.Printf("[INFO] Deleting application server ID: %v\n", d.Id())

	err := removeServerFromGroup(zClient, d.Id())
	if err != nil {
		return err
	}

	if _, err = zClient.appservercontroller.Delete(d.Id()); err != nil {
		return err
	}

	return nil
}

func removeServerFromGroup(zClient *Client, serverID string) error {
	// Remove the reference to this server from server groups.
	id, err := strconv.ParseInt(serverID, 10, 64)
	if err != nil {
		return err
	}
	resp, _, err := zClient.appservercontroller.Get(id)
	if err != nil {
		return err
	}

	if len(resp.AppServerGroupIds) != 0 {
		log.Printf("[INFO] Removing server group ID/s from application server: %s", serverID)
		resp.AppServerGroupIds = make([]string, 0)

		log.Printf("[INFO] Updating server group ID: %s", serverID)
		_, err = zClient.appservercontroller.Update(serverID, *resp)
		if err != nil {
			log.Printf("[ERROR] Failed to update application server ID: %s", serverID)
			return err
		}
	}

	return nil
}

// func removeServer(serversList []common.ApplicationServer, serverID int64) []common.ApplicationServer {
// 	for i, server := range serversList {
// 		if server.ID == serverID {
// 			return append(serversList[:i], serversList[i+1:]...)
// 		}
// 	}
// 	return serversList
// }

func expandCreateAppServerRequest(d *schema.ResourceData) common.ApplicationServer {
	applicationServer := common.ApplicationServer{
		Address:           d.Get("address").(string),
		ConfigSpace:       d.Get("configspace").(string),
		AppServerGroupIds: resourceTypeSetToStringSlice(d.Get("appservergroupids").(*schema.Set)),
		Description:       d.Get("description").(string),
		Enabled:           d.Get("enabled").(bool),
		Name:              d.Get("name").(string),
	}
	return applicationServer
}
