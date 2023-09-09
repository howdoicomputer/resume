package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	src := client.Host().Directory(".")
	tex := client.Container().
		From("mingc/latex").
		WithDirectory("/doc", src).
		WithWorkdir("/doc").
		WithoutMount(".git").
		WithExec([]string{"make"})

	_, err = tex.Export(ctx, "resume.pdf")
	if err != nil {
		panic(err)
	}

	out, err := tex.Stdout(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
