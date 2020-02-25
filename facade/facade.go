package facade

type A struct {
	Member int
}
func (a *A) UpdateBaseVal() {
	a.Member += 1
}

type B struct {
	MemberA *A
	Member int
}
func (b *B) UpdateBaseVal() {
	b.MemberA.Member += b.Member
}

type C struct {
	MemberB *B
	MemberA *A
	Member int
}
func (c *C) UpdateBaseVal() {
	c.MemberA.Member += c.Member
}

type D struct {
	MemberB *B
	MemberC *C
	Member int
}
func (d *D) UpdateBaseVal() {
	d.MemberC.MemberB.MemberA.Member += d.Member
}

// Functions providing the Facade (API) for managing these complex "objects" below
func CreateD(val int) *D {
	constructed := &D{}
	constructed.MemberB = &B{
		MemberA: &A{val},
		Member: val,
	}
	constructed.MemberC = &C{
		MemberB: constructed.MemberB,
		MemberA: constructed.MemberB.MemberA,
		Member: val,
	}

	return constructed
}

// GetBaseVal is a stand-in for complex functionality that requires multiple nested operations starting at D
func GetBaseVal(d *D) int {
	return d.MemberB.MemberA.Member
}

func UpdateBaseVal(d *D) int {
	d.UpdateBaseVal()
	d.MemberB.UpdateBaseVal()
	d.MemberC.UpdateBaseVal()
	d.MemberC.MemberA.UpdateBaseVal()
	return d.MemberC.MemberB.MemberA.Member
}