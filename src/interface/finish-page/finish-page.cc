#include "finish-page.hh"

#include "config.h"

FinishPage::FinishPage(InstallData* data) : Page(data) {
  m_Image.set_from_resource(APP_PREFIX "finish.png");
  m_Image.set_valign(Gtk::ALIGN_START);
  pack_start(m_Image, true, true);

  m_EndMesg.set_markup(
      "<span size=\"large\"><b>rlxos</b> " +
      std::string(_("has been installed. The Computer will\nrestart in a few "
                    "moments to complete the setup")) +
      "</span>");
  m_EndMesg.set_valign(Gtk::ALIGN_START);
  pack_start(m_EndMesg, true, true);
}

void FinishPage::on_load() { m_signal_update_buttons(false, false); }