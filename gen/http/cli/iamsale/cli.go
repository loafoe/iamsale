// Code generated by goa v3.13.2, DO NOT EDIT.
//
// iamsale HTTP client CLI support package
//
// Command:
// $ goa gen github.com/loafoe/iamsale/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	accountc "github.com/loafoe/iamsale/gen/http/account/client"
	aggregatec "github.com/loafoe/iamsale/gen/http/aggregate/client"
	testc "github.com/loafoe/iamsale/gen/http/test/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `test test
aggregate (accounts|groups)
account (create|get|update|delete|group-add|group-remove)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` test test` + "\n" +
		os.Args[0] + ` aggregate accounts --username "Minima et veniam voluptates quaerat minima sunt." --password "Consequuntur impedit necessitatibus neque nesciunt maxime."` + "\n" +
		os.Args[0] + ` account create --body '{
      "email": "amos@hostedzonehere.com",
      "login": "amos",
      "name": "Amos Burton"
   }' --username "Vitae ut modi placeat ut quisquam." --password "Dolores ut doloremque ullam omnis amet."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		testFlags = flag.NewFlagSet("test", flag.ContinueOnError)

		testTestFlags = flag.NewFlagSet("test", flag.ExitOnError)

		aggregateFlags = flag.NewFlagSet("aggregate", flag.ContinueOnError)

		aggregateAccountsFlags        = flag.NewFlagSet("accounts", flag.ExitOnError)
		aggregateAccountsUsernameFlag = aggregateAccountsFlags.String("username", "REQUIRED", "")
		aggregateAccountsPasswordFlag = aggregateAccountsFlags.String("password", "REQUIRED", "")

		aggregateGroupsFlags        = flag.NewFlagSet("groups", flag.ExitOnError)
		aggregateGroupsUsernameFlag = aggregateGroupsFlags.String("username", "REQUIRED", "")
		aggregateGroupsPasswordFlag = aggregateGroupsFlags.String("password", "REQUIRED", "")

		accountFlags = flag.NewFlagSet("account", flag.ContinueOnError)

		accountCreateFlags        = flag.NewFlagSet("create", flag.ExitOnError)
		accountCreateBodyFlag     = accountCreateFlags.String("body", "REQUIRED", "")
		accountCreateUsernameFlag = accountCreateFlags.String("username", "REQUIRED", "")
		accountCreatePasswordFlag = accountCreateFlags.String("password", "REQUIRED", "")

		accountGetFlags         = flag.NewFlagSet("get", flag.ExitOnError)
		accountGetAccountIDFlag = accountGetFlags.String("account-id", "REQUIRED", "Account ID")
		accountGetUsernameFlag  = accountGetFlags.String("username", "REQUIRED", "Username")
		accountGetPasswordFlag  = accountGetFlags.String("password", "REQUIRED", "Password")

		accountUpdateFlags         = flag.NewFlagSet("update", flag.ExitOnError)
		accountUpdateBodyFlag      = accountUpdateFlags.String("body", "REQUIRED", "")
		accountUpdateAccountIDFlag = accountUpdateFlags.String("account-id", "REQUIRED", "Account ID")
		accountUpdateUsernameFlag  = accountUpdateFlags.String("username", "REQUIRED", "Username")
		accountUpdatePasswordFlag  = accountUpdateFlags.String("password", "REQUIRED", "Password")

		accountDeleteFlags         = flag.NewFlagSet("delete", flag.ExitOnError)
		accountDeleteAccountIDFlag = accountDeleteFlags.String("account-id", "REQUIRED", "Account ID")
		accountDeleteUsernameFlag  = accountDeleteFlags.String("username", "REQUIRED", "")
		accountDeletePasswordFlag  = accountDeleteFlags.String("password", "REQUIRED", "")

		accountGroupAddFlags         = flag.NewFlagSet("group-add", flag.ExitOnError)
		accountGroupAddAccountIDFlag = accountGroupAddFlags.String("account-id", "REQUIRED", "Account ID")
		accountGroupAddGroupIDFlag   = accountGroupAddFlags.String("group-id", "REQUIRED", "Group ID")
		accountGroupAddUsernameFlag  = accountGroupAddFlags.String("username", "REQUIRED", "")
		accountGroupAddPasswordFlag  = accountGroupAddFlags.String("password", "REQUIRED", "")

		accountGroupRemoveFlags         = flag.NewFlagSet("group-remove", flag.ExitOnError)
		accountGroupRemoveAccountIDFlag = accountGroupRemoveFlags.String("account-id", "REQUIRED", "Account ID")
		accountGroupRemoveGroupIDFlag   = accountGroupRemoveFlags.String("group-id", "REQUIRED", "Group ID")
		accountGroupRemoveUsernameFlag  = accountGroupRemoveFlags.String("username", "REQUIRED", "")
		accountGroupRemovePasswordFlag  = accountGroupRemoveFlags.String("password", "REQUIRED", "")
	)
	testFlags.Usage = testUsage
	testTestFlags.Usage = testTestUsage

	aggregateFlags.Usage = aggregateUsage
	aggregateAccountsFlags.Usage = aggregateAccountsUsage
	aggregateGroupsFlags.Usage = aggregateGroupsUsage

	accountFlags.Usage = accountUsage
	accountCreateFlags.Usage = accountCreateUsage
	accountGetFlags.Usage = accountGetUsage
	accountUpdateFlags.Usage = accountUpdateUsage
	accountDeleteFlags.Usage = accountDeleteUsage
	accountGroupAddFlags.Usage = accountGroupAddUsage
	accountGroupRemoveFlags.Usage = accountGroupRemoveUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "test":
			svcf = testFlags
		case "aggregate":
			svcf = aggregateFlags
		case "account":
			svcf = accountFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "test":
			switch epn {
			case "test":
				epf = testTestFlags

			}

		case "aggregate":
			switch epn {
			case "accounts":
				epf = aggregateAccountsFlags

			case "groups":
				epf = aggregateGroupsFlags

			}

		case "account":
			switch epn {
			case "create":
				epf = accountCreateFlags

			case "get":
				epf = accountGetFlags

			case "update":
				epf = accountUpdateFlags

			case "delete":
				epf = accountDeleteFlags

			case "group-add":
				epf = accountGroupAddFlags

			case "group-remove":
				epf = accountGroupRemoveFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "test":
			c := testc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "test":
				endpoint = c.Test()
				data = nil
			}
		case "aggregate":
			c := aggregatec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "accounts":
				endpoint = c.Accounts()
				data, err = aggregatec.BuildAccountsPayload(*aggregateAccountsUsernameFlag, *aggregateAccountsPasswordFlag)
			case "groups":
				endpoint = c.Groups()
				data, err = aggregatec.BuildGroupsPayload(*aggregateGroupsUsernameFlag, *aggregateGroupsPasswordFlag)
			}
		case "account":
			c := accountc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = accountc.BuildCreatePayload(*accountCreateBodyFlag, *accountCreateUsernameFlag, *accountCreatePasswordFlag)
			case "get":
				endpoint = c.Get()
				data, err = accountc.BuildGetPayload(*accountGetAccountIDFlag, *accountGetUsernameFlag, *accountGetPasswordFlag)
			case "update":
				endpoint = c.Update()
				data, err = accountc.BuildUpdatePayload(*accountUpdateBodyFlag, *accountUpdateAccountIDFlag, *accountUpdateUsernameFlag, *accountUpdatePasswordFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = accountc.BuildDeletePayload(*accountDeleteAccountIDFlag, *accountDeleteUsernameFlag, *accountDeletePasswordFlag)
			case "group-add":
				endpoint = c.GroupAdd()
				data, err = accountc.BuildGroupAddPayload(*accountGroupAddAccountIDFlag, *accountGroupAddGroupIDFlag, *accountGroupAddUsernameFlag, *accountGroupAddPasswordFlag)
			case "group-remove":
				endpoint = c.GroupRemove()
				data, err = accountc.BuildGroupRemovePayload(*accountGroupRemoveAccountIDFlag, *accountGroupRemoveGroupIDFlag, *accountGroupRemoveUsernameFlag, *accountGroupRemovePasswordFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// testUsage displays the usage of the test command and its subcommands.
func testUsage() {
	fmt.Fprintf(os.Stderr, `Test service availability
Usage:
    %[1]s [globalflags] test COMMAND [flags]

COMMAND:
    test: Test implements test.

Additional help:
    %[1]s test COMMAND --help
`, os.Args[0])
}
func testTestUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] test test

Test implements test.

Example:
    %[1]s test test
`, os.Args[0])
}

// aggregateUsage displays the usage of the aggregate command and its
// subcommands.
func aggregateUsage() {
	fmt.Fprintf(os.Stderr, `Integration APIs
Usage:
    %[1]s [globalflags] aggregate COMMAND [flags]

COMMAND:
    accounts: Account aggregation. Returns all known accounts
    groups: Group aggregation. Returns list of all known groups

Additional help:
    %[1]s aggregate COMMAND --help
`, os.Args[0])
}
func aggregateAccountsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] aggregate accounts -username STRING -password STRING

Account aggregation. Returns all known accounts
    -username STRING: 
    -password STRING: 

Example:
    %[1]s aggregate accounts --username "Minima et veniam voluptates quaerat minima sunt." --password "Consequuntur impedit necessitatibus neque nesciunt maxime."
`, os.Args[0])
}

func aggregateGroupsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] aggregate groups -username STRING -password STRING

Group aggregation. Returns list of all known groups
    -username STRING: 
    -password STRING: 

Example:
    %[1]s aggregate groups --username "Iusto perferendis nihil quam voluptas." --password "Tenetur eaque veritatis."
`, os.Args[0])
}

// accountUsage displays the usage of the account command and its subcommands.
func accountUsage() {
	fmt.Fprintf(os.Stderr, `Accounts APIs
Usage:
    %[1]s [globalflags] account COMMAND [flags]

COMMAND:
    create: Creates an account. Note that the client may choose to create a shadow account or hold the account in a temporary store until the actual account materializes.
    get: Get account details
    update: Update account details
    delete: Delete an account
    group-add: Add an account to a group
    group-remove: Remove an account from a group

Additional help:
    %[1]s account COMMAND --help
`, os.Args[0])
}
func accountCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account create -body JSON -username STRING -password STRING

Creates an account. Note that the client may choose to create a shadow account or hold the account in a temporary store until the actual account materializes.
    -body JSON: 
    -username STRING: 
    -password STRING: 

Example:
    %[1]s account create --body '{
      "email": "amos@hostedzonehere.com",
      "login": "amos",
      "name": "Amos Burton"
   }' --username "Vitae ut modi placeat ut quisquam." --password "Dolores ut doloremque ullam omnis amet."
`, os.Args[0])
}

func accountGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account get -account-id STRING -username STRING -password STRING

Get account details
    -account-id STRING: Account ID
    -username STRING: Username
    -password STRING: Password

Example:
    %[1]s account get --account-id "18ee082f-1d61-40d3-b8a2-f4eee67cefff" --username "Rem velit." --password "Illo accusamus."
`, os.Args[0])
}

func accountUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account update -body JSON -account-id STRING -username STRING -password STRING

Update account details
    -body JSON: 
    -account-id STRING: Account ID
    -username STRING: Username
    -password STRING: Password

Example:
    %[1]s account update --body '{
      "status": "disabled"
   }' --account-id "18ee082f-1d61-40d3-b8a2-f4eee67cefff" --username "Earum voluptatem dolorem dolore deserunt aut." --password "Temporibus et tempora et."
`, os.Args[0])
}

func accountDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account delete -account-id STRING -username STRING -password STRING

Delete an account
    -account-id STRING: Account ID
    -username STRING: 
    -password STRING: 

Example:
    %[1]s account delete --account-id "18ee082f-1d61-40d3-b8a2-f4eee67cefff" --username "Nostrum placeat est reiciendis." --password "Ea dolor est autem."
`, os.Args[0])
}

func accountGroupAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account group-add -account-id STRING -group-id STRING -username STRING -password STRING

Add an account to a group
    -account-id STRING: Account ID
    -group-id STRING: Group ID
    -username STRING: 
    -password STRING: 

Example:
    %[1]s account group-add --account-id "18ee082f-1d61-40d3-b8a2-f4eee67cefff" --group-id "4085f7a1-6956-4003-8a89-68931f31ab12" --username "Doloribus aut et facere unde." --password "Impedit deserunt voluptas."
`, os.Args[0])
}

func accountGroupRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account group-remove -account-id STRING -group-id STRING -username STRING -password STRING

Remove an account from a group
    -account-id STRING: Account ID
    -group-id STRING: Group ID
    -username STRING: 
    -password STRING: 

Example:
    %[1]s account group-remove --account-id "18ee082f-1d61-40d3-b8a2-f4eee67cefff" --group-id "4085f7a1-6956-4003-8a89-68931f31ab12" --username "Doloribus dolore magni nihil rem repellendus qui." --password "Culpa quae vitae delectus eligendi suscipit voluptates."
`, os.Args[0])
}
