package mysql

import "github.com/yangyuqian/genus"

var testImports = genus.Imports{
	"database/sql":  "",
	"bytes":         "",
	"flag":          "",
	"fmt":           "",
	"reflect":       "",
	"math/rand":     "",
	"os":            "",
	"path/filepath": "",
	"testing":       "",
	"time":          "",
	"strings":       "",
	"sync":          "",

	"github.com/kat-co/vala":                 "",
	"github.com/pkg/errors":                  "",
	"github.com/spf13/viper":                 "",
	"github.com/vattle/sqlboiler/boil":       "",
	"github.com/vattle/sqlboiler/queries":    "",
	"github.com/vattle/sqlboiler/queries/qm": "",
	"github.com/vattle/sqlboiler/strmangle":  "",
	"github.com/jinzhu/gorm":                 "",
	"io":                                    "",
	"io/ioutil":                             "",
	"regexp":                                "",
	"os/exec":                               "",
	"github.com/go-sql-driver/mysql":        "_",
	"gopkg.in/nullbio/null.v6":              "null",
	"github.com/jinzhu/gorm/dialects/mysql": "_",
}

var ormImports = genus.Imports{
	"database/sql":  "",
	"bytes":         "",
	"flag":          "",
	"fmt":           "",
	"reflect":       "",
	"math/rand":     "",
	"os":            "",
	"path/filepath": "",
	"testing":       "",
	"time":          "",
	"strings":       "",
	"sync":          "",

	"github.com/jinzhu/gorm":                 "",
	"github.com/kat-co/vala":                 "",
	"github.com/pkg/errors":                  "",
	"github.com/spf13/viper":                 "",
	"github.com/vattle/sqlboiler/boil":       "",
	"github.com/vattle/sqlboiler/queries":    "",
	"github.com/vattle/sqlboiler/queries/qm": "",
	"github.com/vattle/sqlboiler/strmangle":  "",
	"io":                       "",
	"io/ioutil":                "",
	"regexp":                   "",
	"os/exec":                  "",
	"gopkg.in/nullbio/null.v6": "null",
}
