package protorowdf

type hasPluralName interface {
	GetName() string
	GetPluralName() string
}

func getPluralName[T hasPluralName](a T) string {
	pn := a.GetPluralName()
	if pn == "" {
		return a.GetName() + "s"
	}

	return pn
}
