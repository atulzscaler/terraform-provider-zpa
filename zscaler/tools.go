package zscaler

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTypeSetToStringSlice(s *schema.Set) []string {
	valuesList := s.List()
	values := make([]string, len(valuesList))
	for i := 0; i < len(valuesList); i++ {
		values[i] = fmt.Sprint(valuesList[i])
	}

	return values
}

func ListToStringSlice(v []interface{}) []string {
	if len(v) == 0 {
		return []string{}
	}

	ans := make([]string, len(v))
	for i := range v {
		switch x := v[i].(type) {
		case nil:
			ans[i] = ""
		case string:
			ans[i] = x
		}
	}

	return ans
}

// getStringList will convert a TypeList attribute to a slice of string
func getStringList(d *schema.ResourceData, k string) []string {
	var sl []string
	for _, v := range d.Get(k).([]interface{}) {
		sl = append(sl, v.(string))
	}
	return sl
}

// getStringSet will convert a TypeSet attribute to a slice of string
func getStringSet(d *schema.ResourceData, k string) []string {
	var sl []string
	for _, v := range d.Get(k).(*schema.Set).List() {
		sl = append(sl, v.(string))
	}
	return sl
}

// intInSlice checks if the needle is in the haystack
func intInSlice(needle int, haystack []int) bool {
	for _, element := range haystack {
		if element == needle {
			return true
		}
	}
	return false
}

func stringInSlice(needle string, haystack []string) bool {
	for _, element := range haystack {
		if element == needle {
			return true
		}
	}
	return false
}

func ResourceDataInterfaceMap(d *schema.ResourceData, key string) map[string]interface{} {
	if _, ok := d.GetOk(key); ok {
		if v1, ok := d.Get(key).([]interface{}); ok && len(v1) != 0 {
			if v2, ok := v1[0].(map[string]interface{}); ok && v2 != nil {
				return v2
			}
		}
	}

	return map[string]interface{}{}
}

func ToInterfaceMap(m map[string]interface{}, k string) map[string]interface{} {
	if _, ok := m[k]; ok {
		if v1, ok := m[k].([]interface{}); ok && len(v1) != 0 {
			if v2, ok := v1[0].(map[string]interface{}); ok && v2 != nil {
				return v2
			}
		}
	}

	return map[string]interface{}{}
}
