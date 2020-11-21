package tags

func NewFloat64() Float64{
	mutex.Lock()
	defer mutex.Unlock()

	return make(Float64,1)
}
