#include "search-boot.hh"

#include <filesystem>

#include "../../../utils/exec.hh"

bool SearchBootCheckup::process() {
  if (m_Data->isEfi()) {
    auto [status, output] = Exec::output(
        "lsblk -npo PATH,PARTTYPENAME | grep 'EFI System' | awk '{print $1}' | "
        "head -n1");

    if (status != 0) {
      m_Mesg = "failed to detect EFI partition, " + output;
      return false;
    }

    if (output.length() == 0) {
      m_Mesg =
          "no EFI partition detected, please make sure it has PartTypeName as "
          "'EFI System'";
      return false;
    }
    if (std::filesystem::exists(output)) {
      m_Mesg = "Found EFI partition '" + output + "'";
      m_Data->bootDevice(output);

      return true;
    } else {
      m_Mesg =
          "InternalError, detected EFI partition '" + output + "' not exists";
      return false;
    }
  } else {
    std::string disk = m_Data->disk();
    while (isdigit(disk.back())) {
      disk.pop_back();
    }

    if (std::filesystem::exists(disk)) {
      m_Mesg = "Using '" + disk + "' for bootloader";
      m_Data->bootDevice(disk);
      return true;
    } else {
      m_Mesg = "InternalError! detected Primary disk not exists";
      return false;
    }
  }
}