package reply_test

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/vitorbaraujo/buschebot/reply"
)

func TestMain(m *testing.M) {
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
			got := reply.GetReply(&reply.MessagePayload{
				Text:   test.text,
			})

			if diff := pretty.Compare(got, test.want); diff != "" {
			    t.Errorf("post-GetReply diff: (-got +want)\n%v", diff)
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
			name: "question",
			text: "eita, verdade?",
			want: &reply.Response{
				Text: "sim",
			},
		},
		{
			name: "untrimmedQuestion",
			text: "    eita, verdade?    ",
			want: &reply.Response{
				Text: "sim",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got := reply.GetReply(&reply.MessagePayload{
				Text: test.text,
			})

			if diff := pretty.Compare(got, test.want); diff != "" {
				t.Errorf("post-GetReply diff: (-got +want)\n%v", diff)
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
			text: "qual foi o motivo?",
			want: &reply.Response{
				Text:  "sei la",
			},
		},
		{
			name: "upperPrefix",
			text: "Pq foi assim?",
			want: &reply.Response{
				Text:  "sei la",
			},
		},
		{
			name: "untrimmedQuestion",
			text: "           Pq foi assim?  ",
			want: &reply.Response{
				Text:  "sei la",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got := reply.GetReply(&reply.MessagePayload{
				Text:   test.text,
			})

			if diff := pretty.Compare(got, test.want); diff != "" {
				t.Errorf("post-GetReply diff: (-got +want)\n%v", diff)
			}
		})
	}
}


func TestReplyMessage_customReplier(t *testing.T) {
	reply.RegisterReplier(&CustomReplier{})

	want := "custom my message"
	if got := reply.GetReply(&reply.MessagePayload{Text: "my message"}); got.Text != want {
		t.Fatalf("GetReply did not use custom replier got %v want %v", got.Text, want)
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
