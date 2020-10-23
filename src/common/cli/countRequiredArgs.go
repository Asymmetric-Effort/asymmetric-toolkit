package cli
/*
	Enumerate the cli Specification and count the number of terms
	where the Required (bool) field is set (true).
 */
func (o *CommandLine) countRequiredArgs()(n int) {
	for _, arg := range *o.spec {
		if arg.Required {
			n++
		}
	}
	return
}
