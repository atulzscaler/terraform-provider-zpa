package zscaler

import (
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetrule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePolicyTimeout() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyTimeoutCreate,
		Read:   resourcePolicyTimeoutRead,
		Update: resourcePolicyTimeoutUpdate,
		Delete: resourcePolicyTimeoutDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "  This is for providing the rule action.",
				ValidateFunc: validation.StringInSlice([]string{
					"RE_AUTH",
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the name of the policy rule.",
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
			"reauth_default_rule": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"reauth_idle_timeout": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reauth_timeout": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_server_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "ID of the server group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						//  "id": {
						//      Type:     schema.TypeSet,
						//      Optional: true,
						//      Elem:     &schema.Schema{Type: schema.TypeInt},
						//  },
					},
				},
			},
			"app_connector_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "This field is a json array of app-connector-id only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						//  "id": {
						//      Type:     schema.TypeSet,
						//      Optional: true,
						//      Elem:     &schema.Schema{Type: schema.TypeInt},
						//  },
					},
				},
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
											"IDP",
											"POSTURE",
											"SAML",
											"SCIM",
											"SCIM_GROUP",
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

func resourcePolicyTimeoutCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreatePolicyTimeoutRule(d)
	log.Printf("[INFO] Creating zpa policy rule with request\n%+v\n", req)

	policysetrule, _, err := zClient.policysetrule.Create(&req)
	if err != nil {
		return err
	}
	d.SetId(policysetrule.ID)

	return resourcePolicyTimeoutRead(d, m)
}

// Please review read operations. It needs to pull the policySetId and RuleId in order to read a specific rule.
func resourcePolicyTimeoutRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	globalPolicyTimeout, _, err := zClient.policysetglobal.GetReauth()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Getting Policy Set Rule: globalPolicySet:%s id: %s\n", globalPolicyTimeout.ID, d.Id())
	resp, _, err := zClient.policysetrule.Get(globalPolicyTimeout.ID, d.Id())
	if err != nil {
		if obj, ok := err.(*client.ErrorResponse); ok && obj.IsObjectNotFound() {
			log.Printf("[WARN] Removing policy rule %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Got Policy Set Rule:\n%+v\n", resp)
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
	_ = d.Set("conditions", flattenPolicyTimeoutConditions(resp.Conditions))
	_ = d.Set("app_server_groups", flattenPolicyTimeoutServerGroups(resp.AppServerGroups))
	_ = d.Set("app_connector_groups", flattenPolicyTimeoutAppConnectorGroups(resp.AppConnectorGroups))

	return nil
}

func resourcePolicyTimeoutUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicyTimeout, _, err := zClient.policysetglobal.GetReauth()
	if err != nil {
		return err
	}
	ruleId := d.Id()
	log.Printf("[INFO] Updating policy rule ID: %v\n", ruleId)
	req := expandCreatePolicyRule(d)

	if _, err := zClient.policysetrule.Update(globalPolicyTimeout.ID, ruleId, &req); err != nil {
		return err
	}

	return resourcePolicyTimeoutRead(d, m)
}

func resourcePolicyTimeoutDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicyTimeout, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting IP list with id %v\n", d.Id())

	if _, err := zClient.policysetrule.Delete(globalPolicyTimeout.ID, d.Id()); err != nil {
		return err
	}

	return nil

}

// Please review the expand and flattening functions. Condition is actually a slice inside PolicyRule
//https://help.zscaler.com/zpa/api-reference#/policy-set-controller/addRuleToPolicySet
func expandCreatePolicyTimeoutRule(d *schema.ResourceData) policysetrule.PolicyRule {
	policySetID, ok := d.Get("policy_set_id").(string)
	if !ok {
		log.Printf("[ERROR] policy_set_id is not set\n")
	}
	log.Printf("[INFO] action_id:%v\n", d.Get("action_id"))
	return policysetrule.PolicyRule{
		Action:             d.Get("action").(string),
		ActionID:           d.Get("action_id").(string),
		CustomMsg:          d.Get("custom_msg").(string),
		Description:        d.Get("description").(string),
		ID:                 d.Get("id").(string),
		Name:               d.Get("name").(string),
		Operator:           d.Get("operator").(string),
		PolicySetID:        policySetID,
		PolicyType:         d.Get("policy_type").(string),
		Priority:           d.Get("priority").(string),
		ReauthDefaultRule:  d.Get("reauth_default_rule").(bool),
		ReauthIdleTimeout:  d.Get("reauth_idle_timeout").(string),
		ReauthTimeout:      d.Get("reauth_timeout").(string),
		RuleOrder:          d.Get("rule_order").(string),
		Conditions:         expandPolicyTimeoutConditionSet(d),
		AppServerGroups:    expandPolicyTimeoutRuleAppServerGroups(d),
		AppConnectorGroups: expandPolicyTimeoutAppConnectorGroups(d),
	}
}

func expandPolicyTimeoutRuleAppServerGroups(d *schema.ResourceData) []policysetrule.AppServerGroups {
	appServerGroupsInterface, ok := d.GetOk("app_server_groups")
	if ok {
		appServer := appServerGroupsInterface.([]interface{})
		log.Printf("[INFO] app server groups data: %+v\n", appServer)
		var appServerGroups []policysetrule.AppServerGroups
		for _, appServerGroup := range appServer {
			appServerGroup, _ := appServerGroup.(map[string]interface{})
			if appServerGroup != nil {
				appServerGroups = append(appServerGroups, policysetrule.AppServerGroups{
					ID: appServerGroup["id"].(string),
				})
			}
		}
		return appServerGroups
	}

	return []policysetrule.AppServerGroups{}
}

func expandPolicyTimeoutAppConnectorGroups(d *schema.ResourceData) []policysetrule.AppConnectorGroups {
	appConnectorGroupsInterface, ok := d.GetOk("app_connector_groups")
	if ok {
		appConnector := appConnectorGroupsInterface.([]interface{})
		log.Printf("[INFO] app connector groups data: %+v\n", appConnector)
		var appConnectorGroups []policysetrule.AppConnectorGroups
		for _, appConnectorGroup := range appConnector {
			appConnectorGroup, _ := appConnectorGroup.(map[string]interface{})
			if appConnectorGroup != nil {
				appConnectorGroups = append(appConnectorGroups, policysetrule.AppConnectorGroups{
					ID: appConnectorGroup["id"].(string),
				})
			}
		}
		return appConnectorGroups
	}

	return []policysetrule.AppConnectorGroups{}
}

func expandPolicyTimeoutConditionSet(d *schema.ResourceData) []policysetrule.Conditions {
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
					Operands: expandPolicyTimeoutOperandsList(conditionSet["operands"]),
				})
			}
		}
		return conditionSets
	}

	return []policysetrule.Conditions{}
}

func expandPolicyTimeoutOperandsList(ops interface{}) []policysetrule.Operands {
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
func flattenPolicyTimeoutConditions(conditions []policysetrule.Conditions) []interface{} {
	ruleConditions := make([]interface{}, len(conditions))
	for i, ruleConditionItems := range conditions {
		ruleConditions[i] = map[string]interface{}{
			"id":       ruleConditionItems.ID,
			"negated":  ruleConditionItems.Negated,
			"operator": ruleConditionItems.Operator,
			"operands": flattenPolicyRuleOperands(ruleConditionItems.Operands),
		}
	}

	return ruleConditions
}

func flattenPolicyTimeoutOperands(conditionOperand []policysetrule.Operands) []interface{} {
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

func flattenPolicyTimeoutServerGroups(appServerGroup []policysetrule.AppServerGroups) []interface{} {
	policyRuleServerGroups := make([]interface{}, len(appServerGroup))
	for i, serverGroup := range appServerGroup {
		policyRuleServerGroups[i] = map[string]interface{}{
			"id": serverGroup.ID,
		}
	}

	return policyRuleServerGroups
}

func flattenPolicyTimeoutAppConnectorGroups(appConnectorGroups []policysetrule.AppConnectorGroups) []interface{} {
	policyRuleAppConnectorGroups := make([]interface{}, len(appConnectorGroups))
	for i, val := range appConnectorGroups {
		policyRuleAppConnectorGroups[i] = map[string]interface{}{
			"id": val.ID,
		}
	}

	return policyRuleAppConnectorGroups
}
