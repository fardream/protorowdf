include!(concat!(env!("OUT_DIR"), "/pkg.name.rs"));

impl std::convert::TryFrom<&[ARowStruct]> for ARowStructs {
    type Error = &'static str;

    fn try_from(value: &[ARowStruct]) -> Result<Self, Self::Error> {
        let mut r = Self::default();

        for v in value.iter() {
            r.densities.push(v.density);
            r.string_fields.push(v.string_field.clone());
            r.t_uint64s.push(v.t_uint64);
            r.t_uint32s.push(v.t_uint32);
            r.t_int32s.push(v.t_int32);
            r.t_bytess.push(v.t_bytes.clone());
            r.t_doubles.push(v.t_double);
            r.t_sfixed64s.push(v.t_sfixed64);
        }

        Ok(r)
    }
}
impl ARowStructs {
    pub fn data_length(&self) -> usize {
        self.densities.len()
    }

    pub fn validate_length(&self) -> Result<(), &'static str> {
        let n = self.data_length();
        if self.densities.len() != n {
            return Err("density/densities has a different length");
        }
        if self.string_fields.len() != n {
            return Err("string_field/string_fields has a different length");
        }
        if self.t_uint64s.len() != n {
            return Err("t_uint64/t_uint64s has a different length");
        }
        if self.t_uint32s.len() != n {
            return Err("t_uint32/t_uint32s has a different length");
        }
        if self.t_int32s.len() != n {
            return Err("t_int32/t_int32s has a different length");
        }
        if self.t_bytess.len() != n {
            return Err("t_bytes/t_bytess has a different length");
        }
        if self.t_doubles.len() != n {
            return Err("t_double/t_doubles has a different length");
        }
        if self.t_sfixed64s.len() != n {
            return Err("t_sfixed64/t_sfixed64s has a different length");
        }

        Ok(())
    }
}
impl std::convert::TryFrom<&ARowStructs> for Vec<ARowStruct> {
    type Error = &'static str;

    fn try_from(value: &ARowStructs) -> Result<Self, Self::Error> {
        value.validate_length()?;
        let n = value.data_length();

        let mut r = Vec::<ARowStruct>::new();

        for i in 0..n {
            r.push(ARowStruct {
                density: value.densities[i],
                string_field: value.string_fields[i].clone(),
                t_uint64: value.t_uint64s[i],
                t_uint32: value.t_uint32s[i],
                t_int32: value.t_int32s[i],
                t_bytes: value.t_bytess[i].clone(),
                t_double: value.t_doubles[i],
                t_sfixed64: value.t_sfixed64s[i],
            });
        }

        Ok(r)
    }
}
impl std::convert::TryFrom<&[AnotherRowStruct]> for AnotherRowStructs {
    type Error = &'static str;

    fn try_from(value: &[AnotherRowStruct]) -> Result<Self, Self::Error> {
        let mut r = Self::default();

        for v in value.iter() {
            r.adatas.push(v.adata.clone());
            r.anotherdatas.push(v.anotherdata);
        }

        Ok(r)
    }
}
impl AnotherRowStructs {
    pub fn data_length(&self) -> usize {
        self.adatas.len()
    }

    pub fn validate_length(&self) -> Result<(), &'static str> {
        let n = self.data_length();
        if self.adatas.len() != n {
            return Err("adata/adatas has a different length");
        }
        if self.anotherdatas.len() != n {
            return Err("anotherdata/anotherdatas has a different length");
        }

        Ok(())
    }
}
impl std::convert::TryFrom<&AnotherRowStructs> for Vec<AnotherRowStruct> {
    type Error = &'static str;

    fn try_from(value: &AnotherRowStructs) -> Result<Self, Self::Error> {
        value.validate_length()?;
        let n = value.data_length();

        let mut r = Vec::<AnotherRowStruct>::new();

        for i in 0..n {
            r.push(AnotherRowStruct {
                adata: value.adatas[i].clone(),
                anotherdata: value.anotherdatas[i],
            });
        }

        Ok(r)
    }
}
