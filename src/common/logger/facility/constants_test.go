package facility_test

import (
	LogFacility "asymmetric-effort/asymmetric-toolkit/src/common/logger/facility"
	"regexp"
	"testing"
)

func TestLogFacilityConstantsFacRegEx(t *testing.T){
	_ = regexp.MustCompile(LogFacility.FacilityRegExPattern)
}
