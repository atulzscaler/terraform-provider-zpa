package zscaler

import (
	"log"
	"strconv"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetglobal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePolicySetGlobal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicySetGlobalRead,
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
				Type:     schema.TypeString,
				Optional: true,
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
			"policytype": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actionid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bypassdefaultrule": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"creationtime": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"custommsg": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"isolationdefaultrule": {
							Type:     schema.TypeBool,
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
						"operator": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policysetid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"policytype": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reauthdefaultrule": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"reauthidletimeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"reauthtimeout": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ruleorder": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"zpncbiprofileid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"zpninspectionprofileid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"zpninspectionprofilename": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"conditions": {
							Type:     schema.TypeList,
							Computed: true,
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
										Computed: true,
									},
									"operands": {
										Type:     schema.TypeList,
										Computed: true,
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
													Type:     schema.TypeString,
													Computed: true,
												},
												"rhs": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"operator": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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

func dataSourcePolicySetGlobalRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for global policy set %s\n", id)

	resp, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}

	log.Printf("[INFO] Getting Policy Set Global Rules:\n%+v\n", resp)
	d.SetId(strconv.FormatInt(int64(resp.ID), 10))
	_ = d.Set("creationtime", resp.CreationTime)
	_ = d.Set("description", resp.Description)
	_ = d.Set("enabled", resp.Enabled)
	_ = d.Set("modifiedby", resp.ModifiedBy)
	_ = d.Set("modifiedtime", resp.ModifiedTime)
	_ = d.Set("name", resp.Name)
	_ = d.Set("policytype", resp.PolicyType)

	if err := d.Set("rules", flattenPolicySetRules(resp)); err != nil {
		return err
	}

	return nil
}

func flattenPolicySetRules(policySetRules *policysetglobal.PolicySet) []interface{} {
	ruleItems := make([]interface{}, len(policySetRules.Rules))
	for i, ruleItem := range policySetRules.Rules {
		ruleItems[i] = map[string]interface{}{
			"action":                   ruleItem.Action,
			"actionid":                 ruleItem.ActionID,
			"creationtime":             ruleItem.CreationTime,
			"custommsg":                ruleItem.CustomMsg,
			"description":              ruleItem.Description,
			"id":                       ruleItem.ID,
			"isolationdefaultrule":     ruleItem.IsolationDefaultRule,
			"modifiedby":               ruleItem.ModifiedBy,
			"modifiedtime":             ruleItem.ModifiedTime,
			"operator":                 ruleItem.Operator,
			"policysetid":              ruleItem.PolicySetID,
			"policytype":               ruleItem.PolicyType,
			"priority":                 ruleItem.Priority,
			"reauthdefaultrule":        ruleItem.ReauthDefaultRule,
			"reauthidletimeout":        ruleItem.ReauthIdleTimeout,
			"reauthtimeout":            ruleItem.ReauthTimeout,
			"ruleorder":                ruleItem.RuleOrder,
			"zpncbiprofileid":          ruleItem.ZpnCbiProfileID,
			"zpninspectionprofileid":   ruleItem.ZpnInspectionProfileId,
			"zpninspectionprofilename": ruleItem.ZpnInspectionProfileName,
			"conditions":               flattenRuleConditions(ruleItem),
		}
	}

	return ruleItems
}

func flattenRuleConditions(conditions policysetglobal.Rules) []interface{} {
	ruleConditions := make([]interface{}, len(*conditions.Conditions))
	for i, ruleCondition := range *conditions.Conditions {
		ruleConditions[i] = map[string]interface{}{
			"creationtime": ruleCondition.CreationTime,
			"id":           ruleCondition.ID,
			"modifiedby":   ruleCondition.ModifiedBy,
			"modifiedtime": ruleCondition.ModifiedTime,
			"negated":      ruleCondition.Negated,
			"operands":     flattenConditionOperands(ruleCondition),
			// Needs to figure it out how to deal with this parameter. Returning the following error:
			//  Error: Invalid address to set: []string{"rules", "0", "conditions", "0", "operator"}
			// Works fine when removed.
			//"operator":     ruleCondition.Operator,
		}
	}

	return ruleConditions
}

func flattenConditionOperands(operands policysetglobal.Conditions) []interface{} {
	conditionOperands := make([]interface{}, len(*operands.Operands))
	for i, conditionOperand := range *operands.Operands {
		conditionOperands[i] = map[string]interface{}{
			"creationtime": conditionOperand.CreationTime,
			"id":           conditionOperand.ID,
			"idpid":        conditionOperand.IdpID,
			"lhs":          conditionOperand.LHS,
			"modifiedby":   conditionOperand.ModifiedBy,
			"modifiedtime": conditionOperand.ModifiedTime,
			"name":         conditionOperand.Name,
			"objecttype":   conditionOperand.ObjectType,
			"rhs":          conditionOperand.RHS,
		}
	}

	return conditionOperands
}
