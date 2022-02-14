#include "system-info.hh"

#include <filesystem>
#include <fstream>

#include "config.h"

bool SystemInfoCheckup::process() {
  m_Data->version("2200");
  m_Mesg = "Found system version '" + m_Data->version() + "'";
  return true;
  
  if (std::filesystem::exists(ISO_PATH "version")) {
    std::ifstream file(ISO_PATH "version");
    std::string version;
    file >> version;
    file.close();

    m_Data->version(version);
    m_Mesg = "found system version '" + version + "'";
    return true;
  }

  m_Mesg = "InternalError! no system version found";
  return false;
}