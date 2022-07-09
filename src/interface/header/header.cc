#include "header.hh"

#include <libintl.h>
#include <locale.h>

#include "config.h"

Header::Header() {
  set_show_close_button(false);

  m_BackBtn.set_label(_("Back"));
  m_BackBtn.signal_clicked().connect(
      sigc::mem_fun(*this, &Header::on_backbtn_clicked));
  pack_start(m_BackBtn);

  m_NextBtn.set_label(_("Next"));
  auto nextBtn_context = m_NextBtn.get_style_context();
  nextBtn_context->add_class("suggested-action");
  m_NextBtn.signal_clicked().connect(
      sigc::mem_fun(*this, &Header::on_nextbtn_clicked));
  pack_end(m_NextBtn);

  m_HelpBtn.set_image_from_icon_name("help", Gtk::ICON_SIZE_LARGE_TOOLBAR);
  m_HelpBtn.signal_clicked().connect(
      sigc::mem_fun(*this, &Header::on_helpbtn_clicked));
  pack_end(m_HelpBtn);
}

void Header::on_update_button(bool nextBtn, bool backBtn) {
  m_NextBtn.set_sensitive(nextBtn);
  m_BackBtn.set_sensitive(backBtn);
}

void Header::on_update_title(std::string title, std::string subtitle) {
  set_title(title);
  // set_subtitle(subtitle);
}

void Header::on_helpbtn_clicked() { system("xdg-open " HELP_URL " &"); }

void Header::on_nextbtn_clicked() { m_type_signal_next.emit(); }
void Header::on_backbtn_clicked() { m_type_signal_back.emit(); }