package action

import (
	"bytes"
	"context"
	"flag"
	"os"
	"testing"

	"github.com/justwatchcom/gopass/tests/gptest"
	"github.com/justwatchcom/gopass/utils/ctxutil"
	"github.com/justwatchcom/gopass/utils/out"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestGit(t *testing.T) {
	u := gptest.NewUnitTester(t)
	defer u.Remove()

	ctx := context.Background()
	ctx = ctxutil.WithAlwaysYes(ctx, true)
	ctx = ctxutil.WithInteractive(ctx, false)

	act, err := newMock(ctx, u)
	assert.NoError(t, err)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	stdout = buf
	defer func() {
		out.Stdout = os.Stdout
		stdout = os.Stdout
	}()

	app := cli.NewApp()

	// git init
	fs := flag.NewFlagSet("default", flag.ContinueOnError)
	un := cli.StringFlag{
		Name:  "username",
		Usage: "username",
	}
	assert.NoError(t, un.ApplyWithError(fs))
	ue := cli.StringFlag{
		Name:  "useremail",
		Usage: "useremail",
	}
	assert.NoError(t, ue.ApplyWithError(fs))
	assert.NoError(t, fs.Parse([]string{"--username", "foobar", "--useremail", "foo.bar@example.org"}))
	c := cli.NewContext(app, fs, nil)

	assert.NoError(t, act.GitInit(ctx, c))
	buf.Reset()
}
