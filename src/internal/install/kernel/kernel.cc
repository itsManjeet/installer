#include "kernel.hh"

#include <filesystem>

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

  std::error_code err;
  std::filesystem::copy(bootdir, m_Data->workDir() + "/boot", err);
  if (err) {
    m_Mesg = "Failed to install kernel and drivers, " + err.message();
    return false;
  }

  m_Mesg = "Installed kernel successfully";
  return true;
}