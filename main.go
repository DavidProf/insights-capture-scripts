package main

import (
    "fmt"
    "log"
    "os"
    "errors"
    "insights-capture-scripts/lib"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Commands: []*cli.Command{
            {
                Name:    "records",
                Aliases: []string{"r"},
                Usage:   "manage records",
                Subcommands: []*cli.Command{
                    {
                        Name:  "list",
                        Aliases: []string{"ls", "l"},
                        Usage: "list records paths",
                        Flags: []cli.Flag{
                            &cli.StringFlag{
                                Name:  "bkp-file",
                                Aliases: []string{"f"},
                                Usage: "insights bkp file",
                                Required: true,
                            },
                        },
                        Action: func(cCtx *cli.Context) error {
                            // data, _ := lib.GetDataFromFileAsMap(cCtx.String("bkp-file"))

                            // for key, _ := range data {
                            //     // Each value is an `any` type, that is type asserted as a string
                            //     // fmt.Println(key, value.(string))
                            //     fmt.Println(key)
                            // }
                            return errors.New("not implemented yet")
                        },
                    },
                },
            },
            {
                Name:    "backup-file",
                Aliases: []string{"backup", "b"},
                Usage:   "manage backup file",
                Subcommands: []*cli.Command{
                    {
                        Name:  "decode",
                        Aliases: []string{"d"},
                        Usage: "decode backup file",
                        Flags: []cli.Flag{
                            &cli.StringFlag{
                                Name:  "bkp-file",
                                Aliases: []string{"f"},
                                Usage: "insights bkp file",
                                Required: true,
                            },
                        },
                        Action: func(cCtx *cli.Context) error {
                            data, err := lib.GetDataFromFileAsBytes(cCtx.String("bkp-file"))
                            if err != nil {
                                return err
                            }
                            filename := cCtx.String("bkp-file") + ".json"
                            os.WriteFile(filename, data, 0644)
                            fmt.Println("decoded to " + filename)
                            return nil
                        },
                    },
                },
            },
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
