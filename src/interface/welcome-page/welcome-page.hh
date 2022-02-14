#ifndef WELCOME_PAGE_HH
#define WELCOME_PAGE_HH

#include "../list-page/list-page.hh"

class WelcomePage : public Page {
 private:
  Gtk::Label m_Title, m_SubTitle;
  Gtk::Image m_Image;

 public:
  WelcomePage(InstallData* data);

  virtual ~WelcomePage() {}

  void on_load();

  std::tuple<std::string, std::string> title() {
    return {_("welcome"), _("rlxos system installer")};
  }
};

#endif