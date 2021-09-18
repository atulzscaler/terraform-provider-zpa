package zscaler

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/client"
	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler/policysetrule"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type listrules struct {
	orders map[string]int
	sync.Mutex
}

var rules = listrules{
	orders: make(map[string]int),
}

func resourcePolicySetRule() *schema.Resource {
	return &schema.Resource{
		Create:   resourcePolicySetCreate,
		Read:     resourcePolicySetRead,
		Update:   resourcePolicySetUpdate,
		Delete:   resourcePolicySetDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "  This is for providing the rule action.",
				ValidateFunc: validation.StringInSlice([]string{
					"ALLOW",
					"DENY",
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
				Optional: true,
			},
			"reauth_timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_order": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"app_server_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "List of the server group IDs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"app_connector_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "List of app-connector IDs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
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
										Type:        schema.TypeString,
										Optional:    true,
										Description: "This signifies the key for the object type. String ID example: id ",
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
											"USER",
											"USER_GROUP",
											"LOCATION",
											"APP",
											"APP_GROUP",
											"SAML",
											"POSTURE",
											"CLIENT_TYPE",
											"IDP",
											"TRUSTED_NETWORK",
											"EDGE_CONNECTOR_GROUP",
											"MACHINE_GRP",
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

func resourcePolicySetCreate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	req := expandCreatePolicyRule(d)
	log.Printf("[INFO] Creating zpa policy rule with request\n%+v\n", req)
	if validateConditions(req.Conditions, zClient) {
		policysetrule, _, err := zClient.policysetrule.Create(&req)
		if err != nil {
			return err
		}
		d.SetId(policysetrule.ID)
		order, ok := d.GetOk("rule_order")
		if ok {
			reorder(order, policysetrule.PolicySetID, policysetrule.ID, zClient)
		}
		return resourcePolicySetRead(d, m)
	} else {
		return fmt.Errorf("couldn't validate the operands, please make sure you are using valid inputs for APP type, LHS & RHS")
	}
}

func resourcePolicySetRead(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)

	globalPolicySet, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}
	log.Printf("[INFO] Getting Policy Set Rule: globalPolicySet:%s id: %s\n", globalPolicySet.ID, d.Id())
	resp, _, err := zClient.policysetrule.Get(globalPolicySet.ID, d.Id())
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
	_ = d.Set("conditions", flattenPolicyRuleConditions(resp.Conditions))
	_ = d.Set("app_server_groups", flattenPolicyRuleServerGroups(resp.AppServerGroups))
	_ = d.Set("app_connector_groups", flattenPolicyRuleAppConnectorGroups(resp.AppConnectorGroups))

	return nil
}

func resourcePolicySetUpdate(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicySet, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}
	ruleID := d.Id()
	log.Printf("[INFO] Updating policy rule ID: %v\n", ruleID)
	req := expandCreatePolicyRule(d)

	if _, err := zClient.policysetrule.Update(globalPolicySet.ID, ruleID, &req); err != nil {
		return err
	}
	if d.HasChange("rule_order") {
		order, ok := d.GetOk("rule_order")
		if ok {
			reorder(order, globalPolicySet.ID, ruleID, zClient)
		}
	}
	return resourcePolicySetRead(d, m)
}

func validateConditions(conditions []policysetrule.Conditions, zClient *Client) bool {
	for _, condition := range conditions {
		if !validateOperands(condition.Operands, zClient) {
			return false
		}
	}
	return true
}
func validateOperands(operands []policysetrule.Operands, zClient *Client) bool {
	for _, operand := range operands {
		if !validateOperand(operand, zClient) {
			return false
		}
	}
	return true
}
func validateOperand(operand policysetrule.Operands, zClient *Client) bool {
	switch operand.ObjectType {
	case "APP":
		return customValidate(operand, []string{"id"}, "application segment ID", Getter(func(id string) error {
			_, _, err := zClient.applicationsegment.Get(id)
			return err
		}))
	case "APP_GROUP":
		return customValidate(operand, []string{"id"}, "Segment Group ID", Getter(func(id string) error {
			_, _, err := zClient.segmentgroup.Get(id)
			return err
		}))

	case "IDP":
		return customValidate(operand, []string{"id"}, "IDP ID", Getter(func(id string) error {
			_, _, err := zClient.idpcontroller.Get(id)
			return err
		}))
	case "CLOUD_CONNECTOR_GROUP":
		return customValidate(operand, []string{"id"}, "cloud connector group ID", Getter(func(id string) error {
			_, _, err := zClient.cloudconnectorgroup.Get(id)
			return err
		}))
	case "CLIENT_TYPE":
		return customValidate(operand, []string{"id"}, "'zpn_client_type_zapp' or 'zpn_client_type_exporter'", Getter(func(id string) error {
			if id != "zpn_client_type_zapp" && id != "zpn_client_type_exporter" {
				return fmt.Errorf("RHS values must be 'zpn_client_type_zapp' or 'zpn_client_type_exporter' wehn object type is CLIENT_TYPE")
			}
			return nil
		}))
	case "MACHINE_GRP":
		return customValidate(operand, []string{"id"}, "machine group ID", Getter(func(id string) error {
			_, _, err := zClient.machinegroup.Get(id)
			return err
		}))
	case "POSTURE":
		if operand.LHS == "" {
			lhsWarn(operand.ObjectType, "valid posture network ID", operand.LHS, nil)
			return false
		}
		_, _, err := zClient.postureprofile.GetByPostureUDID(operand.LHS)
		if err != nil {
			lhsWarn(operand.ObjectType, "valid posture network ID", operand.LHS, err)
			return false
		}
		if !contains([]string{"true", "false"}, operand.RHS) {
			rhsWarn(operand.ObjectType, "\"true\"/\"false\"", operand.RHS, nil)
			return false
		}
		return true
	case "TRUSTED_NETWORK":
		if operand.LHS == "" {
			lhsWarn(operand.ObjectType, "valid trusted network ID", operand.LHS, nil)
			return false
		}
		_, _, err := zClient.trustednetwork.GetByNetID(operand.LHS)
		if err != nil {
			lhsWarn(operand.ObjectType, "valid trusted network ID", operand.LHS, err)
			return false
		}
		if operand.RHS != "true" {
			rhsWarn(operand.ObjectType, "\"true\"", operand.RHS, nil)
			return false
		}
		return true
	case "SAML":
		if operand.LHS == "" {
			lhsWarn(operand.ObjectType, "valid SAML Attribute ID", operand.LHS, nil)
			return false
		}
		_, _, err := zClient.samlattribute.Get(operand.LHS)
		if err != nil {
			lhsWarn(operand.ObjectType, "valid SAML Attribute ID", operand.LHS, err)
			return false
		}
		if operand.RHS == "" {
			rhsWarn(operand.ObjectType, "SAML Attribute Value", operand.RHS, nil)
			return false
		}
		return true
	case "SCIM":
		if operand.LHS == "" {
			lhsWarn(operand.ObjectType, "valid SCIM Attribute ID", operand.LHS, nil)
			return false
		}
		_, _, err := zClient.scimattributeheader.Get(operand.LHS)
		if err != nil {
			lhsWarn(operand.ObjectType, "valid SCIM Attribute ID", operand.LHS, err)
			return false
		}
		if operand.RHS == "" {
			rhsWarn(operand.ObjectType, "SCIM Attribute Value", operand.RHS, nil)
			return false
		}
		return true
	case "SCIM_GROUP":
		if operand.LHS == "" {
			lhsWarn(operand.ObjectType, "valid IDP Controller ID", operand.LHS, nil)
			return false
		}
		_, _, err := zClient.idpcontroller.Get(operand.LHS)
		if err != nil {
			lhsWarn(operand.ObjectType, "valid IDP Controller ID", operand.LHS, err)
			return false
		}
		if operand.RHS == "" {
			rhsWarn(operand.ObjectType, "SCIM Group ID", operand.RHS, nil)
			return false
		}
		_, _, err = zClient.scimgroup.Get(operand.RHS)
		if err != nil {
			rhsWarn(operand.ObjectType, "SCIM Group ID", operand.RHS, err)
			return false
		}
		return true
	default:
		log.Printf("[WARN] invalid operand object type %s\n", operand.ObjectType)
		return false
	}
}

func contains(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

type Getter func(id string) error

func (g Getter) Get(id string) error {
	return g(id)
}
func customValidate(operand policysetrule.Operands, expectedLHS []string, expectedRHS string, clientRHS Getter) bool {
	if operand.LHS == "" || !contains(expectedLHS, operand.LHS) {
		lhsWarn(operand.ObjectType, expectedLHS, operand.LHS, nil)
		return false
	}
	if operand.RHS == "" {
		rhsWarn(operand.ObjectType, expectedRHS, operand.RHS, nil)
		return false
	}
	err := clientRHS.Get(operand.RHS)
	if err != nil {
		rhsWarn(operand.ObjectType, expectedRHS, operand.RHS, err)
		return false
	}
	return true
}
func rhsWarn(objType, expected, rhs interface{}, err error) {
	log.Printf("[WARN] when operand object type is %v RHS must be %#v, value is \"%v\", %v\n", objType, expected, rhs, err)
}
func lhsWarn(objType, expected, lhs interface{}, err error) {
	log.Printf("[WARN] when operand object type is %v LHS must be %#v value is \"%v\", %v\n", objType, expected, lhs, err)
}
func resourcePolicySetDelete(d *schema.ResourceData, m interface{}) error {
	zClient := m.(*Client)
	globalPolicySet, _, err := zClient.policysetglobal.Get()
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting IP list with id %v\n", d.Id())

	if _, err := zClient.policysetrule.Delete(globalPolicySet.ID, d.Id()); err != nil {
		return err
	}

	return nil

}

func reorder(orderI interface{}, policySetID, id string, zClient *Client) {
	defer reorderAll(policySetID, zClient)
	if orderI == nil {
		log.Printf("[WARN] Invalid order for policy set %s: %v\n", id, orderI)
		return
	}
	order, ok := orderI.(string)
	if !ok || order == "" {
		log.Printf("[WARN] Invalid order for policy set %s: %v\n", id, order)
		return
	}
	orderInt, err := strconv.Atoi(order)
	if err != nil || orderInt < 0 {
		log.Printf("[ERROR] couldn't reorder the policy set, the order may not have taken place:%v %v\n", orderInt, err)
		return
	}
	rules.Lock()
	rules.orders[id] = orderInt
	rules.Unlock()
}

// we keep calling reordering endpoint to reorder all rules after new rule was added
// because the reorder endpoint shifts all order up to replac the new order.
func reorderAll(policySetID string, zClient *Client) {
	rules.Lock()
	defer rules.Unlock()
	count, _, _ := zClient.policysetglobal.RulesCount()
	for k, v := range rules.orders {
		if v <= count {
			_, err := zClient.policysetrule.Reorder(policySetID, k, v)
			if err != nil {
				log.Printf("[ERROR] couldn't reorder the policy set, the order may not have taken place: %v\n", err)
			}
		}
	}
}
func expandCreatePolicyRule(d *schema.ResourceData) policysetrule.PolicyRule {
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
		Conditions:         expandConditionSet(d),
		AppServerGroups:    expandPolicySetRuleAppServerGroups(d),
		AppConnectorGroups: expandPolicySetRuleAppConnectorGroups(d),
	}
}

func expandPolicySetRuleAppServerGroups(d *schema.ResourceData) []policysetrule.AppServerGroups {
	appServerGroupsInterface, ok := d.GetOk("app_server_groups")
	if ok {
		appServer := appServerGroupsInterface.(*schema.Set)
		log.Printf("[INFO] app server groups data: %+v\n", appServer)
		var appServerGroups []policysetrule.AppServerGroups
		for _, appServerGroup := range appServer.List() {
			appServerGroup, _ := appServerGroup.(map[string]interface{})
			if appServerGroup != nil {
				for _, id := range appServerGroup["id"].([]interface{}) {
					appServerGroups = append(appServerGroups, policysetrule.AppServerGroups{
						ID: id.(string),
					})
				}
			}
		}
		return appServerGroups
	}

	return []policysetrule.AppServerGroups{}
}

func expandPolicySetRuleAppConnectorGroups(d *schema.ResourceData) []policysetrule.AppConnectorGroups {
	appConnectorGroupsInterface, ok := d.GetOk("app_connector_groups")
	if ok {
		appConnector := appConnectorGroupsInterface.(*schema.Set)
		log.Printf("[INFO] app connector groups data: %+v\n", appConnector)
		var appConnectorGroups []policysetrule.AppConnectorGroups
		for _, appConnectorGroup := range appConnector.List() {
			appConnectorGroup, _ := appConnectorGroup.(map[string]interface{})
			if appConnectorGroup != nil {
				for _, id := range appConnectorGroup["id"].([]interface{}) {
					appConnectorGroups = append(appConnectorGroups, policysetrule.AppConnectorGroups{
						ID: id.(string),
					})
				}

			}
		}
		return appConnectorGroups
	}

	return []policysetrule.AppConnectorGroups{}
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
					ID:       conditionSet["id"].(string),
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
func flattenPolicyRuleConditions(conditions []policysetrule.Conditions) []interface{} {
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

func flattenPolicyRuleOperands(conditionOperand []policysetrule.Operands) []interface{} {
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

func flattenPolicyRuleServerGroups(appServerGroup []policysetrule.AppServerGroups) []interface{} {
	result := make([]interface{}, 1)
	mapIds := make(map[string]interface{})
	ids := make([]string, len(appServerGroup))
	for i, serverGroup := range appServerGroup {
		ids[i] = serverGroup.ID
	}
	mapIds["id"] = ids
	result[0] = mapIds
	return result
}

func flattenPolicyRuleAppConnectorGroups(appConnectorGroups []policysetrule.AppConnectorGroups) []interface{} {
	result := make([]interface{}, 1)
	mapIds := make(map[string]interface{})
	ids := make([]string, len(appConnectorGroups))
	for i, group := range appConnectorGroups {
		ids[i] = group.ID
	}
	mapIds["id"] = ids
	result[0] = mapIds
	return result
}
