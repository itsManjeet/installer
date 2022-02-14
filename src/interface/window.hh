#ifndef WINDOW_HH
#define WINDOW_HH

#include <gtkmm.h>

#include "../internal/data/data.hh"
#include "checkup-page/checkup-page.hh"
#include "disk-page/disk-page.hh"
#include "finish-page/finish-page.hh"
#include "header/header.hh"
#include "install-page/install-page.hh"
#include "stack/stack.hh"
#include "welcome-page/welcome-page.hh"

class Window : public Gtk::ApplicationWindow {
 private:
  Glib::RefPtr<Gio::Settings> m_Settings;
  Header m_Header;

  // Pages
  WelcomePage m_WelcomePage;
  CheckupPage m_CheckupPage;
  DiskPage m_DiskPage;
  InstallPage m_InstallPage;
  FinishPage m_FinishPage;

  Stack m_Stack;

  InstallData m_InstallData;

  int m_PageIndex = 0;

 public:
  Window();
  ~Window();

  static Window* create();

  void on_next_page();
  void on_previous_page();
};

#endif