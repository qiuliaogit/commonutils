package commonutils

const (
	cny_precision = 100                     // 货币精度
	cny_max_value = int64(9007199254740991) // 24
	cny_min_value = -int64(9007199254740991)

	cny_float_max_value = float64(cny_max_value) / float64(cny_precision)
	cny_float_min_value = -float64(cny_min_value) / float64(cny_precision)

	// // 基本单位
	//
	//	cnIntRadice = ['', '拾', '佰', '仟'];
	//
	// // 对应整数部分扩展单位
	//
	//	cnIntUnits = ['', '万', '亿', '兆'];
	//
	// // 对应小数部分单位
	//
	//	cnDecUnits = ['角', '分', '毫', '厘'];
	//
	// // 整型完以后的单位
	//
	cnIntLast = "元"
)

var (
	cnNums      = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	cnIntRadice = []string{"", "拾", "佰", "仟"}
	cnIntUnits  = []string{"", "万", "亿", "兆"}
	cnDecUnits  = []string{"角", "分", "毫", "厘"}
)

// 货币类型
type Currency struct {
	Value int64
}

func CreateCurrency[T int64 | float64 | string| int32 | uint32 | int8 | uint8 | int | uint](value int64) (Currency, error) {
}
