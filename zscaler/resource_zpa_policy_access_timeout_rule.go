package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetrule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePolicyTimeoutRule() *schema.Resource {
	return &schema.Resource{
		Create:   resourcePolicyTimeoutRuleCreate,
		Read:     resourcePolicyTimeoutRuleRead,
		Update:   resourcePolicyTimeoutRuleUpdate,
		Delete:   resourcePolicyTimeoutRuleDelete,
		Importer: &schema.ResourceImporter{},

		Schema: MergeSchema(
			CommonPolicySchema(), map[string]*schema.Schema{
				"action": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "  This is for providing the rule action.",
					ValidateFunc: validation.StringInSlice([]string{
						"RE_AUTH",
					}, false),
				},

				"conditions": GetPolicyConditionsSchema([]string{
					"APP",
					"APP_GROUP",
					"CLIENT_TYPE",
					"IDP",
					"POSTURE",
					"SAML",
					"SCIM",
					"SCIM_GROUP",
				}),
			},
		),
	}
}

func resourcePolicyTimeoutRuleCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req, err := expandCreatePolicyTimeoutRule(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating zpa policy timeout rule with request\n%+v\n", req)
	if ValidateConditions(req.Conditions, zClient) {
		policysetrule, _, err := zClient.policysetrule.Create(req)
		if err != nil {
			return err
		}
		d.SetId(policysetrule.ID)
		order, ok := d.GetOk("rule_order")
		if ok {
			reorder(order, policysetrule.PolicySetID, policysetrule.ID, zClient)
		}
		return resourcePolicyTimeoutRuleRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the zpa policy timeout rule (%s) operands, please make sure you are using valid inputs for APP type, LHS & RHS", req.Name)
	}

}

func resourcePolicyTimeoutRuleRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	globalPolicyTimeout, _, err := zClient.policysetglobal.GetReauth()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Getting Policy Set Timeout Rule: globalPolicySet:%s id: %s\n", globalPolicyTimeout.ID, d.Id())
	resp, _, err := zClient.policysetrule.Get(globalPolicyTimeout.ID, d.Id())
	if err != nil {
		if obj, ok := err.(*client.ErrorResponse); ok && obj.IsObjectNotFound() {
			log.Printf("[WARN] Removing policy timeout rule %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Got Policy Set Timeout Rule:\n%+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("action", resp.Action)
	_ = d.Set("action_id", resp.ActionID)
	_ = d.Set("custom_msg", resp.CustomMsg)
	_ = d.Set("description", resp.Description)
	_ = d.Set("name", resp.Name)
	_ = d.Set("bypass_default_rule", resp.BypassDefaultRule)
	_ = d.Set("operator", resp.Operator)
	_ = d.Set("policy_set_id", resp.PolicySetID)
	_ = d.Set("policy_type", resp.PolicyType)
	_ = d.Set("priority", resp.Priority)
	_ = d.Set("reauth_default_rule", resp.ReauthDefaultRule)
	_ = d.Set("reauth_idle_timeout", resp.ReauthIdleTimeout)
	_ = d.Set("reauth_timeout", resp.ReauthTimeout)
	_ = d.Set("rule_order", resp.RuleOrder)
	_ = d.Set("conditions", FlattenPolicyConditions(resp.Conditions))

	return nil
}

func resourcePolicyTimeoutRuleUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicyTimeout, _, err := zClient.policysetglobal.GetReauth()
	if err != nil {
		return err
	}
	ruleID := d.Id()
	log.Printf("[INFO] Updating policy rule ID: %v\n", ruleID)
	req, err := expandCreatePolicyRule(d)
	if err != nil {
		return err
	}
	if ValidateConditions(req.Conditions, zClient) {
		if _, err := zClient.policysetrule.Update(globalPolicyTimeout.ID, ruleID, req); err != nil {
			return err
		}
		if d.HasChange("rule_order") {
			order, ok := d.GetOk("rule_order")
			if ok {
				reorder(order, globalPolicyTimeout.ID, ruleID, zClient)
			}
		}
		return resourcePolicyTimeoutRuleRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the zpa policy timeout (%s) operands, please make sure you are using valid inputs for APP type, LHS & RHS", req.Name)
	}
}

func resourcePolicyTimeoutRuleDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicyTimeout, _, err := zClient.policysetglobal.GetReauth()
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting Policy Timeout Rule with id %v\n", d.Id())

	if _, err := zClient.policysetrule.Delete(globalPolicyTimeout.ID, d.Id()); err != nil {
		return err
	}

	return nil

}

func expandCreatePolicyTimeoutRule(d *schema.ResourceData) (*policysetrule.PolicyRule, error) {
	policySetID, ok := d.Get("policy_set_id").(string)
	if !ok {
		log.Printf("[ERROR] policy_set_id is not set\n")
		return nil, fmt.Errorf("policy_set_id is not set")
	}
	log.Printf("[INFO] action_id:%v\n", d.Get("action_id"))
	conditions, err := ExpandPolicyConditions(d)
	if err != nil {
		return nil, err
	}
	return &policysetrule.PolicyRule{
		Action:            d.Get("action").(string),
		ActionID:          d.Get("action_id").(string),
		CustomMsg:         d.Get("custom_msg").(string),
		Description:       d.Get("description").(string),
		ID:                d.Get("id").(string),
		Name:              d.Get("name").(string),
		Operator:          d.Get("operator").(string),
		PolicySetID:       policySetID,
		PolicyType:        d.Get("policy_type").(string),
		Priority:          d.Get("priority").(string),
		ReauthDefaultRule: d.Get("reauth_default_rule").(bool),
		ReauthIdleTimeout: d.Get("reauth_idle_timeout").(string),
		ReauthTimeout:     d.Get("reauth_timeout").(string),
		RuleOrder:         d.Get("rule_order").(string),
		Conditions:        conditions,
	}, nil
}
