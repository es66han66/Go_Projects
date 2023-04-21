package main

/*
Below is a filter which is written in such a way that for every filtering mechanism which will be needed we have to write another function
and attach that function to it, this is not the ideal way because we are continuously modifying the existing functionality of
Filter type, instead of doing this, we should be making Filter type in such a way that we can extend it or our Filter type can
extend some else type so that it can perform the desired comparisons.
*/
type Filter struct {
}

func (f *Filter) filterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}
