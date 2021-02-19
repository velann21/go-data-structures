package solid

//wrong approach
// In this below apprch if we need to add a new filter then we should be
// chnaging the code and we should be testing it again & agian
type Color int
const (
	Green Color = iota
	Red
	Blue
)
type Size int
const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name string
	Color Color
	Size  Size
}

type Filter struct {

}

func (filter *Filter) FilterByColour(prod []Product, color Color)*[]Product{
	result := make([]Product, 0)
	for i, v := range prod{
		if v.Color == color{
			result = append(result, prod[i])
		}
	}

	return &result
}


func (filter *Filter) FilterBySize(prod []Product, size Size)*[]Product{
	result := make([]Product, 0)
	for i, v := range prod{
		if v.Size == size{
			result = append(result, prod[i])
		}
	}

	return &result
}


//correct approach
type Specification interface {
	IsSatisfied(p *Product) bool
}

type SizeSpecification struct {
	Size Size
}

func (sizeSpecification *SizeSpecification) IsSatisfied(p *Product)bool{
	if p.Size == sizeSpecification.Size{
		return true
	}
	return false
}


type ColorSpecification struct {
	Color Color
}
func (colorSpecification *ColorSpecification) IsSatisfied(p *Product)bool{
	if p.Color == colorSpecification.Color{
		return true
	}
	return false
}


type BetterFilter struct {

}

func (bf *BetterFilter) Filter(prod []Product, spec Specification) *[]Product {
	result := make([]Product, 0)
	for i,v := range prod{
		if spec.IsSatisfied(&v){
			result = append(result, prod[i])
		}
	}
	return &result
}