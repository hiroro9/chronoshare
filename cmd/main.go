package main

import (
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/hiroro9/chronoshare/pkg/timer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
	timerMap = map[string]*timer.Timer{}
)

func readTimer(c echo.Context) error {

	timerId := c.Param("id")

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte(
			strconv.Itoa(timerMap[timerId].GetRemain())),
		)
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		// _, msg, err := ws.ReadMessage()
		// if err != nil {
		// 	c.Logger().Error(err)
		// }
		// fmt.Printf("%s\n", msg)
	}
}

func initTimer(c echo.Context) error {
	timerId := c.Param("id")
	remain, err := strconv.Atoi(c.QueryParam("remain"))

	if err != nil {
		remain = 200
	}

	existingTimer, ok := timerMap[timerId]

	if !ok {
		newTimer := *timer.NewTimer(timerId, remain)
		timerMap[timerId] = &newTimer
		return c.JSON(200, newTimer.GetRemain())
	}

	return c.JSON(200, existingTimer.GetRemain())

}

func getAllTimer(c echo.Context) error {
	return c.JSON(200, timerMap)
}

func startTimer(c echo.Context) error {
	timerId := c.Param("id")
	go timerMap[timerId].Start()
	return c.String(200, fmt.Sprintf("start %s \n", timerId))
}

func stopTimer(c echo.Context) error {
	timerId := c.Param("id")
	timerMap[timerId].Stop()
	return c.String(200, fmt.Sprintf("stop %s \n", timerId))
}

func restartTimer(c echo.Context) error {
	timerId := c.Param("id")
	timerMap[timerId].Restart()
	return c.String(200, fmt.Sprintf("restart %s \n", timerId))
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/timer:id", readTimer)
	e.GET("/init:id", initTimer)
	e.GET("/start:id", startTimer)
	e.GET("/stop:id", stopTimer)
	e.GET("/restart:id", restartTimer)
	e.GET("/timerMap", getAllTimer)
	e.Logger.Fatal(e.Start(":8080"))
}
