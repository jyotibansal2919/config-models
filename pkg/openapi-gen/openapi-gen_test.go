// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package openapi_gen

import (
	"github.com/openconfig/goyang/pkg/yang"
	"gotest.tools/assert"
	"math"
	"testing"
)

func Test_floatFromNumber(t *testing.T) {

	n1 := floatFromYnumber(yang.Number{
		Value:          12345,
		FractionDigits: 0,
		Negative:       false,
	})
	assert.Equal(t, 12345.0, n1)

	n2 := floatFromYnumber(yang.Number{
		Value:          12345,
		FractionDigits: 2,
		Negative:       false,
	})
	assert.Equal(t, 123.45, n2)

	n3 := floatFromYnumber(yang.Number{
		Value:          12345,
		FractionDigits: 2,
		Negative:       true,
	})
	assert.Equal(t, -123.45, n3)

	n4 := floatFromYnumber(yang.Number{
		Value:          12345,
		FractionDigits: 0,
		Negative:       true,
	})
	assert.Equal(t, -12345.0, n4)
}

// Test the range min..10 | 20..100
func Test_yangRangeDouble(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Max: yang.Number{
			Value:    10,
			Negative: false,
		},
	})
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Negative: false,
			Value:    20,
		},
		Max: yang.Number{
			Negative: false,
			Value:    100,
		},
	})

	min, max, err := yangRange(testRange1, yang.Yuint16)
	assert.NilError(t, err)
	assert.Assert(t, min != nil)
	if min != nil {
		assert.Equal(t, 0.0, *min)
	}
	assert.Assert(t, max != nil)
	if max != nil {
		assert.Equal(t, 100.0, *max)
	}
}

// Test the range min..10 | 20..max of uint16
func Test_yangRangeDoubleUint8(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Max: yang.Number{
			Value:    10,
			Negative: false,
		},
	})
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Negative: false,
			Value:    20,
		},
		Max: yang.Number{
			Negative: false,
			Value:    65535,
		},
	})

	min, max, err := yangRange(testRange1, yang.Yuint16)
	assert.NilError(t, err)
	assert.Assert(t, min != nil)
	if min != nil {
		assert.Equal(t, 0.0, *min)
	}
	assert.Assert(t, max != nil)
	if max != nil {
		assert.Equal(t, 65535.0, *max)
	}
}

// Test the range -0.02..0.002
func Test_yangRangeDecimal(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Negative:       true,
			Value:          201,
			FractionDigits: 2,
		},
		Max: yang.Number{
			Negative:       false,
			Value:          2005,
			FractionDigits: 2,
		},
	})

	min, max, err := yangRange(testRange1, yang.Ydecimal64)
	assert.NilError(t, err)
	assert.Assert(t, min != nil)
	if min != nil {
		assert.Equal(t, -2.0100000000000002, *min)
	}
	assert.Assert(t, max != nil)
	if max != nil {
		assert.Equal(t, 20.05, *max)
	}
}

// Test the min and max of int32 - range is not needed then
func Test_yangRangeMinMaxInt32(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Negative:       true,
			Value:          math.MinInt32 * -1,
			FractionDigits: 0,
		},
		Max: yang.Number{
			Negative:       false,
			Value:          math.MaxInt32,
			FractionDigits: 0,
		},
	})

	min, max, err := yangRange(testRange1, yang.Yint32)
	assert.NilError(t, err)
	assert.Assert(t, min == nil)
	assert.Assert(t, max == nil)
}

// Test the min and max of int64 - range is not needed
func Test_yangRangeMinMaxInt64(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Negative:       true,
			Value:          math.MinInt64 * -1,
			FractionDigits: 0,
		},
		Max: yang.Number{
			Negative:       false,
			Value:          math.MaxInt64,
			FractionDigits: 0,
		},
	})

	min, max, err := yangRange(testRange1, yang.Yint64)
	assert.NilError(t, err)
	assert.Assert(t, min == nil)
	assert.Assert(t, max == nil)
}

// Test the min and max of uint32 - range is needed
func Test_yangRangeMinMaxUint32(t *testing.T) {
	testRange1 := make(yang.YangRange, 0)
	testRange1 = append(testRange1, yang.YRange{
		Min: yang.Number{
			Negative:       false,
			Value:          0,
			FractionDigits: 0,
		},
		Max: yang.Number{
			Negative:       false,
			Value:          math.MaxUint32,
			FractionDigits: 0,
		},
	})

	min, max, err := yangRange(testRange1, yang.Yuint32)
	assert.NilError(t, err)
	assert.Assert(t, min != nil)
	if min != nil {
		assert.Equal(t, 0.0, *min)
	}
	assert.Assert(t, max == nil)
}

func Test_pathToSchemaName(t *testing.T) {
	schName1 := pathToSchemaName("/qos-profile/qos-profile/{id}/arp")
	assert.Equal(t, "/qos-profile/qos-profile/arp", schName1)

	schName2 := pathToSchemaName("/subscriber/ue/{id}/profiles/access-profile/{access-profile}")
	assert.Equal(t, "/subscriber/ue/profiles/access-profile", schName2)
}

func Test_newPathItem(t *testing.T) {

	targetParameter = targetParam("targettest")

	testDirEntry := yang.Entry{
		Config: yang.TSTrue,
		Parent: &yang.Entry{},
		Type: &yang.YangType{
			Name: "Test1",
		},
	}

	pathItem := newPathItem(&testDirEntry, "/test-1/test-2/{id}/test-3/{id}/test-4",
		"/parent-1/{parent1-name}/parent-2/{parent2-name}", pathTypeContainer, "targettest")
	assert.Assert(t, pathItem != nil)
	if pathItem != nil {
		g := pathItem.Get
		assert.Assert(t, g != nil)
		if g != nil {
			assert.Equal(t, "GET /test-1/test-2/{id}/test-3/{id}/test-4 Container", g.Summary)
		}
		assert.Assert(t, pathItem.Post != nil)
		assert.Equal(t, "POST /test-1/test-2/{id}/test-3/{id}/test-4", pathItem.Post.Summary)
		assert.Equal(t, "postTest-1_Test-2_Test-3_Test-4", pathItem.Post.OperationID)

		assert.Assert(t, pathItem.Delete != nil)
		assert.Equal(t, "DELETE /test-1/test-2/{id}/test-3/{id}/test-4", pathItem.Delete.Summary)
		assert.Equal(t, "deleteTest-1_Test-2_Test-3_Test-4", pathItem.Delete.OperationID)

		assert.Equal(t, 3, len(pathItem.Parameters))
		for _, p := range pathItem.Parameters {
			switch name := p.Value.Name; name {
			case "targettest":
				assert.Equal(t, "targettest (target in onos-config)", p.Value.Description)
			case "parent1-name":
				assert.Equal(t, "key {parent1-name}", p.Value.Description)
			case "parent2-name":
				assert.Equal(t, "key {parent2-name}", p.Value.Description)
			default:
				t.Errorf("unexpected parameter %s", name)
			}
		}
	}
}

func Test_buildSchemaIntegerLeaf(t *testing.T) {

	targetParameter = targetParam("targettest")

	testLeaf1 := yang.Entry{
		Name:        "Leaf1",
		Description: "Leaf1 Description",
		Config:      yang.TSTrue,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yint16,
			Range: []yang.YRange{
				{Min: yang.Number{
					Negative: false,
					Value:    1,
				},
					Max: yang.Number{
						Negative: false,
						Value:    10,
					},
				},
			},
		},
	}

	testLeaf2 := yang.Entry{
		Name:        "Leaf2",
		Description: "Leaf2 Description",
		Config:      yang.TSTrue,
		Type: &yang.YangType{
			Kind:    yang.Ystring,
			Pattern: []string{"^[abc]*"},
			Default: "test default",
			Length: []yang.YRange{
				{
					Min: yang.Number{
						Negative: false,
						Value:    20,
					},
					Max: yang.Number{
						Negative: false,
						Value:    30,
					},
				},
			},
		},
	}

	testDirEntry := yang.Entry{
		Name:   "Test1",
		Parent: &yang.Entry{},
		Kind:   yang.DirectoryEntry,
		Config: yang.TSTrue,
		Dir:    make(map[string]*yang.Entry),
		Type: &yang.YangType{
			Name: "Test1",
		},
	}
	testLeaf1.Parent = &testDirEntry
	testLeaf2.Parent = &testDirEntry
	testDirEntry.Dir["leaf1"] = &testLeaf1
	testDirEntry.Dir["leaf2"] = &testLeaf2

	hasLeafref := false
	paths, components, err := buildSchema(&testDirEntry, yang.TSUnset, "/test", "targettest", &hasLeafref)
	assert.NilError(t, err)
	assert.Equal(t, len(paths), 0)
	assert.Equal(t, len(components.Schemas), 2)
	s, ok := components.Schemas["Test_Leaf1"]
	assert.Assert(t, ok, "expecting Test_Leaf1")
	assert.Equal(t, "Leaf1", s.Value.Title)
	assert.Equal(t, "Leaf1 Description", s.Value.Description)
	assert.Equal(t, "integer", s.Value.Type)
	assert.Equal(t, 1, len(s.Value.Required))
	assert.Equal(t, 1.0, *s.Value.Min)
	assert.Equal(t, 10.0, *s.Value.Max)

	s, ok = components.Schemas["Test_Leaf2"]
	assert.Assert(t, ok, "expecting Test_Leaf2")
	assert.Equal(t, "Leaf2", s.Value.Title)
	assert.Equal(t, "Leaf2 Description", s.Value.Description)
	assert.Equal(t, "string", s.Value.Type)
	assert.Equal(t, "^[abc]*", s.Value.Pattern)
	assert.Equal(t, "test default", s.Value.Default)
	assert.Equal(t, uint64(20), s.Value.MinLength)
	assert.Equal(t, uint64(30), *s.Value.MaxLength)
}

func Test_buildSchemaLeafList(t *testing.T) {

	targetParameter = targetParam("targettest")

	testDirEntry := yang.Entry{
		Name:     "parent-list",
		Kind:     yang.DirectoryEntry,
		Parent:   &yang.Entry{},
		ListAttr: yang.NewDefaultListAttr(),
		Dir:      make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	testList1 := yang.Entry{
		Name:     "list1",
		Kind:     yang.DirectoryEntry,
		Parent:   &testDirEntry,
		ListAttr: yang.NewDefaultListAttr(),
		Dir:      make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	testContainer1 := yang.Entry{
		Name:   "container1",
		Kind:   yang.DirectoryEntry,
		Parent: &testList1,
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	containerIntRefLeafList := yang.Entry{
		Name:   "cont-int-ref-leaf-list",
		Kind:   yang.LeafEntry,
		Parent: &testContainer1,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../../../leaf-uint16",
		},
		ListAttr: yang.NewDefaultListAttr(),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	containerStrRefLeafList := yang.Entry{
		Name:   "cont-str-ref-leaf-list",
		Kind:   yang.LeafEntry,
		Parent: &testContainer1,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../../../leaf-string",
		},
		ListAttr: yang.NewDefaultListAttr(),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	listIntRefLeafList := yang.Entry{
		Name:   "list-int-ref-leaf-list",
		Kind:   yang.LeafEntry,
		Parent: &testList1,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../../leaf-uint16",
		},
		ListAttr: yang.NewDefaultListAttr(),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	listStrRefLeafList := yang.Entry{
		Name:   "list-str-ref-leaf-list",
		Kind:   yang.LeafEntry,
		Parent: &testList1,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../../leaf-string",
		},
		ListAttr: yang.NewDefaultListAttr(),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	leafUint16 := yang.Entry{
		Name:   "leaf-uint16",
		Kind:   yang.LeafEntry,
		Parent: &testDirEntry,
		Type:   &yang.YangType{Kind: yang.Yuint16},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	leafString := yang.Entry{
		Name:   "leaf-string",
		Kind:   yang.LeafEntry,
		Parent: &testDirEntry,
		Type:   &yang.YangType{Kind: yang.Ystring},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	testContainer1.Dir["cont-int-ref-leaf-list"] = &containerIntRefLeafList
	testContainer1.Dir["cont-str-ref-leaf-list"] = &containerStrRefLeafList

	testList1.Dir["list-int-ref-leaf-list"] = &listIntRefLeafList
	testList1.Dir["list-str-ref-leaf-list"] = &listStrRefLeafList
	testList1.Dir["container1"] = &testContainer1

	testDirEntry.Dir["leaf-uint16"] = &leafUint16
	testDirEntry.Dir["leaf-string"] = &leafString
	testDirEntry.Dir["list1"] = &testList1

	hasLeafref := false
	paths, components, err := buildSchema(&testDirEntry, yang.TSUnset, "/test", "targettest", &hasLeafref)
	assert.NilError(t, err)
	assert.Equal(t, len(paths), 7)
	assert.Equal(t, len(components.Schemas), 9)

	// Assert the leaf list with leaf ref to integer inside a Container
	s, ok := components.Schemas["Test_List1_Container1_Cont-int-ref-leaf-list"]
	assert.Assert(t, ok, "expecting Test_List1_Container1_Cont-int-ref-leaf-list")
	assert.Equal(t, "cont-int-ref-leaf-list", s.Value.Title)
	assert.Equal(t, "array", s.Value.Type)
	assert.Equal(t, "integer", s.Value.Items.Value.Type)

	// Assert the leaf list with leaf ref to string inside a Container
	s, ok = components.Schemas["Test_List1_Container1_Cont-str-ref-leaf-list"]
	assert.Assert(t, ok, "expecting Test_List1_Container1_Cont-str-ref-leaf-list")
	assert.Equal(t, "cont-str-ref-leaf-list", s.Value.Title)
	assert.Equal(t, "array", s.Value.Type)
	assert.Equal(t, "string", s.Value.Items.Value.Type)

	// Assert the leaf list with leaf ref to integer inside a List
	s, ok = components.Schemas["Test_List1_List-int-ref-leaf-list"]
	assert.Assert(t, ok, "expecting Test_List1_List-int-ref-leaf-list")
	assert.Equal(t, "list-int-ref-leaf-list", s.Value.Title)
	assert.Equal(t, "array", s.Value.Type)
	assert.Equal(t, "integer", s.Value.Items.Value.Type)

	// Assert the leaf list with leaf ref to string inside a List
	s, ok = components.Schemas["Test_List1_List-str-ref-leaf-list"]
	assert.Assert(t, ok, "expecting Test_List1_List-str-ref-leaf-list")
	assert.Equal(t, "list-str-ref-leaf-list", s.Value.Title)
	assert.Equal(t, "array", s.Value.Type)
	assert.Equal(t, "string", s.Value.Items.Value.Type)

}

func Test_walkPath(t *testing.T) {

	//A tree like:
	// Device
	//  |- test-grand-parent
	//      |- test-parent
	//      |    |- Leaf1   type int16
	//      |    |- Leaf2   Ref to ../Leaf1
	//      |- test-uncle
	//           |- cousin   Ref to /Test:test-grand-parent/Test:test-parent/Test:Leaf1
	//           |- cousin2   Ref to ../cousin
	//           |- cousin2   Ref to ../../test-parent/Leaf2
	testDevice := yang.Entry{
		Name:   "Device",
		Config: yang.TSTrue,
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testGrandParent := yang.Entry{
		Name:   "test-grand-parent",
		Config: yang.TSTrue,
		Parent: &testDevice,
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testParent := yang.Entry{
		Name:   "test-parent",
		Config: yang.TSTrue,
		Parent: &testGrandParent,
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testUncle := yang.Entry{
		Name:   "test-uncle",
		Config: yang.TSTrue,
		Parent: &testGrandParent,
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testLeaf1 := yang.Entry{
		Name:        "Leaf1",
		Description: "Leaf1 Description",
		Config:      yang.TSTrue,
		Parent:      &testParent,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yint16,
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testLeaf2 := yang.Entry{
		Name:        "Leaf2",
		Description: "Leaf2 Description",
		Config:      yang.TSTrue,
		Parent:      &testParent,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../Leaf1",
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testCousin := yang.Entry{
		Name:        "cousin",
		Description: "Cousin Description",
		Config:      yang.TSTrue,
		Parent:      &testUncle,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "/Test:test-grand-parent/Test:test-parent/Test:Leaf1",
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testCousin2 := yang.Entry{
		Name:        "cousin2",
		Description: "Cousin2 Description",
		Config:      yang.TSTrue,
		Parent:      &testUncle,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../cousin",
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}
	testCousin3 := yang.Entry{
		Name:        "cousin3",
		Description: "Cousin3 Description",
		Config:      yang.TSTrue,
		Parent:      &testUncle,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../../test-parent/Leaf2",
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	testDevice.Dir["test-grand-parent"] = &testGrandParent

	testGrandParent.Dir["test-parent"] = &testParent
	testGrandParent.Dir["test-uncle"] = &testUncle

	testParent.Dir["Leaf1"] = &testLeaf1
	testParent.Dir["Leaf2"] = &testLeaf2

	testUncle.Dir["cousin"] = &testCousin
	testUncle.Dir["cousin2"] = &testCousin2
	testUncle.Dir["cousin3"] = &testCousin3

	kindLeaf2 := resolveLeafRefType(&testLeaf2)
	assert.Equal(t, "int16", kindLeaf2.String())

	kindCousin1 := resolveLeafRefType(&testCousin)
	assert.Equal(t, "int16", kindCousin1.String())

	kindCousin2 := resolveLeafRefType(&testCousin2)
	assert.Equal(t, "int16", kindCousin2.String())

	kindCousin3 := resolveLeafRefType(&testCousin3)
	assert.Equal(t, "int16", kindCousin3.String())
}

func Test_ReadOnly(t *testing.T) {
	targetParameter = targetParam("targettest")

	// Configurable parent
	test1Parent := yang.Entry{
		Name:   "test-parent",
		Kind:   yang.DirectoryEntry,
		Config: yang.TSTrue,
		Parent: &yang.Entry{},
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	// Configurable leaf
	test1Leaf1 := yang.Entry{
		Name:        "Leaf1",
		Description: "Leaf1 Description",
		Config:      yang.TSTrue,
		Parent:      &test1Parent,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yint16,
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	// Unconfigurable leaf
	test1Leaf2 := yang.Entry{
		Name:        "Leaf2",
		Description: "Leaf2 Description",
		Config:      yang.TSFalse,
		Parent:      &test1Parent,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../Leaf1",
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	test1Parent.Dir["Leaf1"] = &test1Leaf1
	test1Parent.Dir["Leaf2"] = &test1Leaf2

	hasLeafref := false
	paths, components, err := buildSchema(&test1Parent, yang.TSUnset, "/test", "targettest", &hasLeafref)
	assert.NilError(t, err)
	assert.Equal(t, len(paths), 1)
	assert.Equal(t, len(components.Schemas), 2)

	// Assert the leaf list with leaf ref to integer inside a Container
	s := components.Schemas["Test_Leaf1"]
	assert.Equal(t, false, s.Value.ReadOnly)
	s = components.Schemas["Test_Leaf2"]
	assert.Equal(t, true, s.Value.ReadOnly)

}

func Test_Parent_ReadOnly(t *testing.T) {
	targetParameter = targetParam("targettest")

	// Unconfigurable parent
	testParent := yang.Entry{
		Name:   "test-parent",
		Kind:   yang.DirectoryEntry,
		Config: yang.TSFalse,
		Parent: &yang.Entry{},
		Dir:    make(map[string]*yang.Entry),
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	// Unset Config state
	testLeaf1 := yang.Entry{
		Name:        "Leaf1",
		Description: "Leaf1 Description",
		Parent:      &testParent,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yint16,
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	// Unconfigurable leaf
	testLeaf2 := yang.Entry{
		Name:        "Leaf2",
		Description: "Leaf2 Description",
		Config:      yang.TSFalse,
		Parent:      &testParent,
		Mandatory:   yang.TSTrue,
		Type: &yang.YangType{
			Kind: yang.Yleafref,
			Path: "../Leaf1",
		},
		Prefix: &yang.Value{
			Name: "Test",
		},
	}

	testParent.Dir["Leaf1"] = &testLeaf1
	testParent.Dir["Leaf2"] = &testLeaf2

	hasLeafref := false
	paths, components, err := buildSchema(&testParent, yang.TSUnset, "/test", "targettest", &hasLeafref)
	assert.NilError(t, err)
	assert.Equal(t, len(paths), 1)
	assert.Equal(t, len(components.Schemas), 2)

	// Assert the leaf list with leaf ref to integer inside a Container
	s := components.Schemas["Test_Leaf1"]
	assert.Equal(t, true, s.Value.ReadOnly)
	s = components.Schemas["Test_Leaf2"]
	assert.Equal(t, true, s.Value.ReadOnly)

}
