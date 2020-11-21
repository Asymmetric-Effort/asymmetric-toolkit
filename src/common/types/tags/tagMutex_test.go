package tags

import "testing"

func TestTagMutex(t *testing.T){
	mutex.Lock()
	defer mutex.Unlock()
}