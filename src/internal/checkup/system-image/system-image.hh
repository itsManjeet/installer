#ifndef SYSTEM_IMAGE_CHECKUP_HH
#define SYSTEM_IMAGE_CHECKUP_HH

#include "../checkup.hh"

class SystemImageCheckup : public Process {
 public:
  SystemImageCheckup(InstallData* data) : Process(_("Checking System Image"), data) {}

  bool process();
};

#endif