package cli

func (o *CommandLine) specificationExpectsFlag(flag *string) bool {
	_, ok := (*o.spec)[*flag]
	return ok
}
