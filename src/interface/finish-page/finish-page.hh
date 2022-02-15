#ifndef FINISH_PAGE_HH
#define FINISH_PAGE_HH

#include "../page.hh"

class FinishPage : public Page {
 private:
  Gtk::Label m_EndMesg;
  Gtk::Button m_Reboot;

  void on_reboot();

 public:
  FinishPage(InstallData* data);

  void on_load();
  std::tuple<std::string, std::string> title() {
    return {_("Almost Done!"), ""};
  }
};

#endif