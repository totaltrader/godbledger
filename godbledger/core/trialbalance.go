package core

import ()

type TBAccount struct {
	Account string   `json:"Account"`
	Amount  int      `json:"Amount"`
	Tags    []string `json:"Tags"`
}
