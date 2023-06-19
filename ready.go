package git

func (g *Git) Ready() bool {
	Services.Lock()
	defer Services.Unlock()

	if Services.readyCount < len(Services.services) {
		return false
	}

	return true
}
