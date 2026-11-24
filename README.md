# Check if Two Trees are Mirrors

> Striver SDE Sheet 129 · Day 19 Binary Tree III — the crossed recursion from the symmetric check, applied across two trees, plus the in-place Mirror constructor. Go, stdlib only, tested.

## Problem

Decide whether two binary trees are mirror images of each other — and, constructively, turn a tree into its mirror.

## Approach

The check is #127's helper promoted to the public interface: values equal, a.Left mirrors b.Right, a.Right mirrors b.Left. (Symmetric tree = `AreMirrors(root.Left, root.Right)` — the sheet lists both because candidates routinely conflate "identical" with "mirror"; the test that identical asymmetric trees are *not* mirrors pins the distinction.) The constructive `Mirror` is a postorder swap of children at every node, and Go's parallel assignment `n.Left, n.Right = Mirror(n.Right), Mirror(n.Left)` does the recursion and the swap in one line. Mirroring twice is the identity — a property the test uses.

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
