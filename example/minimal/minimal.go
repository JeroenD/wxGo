package main

import (
	"wx"
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(4)

	app := wx.NewWxGoApp()
	
	frame := wx.NewWxFrame(wx.NullWindow(), -1, "Minimal wxGo App")
	
	fileMenu := wx.NewWxMenu()
	fileMenu.Append( wx.WxID_EXIT, "E&xit\tAlt-X", "Quit this program" )
	
	helpMenu := wx.NewWxMenu()
	helpMenu.Append( wx.WxID_ABOUT, "&About...\tF1", "Show about dialog")
	
	menuBar := wx.NewWxMenuBar()
	menuBar.Append(fileMenu, "&File")
	menuBar.Append(helpMenu, "&Help")
	
	wx.BindMenuEvent( frame, wx.WxID_ABOUT, func(event wx.WxEvent) {
	   fmt.Printf("About Called")
	   wx.WxMessageBox( fmt.Sprintf("Welcome to wxGo!\n\nThis is the minimal wxGo sample\nrunning under %s.",
	   wx.WxGetOsDescription() ),
	   "About wxGo minimal sample\n",
	   wx.WxOK | wx.WxICON_INFORMATION,
	   frame )
	})
	
	wx.BindMenuEvent( frame, wx.WxID_EXIT, func(event wx.WxEvent) {
	   frame.Close();
	})
	
	frame.SetMenuBar( menuBar )
	
	frame.CreateStatusBar(2);
	frame.SetStatusText("Welcome to wxGo!")
	
	frame.Show(true)
	
	app.MainLoop()
}
