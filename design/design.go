package design

/*

[x] Test Connection – To check the connectivity
[x] Account Aggregation – To bring all users data into SailPoint
[x] Group Aggregation – To bring all groups data into SailPoint
[x] Single account Aggregation – To bring the single user data
[x] Create Account – To create the user account
[x] Update Account – update User data like FirstName, LastName, Email.
[x] Add Group/Roles – To add group/roles to user profile
[x] Remove Group/Roles  – To remove group/roles from user profile
[x] Enable Account – To enable (or reactivate) user profile
[x] Disable/Delete Account – To disable/Delete user profile

*/

import (
	. "goa.design/goa/v3/dsl"
)

var BasicAuth = BasicAuthSecurity("BasicAuth", func() {
	Description("Use client ID and client secret to authenticate")
})

var _ = API("iamsale", func() {
	Title("IAM to broker identity integration")
	Version("1")
	Security(BasicAuth)
	Description("An integration API towards identity brokers.")
	Server("iamsale", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var CreateAccount = Type("CreateAccount", func() {
	Attribute("name", String, "Name of user", func() {
		Example("Amos Burton")
	})
	Attribute("login", String, "Login of user", func() {
		Example("amos")
	})
	Attribute("email", String, "Email of user", func() {
		Example("amos@hostedzonehere.com")
	})
	Required("name", "login", "email")
})

var UpdateAccount = Type("UpdateAccount", func() {
	Attribute("status", String, "Status of user", func() {
		Enum("active", "disabled", "pending")
	})
	Required("status")
})

var Account = Type("Account", func() {
	Attribute("id", Int64, "Temporary account identifier", func() {
		Meta("struct:tag:json", "id,omitempty")
		Meta("struct:tag:gorm", "autoIncrement")
	})
	Attribute("guid", String, "IDP account identifier", func() {
		Meta("struct:tag:json", "guid,omitempty")
	})
	Attribute("name", String, "Name of user", func() {
		Example("Amos Burton")
	})
	Attribute("login", String, "Login of user", func() {
		Example("amos")
		Meta("struct:tag:json", "login")
		Meta("struct:tag:gorm", "uniqueIndex")
	})
	Attribute("email", String, "Email of user", func() {
		Example("amos@hostedzonehere.com")
	})
	Attribute("status", String, "Status of account", func() {
		Enum("active", "disabled", "pending")
	})
	Required("name", "login", "email")
})

var Group = Type("group", func() {
	Attribute("name", String, "Name of group", func() {
		Example("GROUPNAME")
	})
	Attribute("id", String, "ID of group", func() {
		Example("example 1", "6dfb5ec6-0c8e-4efa-ae7d-9eca77b73d2b")
		Example("another 2", "ac32dcb3-3a58-4bfc-b801-6609681ec712")
	})
})

// Test
var _ = Service("test", func() {
	Description("Test service availability")
	Method("test", func() {
		NoSecurity()
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/test")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})

// Aggregate
var _ = Service("aggregate", func() {
	Description("Integration APIs")

	Method("accounts", func() {
		Description("Account aggregation. Returns all known accounts")
		Payload(func() {
			Username("username")
			Password("password")
			Required("username", "password")
		})
		Error("PermissionDenied")
		Result(ArrayOf(Account))
		HTTP(func() {
			GET("/aggregate/accounts")
			Response(StatusOK)
			Response("PermissionDenied", StatusForbidden)
		})
	})
	Method("groups", func() {
		Description("Group aggregation. Returns list of all known groups")
		Payload(func() {
			Username("username")
			Password("password")
			Required("username", "password")
		})
		Error("PermissionDenied")
		Result(ArrayOf(Group))
		HTTP(func() {
			GET("/aggregate/groups")
			Response(StatusOK)
			Response("PermissionDenied", StatusForbidden)
		})
	})
})

// Account
var _ = Service("account", func() {
	Description("Accounts APIs")
	Method("create", func() {
		Description("Creates an account. Note that the client may choose to create a shadow account or hold the account in a temporary store until the actual account materializes.")
		Payload(func() {
			Username("username")
			Password("password")
			Field(1, "account", CreateAccount)
			Required("username", "password", "account")
		}, CreateAccount)
		Result(Account)
		HTTP(func() {
			POST("/account")
			Body("account")
			Response(StatusCreated)
		})

	})
	Method("get", func() {
		Description("Get account details")
		Payload(func() {
			UsernameField(1, "username", String, "Username")
			PasswordField(2, "password", String, "Password")
			Field(3, "accountId", String, "Account ID", func() {
				Example("18ee082f-1d61-40d3-b8a2-f4eee67cefff")
			})
			Required("username", "password", "accountId")
		})
		Result(Account)
		HTTP(func() {
			GET("/account/{accountId}")
			Response(StatusOK)
		})
	})
	Method("update", func() {
		Description("Update account details")
		Payload(func() {
			UsernameField(1, "username", String, "Username")
			PasswordField(2, "password", String, "Password")
			Field(3, "accountId", String, "Account ID", func() {
				Example("18ee082f-1d61-40d3-b8a2-f4eee67cefff")
			})
			Field(4, "account", UpdateAccount)
			Required("username", "password", "accountId", "account")
		})
		Error("NotImplemented")
		Result(Account)
		HTTP(func() {
			PUT("/account/{accountId}")
			Body("account")
			Response("NotImplemented", StatusNotImplemented)
			Response(StatusOK)
		})
	})
	Method("delete", func() {
		Description("Delete an account")
		Payload(func() {
			Username("username")
			Password("password")
			Field(1, "accountId", String, "Account ID", func() {
				Example("18ee082f-1d61-40d3-b8a2-f4eee67cefff")
			})
			Required("username", "password", "accountId")
		}, Account)
		HTTP(func() {
			DELETE("/account/{accountId}")
			Response(StatusNoContent)
		})
	})

	Method("groupAdd", func() {
		Description("Add an account to a group")
		Payload(func() {
			Username("username")
			Password("password")
			Field(1, "accountId", String, "Account ID", func() {
				Example("18ee082f-1d61-40d3-b8a2-f4eee67cefff")
			})
			Field(2, "groupId", String, "Group ID", func() {
				Example("4085f7a1-6956-4003-8a89-68931f31ab12")
			})
			Required("username", "password", "accountId", "groupId")
		}, Account)
		HTTP(func() {
			POST("/account/{accountId}/group/{groupId}")
			Response(StatusNoContent)
		})
	})

	Method("groupRemove", func() {
		Description("Remove an account from a group")
		Payload(func() {
			Username("username")
			Password("password")
			Field(1, "accountId", String, "Account ID", func() {
				Example("18ee082f-1d61-40d3-b8a2-f4eee67cefff")
			})
			Field(2, "groupId", String, "Group ID", func() {
				Example("4085f7a1-6956-4003-8a89-68931f31ab12")
			})
			Required("username", "password", "accountId", "groupId")
		}, Account)
		HTTP(func() {
			DELETE("/account/{accountId}/group/{groupId}")
			Response(StatusNoContent)
		})
	})

})
