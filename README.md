Multi format transform

Transform (xform) filter for  json, hcl and yaml via STDIN / STDOUT.

multi call binary `xform [translation]` or `ln xform translation` where
translation is one of

- yaml2json
- yaml2hcl
- json2hcl
- json2yaml
- hcl2json
- hcl2yaml

For example:
```
ln xlate yaml2json
./yaml2json < yamlfile > jsonfile
```

or

```
./xlate yaml2json < yamlfile > jsonfile
```

## Install

```
GO111MODULE=on go get -tags netgo -ldflags '-w -s' github.com/davidwalter0/xform
```

## Use

Here's an example [`fixtures/infra.tf.json`](fixtures/infra.tf.json) being
converted to HCL:

```bash
$ xform json2hcl < fixtures/infra.tf.json
"output" "arn" {
  "value" = "${aws_dynamodb_table.basic-dynamodb-table.arn}"
}
... rest of HCL truncated
```

Typical use would be

```bash
$ xform json2hcl < fixtures/infra.tf.json > fixtures/infra.tf
```

## hcl2json

As a bonus, the conversation the other way around is also supported via the `-reverse` flag:

```bash
$ xform hcl2json < fixtures/infra.tf
{
  "output": [
    {
      "arn": [
        {
          "value": "${aws_dynamodb_table.basic-dynamodb-table.arn}"
        }
      ]
    }, 
  ... rest of JSON truncated
  ]
}
```

## Development

```bash
go get -tags netgo -ldflags '-w -s' github.com/davidwalter0/xform
```

---
## upstream rational
[Upstream README](https://github.com/kvz/json2hcl/blob/master/README.md)

If you don't know HCL, read [Why HCL](https://github.com/hashicorp/hcl#why).

## Contributors
- [Upstream](https://github.com/kvz/json2hcl)
- [Marius Kleidl](https://github.com/Acconut)
- [Kevin van Zonneveld](https://github.com/kvz)
