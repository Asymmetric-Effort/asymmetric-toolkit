package cli

func MergeMaps(lhs *Specification, rhs *Specification){
	for k,v:=range *rhs{
		(*lhs)[k]=v
	}
}