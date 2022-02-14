#ifndef PAGE_HH
#define PAGE_HH

#include <gtkmm.h>

#include "../internal/data/data.hh"
#include "../locale.hh"

class Page : public Gtk::Box {
 protected:
  InstallData* m_InstallData;

 public:
  Page(InstallData* data)
      : Gtk::Box(Gtk::ORIENTATION_VERTICAL, 0), m_InstallData(data) {}

  virtual ~Page() {}
  virtual void on_load() = 0;
  virtual std::tuple<std::string, std::string> title() = 0;

  // Signal
 public:
  typedef sigc::signal<void, bool, bool> type_signal_update_buttons;
  type_signal_update_buttons signal_update_buttons() {
    return m_signal_update_buttons;
  }

  typedef sigc::signal<void> type_signal_next_page;
  type_signal_next_page signal_next_page() { return m_signal_next_page; }

 protected:
  type_signal_update_buttons m_signal_update_buttons;
  type_signal_next_page m_signal_next_page;
};
#endif