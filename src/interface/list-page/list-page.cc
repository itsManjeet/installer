#include "list-page.hh"

ListPage::ListPage(std::string const& title, InstallData* data) : Page(data) {
  m_Title.set_markup("<span size=\"xx-large\" weight=\"ultrabold\">" + title +
                     "</span>");
  m_Title.set_margin_top(27);
  m_Title.set_margin_bottom(17);
  pack_start(m_Title, false, false);
  pack_start(m_ListBox, true, true);
  m_ListBox.set_vexpand(false);

  m_ListBox.get()->signal_row_selected().connect(
      sigc::mem_fun(*this, &ListPage::on_selected));
}

void ListPage::on_load() {
  m_ListBox.clear();
  for (auto i : get()) {
    m_ListBox.append(i.get());
    i->show_all();
  }
  m_signal_update_buttons.emit(false, true);
}