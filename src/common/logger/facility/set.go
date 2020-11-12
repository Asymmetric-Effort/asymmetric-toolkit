package logFacility

import (
	"regexp"
)

func (f *Facility) Set(v string) {
	re:=regexp.MustCompile(facilityRegExPattern)
	if re.MatchString(v){
		*f = Facility(v)
	} else {
		panic("Invalid Log facility in logging/facility Facility::Set()")
	}
}
