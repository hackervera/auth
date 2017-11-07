package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tjgillies/auth"
)

type Params struct {
	Domain    string `form:"domain"`
	Sig       string `form:"sig"`
	PublicKey string `form:"publicKey"`
	Challenge string `form:"challenge"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.POST("/auth", func(c echo.Context) (err error) {
		p := new(Params)
		fmt.Println(c)
		if err = c.Bind(p); err != nil {
			return
		}
		fmt.Println(p)
		v, err := auth.Verify(p.Domain, p.Challenge,
			p.Sig, p.PublicKey)
		if err != nil {
			return
		}
		if !v {
			return
		}
		return c.String(http.StatusOK, "Verified")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
