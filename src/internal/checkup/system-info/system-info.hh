#ifndef SYSTEM_INFO_HH
#define SYSTEM_INFO_HH

#include "../checkup.hh"

class SystemInfoCheckup : public Process {
 public:
  SystemInfoCheckup(InstallData* data)
      : Process(_("Gathering System details"), data) {}

  bool process();
};

#endif