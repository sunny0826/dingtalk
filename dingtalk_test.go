package dingtalk

import (
	"testing"
)

func TestWebHook(t *testing.T) {
	webHook := NewWebHook("foo","bar")

	// test send text message
	err := webHook.SendTextMsg("Test text message", false, "")
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	// test send link message
	err = webHook.SendLinkMsg("A link message", "Click me to baidu search", "", "https://www.baidu.com")
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	// test send markdown message
	err = webHook.SendMarkdownMsg("A markdown message", "# This is title \n > Hello World", false, "13800138000")
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	// test send action card message
	err = webHook.SendActionCardMsg("A action card message", "This is a action card message", []string{}, []string{}, true, true)
	if nil == err {
		t.Error("links and titles cannot be null error should be catch!")
	}

	err = webHook.SendActionCardMsg("A action card message", "This is a action card message", []string{"Title 1"}, []string{}, true, true)
	if nil == err {
		t.Error("links and titles length not equal error should be catch!")
	}

	err = webHook.SendActionCardMsg("A action card message", "This is a action card message", []string{"Baidu Search"}, []string{"https://www.baidu.com"}, true, true)
	if err == nil {
		t.Error("token missing error should be catch!")
	}

	// test send link card message
	err = webHook.SendLinkCardMsg([]LinkMsg{{Title: "Hello Bob", MessageURL: "https://www.google.com", PicURL: ""}})
	if nil == err {
		t.Error("token missing error should be catch!")
	}

	t.Log("All test had pass ..")

}