package zscaler

/*
import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetrule"
)

func resourcePolicySetRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Description: "	This is for providing the rule action.",
			},
			"actionid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "This field defines the description of the server.",
			},
			"servergroups": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "ID of the server group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"appconnectorgroups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "This field is a json array of app-connector-id only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"conditions": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "This is for proviidng the set of conditions for the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"id": {
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
						"negated": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"operator": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"operands": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "This signifies the various policy criteria.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"creationtime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"idpid": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"lhs": {
										Type:     schema.TypeString,
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
									"objecttype": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Description: "	This is for specifying the policy critiera.",
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
			"creationtime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"custommsg": {
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
			"modifiedby": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modifiedtime": {
				Type:     schema.TypeInt,
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
			"policysetid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"policytype": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reauthidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reauthtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ruleorder": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"zpncbiprofileid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
		//Create: resourcePolicySetCreate,
		Read:   resourcePolicySetRead,
		Update: resourcePolicySetUpdate,
		Delete: resourcePolicySetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourcePolicySetCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreatePolicyRule(d)
	log.Printf("[INFO] Creating zpa policy rule with request\n%+v\n", req)

	// Having problems here. Cannot use req (variable of type policysetrule.PolicyRule) as string value in argument to zClient.policysetrule.Create
	policysetrule, _, err := zClient.policysetrule.Create(req)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(int64(policysetrule.ID), 10))

	return resourcePolicySetRead(d, m)
}

// Please review read operations. It needs to pull the policySetId and RuleId in order to read a specific rule.
func resourcePolicySetRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	policySetId := d.Id()
	ruleId := d.Id()
	resp, _, err := zClient.policysetrule.Get(policySetId, ruleId)
	if err != nil {
		if err.(*client.ErrorResponse).IsObjectNotFound() {
			log.Printf("[WARN] Removing policy rule %s from state because it no longer exists in ZPA", d.Id())
			d.SetId("")
			return nil
		}

		return err
	}

	log.Printf("[INFO] Getting Policy Set Global Rules:\n%+v\n", resp)
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("action", resp.Action)
	_ = d.Set("actionid", resp.ActionID)
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("custommsg", resp.CustomMsg)
	_ = d.Set("description", resp.Description)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("operator", resp.Operator)
	_ = d.Set("policysetid", resp.PolicySetID)
	_ = d.Set("policytype", resp.PolicyType)
	_ = d.Set("priority", resp.Priority)
	_ = d.Set("reauthidletimeout", resp.ReauthIdleTimeout)
	_ = d.Set("reauthtimeout", resp.ReauthTimeout)
	_ = d.Set("ruleorder", resp.RuleOrder)
	_ = d.Set("zpncbiprofileid", resp.ZpnCbiProfileID)
	_ = d.Set("conditions", flattenPolicyRuleConditions(resp.Conditions))

	return nil
}

// Please review Update operations. It needs to pull the policySetId and RuleId in order to update a specific rule.
func resourcePolicySetUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	policySetId := d.Id()
	ruleId := d.Id()
	log.Printf("[INFO] Updating policy rule ID: %v\n", ruleId)
	req := expandCreatePolicyRule(d)

	if _, err := zClient.policysetrule.Update(policySetId, ruleId, req); err != nil {
		return err
	}

	return resourcePolicySetRead(d, m)
}

// Please review Delete operations. It needs to pull the policySetId and RuleId in order to delete a specific rule.
func resourcePolicySetDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	policySetId := d.Id()
	ruleId := d.Id()

	log.Printf("[INFO] Deleting rule with id %v\n")

	if _, err := zClient.policysetrule.Delete(policySetId, ruleId); err != nil {
		return err
	}

	return nil
}

// Please review the expand and flattening functions. Condition is actually a slice inside PolicyRule
//https://help.zscaler.com/zpa/api-reference#/policy-set-controller/addRuleToPolicySet
func expandCreatePolicyRule(d *schema.ResourceData) policysetrule.PolicyRule {
	return policysetrule.PolicyRule{
		Action:            d.Get("action").(string),
		ActionID:          d.Get("actionid").(int64),
		CreationTime:      d.Get("creationtime").(int32),
		CustomMsg:         d.Get("custommsg").(string),
		Description:       d.Get("description").(string),
		ID:                d.Get("id").(int64),
		ModifiedBy:        d.Get("modifiedby").(int64),
		ModifiedTime:      d.Get("Modifiedtime").(int32),
		Name:              d.Get("name").(string),
		Operator:          d.Get("operator").(string),
		PolicySetID:       d.Get("policysetid").(int64),
		PolicyType:        d.Get("policytype").(int32),
		Priority:          d.Get("priority").(int32),
		ReauthIdleTimeout: d.Get("reauthidletimeout").(int32),
		ReauthTimeout:     d.Get("reauthtimeout").(int32),
		RuleOrder:         d.Get("ruleorder").(int32),
		ZpnCbiProfileID:   d.Get("zpncbiprofileid").(int64),
		Conditions:        expandConditionSet(d),
	}
}

func expandConditionSet(d *schema.ResourceData) []policysetrule.Conditions {
	var conditionSets []policysetrule.Conditions
	if conditionInterface, ok := d.GetOk("conditions"); ok {
		conditions := conditionInterface.([]interface{})
		conditionSets = make([]policysetrule.Conditions, len(conditions))
		for i, condition := range conditions {
			conditionSet := condition.(map[string]interface{})
			conditionSets[i] = policysetrule.Conditions{
				CreationTime: conditionSet["creationtime"].(int32),
				ID:           conditionSet["id"].(int64),
				ModifiedBy:   conditionSet["modifiedby"].(int64),
				ModifiedTime: conditionSet["modifiedtime"].(int32),
				Negated:      conditionSet["negated"].(bool),
				Operator:     conditionSet["operator"].(string),
			}
		}
	}

	return conditionSets
}

func flattenPolicyRuleConditions(conditions []policysetrule.Conditions) []interface{} {
	ruleConditions := make([]interface{}, len(conditions))
	for i, ruleConditionItems := range conditions {
		ruleConditions[i] = map[string]interface{}{
			"creationtime": ruleConditionItems.CreationTime,
			"id":           ruleConditionItems.ID,
			"modifiedby":   ruleConditionItems.ModifiedBy,
			"modifiedtime": ruleConditionItems.ModifiedTime,
			"negated":      ruleConditionItems.Negated,
			"operator":     ruleConditionItems.Operator,
			"operands":     flattenPolicyRuleOperands(ruleConditionItems.Operands),
		}
	}

	return ruleConditions
}

func flattenPolicyRuleOperands(conditionOperand []policysetrule.Operands) []interface{} {
	conditionOperands := make([]interface{}, len(conditionOperand))
	for i, operandItems := range conditionOperand {
		conditionOperands[i] = map[string]interface{}{
			"creationtime": operandItems.CreationTime,
			"id":           operandItems.ID,
			"idpid":        operandItems.IdpID,
			"lhs":          operandItems.LHS,
			"modifiedby":   operandItems.ModifiedBy,
			"modifiedtime": operandItems.ModifiedTime,
			"name":         operandItems.Name,
			"objecttype":   operandItems.ObjectType,
			"rhs":          operandItems.RHS,
		}
	}

	return conditionOperands
}

// Need to flatten the Operands menu, which is a slice inside the slice Conditions
//https://help.zscaler.com/zpa/api-reference#/policy-set-controller/addRuleToPolicySet
*/