#include "prepare.hh"

#include <sys/mount.h>
#include <unistd.h>

#include <filesystem>

#include "../../../logging.hh"
#include "../../../utils/exec.hh"
#include "../../../utils/temp.hh"

bool Prepare::process() {
  LOG << "preparing process" << std::endl;
  if (!std::filesystem::exists(m_Data->disk())) {
    m_Mesg = "Target disk '" + m_Data->disk() + "' not exists";
    ERROR << m_Mesg << std::endl;
    return false;
  }

  auto [status, output] =
      Exec::output(("mkfs.ext4 -F " + m_Data->disk()).c_str());
  if (status != 0) {
    m_Mesg = "Failed to format target disk '" + m_Data->disk() + "', " + output;
    ERROR << m_Mesg << std::endl;
    return false;
  }

  try {
    m_Data->workDir(tempdir("/tmp/installer"));
  } catch (std::exception const& exc) {
    m_Mesg = "Failed to prepare workdir " + std::string(exc.what());
    ERROR << m_Mesg << std::endl;
    return false;
  }

  if (int status = mount(m_Data->disk().c_str(), m_Data->workDir().c_str(),
                         "ext4", 0, NULL);
      status != 0) {
    m_Mesg = "Failed to mount target disk '" + m_Data->disk() + "', " +
             std::string(strerror(errno));
    ERROR << m_Mesg << std::endl;
    return false;
  }
  m_Mesg = "Prepared target disk at '" + m_Data->workDir() + "'";
  LOG << m_Mesg << std::endl;
  return true;
}