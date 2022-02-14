#include "window.hh"

#include "config.h"

#define CHECKUP_PAGE 1

Window::Window()
    : m_WelcomePage(&m_InstallData),
      m_CheckupPage(&m_InstallData),
      m_DiskPage(&m_InstallData),
      m_InstallPage(&m_InstallData),
      m_FinishPage(&m_InstallData) {
  // m_Settings = Gio::Settings::create(APP_ID);

  set_position(Gtk::WIN_POS_CENTER_ALWAYS);
  // set_deletable(false);
  set_resizable(false);

  set_default_size(800, 600);
  set_icon(
      Gdk::Pixbuf::create_from_resource(APP_PREFIX "icons/scalable/icon.svg"));

  set_titlebar(m_Header);
  m_Header.signal_back().connect(
      sigc::mem_fun(*this, &Window::on_previous_page));
  m_Header.signal_next().connect(sigc::mem_fun(*this, &Window::on_next_page));

  m_Stack.signal_update_buttons().connect(
      sigc::mem_fun(m_Header, &Header::on_update_button));
  m_Stack.signal_update_title().connect(
      sigc::mem_fun(m_Header, &Header::on_update_title));
  add(m_Stack);

  m_WelcomePage.signal_update_buttons().connect(
      sigc::mem_fun(m_Header, &Header::on_update_button));
  m_Stack.add(m_WelcomePage, "welcome-page");

  m_DiskPage.signal_update_buttons().connect(
      sigc::mem_fun(m_Header, &Header::on_update_button));
  m_Stack.add(m_DiskPage, "disk-page");

  m_CheckupPage.signal_update_buttons().connect(
      sigc::mem_fun(m_Header, &Header::on_update_button));
  m_CheckupPage.signal_next_page().connect(
      sigc::mem_fun(m_Stack, &Stack::next));
  m_Stack.add(m_CheckupPage, "checkup-page");

  m_InstallPage.signal_update_buttons().connect(
      sigc::mem_fun(m_Header, &Header::on_update_button));
  m_InstallPage.signal_next_page().connect(
      sigc::mem_fun(m_Stack, &Stack::next));
  m_Stack.add(m_InstallPage, "install-page");

  m_FinishPage.signal_update_buttons().connect(
      sigc::mem_fun(m_Header, &Header::on_update_button));
  m_Stack.add(m_FinishPage, "finish-page");
}

Window::~Window() {}

Window* Window::create() {
  Window* window = new Window();
  return window;
}

void Window::on_next_page() {
  if (m_Stack.index() < m_Stack.size() - 1) {
    m_Stack.next();
  }
}

void Window::on_previous_page() {
  if (m_Stack.index() > 0) {
    m_Stack.back();
  }
}