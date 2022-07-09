#include "disk.hh"

#include <filesystem>
#include <iostream>

#include "../logging.hh"
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

  LOG << "listing avaliable partitions" << std::endl;
  if (!std::filesystem::exists(DISK_PART_UUID_PATH)) {
    LOG << "no parition found, starting garted" << std::endl;
    system("gparted");
  }

  try {
    fs::current_path(DISK_PART_UUID_PATH);
    for (auto const& i : fs::directory_iterator(".")) {
      auto disk_path = fs::canonical(fs::absolute(fs::read_symlink(i.path())));
      LOG << "found " << disk_path << std::endl;
      disks.push_back(Disk(disk_path.string()));
    }
  } catch (std::exception const& exc) {
    ERROR << "Failed to read disk paritition, no " DISK_PART_UUID_PATH " exists"
          << std::endl;
  }

  return disks;
}