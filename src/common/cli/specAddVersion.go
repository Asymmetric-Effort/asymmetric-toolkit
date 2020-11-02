package cli

/*
	Specification::AddVersion() will add a common flag to a given application's specification.
	By making reusable Specification methods for this purpose, we can maintain consistency
	in our toolkit CLI across different programs, reducing the learning curve.

	This specific flag adds version messaging (--version or -v) which generates usage messages based
	on a given program's specification object.

	Note: This can be overridden entirely or in part by the program's other specification definitions.
*/

const (
	versionArgLong  = "version"
	versionArgShort = "v"
)

func (o *Specification) AddVersion() {
	if _, ok := o.Argument[versionArgLong]; ok {
		o.Argument[versionArgLong] = ArgumentDescriptor{
			flagVersion,
			None,
			"",
			"Show version string",
			o.ShowVersion,
			ExpectEnd,
		}
	}
	if _, ok := o.Argument[versionArgShort]; ok {
		o.Argument[versionArgShort] = ArgumentDescriptor{
			flagVersion,
			None,
			"",
			"Show version string",
			o.ShowVersion,
			ExpectEnd,
		}
	}
}
