package facility

import "regexp"

func (f *Facility) String() string {
	re:=regexp.MustCompile(FacilityRegExPattern)
	if re.MatchString(string(*f)){
		return string(*f)
	} else {
		panic("Invalid Log facility in logging/facility Facility::Get()")
	}
}
