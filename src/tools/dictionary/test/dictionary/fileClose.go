package dictionary

import "asymmetric-effort/asymmetric-toolkit/src/common/file"

func (o *File) Close() {
	file.CloseFile(o.Handle)
}
