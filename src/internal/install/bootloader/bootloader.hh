#ifndef BOOTLOADER_HH
#define BOOTLOADER_HH

#include "../install.hh"
class Bootloader : public Process {
 public:
  Bootloader(InstallData* data) : Process(_("Installed Bootloader"), data) {}
  bool process();
};

#endif