// Code generated by scorify, DO NOT EDIT.
// Import generated from "github.com/scorify/ssh@v1.0.0"

package checks

import (
	ssh "github.com/scorify/ssh"
)

func init() {
	Checks["ssh"] = Check{
		Func:   ssh.Run,
		Schema: ConvertSchema(ssh.Schema{}),
	}
}