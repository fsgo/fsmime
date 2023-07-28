// Copyright(C) 2023 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2023/7/28

package fsmime

import (
	"mime"
)

func init() {
	for ext, tp := range apache {
		if mime.TypeByExtension(ext) == "" {
			_ = mime.AddExtensionType(ext, tp)
		}
	}
}
