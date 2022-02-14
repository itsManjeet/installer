#ifndef FINISH_PAGE_HH
#define FINISH_PAGE_HH

#include "../page.hh"

class FinishPage : public Page {
 private:
  Gtk::Image m_Image;
  Gtk::Label m_EndMesg;

 public:
  FinishPage(InstallData* data);

  void on_load();
  std::tuple<std::string, std::string> title() {
    return {_("Almost Done!"), ""};
  }
};

#endif