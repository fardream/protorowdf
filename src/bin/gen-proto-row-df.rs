use clap::Parser;
use protorowdf::ProtoFile;
use std::fs::File;

#[derive(Parser, Debug)]
struct Args {
    #[arg(short, long)]
    input_file: String,
    #[arg(short, long)]
    output_file: String,
}

fn main() {
    let args = Args::parse();

    let filecontent = std::fs::read_to_string(args.input_file).unwrap();

    let mut config: ProtoFile = serde_yaml::from_str(&filecontent).unwrap();

    let mut o = File::create(args.output_file).unwrap();

    config.write_proto_file_definition(&mut o).unwrap();
}
