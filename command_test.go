package shamrock

import "testing"

func TestRun(t *testing.T) {
	stdOut, stdErr, err := Run("ls", []string{"-a"})
	if err != nil {
		t.Log(err)
	}
	t.Logf("%s,%s", stdOut, stdErr)
}
