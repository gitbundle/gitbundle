// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"net/http"

	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/repo/private"

	"github.com/urfave/cli"
)

func runSendMail(c *cli.Context) error {
	ctx, cancel := installSignals()
	defer cancel()

	setting.LoadFromExisting()

	if err := argsSet(c, "title"); err != nil {
		return err
	}

	subject := c.String("title")
	confirmSkiped := c.Bool("force")
	body := c.String("content")

	if !confirmSkiped {
		if len(body) == 0 {
			fmt.Print("warning: Content is empty")
		}

		fmt.Print("Proceed with sending email? [Y/n] ")
		isConfirmed, err := confirm()
		if err != nil {
			return err
		} else if !isConfirmed {
			fmt.Println("The mail was not sent")
			return nil
		}
	}

	status, message := private.SendEmail(ctx, subject, body, nil)
	if status != http.StatusOK {
		fmt.Printf("error: %s\n", message)
		return nil
	}

	fmt.Printf("Success: %s\n", message)

	return nil
}
