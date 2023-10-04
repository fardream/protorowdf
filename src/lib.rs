pub mod copier;

use std::{collections::HashMap, io::Write};

use anyhow::anyhow;
use phf::phf_map;
use serde::{de::Visitor, ser::Error, Deserialize, Serialize, Serializer};

include!(concat!(env!("OUT_DIR"), "/protorowdf.rs"));

static NAME_TO_SUPPORTED_TYPE: phf::Map<&'static str, SupportedType> = phf_map! {
    "bool" => SupportedType::Bool,
    "bytes" => SupportedType::Bytes,
    "double" => SupportedType::Double,
    "fixed32" => SupportedType::Fixed32,
    "fixed64" => SupportedType::Fixed64,
    "float" => SupportedType::Float,
    "int32" => SupportedType::Int32,
    "int64" => SupportedType::Int64,
    "sfixed32" => SupportedType::Sfixed32,
    "sfixed64" => SupportedType::Sfixed64,
    "sint32" => SupportedType::Sint32,
    "sint64" => SupportedType::Sint64,
    "string" => SupportedType::String,
    "uint32" => SupportedType::Uint32,
    "uint64" => SupportedType::Uint64,
    "unknown" => SupportedType::Unknown,
};

impl SupportedType {
    pub fn proto_type_string(&self) -> String {
        match self {
            SupportedType::Bool => "bool".to_owned(),
            SupportedType::Bytes => "bytes".to_owned(),
            SupportedType::Double => "double".to_owned(),
            SupportedType::Fixed32 => "fixed32".to_owned(),
            SupportedType::Fixed64 => "fixed64".to_owned(),
            SupportedType::Float => "float".to_owned(),
            SupportedType::Int32 => "int32".to_owned(),
            SupportedType::Int64 => "int64".to_owned(),
            SupportedType::Sfixed32 => "sfixed32".to_owned(),
            SupportedType::Sfixed64 => "sfixed64".to_owned(),
            SupportedType::Sint32 => "sint32".to_owned(),
            SupportedType::Sint64 => "sint64".to_owned(),
            SupportedType::String => "string".to_owned(),
            SupportedType::Uint32 => "uint32".to_owned(),
            SupportedType::Uint64 => "uint64".to_owned(),
            SupportedType::Unknown => "unknown".to_owned(),
        }
    }

    pub fn from_proto_type_string(s: &str) -> anyhow::Result<Self> {
        let lowered = s.to_lowercase();

        if let Some(d) = NAME_TO_SUPPORTED_TYPE.get(&lowered) {
            Ok(*d)
        } else {
            Err(anyhow!("cannot parse {}", s))
        }
    }
}

fn serialize_supported_type<S: Serializer>(i: &i32, serializer: S) -> Result<S::Ok, S::Error> {
    if let Ok(t) = SupportedType::try_from(*i) {
        t.serialize(serializer)
    } else {
        Err(Error::custom("failed to convert input to string"))
    }
}

impl Serialize for SupportedType {
    fn serialize<S>(&self, serializer: S) -> Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        serializer.serialize_str(&self.proto_type_string())
    }
}

fn deserialize_supported_type<'de, D: serde::Deserializer<'de>>(
    deserializer: D,
) -> Result<i32, D::Error> {
    SupportedType::deserialize(deserializer).map(|x| x as i32)
}

impl<'de> Deserialize<'de> for SupportedType {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        struct SupportedTypeVisitor;

        impl<'de> Visitor<'de> for SupportedTypeVisitor {
            type Value = SupportedType;
            fn expecting(&self, formatter: &mut std::fmt::Formatter) -> std::fmt::Result {
                formatter.write_str("expecting protobuf prelimiary types")
            }

            fn visit_i8<E>(self, v: i8) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_i64(v as i64)
            }

            fn visit_i16<E>(self, v: i16) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_i64(v as i64)
            }

            fn visit_i32<E>(self, v: i32) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                if let Ok(s) = SupportedType::try_from(v) {
                    Ok(s)
                } else {
                    Err(serde::de::Error::invalid_type(
                        serde::de::Unexpected::Signed(v.into()),
                        &self,
                    ))
                }
            }

            fn visit_i64<E>(self, v: i64) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_i32(v as i32)
            }

            fn visit_i128<E>(self, v: i128) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_i32(v as i32)
            }

            fn visit_u8<E>(self, v: u8) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_u64(v as u64)
            }

            fn visit_u16<E>(self, v: u16) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_u64(v as u64)
            }

            fn visit_u32<E>(self, v: u32) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_u64(v as u64)
            }

            fn visit_u64<E>(self, v: u64) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                Err(serde::de::Error::invalid_type(
                    serde::de::Unexpected::Unsigned(v),
                    &self,
                ))
            }

            fn visit_u128<E>(self, v: u128) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                self.visit_u64(v as u64)
            }

            fn visit_str<E>(self, v: &str) -> Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                if let Ok(s) = SupportedType::from_proto_type_string(v) {
                    Ok(s)
                } else {
                    Err(serde::de::Error::invalid_type(
                        serde::de::Unexpected::Str(v),
                        &self,
                    ))
                }
            }
        }

        deserializer.deserialize_any(SupportedTypeVisitor {})
    }
}

impl Field {
    pub fn rust_need_clone(&self) -> &'static str {
        if let Ok(t) = SupportedType::try_from(self.data_type) {
            if t == SupportedType::String || t == SupportedType::Bytes {
                return ".clone()";
            }
        }

        ""
    }
    pub fn write_plural_filed_definition<T: Write>(&self, w: &mut T) -> std::io::Result<()> {
        for aline in pretty_comment(&self.comment).iter() {
            writeln!(w, "{}//{}", self.indent, aline)?;
        }

        let t =
            SupportedType::try_from(self.data_type).map_err(|_| std::io::ErrorKind::InvalidData)?;
        writeln!(
            w,
            "{}repeated {} {} = {};",
            self.indent,
            t.proto_type_string(),
            plural_name(&self.name, &self.plural_name),
            self.tag_num
        )?;

        Ok(())
    }

    pub fn write_field_definition<T: Write>(&self, w: &mut T) -> std::io::Result<()> {
        for aline in pretty_comment(&self.comment).iter() {
            writeln!(w, "{}//{}", self.indent, aline)?;
        }

        let t =
            SupportedType::try_from(self.data_type).map_err(|_| std::io::ErrorKind::InvalidData)?;
        writeln!(
            w,
            "{}{} {} = {};",
            self.indent,
            t.proto_type_string(),
            self.name,
            self.tag_num
        )?;

        Ok(())
    }
}

fn plural_name(name: &str, plural_name: &str) -> String {
    if plural_name.is_empty() {
        format!("{}s", name)
    } else {
        plural_name.to_owned()
    }
}

impl Struct {
    pub fn first_field(&self) -> &Field {
        &self.fields[0]
    }

    pub fn write_message_definition<T: Write>(&self, w: &mut T) -> std::io::Result<()> {
        for aline in pretty_comment(&self.comment).iter() {
            writeln!(w, "//{}", aline)?;
        }

        writeln!(w, "message {} {{", self.name)?;
        for f in self.fields.iter() {
            f.write_field_definition(w)?;
        }

        writeln!(w, "}}")?;

        writeln!(w)?;

        writeln!(w, "// Plural Version of {}.", self.name)?;
        for aline in pretty_comment(&self.comment).iter() {
            writeln!(w, "//{}", aline)?;
        }

        writeln!(
            w,
            "message {} {{",
            plural_name(&self.name, &self.plural_name)
        )?;
        for f in self.fields.iter() {
            f.write_plural_filed_definition(w)?;
        }

        writeln!(w, "}}")?;

        Ok(())
    }

    pub fn valid_tag_nums(&self) -> bool {
        let mut m = HashMap::<i32, Vec<&str>>::new();

        for f in self.fields.iter() {
            if !m.contains_key(&f.tag_num) {
                m.insert(f.tag_num, vec![&f.name]);
            } else {
                m.get_mut(&f.tag_num).unwrap().push(&f.name);
            }
        }

        for (_, v) in m.iter() {
            if v.len() > 1 {
                return false;
            }
        }

        true
    }
}

impl ProtoFile {
    pub fn write_proto_file_definition<T: Write>(&mut self, w: &mut T) -> std::io::Result<()> {
        self.update_indent();
        writeln!(w, "syntax = \"proto3\";")?;
        writeln!(w)?;

        for aline in pretty_comment(&self.comment).iter() {
            writeln!(w, "//{}", aline)?;
        }
        writeln!(w, "package {};", self.package_name)?;
        writeln!(w)?;

        for o in self.options.iter() {
            writeln!(w, "option {} = {};", o.name, o.value)?;
        }

        for s in self.structs.iter() {
            writeln!(w)?;
            s.write_message_definition(w)?;
        }

        Ok(())
    }
    pub fn valid_tag_nums(&self) -> bool {
        self.structs.iter().all(|x| x.valid_tag_nums())
    }

    pub fn update_indent(&mut self) {
        let mut indent = self.field_indent.clone();

        if indent.is_empty() {
            indent = "  ".to_owned();
        }

        for s in self.structs.iter_mut() {
            for f in s.fields.iter_mut() {
                f.indent = indent.clone();
            }
        }
    }
}

fn pretty_comment(s: &str) -> Vec<String> {
    let lines = s
        .trim()
        .split('\n')
        .map(|x| {
            if x.trim().is_empty() {
                ""
            } else {
                x.trim_end()
            }
        })
        .collect::<Vec<_>>();

    if lines.len() == 1 && lines[0].is_empty() {
        return vec![];
    }

    lines
        .iter()
        .map(|x| {
            if x.is_empty() {
                (*x).to_owned()
            } else {
                format!(" {}", *x)
            }
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    static YAMLINPUT: &str = include_str!("../testdata/input.yml");
    static GENERATED_PROTO: &str = include_str!("../testdata/output.proto");

    #[test]
    fn test_parse_yaml() {
        let _: ProtoFile = serde_yaml::from_str(YAMLINPUT).unwrap();
    }

    #[test]
    fn test_generae_proto_file() {
        let mut c: ProtoFile = serde_yaml::from_str(YAMLINPUT).unwrap();

        let mut v = Vec::new();

        c.write_proto_file_definition(&mut v).unwrap();

        let gened = std::str::from_utf8(&v).unwrap();

        assert_eq!(gened, GENERATED_PROTO);
    }
}
