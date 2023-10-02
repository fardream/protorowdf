use std::io::Result;

fn main() -> Result<()> {
    prost_build::Config::new()
        .message_attribute(".protorowdf.Field", "#[derive(Serialize, Deserialize)]")
        .message_attribute(".protorowdf.Struct", "#[derive(Serialize, Deserialize)]")
        .message_attribute(".protorowdf.ProtoFile", "#[derive(Serialize, Deserialize)]")
        .message_attribute(".protorowdf.Option", "#[derive(Serialize, Deserialize)]")
        .field_attribute(
            ".protorowdf.Field.data_type",
            "#[serde(serialize_with=\"serialize_supported_type\", deserialize_with=\"deserialize_supported_type\")]",
        )
        .field_attribute(".protorowdf.Field", "#[serde(default)]")
        .field_attribute(".protorowdf.Option", "#[serde(default)]")
        .field_attribute(".protorowdf.Struct", "#[serde(default)]")
        .field_attribute(".protorowdf.ProtoFile", "#[serde(default)]")
        .compile_protos(&["config.proto"], &[""])?;

    Ok(())
}
