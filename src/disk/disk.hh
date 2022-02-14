#ifndef DISK_HH
#define DISK_HH

#include <string>
#include <vector>

class Disk {
 private:
  std::string m_Label;
  std::string m_Path;
  std::string m_Size;
  std::string m_UUID;

 public:
  Disk(std::string const& path);

  std::string const& label() const { return m_Label; }
  std::string const& path() const { return m_Path; }
  std::string const& size() const { return m_Size; }
  std::string const& uuid() const { return m_UUID; }
  static std::vector<Disk> list();
};

#endif