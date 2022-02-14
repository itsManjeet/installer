#ifndef PREPARE_HH
#define PREPARE_HH

#include "../install.hh"

class Prepare : public Process {
 public:
  Prepare(InstallData* data) : Process(_("Prepared Work dir"), data) {}

  bool process();
};

#endif