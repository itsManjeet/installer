#include "permission.hh"

#include <unistd.h>

bool PermissionCheckup::process() {
  if (geteuid() != 0) {
    m_Mesg = "No super user access";
    return false;
  }
  m_Mesg = "Got super user access";
  return true;
}