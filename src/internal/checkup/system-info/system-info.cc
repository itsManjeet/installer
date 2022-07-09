#include "system-info.hh"

#include <filesystem>
#include <fstream>

#include "config.h"
#include "../../../logging.hh"

bool SystemInfoCheckup::process() {  
  if (std::filesystem::exists(ISO_PATH "version")) {
    std::ifstream file(ISO_PATH "version");
    std::string version;
    file >> version;
    file.close();

    m_Data->version(version);
    m_Mesg = "found system version '" + version + "'";
    LOG << m_Mesg << std::endl;
    return true;
  }
  m_Data->version("2200");
  m_Mesg = "InternalError! no system version found";
  ERROR << m_Mesg << std::endl;
  return false;
}