#include "system-image.hh"

#include <filesystem>
#include "../../../logging.hh"

bool SystemImage::process() {
  LOG << "installing system image" << std::endl;
  if (!std::filesystem::exists(m_Data->workDir())) {
    m_Mesg = "InternalError, '" + m_Data->workDir() + "' dir is missing";
    ERROR << m_Mesg << std::endl;
    return false;
  }

  std::string system_dir = m_Data->workDir() + "/rlxos/system/";

  std::error_code err;
  std::filesystem::create_directories(system_dir, err);
  if (err) {
    m_Mesg = "Failed to create system_dir " + err.message();
    ERROR << m_Mesg << std::endl;
    return false;
  }

  std::filesystem::copy_file(m_Data->systemImage(),
                             system_dir + "/" + m_Data->version(), err);
  if (err) {
    m_Mesg = "Failed to install system image '" + m_Data->systemImage() + "' " + err.message();
    ERROR << m_Mesg << std::endl;
    return false;
  }

  m_Mesg = "installed system image";
  LOG << m_Mesg << std::endl;
  return true;
}