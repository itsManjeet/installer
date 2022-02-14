#ifndef LIST_PAGE_HH
#define LIST_PAGE_HH

#include "../listbox/listbox.hh"
#include "../page.hh"

class ListPage : public Page {
 protected:
  ListBox m_ListBox;
  Gtk::Label m_Title;

 public:
  ListPage(std::string const& title, InstallData* data);

  virtual ~ListPage() {}

  void on_load();
  virtual void on_selected(Gtk::ListBoxRow* row) = 0;
  virtual std::vector<Glib::RefPtr<Gtk::Widget>> get() = 0;
};

#endif