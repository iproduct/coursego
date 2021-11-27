package copymap

func UpdateMap(dest map[interface{}]interface{}, src map[interface{}]interface{}) {
	for k, v := range src {
		dest[k] = v
	}
}
