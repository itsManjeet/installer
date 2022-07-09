#include "kernel.hh"

#include <filesystem>

#include "../../../logging.hh"
#include "../../../utils/exec.hh"

bool Kernel::process() {
  LOG << "installing kernel" << std::endl;
  if (!std::filesystem::exists(m_Data->workDir())) {
    m_Mesg = "InternalError, '" + m_Data->workDir() + "' dir is missing";
    ERROR << m_Mesg << std::endl;
    return false;
  }

  std::string bootdir =
      std::filesystem::path(m_Data->systemImage()).parent_path() / "boot";

  if (!std::filesystem::exists(bootdir)) {
    m_Mesg = "InternalError, '" + bootdir + "' boot dir is missing";
    ERROR << m_Mesg << std::endl;
    return false;
  }

  auto [status, output] = Exec::output(
      ("cp -a " + bootdir + " " + m_Data->workDir() + "/boot").c_str());
  if (status != 0) {
    m_Mesg = "Failed to install kernel and drivers, " + output;
    ERROR << m_Mesg << std::endl;
    return false;
  }

  m_Mesg = "Installed kernel successfully";
  LOG << m_Mesg << std::endl;
  return true;
}