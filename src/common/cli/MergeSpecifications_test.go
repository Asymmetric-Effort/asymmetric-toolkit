package cli_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	func() {
		lhs := cli.Specification{
			"test1": cli.Term{Validation: cli.Noop, Help: "1"},
			"test2": cli.Term{Validation: cli.Noop, Help: "2"},
		}
		rhs := cli.Specification{
			"test3": cli.Term{Validation: cli.Noop, Help: "3"},
			"test4": cli.Term{Validation: cli.Noop, Help: "4"},
		}
		expected := cli.Specification{
			"test1": cli.Term{Validation: cli.Noop, Help: "1"},
			"test2": cli.Term{Validation: cli.Noop, Help: "2"},
			"test3": cli.Term{Validation: cli.Noop, Help: "3"},
			"test4": cli.Term{Validation: cli.Noop, Help: "4"},
		}
		cli.MergeMaps(&lhs, &rhs)
		for k, v := range lhs {
			if expected[k].Help != v.Help {
				t.Errorf("Merge failed at %s, %s", k, v.Help)
			}
		}
	}()
	func() {
		lhs := cli.Specification{
			"test1": cli.Term{Validation: cli.Noop, Help: "1"},
			"test2": cli.Term{Validation: cli.Noop, Help: "2"},
		}
		rhs := cli.Specification{
			"test2": cli.Term{Validation: cli.Noop, Help: "3"},
			"test3": cli.Term{Validation: cli.Noop, Help: "4"},
		}
		expected := cli.Specification{
			"test1": cli.Term{Validation: cli.Noop, Help: "1"},
			"test2": cli.Term{Validation: cli.Noop, Help: "3"},
			"test3": cli.Term{Validation: cli.Noop, Help: "4"},
		}
		cli.MergeMaps(&lhs, &rhs)
		for k, v := range lhs {
			if expected[k].Help != v.Help {
				t.Errorf("Merge failed at %s, %s", k, v.Help)
			}
		}
	}()
}
