package datatypes

type distance float32
type km float64


func (miles distance) toKm() km {
	return km(1.60934 * miles)
}
func (miles km) toMiles() km {
	return km(miles / 1.60934)
}

func test() {
	d := distance(35.5)
	km := d.toKm()
	km.toMiles()
	print(d)

}
