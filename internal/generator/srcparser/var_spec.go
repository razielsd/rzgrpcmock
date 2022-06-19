package srcparser

type VarSpec struct {
	Name      string
	Package   string
	IsPointer bool
}

func NewVarSpec(pkg, name string, isPointer bool) *VarSpec {
	return &VarSpec{
		Name:      name,
		Package:   pkg,
		IsPointer: isPointer,
	}
}

func (v VarSpec) Header() string {
	s := ""
	if v.IsPointer {
		s = "*"
	}
	if v.Package != "" {
		s += v.Package + "."
	}
	s += v.Name
	return s
}

func (v VarSpec) FullName() string {
	s := ""
	if v.Package != "" {
		s += v.Package + "."
	}
	s += v.Name
	return s
}
