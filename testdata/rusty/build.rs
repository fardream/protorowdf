use std::io::Result;

fn main() -> Result<()> {
    prost_build::Config::new().compile_protos(&["../output.proto"], &[".."])?;

    Ok(())
}
