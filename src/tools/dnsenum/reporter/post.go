package Reporter

import (
	"asymmetric-effort/asymmetric-toolkit/tools/dnsenum/response"
	"fmt"
)

func (o *Report) Post(response *response.Response){
	if response == nil {
		fmt.Errorf("nil response encountered in Reporter.Report::Post()")
	}
}