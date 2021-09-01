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

func getSliceFromTerraformTypeList(list interface{}) []string {
	if list == nil {
		return nil
	}
	terraformList, ok := list.([]interface{})
	if !ok {
		terraformSet, ok := list.(*schema.Set)
		if ok {
			terraformList = terraformSet.List()
		} else {
			// It's not a list or set type
			return nil
		}
	}
	var newSlice []string
	for _, v := range terraformList {
		if v != nil {
			newSlice = append(newSlice, v.(string))
		}
	}
	return newSlice
}

func convertStringArrToInterface(strs []string) []interface{} {
	arr := make([]interface{}, len(strs))
	for i, str := range strs {
		arr[i] = str
	}
	return arr
}

// Takes the result of schema.Set of strings and returns a []*int64
func expandInt64Set(configured *schema.Set) []*int64 {
	return expandInt64List(configured.List())
}

// Takes the result of flatmap.Expand for an array of int64
// and returns a []*int64
func expandInt64List(configured []interface{}) []*int64 {
	vs := make([]*int64, 0, len(configured))
	for _, v := range configured {
		vs = append(vs, Int64(int64(v.(int))))
	}
	return vs
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
