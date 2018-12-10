package examples

import "github.com/liuchong/econf"

func init() {
	econf.SetFields(&Postgres)
	econf.SetFields(&Redis)
}
