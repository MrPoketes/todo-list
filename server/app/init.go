package app

import (
	"context"
	"net/http"

	"github.com/revel/revel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	Collection *mongo.Collection

	Disconnect func()
)

func InitDB() {
	uri, found := revel.Config.String("atlas.uri")
	if !found {
		panic("No ATLAS_URI found")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	Collection = client.Database("todolist").Collection("users")

	Disconnect = func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		ValidateOrigin,
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(InitDB)
	revel.OnAppStop(Disconnect)
}

var ValidateOrigin = func(c *revel.Controller, fc []revel.Filter) {
	if c.Request.Method == "OPTIONS" {
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Response.Out.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Response.SetStatus(http.StatusNoContent)
	} else {
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", "*")
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Response.Out.Header().Add("Content-Type", "application/json; charset=UTF-8")
		c.Response.Out.Header().Add("X-Frame-Options", "SAMORIGIN")
		c.Response.Out.Header().Add("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")

		fc[0](c, fc[1:]) // Execute the next filter stage.
	}
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
