package fuzzy_bool

import "fmt"

type FuzzyBool struct {
	value float32
}

func (fuzzy *FuzzyBool)String() string {
	return fmt.Sprintf("%.0f%%",100*fuzzy.value)
}

func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

func (fuzzy *FuzzyBool) And(first *FuzzyBool,
rest ...*FuzzyBool) *FuzzyBool {
	minimum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if minimum > other.value {
			minimum = other.value
		}
	}
	return &FuzzyBool{minimum}
}

func (fuzzy *FuzzyBool) Or(first *FuzzyBool,
rest ...*FuzzyBool) *FuzzyBool {
	maximum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if maximum < other.value {
			maximum = other.value
		}
	}
	return &FuzzyBool{maximum}
}

func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.value >= .5
}

func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.value)
}

func (fuzzy *FuzzyBool) Set(value interface{})(err error) {
	fuzzy.value,err=float32ForValue(value)
	return err
}

func (fuzzy *FuzzyBool)Copy()*FuzzyBool  {
	return &FuzzyBool{fuzzy.value}
}

func New(value interface{})(*FuzzyBool,error)  {
	amount,err:=float32ForValue(value)
	return &FuzzyBool{amount},err
}

func float32ForValue(value interface{})(fuzzy float32,err error)  {

	fmt.Println("debug");
	switch  value := value.(type) {

	case float32:
		fmt.Println(value)
		fuzzy=value
	case float64:
		fmt.Println(value)
		fuzzy=float32(value)
	case int:
		fmt.Println(value)
		fuzzy=float32(value)
	case bool:
		fuzzy=0
		if value{
			fuzzy=1
		}
	default:
		return 0,fmt.Errorf("float32ForValue():%v is not a number or Boolean",value)

	}
	if fuzzy<0{
		fuzzy=0
	}else if fuzzy>1{
		fuzzy=1
	}
	return fuzzy,nil
}