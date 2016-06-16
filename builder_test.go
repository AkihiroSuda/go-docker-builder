package builder

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"testing"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
)

var (
	fSrcPath        string
	fDockerfilePath string
)

func init() {
	flag.StringVar(&fSrcPath, "builder.src", "", "srcPath")
	flag.StringVar(&fDockerfilePath, "builder.dockerfile", "Dockerfile", "dockerfilePath")
	flag.Parse()
}

func TestCreateTarStream(t *testing.T) {
	if fSrcPath == "" {
		t.Skipf("srcPath is not set. Please run `go test` with the -builder.src flag.")
	}
	testCreateTarStream(fSrcPath, fDockerfilePath, t)
}

func ensureDockerClient(t *testing.T) *client.Client {
	c, err := client.NewEnvClient()
	if err != nil {
		t.Skipf("DOCKER_HOST not set?: %v", err)
	}
	return c
}

func testCreateTarStream(srcPath, dockerfilePath string, t *testing.T) {
	// t.Logf does not print a line until the test completion.
	// So we use fmt.Printf here.
	fmt.Printf("srcPath=%q, dockerfilePath=%q\n", srcPath, dockerfilePath)
	c := ensureDockerClient(t)
	tarReader, err := CreateTarStream(srcPath, dockerfilePath)
	if err != nil {
		t.Fatal(err)
	}
	netCtx := context.Background()
	opts := types.ImageBuildOptions{
		Dockerfile: dockerfilePath,
	}
	buildResp, err := c.ImageBuild(netCtx,
		tarReader, opts)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("OSType=%q\n", buildResp.OSType)
	bodyReader := bufio.NewReader(buildResp.Body)
	for {
		line, _, err := bodyReader.ReadLine()
		fmt.Printf("build: %q\n", string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println("done")
}
