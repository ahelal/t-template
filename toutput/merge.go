package toutput

// MergeData merge tempdata in data
func MergeData(data *[]interface{}, tempData []interface{}) *[]interface{} {

	for _, d := range tempData {
		*data = append(*data, d)
	}
	return data
}
