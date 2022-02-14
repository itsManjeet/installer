#ifndef SYSTEM_IMAGE_HH
#define SYSTEM_IMAGE_HH

#include "../install.hh"
class SystemImage : public Process {
 public:
  SystemImage(InstallData* data)
      : Process(_("Installed System Image"), data) {}

  bool process();
};

#endif