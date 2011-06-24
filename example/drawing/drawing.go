package main

import (
	"wx"
	"runtime"
	"fmt"
)

var times int = 0

func main() {
	runtime.GOMAXPROCS(4)

	app := wx.NewWxGoApp()
	
	frame := wx.NewWxFrame(wx.NullWindow(), -1, "Drawing")
		
    wx.BindPaintEvent( frame, func(event wx.WxPaintEvent) {
		pdc := wx.NewWxPaintDC( frame )
		pdc.DrawLine(10, 10, 50, 50 + times % 100)
		times += 1
		fmt.Printf("times: %d", times)
		wx.DeleteWxPaintDC( pdc )
	})
	
	frame.Show(true)
	
	app.MainLoop()
}
