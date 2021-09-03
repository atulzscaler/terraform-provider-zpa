package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetrule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePolicySetRule() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicySetCreate,
		Read:   resourcePolicySetRead,
		Update: resourcePolicySetUpdate,
		Delete: resourcePolicySetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "  This is for providing the rule action.",
			},
			"bypass_default_rule": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"custom_msg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This is for providing a customer message for the user.",
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
			"isolation_default_rule": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the name of the policy rule.",
			},
			"operator": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "This denotes the operation type.",
			},
			"policy_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_type": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reauth_default_rule": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"reauth_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reauth_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rule_order": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// "zpn_cbi_profile_id": {
			//  Type:     schema.TypeInt,
			//  Optional: true,
			// },
			"zpn_inspection_profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// "action_id": {
			//  Type:        schema.TypeInt,
			//  Optional:    true,
			//  Description: "This field defines the description of the server.",
			// },
			// "server_groups": {
			//  Type:        schema.TypeList,
			//  Optional:    true,
			//  Description: "ID of the server group.",
			//  Elem: &schema.Resource{
			//      Schema: map[string]*schema.Schema{
			//          "id": {
			//              Type:     schema.TypeList,
			//              Optional: true,
			//              Elem:     &schema.Schema{Type: schema.TypeInt},
			//          },
			//      },
			//  },
			// },
			// "app_connector_groups": {
			//  Type:        schema.TypeList,
			//  Optional:    true,
			//  Description: "This field is a json array of app-connector-id only.",
			//  Elem: &schema.Resource{
			//      Schema: map[string]*schema.Schema{
			//          "id": {
			//              Type:     schema.TypeList,
			//              Optional: true,
			//              Elem:     &schema.Schema{Type: schema.TypeInt},
			//          },
			//      },
			//  },
			// },
			"conditions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "This is for proviidng the set of conditions for the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"negated": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"operands": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "This signifies the various policy criteria.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"idp_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"lhs": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"object_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "  This is for specifying the policy critiera.",
									},
									"rhs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "This denotes the value for the given object type. Its value depends upon the key.",
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

func resourcePolicySetCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreatePolicyRule(d)
	log.Printf("[INFO] Creating zpa policy rule with request\n%+v\n", req)

	policysetrule, _, err := zClient.policysetrule.Create(&req)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(int64(policysetrule.ID), 10))

	return resourcePolicySetRead(d, m)
}

// Please review read operations. It needs to pull the policySetId and RuleId in order to read a specific rule.
func resourcePolicySetRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	globalPolicySet, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}
	ruleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	resp, _, err := zClient.policysetrule.Get(globalPolicySet.ID, ruleId)
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing policy rule %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting Policy Set Rule:\n%+v\n", resp)
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("action", resp.Action)
	_ = d.Set("action_id", resp.ActionID)
	// _ = d.Set("creation_time", resp.CreationTime)
	_ = d.Set("custom_msg", resp.CustomMsg)
	_ = d.Set("description", resp.Description)
	// _ = d.Set("modifiedby", resp.ModifiedBy)
	// _ = d.Set("modified_time", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("operator", resp.Operator)
	_ = d.Set("policy_set_id", resp.PolicySetID)
	_ = d.Set("policy_type", resp.PolicyType)
	_ = d.Set("priority", resp.Priority)
	//_ = d.Set("reauth_idle_timeout", resp.ReauthIdleTimeout)
	//_ = d.Set("reauth_timeout", resp.ReauthTimeout)
	_ = d.Set("rule_order", resp.RuleOrder)
	//_ = d.Set("zpn_cbi_profile_id", resp.ZpnCbiProfileID)
	_ = d.Set("conditions", flattenPolicyRuleConditions(resp.Conditions))

	return nil
}

// Please review Update operations. It needs to pull the policySetId and RuleId in order to update a specific rule.
func resourcePolicySetUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicySet, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}
	ruleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Updating policy rule ID: %v\n", ruleId)
	req := expandCreatePolicyRule(d)

	if _, err := zClient.policysetrule.Update(globalPolicySet.ID, ruleId, &req); err != nil {
		return err
	}

	return resourcePolicySetRead(d, m)
}

// Please review Delete operations. It needs to pull the policySetId and RuleId in order to delete a specific rule.
func resourcePolicySetDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicySet, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting IP list with id %v\n", id)

	if _, err := zClient.policysetrule.Delete(globalPolicySet.ID, id); err != nil {
		return err
	}

	return nil

}

// Please review the expand and flattening functions. Condition is actually a slice inside PolicyRule
//https://help.zscaler.com/zpa/api-reference#/policy-set-controller/addRuleToPolicySet
func expandCreatePolicyRule(d *schema.ResourceData) policysetrule.PolicyRule {
	policySetID, err := strconv.ParseInt(d.Get("policy_set_id").(string), 10, 64)
	if err != nil {
		log.Printf("[ERROR] policy_set_id is not set, err:%v\n", err)
	}
	return policysetrule.PolicyRule{
		Action: d.Get("action").(string),
		// ActionID:     d.Get("action_id").(int),
		// CreationTime: d.Get("creation_time").(int),
		CustomMsg:   d.Get("custom_msg").(string),
		Description: d.Get("description").(string),
		// ID:          d.Get("id").(int),
		// ModifiedBy:        d.Get("modifiedby").(int),
		// ModifiedTime:      d.Get("Modified_time").(int),
		Name: d.Get("name").(string),
		// Operator:    d.Get("operator").(string),
		PolicySetID: policySetID,
		PolicyType:  d.Get("policy_type").(int),
		Priority:    d.Get("priority").(int),
		// ReauthIdleTimeout: d.Get("reauth_idle_timeout").(int),
		// ReauthTimeout:     d.Get("reauth_timeout").(int),
		RuleOrder: d.Get("rule_order").(int),
		//ZpnCbiProfileID: d.Get("zpn_cbi_profile_id").(int),
		Conditions: expandConditionSet(d),
	}
}

func expandConditionSet(d *schema.ResourceData) []policysetrule.Conditions {
	conditionInterface, ok := d.GetOk("conditions")
	if ok {
		conditions := conditionInterface.([]interface{})
		log.Printf("[INFO] conditions data: %+v\n", conditions)
		var conditionSets []policysetrule.Conditions
		for _, condition := range conditions {
			conditionSet, _ := condition.(map[string]interface{})
			if conditionSet != nil {
				conditionSets = append(conditionSets, policysetrule.Conditions{
					// CreationTime: conditionSet["creation_time"].(int),
					// ID:           conditionSet["id"].(int),
					// ModifiedBy:   conditionSet["modifiedby"].(int),
					// ModifiedTime: conditionSet["modified_time"].(int),
					Negated:  conditionSet["negated"].(bool),
					Operator: conditionSet["operator"].(string),
					Operands: expandOperandsList(conditionSet["operands"]),
				})
			}
		}
		return conditionSets
	}

	return []policysetrule.Conditions{}
}

func expandOperandsList(ops interface{}) []policysetrule.Operands {
	if ops != nil {
		operands := ops.([]interface{})
		log.Printf("[INFO] operands data: %+v\n", operands)
		var operandsSets []policysetrule.Operands
		for _, operand := range operands {
			operandSet, _ := operand.(map[string]interface{})
			id, _ := operandSet["id"].(int64)
			IdpID, _ := operandSet["idp_id"].(int64)
			if operandSet != nil {
				operandsSets = append(operandsSets, policysetrule.Operands{
					ID:         id,
					IdpID:      IdpID,
					LHS:        operandSet["lhs"].(string),
					ObjectType: operandSet["object_type"].(string),
					RHS:        operandSet["rhs"].(string),
				})
			}
		}

		return operandsSets
	}
	return []policysetrule.Operands{}
}
func flattenPolicyRuleConditions(conditions []policysetrule.Conditions) []interface{} {
	ruleConditions := make([]interface{}, len(conditions))
	for i, ruleConditionItems := range conditions {
		ruleConditions[i] = map[string]interface{}{
			"creation_time": ruleConditionItems.CreationTime,
			"id":            ruleConditionItems.ID,
			"modifiedby":    ruleConditionItems.ModifiedBy,
			"modified_time": ruleConditionItems.ModifiedTime,
			"negated":       ruleConditionItems.Negated,
			"operator":      ruleConditionItems.Operator,
			"operands":      flattenPolicyRuleOperands(ruleConditionItems.Operands),
		}
	}

	return ruleConditions
}

func flattenPolicyRuleOperands(conditionOperand []policysetrule.Operands) []interface{} {
	conditionOperands := make([]interface{}, len(conditionOperand))
	for i, operandItems := range conditionOperand {
		conditionOperands[i] = map[string]interface{}{
			"creation_time": operandItems.CreationTime,
			"id":            operandItems.ID,
			"idp_id":        operandItems.IdpID,
			"lhs":           operandItems.LHS,
			"modifiedby":    operandItems.ModifiedBy,
			"modified_time": operandItems.ModifiedTime,
			"object_type":   operandItems.ObjectType,
			"rhs":           operandItems.RHS,
		}
	}

	return conditionOperands
}

// Need to flatten the Operands menu, which is a slice inside the slice Conditions
//https://help.zscaler.com/zpa/api-reference#/policy-set-controller/addRuleToPolicySet
