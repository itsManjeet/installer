#ifndef SEARCH_BOOT_HH
#define SEARCH_BOOT_HH

#include "../checkup.hh"
class SearchBootCheckup : public Process {
 public:
  SearchBootCheckup(InstallData* data)
      : Process(_("Searching Boot Device"), data) {}

  bool process();
};

#endif