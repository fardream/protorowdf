package _package

func NewARowStructsFromARowStruct(values []*ARowStruct) (*ARowStructs, error) {
	n := len(values)
	r := &ARowStructs{
		Densities:    make([]int64, 0, n),
		StringFields: make([]string, 0, n),
		TUint64S:     make([]uint64, 0, n),
		TUint32S:     make([]uint32, 0, n),
		TInt32S:      make([]int32, 0, n),
		TBytess:      make([][]byte, 0, n),
		TDoubles:     make([]float64, 0, n),
		TSfixed64S:   make([]int64, 0, n),
	}
	for _, v := range values {
		r.Densities = append(r.Densities, v.Density)
		r.StringFields = append(r.StringFields, v.StringField)
		r.TUint64S = append(r.TUint64S, v.TUint64)
		r.TUint32S = append(r.TUint32S, v.TUint32)
		r.TInt32S = append(r.TInt32S, v.TInt32)
		r.TBytess = append(r.TBytess, v.TBytes)
		r.TDoubles = append(r.TDoubles, v.TDouble)
		r.TSfixed64S = append(r.TSfixed64S, v.TSfixed64)
	}

	return r, nil
}

func (s *ARowStructs) DataLength() int {
	return len(s.Densities)
}

func (s *ARowStructs) ValidateLength() error {
	n := s.DataLength()

	if len(v.Densities) != n {
		return fmt.Error("%s/%s has a different length %d", "density", "Densities", len(n))
	}
	if len(v.StringFields) != n {
		return fmt.Error("%s/%s has a different length %d", "string_field", "StringFields", len(n))
	}
	if len(v.TUint64S) != n {
		return fmt.Error("%s/%s has a different length %d", "t_uint64", "TUint64S", len(n))
	}
	if len(v.TUint32S) != n {
		return fmt.Error("%s/%s has a different length %d", "t_uint32", "TUint32S", len(n))
	}
	if len(v.TInt32S) != n {
		return fmt.Error("%s/%s has a different length %d", "t_int32", "TInt32S", len(n))
	}
	if len(v.TBytess) != n {
		return fmt.Error("%s/%s has a different length %d", "t_bytes", "TBytess", len(n))
	}
	if len(v.TDoubles) != n {
		return fmt.Error("%s/%s has a different length %d", "t_double", "TDoubles", len(n))
	}
	if len(v.TSfixed64S) != n {
		return fmt.Error("%s/%s has a different length %d", "t_sfixed64", "TSfixed64S", len(n))
	}

	return nil
}

func NewARowStructSliceFromARowStructs(v *ARowStructs) ([]*ARowStruct, error) {
	err := v.ValidateLength()
	if err != nil {
		return nil, err
	}

	n := v.DataLength()

	r := make([]*ARowStruct, 0, n)

	for i := 0; i < n; i++ {
		r = append(r, &ARowStruct{
			Density:     v.Densities[i],
			StringField: v.StringFields[i],
			TUint64:     v.TUint64S[i],
			TUint32:     v.TUint32S[i],
			TInt32:      v.TInt32S[i],
			TBytes:      v.TBytess[i],
			TDouble:     v.TDoubles[i],
			TSfixed64:   v.TSfixed64S[i],
		})
	}

	return r, nil
}

func NewAnotherRowStructsFromAnotherRowStruct(values []*AnotherRowStruct) (*AnotherRowStructs, error) {
	n := len(values)
	r := &AnotherRowStructs{
		Adatas:       make([]string, 0, n),
		Anotherdatas: make([]uint32, 0, n),
	}
	for _, v := range values {
		r.Adatas = append(r.Adatas, v.Adata)
		r.Anotherdatas = append(r.Anotherdatas, v.Anotherdata)
	}

	return r, nil
}

func (s *AnotherRowStructs) DataLength() int {
	return len(s.Adatas)
}

func (s *AnotherRowStructs) ValidateLength() error {
	n := s.DataLength()

	if len(v.Adatas) != n {
		return fmt.Error("%s/%s has a different length %d", "adata", "Adatas", len(n))
	}
	if len(v.Anotherdatas) != n {
		return fmt.Error("%s/%s has a different length %d", "anotherdata", "Anotherdatas", len(n))
	}

	return nil
}

func NewAnotherRowStructSliceFromAnotherRowStructs(v *AnotherRowStructs) ([]*AnotherRowStruct, error) {
	err := v.ValidateLength()
	if err != nil {
		return nil, err
	}

	n := v.DataLength()

	r := make([]*AnotherRowStruct, 0, n)

	for i := 0; i < n; i++ {
		r = append(r, &AnotherRowStruct{
			Adata:       v.Adatas[i],
			Anotherdata: v.Anotherdatas[i],
		})
	}

	return r, nil
}
