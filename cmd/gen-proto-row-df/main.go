package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"

	"github.com/fardream/protorowdf"
)

func parseInput(data []byte, extension string) (*protorowdf.ProtoFile, error) {
	switch extension {
	case ".yaml", ".yml":
		return protorowdf.ParseYAML(data)
	case ".json":
		return protorowdf.ParseJSON(data)
	case ".txt", ".txtpb":
		r := &protorowdf.ProtoFile{}
		if err := prototext.Unmarshal(data, r); err != nil {
			return nil, err
		}
		return r, nil
	case ".binpb", ".pb":
		r := &protorowdf.ProtoFile{}
		if err := proto.Unmarshal(data, r); err != nil {
			return nil, err
		}
		return r, nil
	default:
		return nil, fmt.Errorf("unrecoganized extension %s", extension)
	}
}

func main() {
	newRootCmd().Execute()
}

type rootCmd struct {
	*cobra.Command

	input  string
	output string

	genProtoCmd *cobra.Command
	genGoCmd    *cobra.Command
	gopackage   string
	genRustCmd  *cobra.Command
	rustconfig  protorowdf.RustCopyConfig
}

func newRootCmd() *rootCmd {
	c := &rootCmd{
		Command: &cobra.Command{
			Use:   "gen-proto-row-df",
			Short: "generate proto defintions for rows of data and the corresponding data frame.",
			Long:  "generate proto defintions for rows of data and the corresponding data frame.",
			Args:  cobra.NoArgs,
		},
	}

	c.Run = c.genProto

	c.PersistentFlags().StringVarP(&c.input, "input", "i", c.input, "input file, the format of the input is decided by the extension (yaml/yml for yaml, json for json, txt for text etc.)")
	c.MarkPersistentFlagRequired("input")
	c.MarkPersistentFlagFilename("input", "yaml", "yml", "json", "txt", "txtpb", "pb", "binpb")
	c.PersistentFlags().StringVarP(&c.output, "output", "o", c.output, "output file")
	c.MarkPersistentFlagFilename("output")
	c.MarkPersistentFlagRequired("output")

	c.genProtoCmd = &cobra.Command{
		Use:   "proto",
		Short: c.Short,
		Long:  c.Long,
		Args:  cobra.NoArgs,
		Run:   c.genProto,
	}
	c.genGoCmd = &cobra.Command{
		Use:   "go",
		Short: "generate copy data code for go",
		Long:  "generate copy data code for go",
		Args:  cobra.NoArgs,
		Run:   c.genGo,
	}

	c.genGoCmd.Flags().StringVarP(&c.gopackage, "package", "p", c.gopackage, "go package name, leave empty to use the last part of the package name in proto package.")

	c.genRustCmd = &cobra.Command{
		Use:   "rust",
		Short: "generate copy data code for rust",
		Long:  "generate copy data code for rust",
		Args:  cobra.NoArgs,
		Run:   c.genRust,
	}

	c.genRustCmd.Flags().BoolVar(&c.rustconfig.GenInclude, "add-include", c.rustconfig.GenInclude, "add rust include line for generated proto file")
	c.genRustCmd.Flags().BoolVar(&c.rustconfig.NoCargoOut, "dont-use-cargo-out", c.rustconfig.NoCargoOut, "dont use cargo out dir for include")
	c.genRustCmd.Flags().StringVar(&c.rustconfig.FileName, "file-for-include", c.rustconfig.FileName, "file name for include, default will use the package name + .rs")

	c.AddCommand(c.genProtoCmd, c.genGoCmd, c.genRustCmd)

	return c
}

func (c *rootCmd) genProto(*cobra.Command, []string) {
	data, err := os.ReadFile(c.input)
	if err != nil {
		log.Panic(err)
	}
	input, err := parseInput(data, filepath.Ext(c.input))
	if err != nil {
		log.Panic(err)
	}

	content, err := input.ProtoFileDefinition()
	if err != nil {
		log.Panic(err)
	}

	if c.output == "-" {
		if _, err := os.Stdout.Write([]byte(content)); err != nil {
			log.Panic(err)
		}
	} else {
		if err := os.WriteFile(c.output, []byte(content), 0o666); err != nil {
			log.Panic(err)
		}
	}
}

func (c *rootCmd) genGo(*cobra.Command, []string) {
	data, err := os.ReadFile(c.input)
	if err != nil {
		log.Panic(err)
	}
	g := &protorowdf.GoProtoFile{ManualGoPackage: c.gopackage}
	g.ProtoFile, err = parseInput(data, filepath.Ext(c.input))
	if err != nil {
		log.Panic(err)
	}

	content, err := g.CopyDataCode()
	if err != nil {
		log.Panic(err)
	}

	if c.output == "-" {
		if _, err := os.Stdout.Write([]byte(content)); err != nil {
			log.Panic(err)
		}
	} else {
		if err := os.WriteFile(c.output, []byte(content), 0o666); err != nil {
			log.Panic(err)
		}
	}
}

func (c *rootCmd) genRust(*cobra.Command, []string) {
	data, err := os.ReadFile(c.input)
	if err != nil {
		log.Panic(err)
	}
	r := &protorowdf.RustProtoFile{RustCopyConfig: &c.rustconfig}

	r.ProtoFile, err = parseInput(data, filepath.Ext(c.input))
	if err != nil {
		log.Panic(err)
	}

	content, err := r.CopyDataCode()
	if err != nil {
		log.Panic(err)
	}

	if c.output == "-" {
		if _, err := os.Stdout.Write([]byte(content)); err != nil {
			log.Panic(err)
		}
	} else {
		if err := os.WriteFile(c.output, []byte(content), 0o666); err != nil {
			log.Panic(err)
		}
	}
}
