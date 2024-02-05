package web

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberHandler func(c *fiber.Ctx) error

const protectedPrefix = "private"

type IWebServer interface {
	Get(path string, handler FiberHandler)
	Post(path string, handler FiberHandler)
	Delete(path string, handler FiberHandler)
	Put(path string, handler FiberHandler)
	Patch(path string, handler FiberHandler)
}

type WebServer struct {
	app            *fiber.App
	getHandlers    map[string]FiberHandler
	postHandlers   map[string]FiberHandler
	deleteHandlers map[string]FiberHandler
	putHandlers    map[string]FiberHandler
	patchHandlers  map[string]FiberHandler
	WebServerPort  string
	jwtSecretKey   string
}

func NewWebServer(port, appName, jwtSecretKey string) *WebServer {
	return &WebServer{
		app: fiber.New(fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
			AppName:     appName,
		}),
		getHandlers:    make(map[string]FiberHandler),
		postHandlers:   make(map[string]FiberHandler),
		deleteHandlers: make(map[string]FiberHandler),
		putHandlers:    make(map[string]FiberHandler),
		patchHandlers:  make(map[string]FiberHandler),
		WebServerPort:  port,
		jwtSecretKey:   jwtSecretKey,
	}
}

type ProtectedRoutes struct {
	server IWebServer
}

func (wb *WebServer) Protected() *ProtectedRoutes {
	return &ProtectedRoutes{
		server: wb,
	}
}
func (pr *ProtectedRoutes) Get(path string, handler FiberHandler) {
	pr.server.Get(fmt.Sprintf("/%s%s", protectedPrefix, path), handler)
}

func (pr *ProtectedRoutes) Post(path string, handler FiberHandler) {
	pr.server.Post(fmt.Sprintf("/%s%s", protectedPrefix, path), handler)
}

func (pr *ProtectedRoutes) Delete(path string, handler FiberHandler) {
	pr.server.Delete(fmt.Sprintf("/%s%s", protectedPrefix, path), handler)
}

func (pr *ProtectedRoutes) Put(path string, handler FiberHandler) {
	pr.server.Put(fmt.Sprintf("/%s%s", protectedPrefix, path), handler)
}

func (pr *ProtectedRoutes) Patch(path string, handler FiberHandler) {
	pr.server.Patch(fmt.Sprintf("/%s%s", protectedPrefix, path), handler)
}

func (wb *WebServer) Get(path string, handler FiberHandler) {
	wb.getHandlers[path] = handler
}

func (wb *WebServer) Post(path string, handler FiberHandler) {
	wb.postHandlers[path] = handler
}

func (wb *WebServer) Delete(path string, handler FiberHandler) {
	wb.deleteHandlers[path] = handler
}

func (wb *WebServer) Put(path string, handler FiberHandler) {
	wb.putHandlers[path] = handler
}

func (wb *WebServer) Patch(path string, handler FiberHandler) {
	wb.patchHandlers[path] = handler
}

func (wb *WebServer) GetApp() *fiber.App {
	return wb.app
}

func (wb *WebServer) Start() error {
	wb.app.Use(logger.New())
	apiR := wb.app.Group("/api")
	apiR.Use("/"+protectedPrefix, Protected(wb.jwtSecretKey))

	for path, hdl := range wb.getHandlers {
		apiR.Get(path, hdl)
	}
	for path, hdl := range wb.postHandlers {
		apiR.Post(path, hdl)
	}
	for path, hdl := range wb.deleteHandlers {
		apiR.Delete(path, hdl)
	}
	for path, hdl := range wb.putHandlers {
		apiR.Put(path, hdl)
	}
	for path, hdl := range wb.patchHandlers {
		apiR.Patch(path, hdl)
	}

	return wb.app.Listen(fmt.Sprintf(":%s", wb.WebServerPort))
}
