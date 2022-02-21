package main

import (
	"encoding/json"

	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type props struct {
	Title  string `json:"Title"`
	Server string `json:"Server"`
}

func main() {
	engine := html.New("./svelte/dist", ".html")
	engine.AddFunc("toJSON", func(p props) template.HTML {
		j, _ := json.Marshal(p)
		return template.HTML(string(j))
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	p := props{Title: "SSR Test", Server: "GoFiber"}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", p)
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
			"satus": "ok",
		})
	})

	//this works too if you do not need to pass any props to svelte
	//statis serve can be used with fetch functions for api
	//app.Static("/", "./svelte/dist")

	app.Listen(":3000")
}
