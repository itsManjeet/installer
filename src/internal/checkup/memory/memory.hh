#ifndef MEMORY_HH
#define MEMORY_HH

#include "../checkup.hh"

class MemoryCheckup : public Process {
 public:
  MemoryCheckup(InstallData* data) : Process(_("Checking System Memory"), data) {}

  bool process();
};

#endif