package reply_test

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/vitorbaraujo/buschebot/reply"
	"github.com/vitorbaraujo/buschebot/storage"
)

func TestMain(m *testing.M) {
	rootPath, err := relativePathToProjectRoot()
	if err != nil {
		panic(err)
	}

	// Replace storage with test database.
	oldDataFile := storage.DataFile
	storage.DataFile = path.Join(rootPath, "storage/data_test.json")
	defer func() { storage.DataFile = oldDataFile }()

	randIntFn := reply.RandInt
	defer func() { reply.RandInt = randIntFn }()
	reply.RandInt = func() int { return 1 } // aka should never reply

	m.Run()
}

func TestReplyMessage_noReply(t *testing.T) {
	tests := []struct{
		name string
		text string
		want *reply.Response
	} {
		{
			name: "noReply",
			text: "hello there",
			want: &reply.Response{},
		},
		{
			name: "notAQuestion",
			text: "why? really.",
			want: &reply.Response{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got, err := reply.GetReply(&reply.MessagePayload{
				Text:   test.text,
			})
			if err != nil {
				t.Fatalf("GetReply returned err = %v", err)
			}

			if diff := pretty.Compare(got, test.want); diff != "" {
			    t.Fatalf("post-GetReply diff: (-got +want)\n%v", diff)
			}
		})
	}
}

func TestReplyMessage_regularQuestion(t *testing.T) {
	tests := []struct{
		name string
		text string
		want *reply.Response
	} {
		{
			name: "regularQuestion",
			text: "oh, really?",
			want: &reply.Response{
				Text: "yes",
			},
		},
		{
			name: "trailingSpaces",
			text: "    oh, really?    ",
			want: &reply.Response{
				Text: "yes",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got, err := reply.GetReply(&reply.MessagePayload{
				Text: test.text,
			})
			if err != nil {
				t.Fatalf("GetReply returned err = %v", err)
			}

			if diff := pretty.Compare(got, test.want); diff != "" {
				t.Fatalf("post-GetReply diff: (-got +want)\n%v", diff)
			}
		})
	}
}

func TestReplyMessage_indagation(t *testing.T) {
	tests := []struct{
		name string
		text string
		want *reply.Response
	} {
		{
			name: "indagation_OK",
			text: "what was the reason?",
			want: &reply.Response{
				Text:  "I don't know",
			},
		},
		{
			name: "upperPrefix",
			text: "WhY is that?",
			want: &reply.Response{
				Text:  "I don't know",
			},
		},
		{
			// unrecognized indagations are treated as regular questions
			name: "unrecognizedIndagation",
			text: "where is the key?",
			want: &reply.Response{
				Text: "yes",
			},
		},
		{
			name: "untrimmedQuestion",
			text: "           why is that?  ",
			want: &reply.Response{
				Text:  "I don't know",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got, err := reply.GetReply(&reply.MessagePayload{
				Text:   test.text,
			})
			if err != nil {
				t.Fatalf("GetReply returned err = %v", err)
			}

			if diff := pretty.Compare(got, test.want); diff != "" {
				t.Fatalf("post-GetReply diff: (-got +want)\n%v", diff)
			}
		})
	}
}


func TestReplyMessage_customReplier(t *testing.T) {
	reply.RegisterReplier(&CustomReplier{})

	want := "custom my message"
	got, err := reply.GetReply(&reply.MessagePayload{Text: "my message"})
	if err != nil {
		t.Fatalf("GetReply returned err = %v", err)
	}

	if  got.Text != want {
		t.Fatalf("GetReply did not use custom replier got %v want %v", got.Text, want)
	}
}

func relativePathToProjectRoot() (string, error) {
	path := "."
	for {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			return "", err
		}

		// The dir containing go.mod is the project root.
		for _, f := range files {
			if f.Name() == "go.mod" {
				return path, nil
			}
		}

		path = "../" + path
	}
}

// CustomReplier implements a custom replier for testing purposes.
type CustomReplier struct{}
func (c CustomReplier) Reply(s string) string {
	return "custom " + s
}

func (CustomReplier) CanReadMessage(*reply.MessagePayload) bool {
	return true
}
