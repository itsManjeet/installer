#include "stack.hh"

Stack::Stack() {
  set_transition_duration(350);
  set_transition_type(Gtk::STACK_TRANSITION_TYPE_SLIDE_LEFT_RIGHT);
}

void Stack::next() {
  set_visible_child(*get_children()[++m_Index]);
  update_buttons();

  (*this)[m_Index]->on_load();
}
void Stack::back() {
  set_visible_child(*get_children()[--m_Index]);
  update_buttons();
}

void Stack::update_buttons() {
  m_signal_update_buttons.emit(index() < size() - 1, index() > 0);
  auto [title, subtitle] = (*this)[index()]->title();
  m_signal_update_title.emit(title, subtitle);
}
