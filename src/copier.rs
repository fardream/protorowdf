use crate::{plural_name, ProtoFile, Struct};

#[derive(Debug, Clone, Default)]
pub struct RustCopyConfig {
    pub no_cargo_out: bool,
    pub file_name: String,
    pub gen_include: bool,
}

impl RustCopyConfig {
    pub fn get_include_line(&self, file: &ProtoFile) -> String {
        if !self.gen_include {
            return "".to_owned();
        }

        let mut file_name = self.file_name.clone();
        if file_name.is_empty() {
            file_name = format!("{}.rs", file.package_name);
        }

        if self.no_cargo_out {
            format!("include!(\"{}\");\n\n", file_name)
        } else {
            format!(
                "include!(concat!(env!(\"OUT_DIR\"), \"/{}\"));\n\n",
                file_name
            )
        }
    }
    pub fn write_rust_copy<T: std::io::Write>(
        &self,
        file: &ProtoFile,
        w: &mut T,
    ) -> std::io::Result<()> {
        write!(w, "{}", self.get_include_line(file))?;
        for s in file.structs.iter() {
            self.write_struct_copy(s, w)?;
        }
        Ok(())
    }

    pub fn write_struct_copy<T: std::io::Write>(
        &self,
        s: &Struct,
        w: &mut T,
    ) -> std::io::Result<()> {
        let struct_plural_name = plural_name(&s.name, &s.plural_name);
        // convert from a slice of rows to dataframe
        writeln!(
            w,
            "impl std::convert::TryFrom<&[{}]> for {} {{",
            s.name, struct_plural_name,
        )?;
        writeln!(w, "    type Error = &'static str;")?;
        writeln!(w)?;
        writeln!(
            w,
            "    fn try_from(value: &[{}]) -> Result<Self, Self::Error> {{",
            s.name
        )?;
        writeln!(w, "        let mut r = Self::default();")?;
        writeln!(w)?;
        writeln!(w, "        for v in value.iter() {{")?;
        for f in s.fields.iter() {
            writeln!(
                w,
                "            r.{}.push(v.{}{});",
                plural_name(&f.name, &f.plural_name),
                f.name,
                f.rust_need_clone()
            )?;
        }
        writeln!(w, "        }}")?;
        writeln!(w)?;
        writeln!(w, "        Ok(r)")?;
        writeln!(w, "    }}")?;
        writeln!(w, "}}")?;

        // DF specific methods for data length and validate data length
        let first_field = s.first_field();
        writeln!(w, "impl {} {{", struct_plural_name)?;

        writeln!(w, "    pub fn data_length(&self) -> usize {{")?;
        writeln!(
            w,
            "        self.{}.len()",
            plural_name(&first_field.name, &first_field.plural_name)
        )?;
        writeln!(w, "    }}")?;
        writeln!(w)?;
        writeln!(
            w,
            "    pub fn validate_length(&self) -> Result<(), &'static str> {{"
        )?;
        writeln!(w, "        let n = self.data_length();")?;
        for f in s.fields.iter() {
            let plural_name = plural_name(&f.name, &f.plural_name);
            writeln!(w, "        if self.{}.len() != n {{", plural_name)?;
            writeln!(
                w,
                "            return Err(\"{}/{} has a different length\");",
                f.name, plural_name
            )?;
            writeln!(w, "        }}")?;
        }
        writeln!(w)?;
        writeln!(w, "        Ok(())")?;
        writeln!(w, "    }}")?;
        writeln!(w, "}}")?;

        // convert df to vector of rows
        writeln!(
            w,
            "impl std::convert::TryFrom<&{}> for Vec<{}> {{",
            struct_plural_name, s.name
        )?;
        writeln!(w, "    type Error = &'static str;")?;
        writeln!(w)?;
        writeln!(
            w,
            "    fn try_from(value: &{}) -> Result<Self, Self::Error> {{",
            struct_plural_name,
        )?;
        writeln!(w, "        value.validate_length()?;")?;
        writeln!(w, "        let n = value.data_length();")?;
        writeln!(w)?;
        writeln!(w, "        let mut r = Vec::<{}>::new();", s.name)?;
        writeln!(w)?;
        writeln!(w, "        for i in 0..n {{")?;
        writeln!(w, "            r.push({} {{", s.name)?;
        for f in s.fields.iter() {
            let f_plural_name = plural_name(&f.name, &f.plural_name);
            writeln!(
                w,
                "                {}: value.{}[i]{},",
                f.name,
                f_plural_name,
                f.rust_need_clone()
            )?;
        }
        writeln!(w, "            }});")?;
        writeln!(w, "        }}")?;
        writeln!(w)?;
        writeln!(w, "        Ok(r)")?;
        writeln!(w, "    }}")?;
        writeln!(w, "}}")?;

        // we are done
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    static YAMLINPUT: &str = include_str!("../testdata/input.yml");
    static GENERATED_RS: &str = include_str!("../testdata/rusty/src/lib.rs");

    #[test]
    fn test_gen_rust_copy() {
        let c: ProtoFile = serde_yaml::from_str(YAMLINPUT).unwrap();

        let config = RustCopyConfig {
            gen_include: true,
            ..Default::default()
        };

        let mut v = Vec::new();

        config.write_rust_copy(&c, &mut v).unwrap();

        let gened = std::str::from_utf8(&v).unwrap();

        println!("{}", gened);
        assert_eq!(gened, GENERATED_RS);
    }
}
