package Reporter

import (
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/response"
)

func (o *Report) Post(response *response.Response){
	if response == nil {
		panic("nil response encountered in Reporter.Report::Post()")
	}
}