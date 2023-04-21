package main

/*
	Now in below approach we are making a specification interface which defines a function IsSatisfied and takes a certain type
	in this case Product, now when some type implements this interface we can compare the properties and pass the result from
	IsSatisfied function, this way we are not modifying any type rather we are extending its functionality
*/

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

/*
AndSpecification is the type which demonstrates how can we implement a functionality where we need to filter on multiple conditional
basis, to be noted that this AndSpecification is also extending functionality of other specifications and not implementing them
again or trying to modify anything in previous written specifications
*/
type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

/*
	Below is the better filter Type which takes in itself the data to work upon and specification to be satisfied
*/
type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
