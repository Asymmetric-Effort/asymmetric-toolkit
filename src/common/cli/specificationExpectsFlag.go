package cli

func (o *Configuration) specificationExpectsFlag(flag *string) bool {
	_, ok := (*o.spec)[*flag]
	return ok
}
