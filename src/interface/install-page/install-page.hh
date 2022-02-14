#ifndef INSTALL_PAGE_HH
#define INSTALL_PAGE_HH

#include "../../internal/install/install.hh"
#include "../../worker/worker.hh"
#include "../process-page/process-page.hh"

class InstallPage : public ProcessPage {
 public:
  InstallPage(InstallData* data)
      : ProcessPage(_("Installation"), Install::list(data), data, true) {}

  std::tuple<std::string, std::string> title() {
    return {_("Installing rlxos"), ""};
  }
};

#endif