package models

type DataStruct struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	Flash     string
	Warning   string
	Error     string
	CSRFToken string
}
