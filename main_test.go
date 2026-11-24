package main

import "testing"

func TestAreMirrors(t *testing.T) {
	a := tn(1, tn(2, tn(4, nil, nil), nil), tn(3, nil, nil))
	b := tn(1, tn(3, nil, nil), tn(2, nil, tn(4, nil, nil)))
	if !AreMirrors(a, b) {
		t.Error("expected mirrors")
	}
	// Equal trees are NOT mirrors unless symmetric themselves.
	c := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	d := tn(1, tn(2, nil, nil), tn(3, nil, nil))
	if AreMirrors(c, d) {
		t.Error("identical asymmetric trees are not mirrors")
	}
	if !AreMirrors(nil, nil) {
		t.Error("two nils mirror")
	}
	if AreMirrors(c, nil) {
		t.Error("tree vs nil cannot mirror")
	}
}

func TestMirrorConstruction(t *testing.T) {
	a := tn(1, tn(2, tn(4, nil, nil), nil), tn(3, nil, nil))
	copyA := tn(1, tn(2, tn(4, nil, nil), nil), tn(3, nil, nil))
	Mirror(copyA)
	if !AreMirrors(a, copyA) {
		t.Error("Mirror(t) must mirror the original")
	}
	// Mirroring twice restores the original.
	Mirror(copyA)
	var same func(x, y *TreeNode) bool
	same = func(x, y *TreeNode) bool {
		if x == nil || y == nil {
			return x == y
		}
		return x.Val == y.Val && same(x.Left, y.Left) && same(x.Right, y.Right)
	}
	if !same(a, copyA) {
		t.Error("double mirror must restore the tree")
	}
}
