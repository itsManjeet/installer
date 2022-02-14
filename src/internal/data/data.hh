#ifndef INSTALL_DATA_HH
#define INSTALL_DATA_HH

#include <string>
#include <vector>

class InstallData {
 private:
  std::string m_Disk;
  bool m_IsEfi;
  std::string m_BootDevice;
  std::string m_SystemImage;
  std::string m_Version;
  std::string m_WorkDir;

 public:
  std::string const& disk() const { return m_Disk; }
  void disk(std::string const& d) { m_Disk = d; }

  bool isEfi() const { return m_IsEfi; }
  void isEfi(bool i) { m_IsEfi = i; }

  std::string const& bootDevice() const { return m_BootDevice; }
  void bootDevice(std::string const& boot) { m_BootDevice = boot; }

  std::string const& systemImage() const { return m_SystemImage; }
  void systemImage(std::string const& img) { m_SystemImage = img; }

  std::string const& version() const { return m_Version; }
  void version(std::string const& ver) { m_Version = ver; }

  std::string const& workDir() const { return m_WorkDir; }
  void workDir(std::string const& work) { m_WorkDir = work; }
};

#endif