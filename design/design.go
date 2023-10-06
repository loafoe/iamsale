package design

/*

[x] Test Connection – To check the connectivity
[ ] Account Aggregation – To bring all users data into SailPoint
[ ] Group Aggregation – To bring all groups data into SailPoint
[ ] Single account Aggregation – To bring the single user data
[ ] Create Account – To create the user account
[ ] Update Account – update User data like FirstName, LastName, Email.
[ ] Add Group/Roles – To add group/roles to user profile
[ ] Remove Group/Roles  – To remove group/roles from user profile
[ ] Enable Account – To enable (or reactivate) user profile
[ ] Disable/Delete Account – To disable/Delete user profile

*/

import (
	. "goa.design/goa/v3/dsl"
)

var BasicAuth = BasicAuthSecurity("BasicAuth", func() {
	Description("Use client ID and client secret to authenticate")
})

var _ = API("sailpoint", func() {
	Title("Sailpoint integration")
	Description("Implements the integration API for sailpoint")
	Server("sailpoint", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var Account = Type("Account", func() {
	Attribute("name", String, "Name of user", func() {
		Example("Amos Burton")
	})
	Attribute("login", String, "Login of user", func() {
		Example("amos")
	})
	Attribute("email", String, "Email of user", func() {
		Example("amos@hostedzonehere.com")
	})
})

var Group = Type("group", func() {
	Attribute("name", String, "Name of group", func() {
		Example("GROUPNAME")
	})
	Attribute("guid", String, "GUID of group", func() {
		Example("example 1", "6dfb5ec6-0c8e-4efa-ae7d-9eca77b73d2b")
		Example("another 2", "ac32dcb3-3a58-4bfc-b801-6609681ec712")
	})
})

var _ = Service("test", func() {
	Description("Test service availability")
	Method("test", func() {
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

var _ = Service("aggregate", func() {
	Description("Integration APIs")

	Method("accounts", func() {
		Description("Account aggregation. Returns all known accounts")
		Payload(func() {
			Username("username")
			Password("password")
			Required("username", "password")
		})
		Security(BasicAuth)
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
		Security(BasicAuth)
		Error("PermissionDenied")
		Result(ArrayOf(Group))
		HTTP(func() {
			GET("/aggregate/groups")
			Response(StatusOK)
			Response("PermissionDenied", StatusForbidden)
		})
	})
})

var _ = Service("account", func() {
	Description("Accounts APIs")
	Method("create", func() {
		Description("Create an account")
		Payload(func() {
			UsernameField(1, "username", String, "Username")
			PasswordField(2, "password", String, "Password")
			Field(3, "account", Account)
			Required("username", "password", "account")
		}, Account)
		Security(BasicAuth)
		HTTP(func() {
			POST("/account")
			Body("account")
			Response(StatusCreated)
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
		Security(BasicAuth)
		HTTP(func() {
			DELETE("/account/{accountId}")
			Response(StatusNoContent)
		})
	})
})
