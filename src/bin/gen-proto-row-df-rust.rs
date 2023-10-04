use clap::Parser;
use protorowdf::{copier::RustCopyConfig, ProtoFile};
use std::fs::File;

#[derive(Parser, Debug)]
struct Args {
    #[arg(short, long)]
    input_file: String,
    #[arg(short, long)]
    output_file: String,
    #[arg(long, default_value_t = false)]
    add_include: bool,
    #[arg(long, default_value_t = false)]
    dont_use_cargo_out: bool,
    #[arg(long, default_value = "")]
    file_for_include: String,
}

fn main() {
    let args = Args::parse();

    let filecontent = std::fs::read_to_string(args.input_file).unwrap();

    let config: ProtoFile = serde_yaml::from_str(&filecontent).unwrap();

    let rust_config = RustCopyConfig {
        no_cargo_out: args.dont_use_cargo_out,
        file_name: args.file_for_include,
        gen_include: args.add_include,
    };

    let mut o = File::create(args.output_file).unwrap();

    rust_config.write_rust_copy(&config, &mut o).unwrap();
}
