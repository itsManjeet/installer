#ifndef DISK_CHECKUP_HH
#define DISK_CHECKUP_HH

#include "../checkup.hh"

class DiskCheckup : public Process {
 public:
  DiskCheckup(InstallData* data) : Process(_("Checking root parition"), data) {}

  bool process();
};

#endif