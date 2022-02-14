#include "disk.hh"

#include <filesystem>
#include <iostream>

#include "../utils/exec.hh"

namespace fs = std::filesystem;

Disk::Disk(std::string const& path) : m_Path(path) {
  m_Label = std::get<1>(Exec::output(("lsblk -dno LABEL " + m_Path).c_str()));
  m_UUID = std::get<1>(Exec::output(("lsblk -dno UUID " + m_Path).c_str()));
  m_Size = std::get<1>(Exec::output(("lsblk -dno SIZE " + m_Path).c_str()));
}

std::vector<Disk> Disk::list() {
  std::vector<Disk> disks;
  fs::current_path("/dev/disk/by-partuuid");
  for (auto const& i : fs::directory_iterator(".")) {
    auto disk_path = fs::canonical(fs::absolute(fs::read_symlink(i.path())));
    disks.push_back(Disk(disk_path.string()));
  }

  return disks;
}