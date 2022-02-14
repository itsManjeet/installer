#ifndef DISK_PAGE_HH
#define DISK_PAGE_HH

#include "../../disk/disk.hh"
#include "../list-page/list-page.hh"

class DiskPage : public ListPage {
 public:
  DiskPage(InstallData* data) : ListPage(_("Select Target Device"), data) {}

  virtual ~DiskPage() {}

  void on_selected(Gtk::ListBoxRow* row);

  std::vector<Glib::RefPtr<Gtk::Widget>> get();

  std::tuple<std::string, std::string> title() {
    return {_("Disk"), _("select root partition")};
  }
};

#endif