#ifndef STACK_HH
#define STACK_HH

#include <gtkmm.h>

#include "../page.hh"

class Stack : public Gtk::Stack {
 private:
  int m_Index = 0;

 public:
  Stack();

  int index() const { return m_Index; }

  Page* operator[](size_t idx) {
    return dynamic_cast<Page*>(get_children()[idx]);
  }

  size_t size() const { return get_children().size(); }

  void next();
  void back();

  void update_buttons();
  // Signal
 public:
  typedef sigc::signal<void, bool, bool> type_signal_update_buttons;
  type_signal_update_buttons signal_update_buttons() {
    return m_signal_update_buttons;
  }

  typedef sigc::signal<void, std::string, std::string> type_signal_update_title;
  type_signal_update_title signal_update_title() {
    return m_signal_update_title;
  }

 protected:
  type_signal_update_buttons m_signal_update_buttons;
  type_signal_update_title m_signal_update_title;
};

#endif