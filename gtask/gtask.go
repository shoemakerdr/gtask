package gtask

type GtaskConfig struct {
	meta  *MetaData
	tasks []Task
}

type MetaData struct {
	name string
}

type Task struct {
	name string
	cmd  string
}
