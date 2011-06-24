package main

import ( . "wx"
    "runtime"
    "fmt"
    "strings"
)

type MyItemData struct {
    desc string
}

func (id MyItemData) ShowInfo(tree WxTreeCtrl) {
  fmt.Printf("Item '%s': selected", id.desc)
}

func NewTreeItemData( desc string ) WxTreeItemData {
	myItemData := new(MyItemData)
	tid := NewDirectorWxTreeItemData(myItemData)
	myItemData.desc = desc
	return tid
}

func addItemsRecursively( treeCtrl WxGoTreeCtrl, idParent WxTreeItemId,
    numChildren int,
    depth int,
    folder int ) {
    if depth > 0 {
    	hasChildren := depth > 1
        var str string
        for n := 0; n < numChildren; n++ {
            if hasChildren {
               str = fmt.Sprintf("%s child %d", "Folder", n + 1)
            } else {
               str = fmt.Sprintf("%s child %d.%d", "File", folder, n + 2)
            }
            
            id := treeCtrl.AppendItemFull( idParent, str, -1, -1, NewTreeItemData(str) )
            
            addItemsRecursively( treeCtrl, id, numChildren, depth -1, n + 1 )
            
        }
    }
}

func addTestItemsToTree( treeCtrl WxGoTreeCtrl,
    numChildren int, depth int ) {
    rootId := treeCtrl.AddRoot( "Root", -1 , -1,
       NewTreeItemData("Root item"))
       
    addItemsRecursively( treeCtrl, rootId, numChildren, depth, 0) 
    
    treeCtrl.SetItemFont( rootId, GetItalicFont() )
    
    var cookie uintptr
    id := treeCtrl.GetFirstChild(rootId, &cookie)
    id = treeCtrl.GetNextChild(rootId, &cookie)
    treeCtrl.SetItemTextColour(id, GetBlue() )
    id = treeCtrl.GetNextChild(rootId, &cookie)
    id = treeCtrl.GetNextChild(rootId, &cookie)
    treeCtrl.SetItemBackgroundColour(id, GetLightGrey() )
}

func resize( frame WxFrame, treeCtrl WxTreeCtrl ) {
    size := frame.GetClientSize()
    fmt.Printf("size.x: %d\n", size.GetX() )
    treeCtrl.SetSize( 0, 0, size.GetX(), size.GetY() )
    fmt.Printf("tc size: %d\n", treeCtrl.GetSize().GetX() )
}

// menu item constrants
const (
    FileAbout int = 6000 + iota
    TreeAddItem
    TreeInsertItem
    TreeCount
    TreeCountRec
    TreeSort
    TreeSortRev
    ItemRename
    ItemBold
    ItemNotBold
    ItemToggle
)

func checkItem( frame WxFrame, item WxTreeItemId ) bool {
    if !item.IsOk() {
        WxMessageBox( "Please select some item first!",
            "Tree sample error",
            WxOK | WxICON_EXCLAMATION,
            frame )            
        return false;
    }
    return true

}

var item_num int = 1

type MyTree struct{
}
var reverse bool = false
func (t *MyTree) OnCompareItems( tree WxGoTreeCtrl, item1 WxTreeItemId, item2 WxTreeItemId ) int {
   // custom ordering example
   if reverse {
     if tree.GetItemText(item1) > tree.GetItemText(item2) { return -1 }
     if tree.GetItemText(item1) < tree.GetItemText(item2) { return 1 }
   } else {
     if tree.GetItemText(item1) < tree.GetItemText(item2) { return -1 }
     if tree.GetItemText(item1) > tree.GetItemText(item2) { return 1 }
   }
   return 0
}

func BuildMyTreeCb() WxTreeCtrlCb {
    tree := new(MyTree)
    dtree := NewDirectorWxTreeCtrlCb(tree)
    return dtree
}

var g_draggedItem WxTreeItemId = nil

func IsTestItem( treeCtrl WxTreeCtrl, itemId WxTreeItemId ) bool {
    if treeCtrl.GetItemParent(itemId) == treeCtrl.GetRootItem() {
       if treeCtrl.GetPrevSibling(itemId) == nil {
          return true
       }
    }
    return false
}

func main() {
    runtime.GOMAXPROCS(4)
    app := NewWxGoApp()
    
    frame := NewWxFrame( NullWindow(), -1, "WxTreeCtrl Test")
    
    fileMenu := NewWxMenu()
    fileMenu.Append( FileAbout, "&About...")
    
    itemMenu := NewWxMenu()
    itemMenu.Append( ItemToggle, "Toggle the item's &icon" )
    
    treeMenu := NewWxMenu()
    
    menuBar := NewWxMenuBar()
    menuBar.Append( fileMenu, "&File" )
    menuBar.Append( treeMenu, "&Tree" )
    menuBar.Append( itemMenu, "&Item" )
    frame.SetMenuBar( menuBar )
    
    panel := NewWxPanel( frame )
    
    treeStyle := WxTR_DEFAULT_STYLE | WxTR_EDIT_LABELS | WxBORDER_SUNKEN
    treeCtrl := NewWxGoTreeCtrl( panel, -1, GetWxDefaultPosition(), GetWxDefaultSize(), treeStyle )
    treeCtrl.SetCb( BuildMyTreeCb() )
    addTestItemsToTree( treeCtrl, 5, 2 )
    
    resize( frame, treeCtrl )
    BindSizeEvent( frame, func(event WxSizeEvent) {
        resize( frame, treeCtrl )
        event.Skip() // skip so panel gets to resize itself
    } )
      
    treeMenu.Append( TreeAddItem, "Append a &new item" )
    BindMenuEvent( frame, TreeAddItem, func(event WxEvent) {
        treeCtrl.AppendItem( treeCtrl.GetRootItem(), fmt.Sprintf("Item #%d", item_num ) )
        item_num++
    })
    
    treeMenu.Append( TreeInsertItem, "&Insert a new item" )
    BindMenuEvent( frame, TreeInsertItem, func(event WxEvent) {
        treeCtrl.InsertItem( treeCtrl.GetRootItem(), 1, "2nd item" )
    })
    
    treeMenu.Append( TreeCount, "Count children of current item" )
    BindMenuEvent( frame, TreeCount, func(event WxEvent) {
    	item := treeCtrl.GetSelection();
    	
    	if !checkItem( frame, item ) { return }
    	
    	i := treeCtrl.GetChildrenCount( item, false )
    	
    	WxLogMessage( fmt.Sprintf("%d children",i) )
    })
    
    treeMenu.Append( TreeCountRec, "Count children of current item (recursively)" )
    BindMenuEvent( frame, TreeCountRec, func(event WxEvent) {
    	item := treeCtrl.GetSelection();
    	
    	if !checkItem( frame, item ) { return }
    	
    	i := treeCtrl.GetChildrenCount( item, true )
    	
    	WxLogMessage( fmt.Sprintf("%d children",i) )
    })
    
    treeMenu.Append( TreeSort, "Sort children of current item" )
    BindMenuEvent( frame, TreeSort, func(event WxEvent) {
    	item := treeCtrl.GetSelection();
    	
    	if !checkItem( frame, item ) { return }
    	reverse = false
    	treeCtrl.SortChildren(item)
    })

    treeMenu.Append( TreeSortRev, "Sort children of current item reversed" )
    BindMenuEvent( frame, TreeSortRev, func(event WxEvent) {
    	item := treeCtrl.GetSelection();
    	
    	if !checkItem( frame, item ) { return }
    	reverse = true
    	treeCtrl.SortChildren(item)
    })

    
    itemMenu.Append( ItemRename, "&Rename item..." )
    BindMenuEvent( frame, ItemRename, func(event WxEvent) {
       item := treeCtrl.GetSelection()
       
       if !checkItem( frame, item ) { return }
    
       treeCtrl.EditLabel(item)
    })    
    itemMenu.Append( ItemBold, "Make item &bold" )
    BindMenuEvent( frame, ItemBold, func(event WxEvent) {
       item := treeCtrl.GetSelection()
       
       if !checkItem( frame, item ) { return }
    
       treeCtrl.SetItemBold( item, true )
       
    })
    itemMenu.Append( ItemNotBold, "Make item &not bold")
    BindMenuEvent( frame, ItemNotBold, func(event WxEvent) {
       item := treeCtrl.GetSelection()
       
       if !checkItem( frame, item ) { return }
    
       treeCtrl.SetItemBold( item, false )
    })
    
    BindTreeEvent( treeCtrl, GetWxEVT_COMMAND_TREE_BEGIN_DRAG(), func(event WxTreeEvent) {
       if event.GetItem() != treeCtrl.GetRootItem() {
          g_draggedItem = event.GetItem()
          //clientpt := event.GetPoint()
          //screenpt := treeCtrl.ClientToScreen( clientpt )
          // WxLogMessage( fmt.Sprintf("OnBeginDrag: started dragging %s at screen coords (%d,%d)", treeCtrl.GetItemText( g_draggedItem ), screenpt.GetX(), screenpt.GetY() ) )
          event.Allow()
       } else {
          WxLogMessage( "OnBeginDrag: this item can't be dragged." )
       }
    })
    
    BindTreeEvent( treeCtrl, GetWxEVT_COMMAND_TREE_END_DRAG(), func(event WxTreeEvent) {
        itemSrc := g_draggedItem
        itemDst := event.GetItem()
        g_draggedItem = nil
        
        if itemDst.IsOk() && !treeCtrl.ItemHasChildren(itemDst) {
           itemDst = treeCtrl.GetItemParent(itemDst)
        }
        
        if !itemDst.IsOk() {
           WxLogMessage("OnEndDrag: can't drop here.")
           return
        }
        
        text := treeCtrl.GetItemText( itemSrc )
        treeCtrl.AppendItem( itemDst, text )
    })
    
    BindTreeEvent( treeCtrl, GetWxEVT_COMMAND_TREE_END_LABEL_EDIT(), func( event WxTreeEvent ) {
        if strings.Contains(event.GetLabel(), " ") {
            WxMessageBox( "The new label should be a single word." )
            event.Veto()
        }
    })
    
    BindTreeEvent( treeCtrl, GetWxEVT_COMMAND_TREE_ITEM_COLLAPSING(), func( event WxTreeEvent ) {
        itemId := event.GetItem()
        if IsTestItem(treeCtrl, itemId) {
            WxMessageBox( "You can't collapse this item." )
            event.Veto()
        }
    })
    
    frame.Show(true)
    app.MainLoop()    
}
