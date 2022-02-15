#include "finish-page.hh"

#include "config.h"

FinishPage::FinishPage(InstallData* data) : Page(data) {
  set_valign(Gtk::ALIGN_CENTER);
  m_EndMesg.set_markup("<span size=\"large\"><b>rlxos</b> " +
                       std::string(_("has been installed. Click 'Restart'\n "
                                     "button to complete the setup")) +
                       "</span>");
  m_EndMesg.set_valign(Gtk::ALIGN_START);
  pack_start(m_EndMesg);

  m_Reboot.set_margin_top(15);
  m_Reboot.set_halign(Gtk::ALIGN_CENTER);
  m_Reboot.set_label(_("Reboot"));
  m_Reboot.get_style_context()->add_class("suggested-action");
  m_Reboot.signal_clicked().connect(
      sigc::mem_fun(*this, &FinishPage::on_reboot));
  pack_start(m_Reboot);
}

// TODO: check if its ok to do this
void FinishPage::on_reboot() { system("reboot"); }

void FinishPage::on_load() { m_signal_update_buttons(false, false); }