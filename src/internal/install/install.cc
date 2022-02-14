#include "install.hh"

#include "bootloader/bootloader.hh"
#include "kernel/kernel.hh"
#include "prepare/prepare.hh"
#include "system-image/system-image.hh"

std::vector<std::shared_ptr<Process>> Install::list(InstallData* data) {
  std::vector<std::shared_ptr<Process>> res;

  res.push_back(std::make_shared<Prepare>(data));
  res.push_back(std::make_shared<SystemImage>(data));
  res.push_back(std::make_shared<Kernel>(data));
  res.push_back(std::make_shared<Bootloader>(data));

  return res;
}