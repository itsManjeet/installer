#ifndef HEADER_HH
#define HEADER_HH

#include <gtkmm.h>

#include "../../locale.hh"

class Header : public Gtk::HeaderBar {
 private:
  Gtk::Button m_BackBtn, m_NextBtn, m_HelpBtn;

 public:
  Header();

  virtual ~Header() {}

  void on_backbtn_clicked();
  void on_nextbtn_clicked();

  void on_helpbtn_clicked();

  void on_update_button(bool, bool);
  void on_update_title(std::string, std::string);

  // signal
 public:
  typedef sigc::signal<void> type_signal_next;
  typedef sigc::signal<void> type_signal_back;

  type_signal_next signal_next() { return m_type_signal_next; }
  type_signal_back signal_back() { return m_type_signal_back; }

 protected:
  type_signal_next m_type_signal_next;
  type_signal_back m_type_signal_back;
};

#endif