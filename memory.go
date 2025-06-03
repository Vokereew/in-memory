package memory

type Memory struct {
	Data *map[string]interface{}
}

func NewMemory() *Memory {
	data := make(map[string]interface{})
	return &Memory{
		Data: &data,
	}
}

func (m *Memory) Set(key string, value interface{}) *Memory {
	(*m.Data)[key] = value
	return m
}

func (m *Memory) Get(key string) interface{} {
	if value, exist := (*m.Data)[key]; exist {
		return value
	}
	return nil
}

func (m *Memory) Delete(key string) *Memory {
	if _, exist := (*m.Data)[key]; exist {
		delete(*m.Data, key)
		return m
	}
	return nil
}
