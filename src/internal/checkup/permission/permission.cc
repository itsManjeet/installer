#include "permission.hh"

#include <unistd.h>
#include "../../../logging.hh"

bool PermissionCheckup::process() {
  if (geteuid() != 0) {
    m_Mesg = "No super user access";
    ERROR << m_Mesg << std::endl;
    return false;
  }
  m_Mesg = "Got super user access";
  LOG << m_Mesg << std::endl;
  return true;
}