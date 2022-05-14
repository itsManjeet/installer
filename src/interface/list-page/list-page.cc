#include "list-page.hh"

#include <handy.h>

ListPage::ListPage(std::string const& title, InstallData* data) : Page(data) {
  m_Title.set_markup("<span size=\"xx-large\" weight=\"ultrabold\">" + title +
                     "</span>");
  m_Title.set_margin_top(27);
  m_Title.set_margin_bottom(17);
  pack_start(m_Title, false, false);

  m_Clamp = hdy_clamp_new();
  gtk_container_add(GTK_CONTAINER(m_Clamp), GTK_WIDGET(m_ListBox.gobj()));
  gtk_container_add(GTK_CONTAINER(this->gobj()), m_Clamp);

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