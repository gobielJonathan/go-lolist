package list

import (
	"reflect"
	"sort"
)

type List struct {
	data []interface{}
}

func (d *List) ToGeneric(data ...interface{}) []interface{} {
	return data
}

func (d *List) Gets() []interface{} {
	return d.data
}

func (d *List) Length() int {
	return len(d.data)
}

func (d *List) Capacity() int {
	return cap(d.data)
}

func (d *List) Add(s interface{}) {
	if d.data == nil {
		d.data = []interface{}{}
	}

	d.data = append(d.data, s)
}

func (d *List) Clear() {
	d.data = nil
}

func (d *List) RemoveAt(i int) {
	d.data = append(d.data[:i], d.data[i+1:]...)
}

func (d *List) Sort(callback func(i, j interface{}) bool) {
	sort.SliceStable(d.data, func(i, j int) bool {
		return callback(d.data[i], d.data[j])
	})
}

func (d *List) ForEach(callback func(i interface{})) {
	for _, data := range d.data {
		callback(data)
	}
}

func (d *List) Reverse() []interface{} {
	for i, j := 0, len(d.data)-1; i < j; i, j = i+1, j-1 {
		d.data[i], d.data[j] = d.data[j], d.data[i]
	}
	return d.data
}

func (d *List) First() interface{} {
	if d.data == nil {
		panic("data is empty, try access first index")
	}
	return d.data[0]
}

func (d *List) FirstOrDefault() interface{} {
	if d.data == nil {
		return nil
	}
	return d.First()
}

func (d *List) Last() interface{} {
	if d.data == nil {
		panic("data is empty, try access last index")
	}
	return d.data[d.Length()-1]
}

func (d *List) LastOrDefault() interface{} {
	if d.data == nil {
		return nil
	}
	return d.data[d.Length()-1]
}

func (d List) Exists(callback func(s interface{}) bool) bool {
	for _, data := range d.data {
		if valid := callback(data); valid {
			return true
		}
	}
	return false
}

func (d *List) Find(callback func(s interface{}) bool) interface{} {
	for _, data := range d.data {
		if valid := callback(data); valid {
			return data
		}
	}
	return nil
}
func (d List) FindIndex(callback func(s interface{}) bool) int {
	for idx, data := range d.data {
		if valid := callback(data); valid {
			return idx
		}
	}
	return -1
}

func (d List) FindLastIndex(callback func(s interface{}) bool) int {
	length := len(d.data) - 1
	for i := length; i >= 0; i-- {
		if valid := callback(d.data[i]); valid {
			return i
		}
	}
	return -1
}

func (d List) ToDictionary(getKey func(s interface{}) string, getValue func(s interface{}) interface{}) map[string]interface{} {
	var dict = map[string]interface{}{}
	for _, data := range d.data {
		key := getKey(data)
		dict[key] = getValue(data)
	}
	return dict
}

func (d List) GroupBy(getKey func(s interface{}) interface{}) map[interface{}][]interface{} {
	var dict = map[interface{}][]interface{}{}

	for _, data := range d.data {
		key := getKey(data)
		reflectValue := reflect.ValueOf(data)

		if reflectValue.Kind() == reflect.Struct {
			for i := 0; i < reflectValue.NumField(); i++ {

				if reflect.DeepEqual(reflectValue.Field(i).Interface(), key) {
					dict[key] = append(dict[key], data)
					break
				}
			}

		} else if reflectValue.Kind() == reflect.String {
			str, ok := data.(string)
			if ok {
				if str == key {
					dict[key] = append(dict[key], data)
					continue
				}
				keyInt, ok := key.(int)
				if ok {
					if len(str) == keyInt {
						dict[key] = append(dict[key], data)
						continue
					}
				}
			}
		}

	}
	return dict
}
