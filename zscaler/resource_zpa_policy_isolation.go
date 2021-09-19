package zscaler

import (
	"fmt"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetrule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePolicyIsolation() *schema.Resource {
	return &schema.Resource{
		Create:   resourcePolicyIsolationCreate,
		Read:     resourcePolicyIsolationRead,
		Update:   resourcePolicyIsolationUpdate,
		Delete:   resourcePolicyIsolationDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "  This is for providing the rule action.",
				ValidateFunc: validation.StringInSlice([]string{
					"ISOLATE",
					"BYPASS_ISOLATE",
				}, false),
			},
			"action_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field defines the description of the server.",
			},
			"bypass_default_rule": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"isolation_default_rule": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This is the description of the access policy rule.",
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the name of the policy isolation rule.",
			},
			"operator": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AND",
					"OR",
				}, false),
			},
			"policy_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zpn_cbi_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"conditions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "This is for proviidng the set of conditions for the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"negated": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"AND",
								"OR",
							}, false),
						},
						"operands": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "This signifies the various policy criteria.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"idp_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"lhs": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"rhs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "This denotes the value for the given object type. Its value depends upon the key.",
									},
									"object_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "  This is for specifying the policy critiera.",
										ValidateFunc: validation.StringInSlice([]string{
											"APP",
											"APP_GROUP",
											"CLIENT_TYPE",
											"CLOUD_CONNECTOR_GROUP",
											"IDP",
											"POSTURE",
											"SAML",
											"SCIM",
											"SCIM_GROUP",
											"TRUSTED_NETWORK",
										}, false),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourcePolicyIsolationCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreatePolicyIsolationRule(d)
	log.Printf("[INFO] Creating zpa policy Isolation rule with request\n%+v\n", req)
	if ValidateConditions(req.Conditions, zClient) {
		policysetrule, _, err := zClient.policysetrule.Create(&req)
		if err != nil {
			return err
		}
		d.SetId(policysetrule.ID)
		order, ok := d.GetOk("rule_order")
		if ok {
			reorder(order, policysetrule.PolicySetID, policysetrule.ID, zClient)
		}
		return resourcePolicyIsolationRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the zpa policy Isolation (%s) operands, please make sure you are using valid inputs for APP type, LHS & RHS", req.Name)
	}

}

func resourcePolicyIsolationRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	globalPolicyIsolation, _, err := zClient.policysetglobal.GetIsolationPolicy()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Getting Policy Set Isolation Rule: globalPolicySet:%s id: %s\n", globalPolicyIsolation.ID, d.Id())
	resp, _, err := zClient.policysetrule.Get(globalPolicyIsolation.ID, d.Id())
	if err != nil {
		if obj, ok := err.(*client.ErrorResponse); ok && obj.IsObjectNotFound() {
			log.Printf("[WARN] Removing policy Isolation rule %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Got Policy Set Isolation Rule:\n%+v\n", resp)
	d.SetId(resp.ID)
	_ = d.Set("action", resp.Action)
	_ = d.Set("action_id", resp.ActionID)
	_ = d.Set("isolation_default_rule", resp.IsolationDefaultRule)
	_ = d.Set("description", resp.Description)
	_ = d.Set("name", resp.Name)
	_ = d.Set("bypass_default_rule", resp.BypassDefaultRule)
	_ = d.Set("operator", resp.Operator)
	_ = d.Set("policy_set_id", resp.PolicySetID)
	_ = d.Set("policy_type", resp.PolicyType)
	_ = d.Set("priority", resp.Priority)
	_ = d.Set("rule_order", resp.RuleOrder)
	_ = d.Set("zpn_cbi_profile_id", resp.ZpnCbiProfileId)
	_ = d.Set("conditions", flattenPolicyIsolationConditions(resp.Conditions))

	return nil
}

func resourcePolicyIsolationUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicyIsolation, _, err := zClient.policysetglobal.GetIsolationPolicy()
	if err != nil {
		return err
	}
	ruleID := d.Id()
	log.Printf("[INFO] Updating policy Isolation rule ID: %v\n", ruleID)
	req := expandCreatePolicyRule(d)
	if ValidateConditions(req.Conditions, zClient) {
		if _, err := zClient.policysetrule.Update(globalPolicyIsolation.ID, ruleID, &req); err != nil {
			return err
		}
		if d.HasChange("rule_order") {
			order, ok := d.GetOk("rule_order")
			if ok {
				reorder(order, globalPolicyIsolation.ID, ruleID, zClient)
			}
		}
		return resourcePolicyIsolationRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the zpa policy Isolation (%s) operands, please make sure you are using valid inputs for APP type, LHS & RHS", req.Name)
	}

}

func resourcePolicyIsolationDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicyIsolation, _, err := zClient.policysetglobal.GetIsolationPolicy()
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting policy Isolation rule with id %v\n", d.Id())

	if _, err := zClient.policysetrule.Delete(globalPolicyIsolation.ID, d.Id()); err != nil {
		return err
	}

	return nil

}

func expandCreatePolicyIsolationRule(d *schema.ResourceData) policysetrule.PolicyRule {
	policySetID, ok := d.Get("policy_set_id").(string)
	if !ok {
		log.Printf("[ERROR] policy_set_id is not set\n")
	}
	log.Printf("[INFO] action_id:%v\n", d.Get("action_id"))
	return policysetrule.PolicyRule{
		Action:   d.Get("action").(string),
		ActionID: d.Get("action_id").(string),
		// CustomMsg:       d.Get("custom_msg").(string),
		Description:     d.Get("description").(string),
		ID:              d.Get("id").(string),
		Name:            d.Get("name").(string),
		Operator:        d.Get("operator").(string),
		PolicySetID:     policySetID,
		PolicyType:      d.Get("policy_type").(string),
		Priority:        d.Get("priority").(string),
		RuleOrder:       d.Get("rule_order").(string),
		ZpnCbiProfileId: d.Get("zpn_cbi_profile_id").(string),
		Conditions:      expandPolicyIsolationConditionSet(d),
	}
}

func expandPolicyIsolationConditionSet(d *schema.ResourceData) []policysetrule.Conditions {
	conditionInterface, ok := d.GetOk("conditions")
	if ok {
		conditions := conditionInterface.([]interface{})
		log.Printf("[INFO] conditions data: %+v\n", conditions)
		var conditionSets []policysetrule.Conditions
		for _, condition := range conditions {
			conditionSet, _ := condition.(map[string]interface{})
			if conditionSet != nil {
				conditionSets = append(conditionSets, policysetrule.Conditions{
					ID:       conditionSet["id"].(string),
					Negated:  conditionSet["negated"].(bool),
					Operator: conditionSet["operator"].(string),
					Operands: expandPolicyIsolationOperandsList(conditionSet["operands"]),
				})
			}
		}
		return conditionSets
	}

	return []policysetrule.Conditions{}
}

func expandPolicyIsolationOperandsList(ops interface{}) []policysetrule.Operands {
	if ops != nil {
		operands := ops.([]interface{})
		log.Printf("[INFO] operands data: %+v\n", operands)
		var operandsSets []policysetrule.Operands
		for _, operand := range operands {
			operandSet, _ := operand.(map[string]interface{})
			id, _ := operandSet["id"].(string)
			IdpID, _ := operandSet["idp_id"].(string)
			if operandSet != nil {
				operandsSets = append(operandsSets, policysetrule.Operands{
					ID:         id,
					IdpID:      IdpID,
					LHS:        operandSet["lhs"].(string),
					ObjectType: operandSet["object_type"].(string),
					RHS:        operandSet["rhs"].(string),
					Name:       operandSet["name"].(string),
				})
			}
		}

		return operandsSets
	}
	return []policysetrule.Operands{}
}
func flattenPolicyIsolationConditions(conditions []policysetrule.Conditions) []interface{} {
	ruleConditions := make([]interface{}, len(conditions))
	for i, ruleConditionItems := range conditions {
		ruleConditions[i] = map[string]interface{}{
			"id":       ruleConditionItems.ID,
			"negated":  ruleConditionItems.Negated,
			"operator": ruleConditionItems.Operator,
			"operands": flattenPolicyIsolationOperands(ruleConditionItems.Operands),
		}
	}

	return ruleConditions
}

func flattenPolicyIsolationOperands(conditionOperand []policysetrule.Operands) []interface{} {
	conditionOperands := make([]interface{}, len(conditionOperand))
	for i, operandItems := range conditionOperand {
		conditionOperands[i] = map[string]interface{}{
			"id":          operandItems.ID,
			"idp_id":      operandItems.IdpID,
			"lhs":         operandItems.LHS,
			"object_type": operandItems.ObjectType,
			"rhs":         operandItems.RHS,
		}
	}

	return conditionOperands
}
