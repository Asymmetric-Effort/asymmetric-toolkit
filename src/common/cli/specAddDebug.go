package cli
/*
	Specification::AddDebug() will add a common flag to a given application's specification.
	By making reusable Specification methods for this purpose, we can maintain consistency
	in our toolkit CLI across different programs, reducing the learning curve.

	This specific flag adds the debug (--debug) flag to enable verbose logging.

	Note: This can be overridden entirely or in part by the program's other specification definitions.

 */
const (
	debugArgLong = "debug"
)

func (o *Specification) AddDebug(){
	if _, ok:=o.Argument[debugArgLong]; ok {
		o.Argument[debugArgLong] = ArgumentDescriptor{
			flagDebug,
			Boolean,
			"false",
			"Enable debug logging.",
			ParserFlag,
			ExpectFlag,
		}
	}
}
