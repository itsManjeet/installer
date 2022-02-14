#ifndef KERNEL_HH
#define KERNEL_HH

#include "../install.hh"

class Kernel : public Process {
 public:
  Kernel(InstallData* data) : Process(_("Installed Kernel Layer"), data) {}

  bool process();
};

#endif