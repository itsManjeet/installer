#include "disk.hh"

#include <sys/mount.h>
#include <fcntl.h>
#include <linux/fs.h>
#include <string.h>
#include <sys/ioctl.h>
#include <unistd.h>

#include <filesystem>

#include "../../../utils/humanize.hh"
#include "config.h"

bool DiskCheckup::process() {
  std::string disk = m_Data->disk();

  int fd = open(disk.c_str(), O_RDONLY);
  if (fd == 0) {
    m_Mesg = "Failed to load " + disk + ", " + std::string(strerror(errno));
    return false;
  }
  size_t size;
  if (ioctl(fd, BLKGETSIZE64, &size) == -1) {
    m_Mesg =
        "Failed to read size of " + disk + ", " + std::string(strerror(errno));
    return false;
  }

  close(fd);

  if (size < MINIMUM_DISKSPACE) {
    m_Mesg =
        "Need atleast " + humanize(MINIMUM_DISKSPACE) + " on root partition";
    return false;
  }

  umount(m_Data->disk().c_str());

  m_Mesg = "Selected device '" + m_Data->disk() + "' for root partition has " +
           humanize(size) + " space";
  return true;
}