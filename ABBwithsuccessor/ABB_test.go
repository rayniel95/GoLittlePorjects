package ABB

import (
	"testing"
)

// func TestABB_isLeaf(t *testing.T) {
// 	type fields struct {
// 		left      *ABB
// 		right     *ABB
// 		successor *ABB
// 		value     int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			node := &ABB{
// 				left:      tt.fields.left,
// 				right:     tt.fields.right,
// 				successor: tt.fields.successor,
// 				value:     tt.fields.value,
// 			}
// 			if got := node.isLeaf(); got != tt.want {
// 				t.Errorf("ABB.isLeaf() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_searchAnteccessorLeftSon(t *testing.T) {
// 	type args struct {
// 		root *ABB
// 	}
// 	node1 := &ABB{
// 		value: 4000,
// 		left: &ABB{
// 			value: 3500,
// 		},
// 		right: &ABB{
// 			value: 4500,
// 		},
// 	}
// 	node2 := &ABB{
// 		value: 3000,
// 		left: &ABB{
// 			value: 2000,
// 			left: &ABB{
// 				value: 500,
// 			},
// 		},
// 		right: node1,
// 	}
// 	node3 := &ABB{
// 		value: 5000,
// 		left:  node2,
// 		right: &ABB{
// 			value: 20000,
// 			left: &ABB{
// 				value: 15000,
// 			},
// 			right: &ABB{
// 				value: 50000,
// 			},
// 		},
// 	}
// 	tree := &ABB{
// 		value: 1000,
// 		left: &ABB{
// 			value: 500,
// 			left: &ABB{
// 				value: 200,
// 			},
// 			right: &ABB{
// 				value: 600,
// 			},
// 		},
// 		right: node3,
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"test1", args{
// 				root: tree.left,
// 			},
// 			600,
// 		},
// 		{
// 			"test2", args{
// 				root: (*node1).left,
// 			},
// 			3500,
// 		},
// 		{
// 			"test3", args{
// 				root: (*node2).left,
// 			},
// 			2000,
// 		},
// 		{
// 			"test3", args{
// 				root: (*node3).left,
// 			},
// 			4500,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := searchAnteccessorLeftSon(tt.args.root); !reflect.DeepEqual((*got).value, tt.want) {
// 				t.Errorf("searchAnteccessorLeftSon() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_searchAnteccessorNotLeftSon(t *testing.T) {c
// 		null = nil
// 		anteccessor = &null
// 	}
// }

// func Test_searchParent(t *testing.T) {

// 	node1 := &ABB{
// 		value: 4000,
// 		left: &ABB{
// 			value: 3500,
// 			right: &ABB{
// 				value: 3700,
// 			},
// 		},
// 		right: &ABB{
// 			value: 4500,
// 		},
// 	}
// 	node2 := &ABB{
// 		value: 3000,
// 		left: &ABB{
// 			value: 2000,
// 			left: &ABB{
// 				value: 1500,
// 			},
// 		},
// 		right: node1,
// 	}

// 	node5 := &ABB{
// 		value: 18000,
// 		right: &ABB{
// 			value: 19000,
// 		},
// 	}

// 	node4 := &ABB{
// 		value: 15000,
// 		right: &ABB{
// 			value: 17000,
// 			right: node5,
// 		},
// 	}

// 	node3 := &ABB{
// 		value: 5000,
// 		left:  node2, // 11
// 		right: &ABB{
// 			value: 20000,
// 			left:  node4,
// 			right: &ABB{
// 				value: 50000,
// 			},
// 		},
// 	}
// 	tree := &ABB{
// 		value: 1000,
// 		left: &ABB{
// 			value: 500,
// 			left: &ABB{
// 				value: 250,
// 			},
// 			right: &ABB{
// 				value: 600,
// 			},
// 		},
// 		right: node3,
// 	}
// 	var null *ABB = nil
// 	var parent *(*ABB) = &null

// 	type args struct {
// 		root   *ABB
// 		value  int
// 		parent *(*ABB)
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"test1", args{
// 				root:   tree,
// 				value:  1500,
// 				parent: parent,
// 			},
// 			2000,
// 		},
// 		{
// 			"test2", args{
// 				root:   tree,
// 				value:  4500,
// 				parent: parent,
// 			},
// 			4000,
// 		},
// 		{
// 			"test3", args{
// 				root:   tree,
// 				value:  5000,
// 				parent: parent,
// 			},
// 			1000,
// 		},
// 		{
// 			"test4", args{
// 				root:   tree,
// 				value:  15000,
// 				parent: parent,
// 			},
// 			20000,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if searchParent(tt.args.root, tt.args.value, tt.args.parent); !reflect.DeepEqual((*(*tt.args.parent)).value, tt.want) {
// 				t.Errorf("searchParent() = %v, want %v", (*(*tt.args.parent)).value, tt.want)
// 			}
// 		})
// 		null = nil
// 		parent = &null
// 	}
// }

// func Test_search(t *testing.T) {

// 	node1 := &ABB{
// 		value: 4000,
// 		left: &ABB{
// 			value: 3500,
// 			right: &ABB{
// 				value: 3700,
// 			},
// 		},
// 		right: &ABB{
// 			value: 4500,
// 		},
// 	}
// 	node2 := &ABB{
// 		value: 3000,
// 		left: &ABB{
// 			value: 2000,
// 			left: &ABB{
// 				value: 1500,
// 			},
// 		},
// 		right: node1,
// 	}

// 	node5 := &ABB{
// 		value: 18000,
// 		right: &ABB{
// 			value: 19000,
// 		},
// 	}

// 	node4 := &ABB{
// 		value: 15000,
// 		right: &ABB{
// 			value: 17000,
// 			right: node5,
// 		},
// 	}

// 	node3 := &ABB{
// 		value: 5000,
// 		left:  node2,
// 		right: &ABB{
// 			value: 20000,
// 			left:  node4,
// 			right: &ABB{
// 				value: 50000,
// 			},
// 		},
// 	}
// 	tree := &ABB{
// 		value: 1000,
// 		left: &ABB{
// 			value: 500,
// 			left: &ABB{
// 				value: 250,
// 			},
// 			right: &ABB{
// 				value: 600,
// 			},
// 		},
// 		right: node3,
// 	}

// 	type args struct {
// 		root  *ABB
// 		value int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"test1", args{
// 				root:  tree,
// 				value: 3500,
// 			},
// 			3500,
// 		},
// 		{
// 			"test2", args{
// 				root:  tree,
// 				value: 4500,
// 			},
// 			4500,
// 		},
// 		{
// 			"test3", args{
// 				root:  tree,
// 				value: 500,
// 			},
// 			500,
// 		},
// 		{
// 			"test4", args{
// 				root:  tree,
// 				value: 3000,
// 			},
// 			3000,
// 		},
// 		{
// 			"test5", args{
// 				root:  tree,
// 				value: 15000,
// 			},
// 			15000,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := search(tt.args.root, tt.args.value); !reflect.DeepEqual((*got).value, tt.want) {
// 				t.Errorf("search() = %v, want %v", (*got).value, tt.want)
// 			}
// 		})
// 	}
// }

func Test_delete(t *testing.T) {

	tree := &ABB{
		value: 1000,
		right: &ABB{
			value: 5000,
			left: &ABB{
				value: 3000,
				left: &ABB{
					value: 2000,
					left: &ABB{
						value: 1500,
					},
				},
				right: &ABB{
					value: 4000,
					left: &ABB{
						value: 3500,
						right: &ABB{
							value: 3700,
						},
					},
					right: &ABB{
						value: 4500,
					},
				},
			},
			right: &ABB{
				value: 20000,
				left: &ABB{
					value: 15000,
					right: &ABB{
						value: 17000,
						right: &ABB{
							value: 18000,
							right: &ABB{
								value: 19000,
							},
						},
					},
				},
				right: &ABB{
					value: 50000,
				},
			},
		},
		left: &ABB{
			value: 500,
			left: &ABB{
				value: 250,
			},
			right: &ABB{
				value: 600,
			},
		},
	}

	setSuccessors(tree)

	type args struct {
		root  *(*ABB)
		value int
	}
	tests := []struct {
		name string
		args args
		want *ABB
	}{
		{
			"test1", args{
				root:  &tree,
				value: 3700,
			},
			&ABB{
				value: 1000,
				right: &ABB{
					value: 5000,
					left: &ABB{
						value: 3000,
						left: &ABB{
							value: 2000,
							left: &ABB{
								value: 1500,
							},
						},
						right: &ABB{
							value: 4000,
							left: &ABB{
								value: 3500,
							},
							right: &ABB{
								value: 4500,
							},
						},
					},
					right: &ABB{
						value: 20000,
						left: &ABB{
							value: 15000,
							right: &ABB{
								value: 17000,
								right: &ABB{
									value: 18000,
									right: &ABB{
										value: 19000,
									},
								},
							},
						},
						right: &ABB{
							value: 50000,
						},
					},
				},
				left: &ABB{
					value: 500,
					left: &ABB{
						value: 250,
					},
					right: &ABB{
						value: 600,
					},
				},
			},
		},
		{
			"test2", args{
				root:  &tree,
				value: 5000,
			},
			&ABB{
				value: 1000,
				right: &ABB{
					value: 3000,
					left: &ABB{
						value: 2000,
						left: &ABB{
							value: 1500,
						},
					},
					right: &ABB{
						value: 4000,
						left: &ABB{
							value: 3500,
						},
						right: &ABB{
							value: 4500,
							right: &ABB{
								value: 20000,
								left: &ABB{
									value: 15000,
									right: &ABB{
										value: 17000,
										right: &ABB{
											value: 18000,
											right: &ABB{
												value: 19000,
											},
										},
									},
								},
								right: &ABB{
									value: 50000,
								},
							},
						},
					},
				},
				left: &ABB{
					value: 500,
					left: &ABB{
						value: 250,
					},
					right: &ABB{
						value: 600,
					},
				},
			},
		},
		{
			"test3", args{
				root:  &tree,
				value: 1000,
			},
			&ABB{
				value: 500,
				left: &ABB{
					value: 250,
				},
				right: &ABB{
					value: 600,
					right: &ABB{
						value: 3000,
						left: &ABB{
							value: 2000,
							left: &ABB{
								value: 1500,
							},
						},
						right: &ABB{
							value: 4000,
							left: &ABB{
								value: 3500,
							},
							right: &ABB{
								value: 4500,
								right: &ABB{
									value: 20000,
									left: &ABB{
										value: 15000,
										right: &ABB{
											value: 17000,
											right: &ABB{
												value: 18000,
												right: &ABB{
													value: 19000,
												},
											},
										},
									},
									right: &ABB{
										value: 50000,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"test4", args{
				root:  &tree,
				value: 15000,
			},
			&ABB{
				value: 500,
				left: &ABB{
					value: 250,
				},
				right: &ABB{
					value: 600,
					right: &ABB{
						value: 3000,
						left: &ABB{
							value: 2000,
							left: &ABB{
								value: 1500,
							},
						},
						right: &ABB{
							value: 4000,
							left: &ABB{
								value: 3500,
							},
							right: &ABB{
								value: 4500,
								right: &ABB{
									value: 20000,
									left: &ABB{
										value: 17000,
										right: &ABB{
											value: 18000,
											right: &ABB{
												value: 19000,
											},
										},
									},
									right: &ABB{
										value: 50000,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			setSuccessors(tt.want)

			delete(tt.args.root, tt.args.value)
			if !equal((*tt.args.root), tt.want) {
				t.Errorf("delete() = %v, want %v in equal", tree, tt.want)
			}
			if !successorCheck((*tt.args.root)) {
				t.Errorf("delete() = %v, want %v in successor check", tree, tt.want)
			}
		})
	}
}

// func Test_insert(t *testing.T) {

// 	type args struct {
// 		root  *ABB
// 		value int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *ABB
// 	}{
// 		{
// 			"test1", args{
// 				root: &ABB{
// 					value: 1000,
// 				},
// 				value: 5000,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 				},
// 			},
// 		},
// 		{
// 			"test2", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 					},
// 				},
// 				value: 3000,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"test3", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 						},
// 					},
// 				},
// 				value: 4000,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						right: &ABB{
// 							value: 4000,
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"test4", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 							right: &ABB{
// 								value: 4000,
// 							},
// 						},
// 					},
// 				},
// 				value: 4500,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						right: &ABB{
// 							value: 4000,
// 							right: &ABB{
// 								value: 4500,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"test5", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 							right: &ABB{
// 								value: 4000,
// 								right: &ABB{
// 									value: 4500,
// 								},
// 							},
// 						},
// 					},
// 				},
// 				value: 500,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						right: &ABB{
// 							value: 4000,
// 							right: &ABB{
// 								value: 4500,
// 							},
// 						},
// 					},
// 				},
// 				left: &ABB{
// 					value: 500,
// 				},
// 			},
// 		},
// 		{
// 			"test6", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 							right: &ABB{
// 								value: 4000,
// 								right: &ABB{
// 									value: 4500,
// 								},
// 							},
// 						},
// 					},
// 					left: &ABB{
// 						value: 500,
// 					},
// 				},
// 				value: 250,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						right: &ABB{
// 							value: 4000,
// 							right: &ABB{
// 								value: 4500,
// 							},
// 						},
// 					},
// 				},
// 				left: &ABB{
// 					value: 500,
// 					left: &ABB{
// 						value: 250,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"test7", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 							right: &ABB{
// 								value: 4000,
// 								right: &ABB{
// 									value: 4500,
// 								},
// 							},
// 						},
// 					},
// 					left: &ABB{
// 						value: 500,
// 						left: &ABB{
// 							value: 250,
// 						},
// 					},
// 				},
// 				value: 600,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						right: &ABB{
// 							value: 4000,
// 							right: &ABB{
// 								value: 4500,
// 							},
// 						},
// 					},
// 				},
// 				left: &ABB{
// 					value: 500,
// 					left: &ABB{
// 						value: 250,
// 					},
// 					right: &ABB{
// 						value: 600,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"test8", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 							right: &ABB{
// 								value: 4000,
// 								right: &ABB{
// 									value: 4500,
// 								},
// 							},
// 						},
// 					},
// 					left: &ABB{
// 						value: 500,
// 						left: &ABB{
// 							value: 250,
// 						},
// 						right: &ABB{
// 							value: 600,
// 						},
// 					},
// 				},
// 				value: 20000,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						right: &ABB{
// 							value: 4000,
// 							right: &ABB{
// 								value: 4500,
// 							},
// 						},
// 					},
// 					right: &ABB{
// 						value: 20000,
// 					},
// 				},
// 				left: &ABB{
// 					value: 500,
// 					left: &ABB{
// 						value: 250,
// 					},
// 					right: &ABB{
// 						value: 600,
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"test9", args{
// 				root: &ABB{
// 					value: 1000,
// 					right: &ABB{
// 						value: 5000,
// 						left: &ABB{
// 							value: 3000,
// 							right: &ABB{
// 								value: 4000,
// 								right: &ABB{
// 									value: 4500,
// 								},
// 							},
// 						},
// 						right: &ABB{
// 							value: 20000,
// 						},
// 					},
// 					left: &ABB{
// 						value: 500,
// 						left: &ABB{
// 							value: 250,
// 						},
// 						right: &ABB{
// 							value: 600,
// 						},
// 					},
// 				},
// 				value: 2000,
// 			},
// 			&ABB{
// 				value: 1000,
// 				right: &ABB{
// 					value: 5000,
// 					left: &ABB{
// 						value: 3000,
// 						left: &ABB{
// 							value: 2000,
// 						},
// 						right: &ABB{
// 							value: 4000,
// 							right: &ABB{
// 								value: 4500,
// 							},
// 						},
// 					},
// 					right: &ABB{
// 						value: 20000,
// 					},
// 				},
// 				left: &ABB{
// 					value: 500,
// 					left: &ABB{
// 						value: 250,
// 					},
// 					right: &ABB{
// 						value: 600,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			setSuccessors(tt.args.root)
// 			setSuccessors(tt.want)

// 			got := insert(tt.args.root, tt.args.value)
// 			if !equal(tt.args.root, tt.want) {
// 				t.Errorf("insert() = %v, want %v in equal", got, tt.want)
// 			}
// 			if !successorCheck(tt.args.root) {
// 				t.Errorf("insert() = %v, want %v in successor check", got, tt.want)
// 			}
// 		})
// 	}
// }
