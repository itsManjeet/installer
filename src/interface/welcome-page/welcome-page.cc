#include "welcome-page.hh"

WelcomePage::WelcomePage(InstallData* data) : Page(data) {
  set_valign(Gtk::ALIGN_CENTER);

  m_Image.set_from_icon_name("rlxos", Gtk::ICON_SIZE_BUTTON);
  m_Image.set_pixel_size(128);
  pack_start(m_Image, false, false);

  m_Title.set_markup("<span size=\"xx-large\" weight=\"ultrabold\"> " _(
      "rlxos GNU/Linux") "</span>");
  m_Title.set_margin_top(12);
  pack_start(m_Title, false, false);

  m_SubTitle.set_margin_top(10);
  m_SubTitle.set_markup(
      "<span size=\"large\" color=\"gray\">Welcome to rlxos system "
      "installer, click 'Next' to start the process</span>");
  pack_start(m_SubTitle, false, false);
  m_SubTitle.set_margin_bottom(50);
}

void WelcomePage::on_load() {}