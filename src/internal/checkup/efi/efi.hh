#ifndef EFI_HH
#define EFI_HH

#include "../checkup.hh"
class EfiCheckup : public Process {
 public:
  EfiCheckup(InstallData* data)
      : Process(_("Checking bootloader type"), data) {}

  bool process();
};

#endif