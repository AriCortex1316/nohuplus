package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.io/nohuplus/internal/commands"
	"github.io/nohuplus/internal/core"
)

func Execute() {
	app := &cli.App{
		Name:  "nohuplus",
		Usage: "Run and manage background processes",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "Show all running background tasks",
				Action: func(c *cli.Context) error {
					if c.Args().Present() {
						return fmt.Errorf("error: 'list' does not take arguments")
					}
					list, err := core.ListTasks()
					if err != nil {
						return err
					}
					for i, v := range list {
						fmt.Printf("[%d] pid=%d cmd=%s args=%v\n", i+1, v.PID, v.Cmd, v.Args)
					}

					return nil
				},
			},
			{
				Name:      "kill",
				Usage:     "Stop a specific task",
				ArgsUsage: "<id>",
				Action: func(c *cli.Context) error {
					if c.NArg() != 1 {
						return fmt.Errorf("error: 'kill' requires exactly one <id>")
					}

					id, err := strconv.Atoi(c.Args().First())
					if err != nil {
						return fmt.Errorf("invalid task id: %q (must be a number)", c.Args().First())
					}
					resolvedID, err := commands.ResolveTaskID(id)
					if err != nil {
						return err
					}
					if err := commands.KillTask(resolvedID); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:      "log",
				Usage:     "View logs of a task",
				ArgsUsage: "<id>",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "f", Usage: "Follow logs"},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 1 {
						return fmt.Errorf("error: 'log' requires exactly one <id>")
					}
					id, err := strconv.Atoi(c.Args().First())
					if err != nil {
						return fmt.Errorf("invalid task id: %q (must be a number)", c.Args().First())
					}
					resolvedID, err := commands.ResolveTaskID(id)
					if err != nil {
						return err
					}
					if c.Bool("f") {
						commands.FollowLog(resolvedID)
					} else {
						commands.ShowLog(resolvedID)
					}
					return nil
				},
			},
		},

		// ✅ 捕获所有未匹配命令 —— 默认行为
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return cli.ShowAppHelp(c)
			}

			cmd := c.Args().First()
			args := c.Args().Slice()[1:]
			task, err := commands.RunCommand(cmd, args)
			if err != nil {
				return err
			}

			fmt.Printf("Started %s (PID %d)\nLog: %s\n", task.Cmd, task.PID, task.Log)
			return nil

		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
