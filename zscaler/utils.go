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

func convertStringArrToInterface(strs []string) []interface{} {
	arr := make([]interface{}, len(strs))
	for i, str := range strs {
		arr[i] = str
	}
	return arr
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

// Int64Slice converts a slice of int64 values into a slice of
// int64 pointers
func Int64Slice(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}
