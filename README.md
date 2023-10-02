# protorowdf (ProtoBuf Row and DataFrame)

Sometimes we want to represent data as a series of rows, and some times the same data should be in columns. This library is designed to generate two proto message definitions, one of row and one for column, and provides the conversion functions between the two.

## Example

Provided below input:

```yaml
comment: |
  Test package for row/dataframe generation for protobuf.
    comment can start with some spaces/indents.

  blank lines will be conserved if they are in between non-empty lines.
package_name: pkg.name
options:
  - name: go_package
    value: '"example.com/what/a/good/package"'
  - name: optimize_for
    value: CODE_SIZE
  - name: cc_enable_arenas
    value: true
structs:
  - name: ARowStruct
    comment: |
      Struct comments are similar to file comments.
    fields:
      - name: density
        plural_name: densities
        data_type: Int64
        comment: |
          comments on data field.
          can also be multiple lines.
        tag_num: 1
      - name: string_field
        data_type: String
        tag_num: 2
        comment: | # notice leading blank lines will be trimmed.



          more comment on another field.

  - name: AnotherRowStruct
    comment: |
      This is another struct's comment.

      And you should be able to put a blank line here too.
    fields:
      - name: adata
        data_type: string
        tag_num: 2
      - name: anotherdata
        data_type: fixed32
        tag_num: 3
```

The below output will be generated

```proto
syntax = "proto3";

// Test package for row/dataframe generation for protobuf.
//   comment can start with some spaces/indents.
//
// blank lines will be conserved if they are in between non-empty lines.
package pkg.name;

option go_package = "example.com/what/a/good/package";
option optimize_for = CODE_SIZE;
option cc_enable_arenas = true;

// Struct comments are similar to file comments.
message ARowStruct {
  // comments on data field.
  // can also be multiple lines.
  int64 density = 1;
  // more comment on another field.
  string string_field = 2;
}

// Plural Version of ARowStruct.
// Struct comments are similar to file comments.
message ARowStructs {
  // comments on data field.
  // can also be multiple lines.
  repeated int64 densities = 1;
  // more comment on another field.
  repeated string string_fields = 2;
}

// This is another struct's comment.
//
// And you should be able to put a blank line here too.
message AnotherRowStruct {
  string adata = 2;
  fixed32 anotherdata = 3;
}

// Plural Version of AnotherRowStruct.
// This is another struct's comment.
//
// And you should be able to put a blank line here too.
message AnotherRowStructs {
  repeated string adatas = 2;
  repeated fixed32 anotherdatas = 3;
}
```
