package responder_test

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/vitorbaraujo/buschebot/responder"
)

func TestMain(m *testing.M) {
	randIntFn := responder.RandInt
	defer func() { responder.RandInt = randIntFn }()
	responder.RandInt = func() int { return 1 } // aka should never reply

	m.Run()
}

func TestReplyMessage_noReply(t *testing.T) {
	tests := []struct{
		name string
		text string
		want *responder.Response
	} {
		{
			name: "noReply",
			text: "hello there",
			want: &responder.Response{},
		},
		{
			name: "notAQuestion",
			text: "why? really.",
			want: &responder.Response{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got := responder.ReplyMessage(test.text)

			if diff := pretty.Compare(got, test.want); diff != "" {
			    t.Errorf("post-ReplyMessage diff: (-got +want)\n%v", diff)
			}
		})
	}
}

func TestReplyMessage_regularQuestion(t *testing.T) {
	tests := []struct{
		name string
		text string
		want *responder.Response
	} {
		{
			name: "question",
			text: "eita, verdade?",
			want: &responder.Response{
				Text: "sim",
			},
		},
		{
			name: "untrimmedQuestion",
			text: "    eita, verdade?    ",
			want: &responder.Response{
				Text: "sim",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got := responder.ReplyMessage(test.text)

			if diff := pretty.Compare(got, test.want); diff != "" {
				t.Errorf("post-ReplyMessage diff: (-got +want)\n%v", diff)
			}
		})
	}
}

func TestReplyMessage_indagation(t *testing.T) {
	tests := []struct{
		name string
		text string
		want *responder.Response
	} {
		{
			name: "indagation_OK",
			text: "qual foi o motivo?",
			want: &responder.Response{
				Text:  "sei la",
			},
		},
		{
			name: "upperPrefix",
			text: "Pq foi assim?",
			want: &responder.Response{
				Text:  "sei la",
			},
		},
		{
			name: "untrimmedQuestion",
			text: "           Pq foi assim?  ",
			want: &responder.Response{
				Text:  "sei la",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func (t *testing.T) {
			got := responder.ReplyMessage(test.text)

			if diff := pretty.Compare(got, test.want); diff != "" {
				t.Errorf("post-ReplyMessage diff: (-got +want)\n%v", diff)
			}
		})
	}
}