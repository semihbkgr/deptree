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
     ├───github.com/pmezard/go-difflib@v1.0.0
     ├───gopkg.in/yaml.v3@v3.0.1
     │    └───gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
     ├───github.com/google/shlex@v0.0.0-20191202100458-e7afc7fbc510
     ├───github.com/gregjones/httpcache@v0.0.0-20180305231024-9cad4c3443a7
     ├───golang.org/x/text@v0.3.7
     │    └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     ├───gopkg.in/inf.v0@v0.9.1
     ├───github.com/chai2010/gettext-go@v1.0.2
     ├───github.com/davecgh/go-spew@v1.1.1
     ├───github.com/daviddengcn/go-colortext@v1.0.0
     │    ├───github.com/golangplus/bytes@v1.0.0
     │    │    └───github.com/golangplus/testing@v1.0.0
     │    │         └───2 more ...
     │    └───github.com/golangplus/testing@v1.0.0
     │         ├───github.com/golangplus/bytes@v0.0.0-20160111154220-45c989fe5450
     │         └───github.com/golangplus/fmt@v1.0.0
     ├───github.com/mitchellh/go-wordwrap@v1.0.0
     ├───github.com/moby/spdystream@v0.2.0
     │    └───github.com/gorilla/websocket@v1.4.2
     ├───github.com/spf13/pflag@v1.0.5
     .
     .
     .
```
