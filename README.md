# deptree

`deptree` is a utility tool that displays the dependency tree of go modules.

It is inspired by maven `dependency:tree` plugin

```bash
$ go install github.com/semihbkgr/deptree@v0.1.2
```

Arg `-d` is used to set depth of the tree. It is 5 by default.

__Example__

```bash
$ deptree -d 3
```

Output

```
k8s.io/kubectl
     ├───github.com/go-openapi/jsonpointer@v0.19.5
     │    ├───github.com/go-openapi/swag@v0.19.5
     │    │    ├───github.com/mailru/easyjson@v0.0.0-20190614124828-94de47d64c63
     │    │    ├───github.com/stretchr/testify@v1.3.0
     │    │    │    └───3 more ...
     │    │    ├───gopkg.in/check.v1@v1.0.0-20180628173108-788fd7840127
     │    │    ├───gopkg.in/yaml.v2@v2.2.2
     │    │    │    └───1 more ...
     │    │    ├───github.com/davecgh/go-spew@v1.1.1
     │    │    └───github.com/kr/pretty@v0.1.0
     │    │         └───1 more ...
     │    ├───github.com/mailru/easyjson@v0.0.0-20190626092158-b2ccc519800e
     │    └───github.com/stretchr/testify@v1.3.0
     │         ├───github.com/pmezard/go-difflib@v1.0.0
     │         ├───github.com/stretchr/objx@v0.1.0
     │         └───github.com/davecgh/go-spew@v1.1.0
     ├───github.com/go-openapi/jsonreference@v0.20.0
     │    ├───github.com/go-openapi/jsonpointer@v0.19.3
     │    │    ├───github.com/stretchr/testify@v1.3.0
     │    │    ├───github.com/go-openapi/swag@v0.19.5
     │    │    └───github.com/mailru/easyjson@v0.0.0-20190626092158-b2ccc519800e
     │    └───github.com/stretchr/testify@v1.3.0
     │         ├───github.com/pmezard/go-difflib@v1.0.0
     │         ├───github.com/stretchr/objx@v0.1.0
     │         └───github.com/davecgh/go-spew@v1.1.0
     ├───github.com/onsi/gomega@v1.20.1
     │    ├───gopkg.in/yaml.v3@v3.0.1
     │    │    └───gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
     │    ├───golang.org/x/sys@v0.0.0-20220422013727-9388b58f7150
     │    ├───golang.org/x/text@v0.3.7
     │    │    └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     │    ├───google.golang.org/protobuf@v1.28.0
     │    ├───github.com/golang/protobuf@v1.5.2
     │    ├───github.com/google/go-cmp@v0.5.8
     │    ├───github.com/onsi/ginkgo/v2@v2.1.4
     │    └───golang.org/x/net@v0.0.0-20220425223048-2871e0cb64e4
     .
     .
     .
```

Input is accepted with `-` arg.

```bash
$ deptree - < graph.txt
```
