package designer

import (
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
)

// /DumpIB <имя файла>
//— выгрузка информационной базы в командном режиме.
type DumpIBOptions struct {
	Designer `v8:",inherit" json:"designer"`

	File string `v8:"/DumpIB" json:"file"`
}

func (d DumpIBOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v
}

// /RestoreIB <имя файла>
// — загрузка информационной базы в командном режиме.
// Если файл информационной базы отсутствует в указанном каталоге, будет создана новая информационная база.
type RestoreIBOptions struct {
	Designer `v8:",inherit" json:"designer"`

	File string `v8:"/RestoreIB" json:"file"`
}

func (d RestoreIBOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v
}
