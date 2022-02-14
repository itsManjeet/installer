#ifndef CHECKUP_PAGE_HH
#define CHECKUP_PAGE_HH

#include <thread>

#include "../../internal/checkup/checkup.hh"
#include "../../worker/worker.hh"
#include "../process-page/process-page.hh"

class CheckupPage : public ProcessPage {
 public:
  CheckupPage(InstallData* data)
      : ProcessPage(_("Compatibality Checkup"), Checkup::list(data), data, false) {}

  std::tuple<std::string, std::string> title() {
    return {_("checkup"), _("checking system compatibality")};
  }
};

#endif