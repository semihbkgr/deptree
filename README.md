# deptree

`deptree` is a utility command that displays the dependency tree of this module.

It is inspired by maven `dependency:tree` plugin

```bash
go install github.com/semihbkgr/deptree:v0.1.1
```

Flag `-d` is used to set depth of the tree. It is 5 by default.

__Example__

```bash
deptree -d 3
```

Output

```
github.com/gin-gonic/gin
     ├───github.com/goccy/go-json@v0.9.11
     ├───golang.org/x/crypto@v0.0.0-20210711020723-a769d52b0f97
     │    ├───golang.org/x/net@v0.0.0-20210226172049-e18ecbb05110
     │    │    ├───golang.org/x/sys@v0.0.0-20201119102817-f84b799fce68
     │    │    ├───golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
     │    │    │    └───1 more ...
     │    │    └───golang.org/x/text@v0.3.3
     │    │         └───1 more ...
     │    ├───golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
     │    ├───golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
     │    │    └───golang.org/x/sys@v0.0.0-20201119102817-f84b799fce68
     │    └───golang.org/x/text@v0.3.3
     │         └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     ├───github.com/go-playground/validator/v10@v10.10.0
     │    ├───golang.org/x/text@v0.3.6
     │    │    └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     │    ├───gopkg.in/check.v1@v1.0.0-20201130134442-10cb98267c6c
     │    │    └───github.com/kr/pretty@v0.2.1
     │    │         └───1 more ...
     │    ├───github.com/go-playground/universal-translator@v0.18.0
     │    ├───github.com/kr/pretty@v0.3.0
     │    │    ├───github.com/kr/text@v0.2.0
     │    │    │    └───1 more ...
     │    │    └───github.com/rogpeppe/go-internal@v1.6.1
     │    │         └───1 more ...
     │    ├───github.com/leodido/go-urn@v1.2.1
     │    │    └───github.com/stretchr/testify@v1.6.1
     │    │         ├───4 more ...
     │    ├───github.com/rogpeppe/go-internal@v1.8.0
     │    │    ├───github.com/pkg/diff@v0.0.0-20210226163009-20ebb0f2a09e
     │    │    └───gopkg.in/errgo.v2@v2.1.0
     │    │         ├───2 more ...
     │    ├───github.com/stretchr/testify@v1.7.0
     │    │    ├───github.com/davecgh/go-spew@v1.1.0
     │    │    ├───github.com/pmezard/go-difflib@v1.0.0
     │    │    ├───github.com/stretchr/objx@v0.1.0
     │    │    └───gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
     │    │         └───1 more ...
     │    ├───golang.org/x/crypto@v0.0.0-20210711020723-a769d52b0f97
     │    │    ├───golang.org/x/net@v0.0.0-20210226172049-e18ecbb05110
     │    │    │    ├───3 more ...
     │    │    ├───golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
     │    │    ├───golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
     │    │    │    └───1 more ...
     │    │    └───golang.org/x/text@v0.3.3
     │    │         └───1 more ...
     │    ├───gopkg.in/yaml.v3@v3.0.0-20210107192922-496545a6307b
     │    │    └───gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
     │    ├───github.com/davecgh/go-spew@v1.1.1
     │    ├───github.com/go-playground/assert/v2@v2.0.1
     │    ├───github.com/go-playground/locales@v0.14.0
     │    └───golang.org/x/sys@v0.0.0-20210806184541-e5e7981a1069
     ├───github.com/go-playground/universal-translator@v0.18.0
     │    └───github.com/go-playground/locales@v0.14.0
     ├───github.com/modern-go/reflect2@v1.0.2
     ├───github.com/pmezard/go-difflib@v1.0.0
     ├───github.com/davecgh/go-spew@v1.1.1
     ├───github.com/mattn/go-isatty@v0.0.16
     │    └───golang.org/x/sys@v0.0.0-20220811171246-fbc7d0a398ab
     ├───github.com/pelletier/go-toml/v2@v2.0.2
     │    └───github.com/stretchr/testify@v1.7.2
     │         ├───github.com/stretchr/objx@v0.1.0
     │         ├───gopkg.in/yaml.v3@v3.0.1
     │         ├───github.com/davecgh/go-spew@v1.1.0
     │         └───github.com/pmezard/go-difflib@v1.0.0
     ├───github.com/twitchyliquid64/golang-asm@v0.15.1
     ├───github.com/ugorji/go/codec@v1.2.7
     │    └───github.com/ugorji/go@v1.2.7
     │         └───github.com/ugorji/go/codec@v1.2.7
     ├───github.com/bytedance/sonic@v1.4.0
     │    ├───github.com/goccy/go-json@v0.9.4
     │    ├───github.com/klauspost/cpuid/v2@v2.0.9
     │    ├───github.com/tidwall/sjson@v1.2.4
     │    │    ├───github.com/tidwall/pretty@v1.2.0
     │    │    └───github.com/tidwall/gjson@v1.12.1
     │    │         ├───2 more ...
     │    ├───golang.org/x/arch@v0.0.0-20210923205945-b76863e36670
     │    │    └───rsc.io/pdf@v0.1.1
     │    ├───github.com/chenzhuoyu/base64x@v0.0.0-20211019084208-fb5309c8db06
     │    │    └───github.com/klauspost/cpuid/v2@v2.0.9
     │    ├───github.com/davecgh/go-spew@v1.1.1
     │    ├───github.com/json-iterator/go@v1.1.12
     │    │    ├───github.com/davecgh/go-spew@v1.1.1
     │    │    ├───github.com/google/gofuzz@v1.0.0
     │    │    ├───github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
     │    │    ├───github.com/modern-go/reflect2@v1.0.2
     │    │    └───github.com/stretchr/testify@v1.3.0
     │    │         ├───3 more ...
     │    ├───github.com/stretchr/testify@v1.7.0
     │    │    ├───github.com/stretchr/objx@v0.1.0
     │    │    ├───gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
     │    │    │    └───1 more ...
     │    │    ├───github.com/davecgh/go-spew@v1.1.0
     │    │    └───github.com/pmezard/go-difflib@v1.0.0
     │    ├───github.com/tidwall/gjson@v1.13.0
     │    │    ├───github.com/tidwall/match@v1.1.1
     │    │    └───github.com/tidwall/pretty@v1.2.0
     │    └───github.com/twitchyliquid64/golang-asm@v0.15.1
     ├───google.golang.org/protobuf@v1.28.1
     │    ├───github.com/golang/protobuf@v1.5.0
     │    │    ├───github.com/google/go-cmp@v0.5.5
     │    │    │    └───1 more ...
     │    │    └───google.golang.org/protobuf@v1.26.0-rc.1
     │    │         └───1 more ...
     │    └───github.com/google/go-cmp@v0.5.5
     │         └───golang.org/x/xerrors@v0.0.0-20191204190536-9bdfabe68543
     ├───gopkg.in/yaml.v3@v3.0.1
     │    └───gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
     ├───golang.org/x/text@v0.3.6
     │    └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     ├───github.com/leodido/go-urn@v1.2.1
     │    └───github.com/stretchr/testify@v1.6.1
     │         ├───github.com/davecgh/go-spew@v1.1.0
     │         ├───github.com/pmezard/go-difflib@v1.0.0
     │         ├───github.com/stretchr/objx@v0.1.0
     │         └───gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
     │              └───1 more ...
     ├───github.com/stretchr/testify@v1.8.0
     │    ├───github.com/davecgh/go-spew@v1.1.1
     │    ├───github.com/pmezard/go-difflib@v1.0.0
     │    ├───github.com/stretchr/objx@v0.4.0
     │    │    ├───github.com/davecgh/go-spew@v1.1.1
     │    │    └───github.com/stretchr/testify@v1.7.1
     │    │         ├───4 more ...
     │    └───gopkg.in/yaml.v3@v3.0.1
     │         └───gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
     ├───github.com/go-playground/locales@v0.14.0
     │    └───golang.org/x/text@v0.3.6
     │         └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     ├───github.com/klauspost/cpuid/v2@v2.0.14
     ├───github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
     ├───golang.org/x/sys@v0.0.0-20220811171246-fbc7d0a398ab
     ├───gopkg.in/yaml.v2@v2.4.0
     │    └───gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
     ├───github.com/chenzhuoyu/base64x@v0.0.0-20220526154910-8bf9453eb81a
     │    └───github.com/klauspost/cpuid/v2@v2.0.9
     ├───github.com/gin-contrib/sse@v0.1.0
     │    └───github.com/stretchr/testify@v1.3.0
     │         ├───github.com/davecgh/go-spew@v1.1.0
     │         ├───github.com/pmezard/go-difflib@v1.0.0
     │         └───github.com/stretchr/objx@v0.1.0
     ├───golang.org/x/arch@v0.0.0-20220412001346-fc48f9fe4c15
     │    └───rsc.io/pdf@v0.1.1
     ├───golang.org/x/net@v0.0.0-20210226172049-e18ecbb05110
     │    ├───golang.org/x/sys@v0.0.0-20201119102817-f84b799fce68
     │    ├───golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
     │    │    └───golang.org/x/sys@v0.0.0-20201119102817-f84b799fce68
     │    └───golang.org/x/text@v0.3.3
     │         └───golang.org/x/tools@v0.0.0-20180917221912-90fa682c2a6e
     └───github.com/json-iterator/go@v1.1.12
          ├───github.com/davecgh/go-spew@v1.1.1
          ├───github.com/google/gofuzz@v1.0.0
          ├───github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
          ├───github.com/modern-go/reflect2@v1.0.2
          └───github.com/stretchr/testify@v1.3.0
               ├───github.com/stretchr/objx@v0.1.0
               ├───github.com/davecgh/go-spew@v1.1.0
               └───github.com/pmezard/go-difflib@v1.0.0
```
