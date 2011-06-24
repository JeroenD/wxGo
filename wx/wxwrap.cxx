#include "wxwrap.h"
#include <wx/toplevel.h>
#include <iostream>

using namespace std;

wxGoApp::wxGoApp()
{
    argc = 0;
    argv = 0;
    wxInitialize(argc, argv);
    wxApp::OnInit();
    SetExitOnFrameDelete( true );
}

wxGoApp::~wxGoApp()
{}

int wxGoApp::MainLoop() {
    int retval = 0;
    bool initialized = wxTopLevelWindows.GetCount() != 0;
    if (initialized) {
        retval = wxApp::MainLoop();
    }
    return retval;
}

ESink::~ESink()
{
}

void ESink::Thunk( wxEvent& event )
{
	ESink* cb = (ESink*)event.m_callbackUserData;
	cb->GoEvent( event );
}

IMPLEMENT_DYNAMIC_CLASS(wxGoTreeCtrl, wxTreeCtrl)

wxGoTreeCtrl::wxGoTreeCtrl( wxWindow *parent, wxWindowID id,
               const wxPoint& pos,
               const wxSize& size,
               int style,
               const wxValidator& validator,
               const wxString& name )
    :wxTreeCtrl( parent, id, pos, size, style, validator, name ),
    mpCb(0)
{
}

int wxGoTreeCtrl::OnCompareItems(const wxTreeItemId& item1,
    const wxTreeItemId& item2)
{
    if ( mpCb ) return mpCb->OnCompareItems( *this, item1, item2 );
    return wxTreeCtrl::OnCompareItems( item1, item2 );
}

wxTreeCtrlCb::~wxTreeCtrlCb()
{
}

extern const wxFont* GetItalicFont() {
   return wxStockGDI::instance().GetFont(wxStockGDI::FONT_ITALIC);
}
extern const wxFont* GetNormalFont() {
    return wxStockGDI::instance().GetFont(wxStockGDI::FONT_NORMAL);
}
extern const wxFont* GetSmallFont() {
    return wxStockGDI::instance().GetFont(wxStockGDI::FONT_SMALL);
}
extern const wxFont* GetSwissFont() {
    return wxStockGDI::instance().GetFont(wxStockGDI::FONT_SWISS);
}
extern const wxColour* GetBlack() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_BLACK);
}
extern const wxColour* GetBlue() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_BLUE);
}
extern const wxColour* GetCyan() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_CYAN);
}
extern const wxColour* GetGreen() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_GREEN);
}
extern const wxColour* GetLightGrey() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_LIGHTGREY);
}
extern const wxColour* GetRed() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_RED);
}
extern const wxColour* GetWhite() {
    return wxStockGDI::instance().GetColour(wxStockGDI::COLOUR_WHITE);
}

IMPLEMENT_APP(wxGoApp)

