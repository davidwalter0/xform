xform: main.go
	GO111MODULE=on go build -tags netgo -ldflags '-w -s' -o xform

install:
	GO111MODULE=on go install -tags netgo -ldflags '-w -s'

.PHONY: test
test-data:
	mkdir -p $@
test: test-data xform
	./xform json2hcl  < fixtures/infra.tf.json > test-data/json2.hcl
	diff -wui fixtures/infra.tf test-data/json2.hcl
	./xform json2yaml < fixtures/infra.tf.json > test-data/json2.yaml
	./xform hcl2yaml  < fixtures/infra.tf      > test-data/hcl2.yaml
	diff -wui test-data/json2.yaml test-data/hcl2.yaml
	./xform hcl2json  < fixtures/infra.tf      > test-data/hcl2.json
	diff -wui fixtures/infra.tf.json test-data/hcl2.json
	./xform yaml2hcl  < test-data/hcl2.yaml         > test-data/yaml2.hcl
	diff -wui fixtures/infra.tf      test-data/yaml2.hcl
	./xform yaml2json < test-data/hcl2.yaml         > test-data/yaml2.json
	diff -wui fixtures/infra.tf.json test-data/yaml2.json
clean:
	rm -f test-data/* xform
