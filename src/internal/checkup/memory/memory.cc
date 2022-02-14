#include "memory.hh"

#include <string.h>
#include <sys/sysinfo.h>

#include "../../../utils/humanize.hh"
#include "config.h"

bool MemoryCheckup::process() {
  struct sysinfo info;

  if (sysinfo(&info) != 0) {
    m_Mesg = "failed to read system info " + std::string(strerror(errno));
    return false;
  }

  if (info.totalram < MINIMUM_MEMORY) {
    m_Mesg = "need alteast " + humanize(MINIMUM_MEMORY) + "";
    return false;
  }

  m_Mesg = "system has " + humanize(info.totalram);
  return true;
}