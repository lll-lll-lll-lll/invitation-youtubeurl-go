package container

import (
	"fmt"
	"testing"

	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
	aes "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()
	testcase := map[string]struct {
		input Input
		want  string
	}{
		"success": {
			input: Input{ID: "testid", Password: "testpassword", URL: "https://www.youtube.com/watch?v=lgKMrOCJTHo"},
			want:  "testidtestpassword",
		},
		"faild": {
			input: Input{ID: "testidfaild", Password: "testpassword", URL: "https://www.youtube.com"},
			want:  "faefewafewew",
		},
	}

	for name, tt := range testcase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			container, err := New(tt.input)
			if err != nil {
				t.Log(err)
			}
			cipher, err := aes.NewAES(container.Key)
			if err != nil {
				t.Log(err.Error())
			}
			got := crypto.Decrypt(cipher, []byte(container.IV), tt.input.String(), []byte(container.EncryptedText))
			fmt.Println(container.IV.ToHexString())
			if string(got) != tt.want {
				t.Logf("got %v, want %v", string(got), tt.want)
			}
		})
	}

}
