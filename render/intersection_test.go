package render

import (
	"testing"
)

func TestTheHitWhenAllIntersectionsHavePositiveT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2)
	if xs.Hit() != i1 {
		t.Errorf("Hit was incorrect")
	}
}

func TestTheHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := NewIntersections(i1, i2)
	if xs.Hit() != i2 {
		t.Errorf("Hit was incorrect")
	}
}

func TestTheHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := NewIntersections(i1, i2)
	if xs.Hit() != nil {
		t.Errorf("Hit was incorrect")
	}
}

func TestTheHitIsAlwaysTheLowestNonnegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2, i3, i4)
	if xs.Hit() != i4 {
		t.Errorf("Hit was incorrect")
	}
}
