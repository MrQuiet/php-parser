package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestSimpleClass(t *testing.T) {
	src := `<? class foo{ }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Stmts:     []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestAbstractClass(t *testing.T) {
	src := `<? abstract class foo{ }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "abstract"},
				},
				Stmts: []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestClassExtends(t *testing.T) {
	src := `<? final class foo extends bar { }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "final"},
				},
				Extends: &name.Name{
					Parts: []node.Node{
						&name.NamePart{Value: "bar"},
					},
				},
				Stmts: []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestClassImplement(t *testing.T) {
	src := `<? final class foo implements bar { }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "final"},
				},
				Implements: []node.Node{
					&name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "bar"},
						},
					},
				},
				Stmts: []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestClassImplements(t *testing.T) {
	src := `<? final class foo implements bar, baz { }`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Class{
				ClassName: &node.Identifier{Value: "foo"},
				Modifiers: []node.Node{
					&node.Identifier{Value: "final"},
				},
				Implements: []node.Node{
					&name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "bar"},
						},
					},
					&name.Name{
						Parts: []node.Node{
							&name.NamePart{Value: "baz"},
						},
					},
				},
				Stmts: []node.Node{},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}

func TestAnonimousClass(t *testing.T) {
	src := `<? new class() extends foo implements bar, baz { };`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.New{
					Class: &stmt.Class{
						ArgumentList: &node.ArgumentList{},
						Extends: &name.Name{
							Parts: []node.Node{
								&name.NamePart{Value: "foo"},
							},
						},
						Implements: []node.Node{
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "bar"},
								},
							},
							&name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "baz"},
								},
							},
						},
						Stmts: []node.Node{},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)
}
