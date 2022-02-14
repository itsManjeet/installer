#include "disk-page.hh"

#include <iostream>

std::vector<Glib::RefPtr<Gtk::Widget>> DiskPage::get() {
  std::vector<Glib::RefPtr<Gtk::Widget>> list;

  for (auto i : Disk::list()) {
    auto disk = ListButton<Disk>::create(
        "drive-harddisk", i.size() + " " + i.label(), i.path(), i);
    list.push_back(disk);
  }
  return list;
}

void DiskPage::on_selected(Gtk::ListBoxRow* row) {
  if (row == nullptr) {
    return;
  }

  auto selected = dynamic_cast<ListButton<Disk>*>(row->get_child());
  if (selected == nullptr) {
    return;
  }

  m_InstallData->disk(selected->data().path());

  m_signal_update_buttons.emit(true, true);
}