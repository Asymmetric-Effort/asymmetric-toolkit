package cli

import "testing"

func TestSpecification_EnsureUniqueFlagId(t *testing.T) {
	const simpleTestData="testdata"
	func() {
		var o Specification = Specification{
			Author:      simpleTestData,
			AuthorEmail: simpleTestData,
			ProgramName: simpleTestData,
			Copyright:   simpleTestData,
			Version:     simpleTestData,
			Argument: map[string]ArgumentDescriptor{
				"test1": {
					1,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
				"test2": {
					2,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
				"test3": {
					3,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
				"test4": {
					4,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
			},
		}
		if err := o.EnsureUniqueFlagId(); err != nil {
			t.Fatal(err)
		}
	}()
	func() {
		defer func(){recover()}()
		var o Specification = Specification{
			Author:      simpleTestData,
			AuthorEmail: simpleTestData,
			ProgramName: simpleTestData,
			Copyright:   simpleTestData,
			Version:     simpleTestData,
			Argument: map[string]ArgumentDescriptor{
				"test1": {
					1,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
				"test2": {
					2,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
				"test3": {
					3,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
				"test4": {
					3,
					None,
					simpleTestData,
					simpleTestData,
					ParserNoop(),
					ExpectFlag,
				},
			},
		}
		err := o.EnsureUniqueFlagId()
		if err != nil {
			panic(err)
		}
	}()
}
