#include "checkup.hh"

#include "disk/disk.hh"
#include "efi/efi.hh"
#include "memory/memory.hh"
#include "permission/permission.hh"
#include "search-boot/search-boot.hh"
#include "system-image/system-image.hh"
#include "system-info/system-info.hh"

std::vector<std::shared_ptr<Process>> Checkup::list(InstallData* data) {
  std::vector<std::shared_ptr<Process>> res;

  res.push_back(std::make_shared<PermissionCheckup>(data));
  res.push_back(std::make_shared<SystemImageCheckup>(data));
  res.push_back(std::make_shared<SystemInfoCheckup>(data));
  res.push_back(std::make_shared<MemoryCheckup>(data));
  res.push_back(std::make_shared<DiskCheckup>(data));
  res.push_back(std::make_shared<EfiCheckup>(data));
  res.push_back(std::make_shared<SearchBootCheckup>(data));

  return res;
}