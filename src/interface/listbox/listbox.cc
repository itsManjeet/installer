#include "listbox.hh"

ListBox::ListBox() {
  property_propagate_natural_height() = true;
  m_ListBox.set_header_func(sigc::mem_fun(*this, &ListBox::header_func));
  add(m_ListBox);
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