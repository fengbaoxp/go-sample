// Sample test to show how to write a basic unit test.
package ch0911

import (
	"testing"
	"net/http"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload validates the http Get function can to download content.
func TestDownload(t *testing.T) {
	url := "http://www.baidu.com"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking\"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould revceive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
