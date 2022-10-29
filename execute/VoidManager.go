package execute

type VoidManager struct{}

func (v *VoidManager) BringUpMachines()     {}
func (v *VoidManager) Shutdown(chan<- bool) {}
