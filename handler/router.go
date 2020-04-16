package handler

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/tpphu/whitewalker/repo"
	"github.com/tpphu/whitewalker/service"
	"github.com/urfave/cli"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
)

func setContext(ctx *apm.Context, c iris.Context, body *apm.BodyCapturer) {
	ctx.SetFramework("iris", iris.Version)
	ctx.SetHTTPRequest(c.Request())
	ctx.SetHTTPRequestBody(body)
	ctx.SetHTTPStatusCode(c.ResponseWriter().StatusCode())
	ctx.SetHTTPResponseHeaders(c.ResponseWriter().Header())
}

// BuildEngine returns a *iris.Application
func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel(appContext.GlobalString("loglevel"))

	type routeInfo struct {
		transactionName string // e.g. "GET /foo"
	}

	type middleware struct {
		tracer         *apm.Tracer
		requestIgnorer apmhttp.RequestIgnorerFunc

		setRouteMapOnce sync.Once
		routeMap        map[string]map[string]routeInfo
	}

	m := &middleware{
		tracer:         apm.DefaultTracer,
		requestIgnorer: apmhttp.DefaultServerRequestIgnorer(),
	}

	app.Use(func(c iris.Context) {
		// c.Next()
		// return
		if m.tracer.Active() && m.requestIgnorer(c.Request()) {
			c.Next()
			return
		}
		m.setRouteMapOnce.Do(func() {
			routes := app.GetRoutes()
			rm := make(map[string]map[string]routeInfo)
			for _, r := range routes {
				mm := rm[r.Method]
				if mm == nil {
					mm = make(map[string]routeInfo)
					rm[r.Method] = mm
				}
				mm[r.MainHandlerName] = routeInfo{
					transactionName: r.Method + " " + r.Path,
				}
			}
			m.routeMap = rm
		})

		var requestName string
		handlerName := c.GetCurrentRoute().MainHandlerName()
		if routeInfo, ok := m.routeMap[c.Request().Method][handlerName]; ok {
			requestName = routeInfo.transactionName
		} else {
			requestName = apmhttp.UnknownRouteRequestName(c.Request())
		}
		tx, _ := apmhttp.StartTransaction(m.tracer, requestName, c.Request())
		// nc := context.NewContext(app)
		// nc.BeginRequest(c.ResponseWriter(), req)
		// c = nc
		defer tx.End()

		body := m.tracer.CaptureHTTPRequestBody(c.Request())
		defer func() {
			if v := recover(); v != nil {
				if c.ResponseWriter().Written() == 0 {
					c.EndRequest()
				} else {
					c.EndRequest()
				}
				e := m.tracer.Recovered(v)
				e.SetTransaction(tx)
				setContext(&e.Context, c, body)
				e.Send()
			}
			//c.ResponseWriter().WriteHeader(500)
			tx.Result = apmhttp.StatusCodeResult(c.GetStatusCode())

			if tx.Sampled() {
				setContext(&tx.Context, c, body)
			}

			// for _, err := range c.Errors {
			// 	e := m.tracer.NewError(err.Err)
			// 	e.SetTransaction(tx)
			// 	setContext(&e.Context, c, body)
			// 	e.Handled = true
			// 	e.Send()
			// }
			c.EndRequest()
			body.Discard()
		}()
		c.Next()
	})

	healthCheckHandler := healthCheckHandlerImpl{
		log: logger,
		db:  db,
	}
	healthCheckHandler.inject(app)
	// Note handler
	noteHanler := noteHandlerImpl{
		noteRepo: repo.NoteRepoImpl{
			DB: db,
		},
		log: logger,
	}
	noteHanler.inject(app)
	// User handler
	userHanler := userHandlerImpl{
		userRepo: repo.UserRepoImpl{
			DB: db,
		},
		log: logger,
	}
	userHanler.inject(app)
	// HTTP ReqResIn service
	reqResInHanler := reqResInHandlerImpl{
		reqResService: service.NewReqResIn("https://reqres.in"),
		log:           logger,
	}
	reqResInHanler.inject(app)
	return app
}

func simpleReturnHandler(c iris.Context, result interface{}, err Error) {
	if err != nil {
		c.StatusCode(err.Status())
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	c.JSON(result)
}
