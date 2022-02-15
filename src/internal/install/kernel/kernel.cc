#include "kernel.hh"

#include <filesystem>

#include "../../../utils/exec.hh"

bool Kernel::process() {
  if (!std::filesystem::exists(m_Data->workDir())) {
    m_Mesg = "InternalError, '" + m_Data->workDir() + "' dir is missing";
    return false;
  }

  std::string bootdir =
      std::filesystem::path(m_Data->systemImage()).parent_path() / "boot";

  if (!std::filesystem::exists(bootdir)) {
    m_Mesg = "InternalError, '" + bootdir + "' boot dir is missing";
    return false;
  }

  auto [status, output] = Exec::output(
      ("cp -a " + bootdir + " " + m_Data->workDir() + "/boot").c_str());
  if (status != 0) {
    m_Mesg = "Failed to install kernel and drivers, " + output;
    return false;
  }

  m_Mesg = "Installed kernel successfully";
  return true;
}