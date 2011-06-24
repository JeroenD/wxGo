#ifndef wxwrap_h
#define wxwrap_h

#include <wx/app.h>
#include <wx/frame.h>
#include <wx/object.h>
#include <wx/treectrl.h>

class ESink : public wxObject {
public:
	ESink()
	:someNr(42)
	{}
	virtual ~ESink();
	void Thunk( wxEvent& event );
	virtual void GoEvent( wxEvent& event ) {}
	int someNr;
};

class wxGoApp : public wxApp {
public:
    wxGoApp();
    ~wxGoApp();
    virtual int MainLoop();
};

class wxTreeCtrlCb;

class wxGoTreeCtrl : public wxTreeCtrl {
public:
    wxGoTreeCtrl() {}
    wxGoTreeCtrl(wxWindow *parent, wxWindowID id = wxID_ANY,
               const wxPoint& pos = wxDefaultPosition,
               const wxSize& size = wxDefaultSize,
               int style = wxTR_HAS_BUTTONS | wxTR_LINES_AT_ROOT,
               const wxValidator& validator = wxDefaultValidator,
               const wxString& name = wxTreeCtrlNameStr);
    
    virtual int OnCompareItems(const wxTreeItemId& item1,
                   const wxTreeItemId& item2);
    void SetCb( wxTreeCtrlCb* pCb ) { mpCb = pCb; }
    DECLARE_DYNAMIC_CLASS(wxGoTreeCtrl)

private:
    wxTreeCtrlCb* mpCb;
};

class wxTreeCtrlCb 
{
public:
    virtual ~wxTreeCtrlCb();
    virtual int OnCompareItems( 
        const wxGoTreeCtrl& rTree,
        const wxTreeItemId& item1,
        const wxTreeItemId& item2 ){ return -1;};
};


#endif // wxwrap_h
