#include "listbox.hh"

ListBox::ListBox() {
  set_policy(Gtk::PolicyType::POLICY_NEVER, Gtk::PolicyType::POLICY_AUTOMATIC);

  add(m_Frame);

  m_ListBox.set_header_func(sigc::mem_fun(*this, &ListBox::header_func));
  m_Frame.add(m_ListBox);

  set_margin_top(27);
  set_margin_bottom(27);
  set_margin_start(250);
  set_margin_end(250);

  set_hexpand();
}

void ListBox::clear() {
  for (auto i : m_ListBox.get_children()) {
    m_ListBox.remove(*i);
  }
}

void ListBox::append(Gtk::Widget* widget) {
  auto adj = get_vadjustment();
  adj->set_value(adj->get_upper() - adj->get_page_size());
  m_ListBox.append(*widget);
}

void ListBox::header_func(Gtk::ListBoxRow* row, Gtk::ListBoxRow* before) {
  if (before != nullptr) {
    auto sep = Glib::RefPtr<Gtk::Separator>(new Gtk::Separator());
    sep->show();
    row->set_header(*sep.get());
  }
}

Gtk::Widget* ListBox::get_selected() {
  auto selected = m_ListBox.get_selected_row();
  return selected->get_child();
}