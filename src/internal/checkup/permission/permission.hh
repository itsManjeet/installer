#ifndef PERMISSION_HH
#define PERMISSION_HH

#include "../checkup.hh"
class PermissionCheckup : public Process {
 public:
  PermissionCheckup(InstallData* data)
      : Process(_("Checking Permissions"), data) {}

  bool process();
};

#endif