package utils

import "reflect"

func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		vField := vVal.Field(i)
		// 在目标结构体中查找同名字段
		bField := bVal.FieldByName(name)
		// 如果目标结构体中存在同名字段并且类型匹配，则进行赋值
		if bField.IsValid() && bField.Type() == vField.Type() {
			bField.Set(vField)
		}
	}
}
