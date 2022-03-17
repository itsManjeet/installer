#include "disk.hh"

#include <filesystem>
#include <iostream>

#include "../utils/exec.hh"
#define DISK_PART_UUID_PATH "/dev/disk/by-partuuid"
namespace fs = std::filesystem;

Disk::Disk(std::string const& path) : m_Path(path) {
  m_Label = std::get<1>(Exec::output(("lsblk -dno LABEL " + m_Path).c_str()));
  m_UUID = std::get<1>(Exec::output(("lsblk -dno UUID " + m_Path).c_str()));
  m_Size = std::get<1>(Exec::output(("lsblk -dno SIZE " + m_Path).c_str()));
}

std::vector<Disk> Disk::list() {
  std::vector<Disk> disks;
  if (!std::filesystem::exists(DISK_PART_UUID_PATH)) {
    system("gparted");
  }

  try {
    fs::current_path(DISK_PART_UUID_PATH);
    for (auto const& i : fs::directory_iterator(".")) {
      auto disk_path = fs::canonical(fs::absolute(fs::read_symlink(i.path())));
      disks.push_back(Disk(disk_path.string()));
    }
  } catch (std::exception const& exc) {
    std::cerr << "Failed to read disk paritition, no " DISK_PART_UUID_PATH
                 " exists"
              << std::endl;
  }

  return disks;
}