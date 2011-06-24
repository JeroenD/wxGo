package wx

import ( "fmt"
)

func NullWindow() WxWindow {
	return SwigcptrWxWindow(0)
}

type EventFunc func(WxEvent)
type CommandEventFunc func(WxCommandEvent)
type SizeEventFunc func(WxSizeEvent)
type PaintEventFunc func(WxPaintEvent)

type DirESink struct{
  ef EventFunc
  cef CommandEventFunc
  sef SizeEventFunc
  pef PaintEventFunc
  id int
}

func (p *DirESink) GoEvent( event WxEvent ) {
   fmt.Printf("In Go land for %d\n", p.id)
   if p.cef != nil {
     cmd_event := SwigcptrWxCommandEvent(event.Swigcptr())
     p.cef( cmd_event )
   } else if p.ef != nil {
     p.ef( event )
   } else if p.sef != nil {
     size_event := SwigcptrWxSizeEvent(event.Swigcptr())
     p.sef( size_event )
   } else if p.pef != nil {
   	 paint_event := SwigcptrWxPaintEvent(event.Swigcptr())
   	 p.pef( paint_event )
   }
}

type Actives []interface{}
var actives Actives

func BindMenuEvent( eh WxEvtHandler, id int, ef EventFunc ) {
   des := new(DirESink)
   des.ef = ef
   des.id = id
   cb := NewDirectorESink(des)
   actives = append(actives, cb)
   eh.Connect( id, WxID_ANY, GetWxEVT_COMMAND_MENU_SELECTED(), cb )
}

func BindCommandEvent( eh WxEvtHandler, eventType int, cef CommandEventFunc ) {
   des := new(DirESink)
   des.cef = cef
   cb := NewDirectorESink(des)
   actives = append(actives, cb)
   eh.Connect( WxID_ANY, WxID_ANY, eventType, cb )
}

func BindSizeEvent( eh WxEvtHandler, sef SizeEventFunc ) {
   des := new(DirESink)
   des.sef = sef
   cb := NewDirectorESink(des)
   actives = append(actives, cb)
   eh.Connect( WxID_ANY, WxID_ANY, GetWxEVT_SIZE(), cb )
}

type TreeEventFunc func(WxTreeEvent)
type DirTreeESink struct {
   tef TreeEventFunc
}
func (p *DirTreeESink) GoEvent( event WxEvent ) {
     treeEvent := SwigcptrWxTreeEvent(event.Swigcptr())
     p.tef( treeEvent )
}

func BindTreeEvent( eh WxEvtHandler, eventType int, tef TreeEventFunc ) {
   des := new(DirTreeESink)
   des.tef = tef
   cb := NewDirectorESink(des)
   actives = append(actives, cb)
   eh.Connect( WxID_ANY, WxID_ANY, eventType, cb )
}

func BindPaintEvent( eh WxEvtHandler, pef PaintEventFunc ) {
   des := new(DirESink)
   des.pef = pef
   cb := NewDirectorESink(des)
   actives = append(actives, cb)
   eh.Connect( WxID_ANY, WxID_ANY, GetWxEVT_PAINT(), cb )
}
