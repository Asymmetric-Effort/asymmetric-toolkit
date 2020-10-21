package reporter

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/response"
)

func (o *Report) Post(response *response.Response){
	if response == nil {
		errors.Fatal(1, "nil response encountered in Reporter.Report::Post()")
	}
}