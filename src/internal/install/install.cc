#include "install.hh"

#include "bootloader/bootloader.hh"
#include "kernel/kernel.hh"
#include "prepare/prepare.hh"
#include "system-image/system-image.hh"

std::vector<std::shared_ptr<Process>> Install::list(InstallData* data) {
  std::vector<std::shared_ptr<Process>> res = {
      std::make_shared<Prepare>(data),
      std::make_shared<SystemImage>(data),
      std::make_shared<Kernel>(data),
      std::make_shared<Bootloader>(data),
  };

  return res;
}