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

func main() {
	newCmd().Execute()
}

type Cmd struct {
	*cobra.Command

	input  string
	output string
}

func newCmd() *Cmd {
	c := &Cmd{
		Command: &cobra.Command{
			Use:   "gen-proto-row-df",
			Short: "generate proto defintions for rows of data and the corresponding data frame.",
			Long:  "generate proto defintions for rows of data and the corresponding data frame.",
		},
	}

	c.Run = c.runit

	c.Flags().StringVarP(&c.input, "input", "i", c.input, "input file, the format of the input is decided by the extension (yaml/yml for yaml, json for json, txt for text etc.)")
	c.MarkFlagRequired("input")
	c.MarkFlagFilename("input", "yaml", "yml", "json", "txt", "txtpb", "pb", "binpb")
	c.Flags().StringVarP(&c.output, "output", "o", c.output, "output file")
	c.MarkFlagFilename("output")
	c.MarkFlagRequired("output")

	return c
}

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

func (c *Cmd) runit(*cobra.Command, []string) {
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
