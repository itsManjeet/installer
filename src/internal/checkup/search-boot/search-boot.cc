#include "search-boot.hh"

#include <filesystem>

#include "../../../logging.hh"
#include "../../../utils/exec.hh"

bool SearchBootCheckup::process() {
  LOG << "searching boot" << std::endl;
  if (m_Data->isEfi()) {
    LOG << "Boot mode : efi" << std::endl;
    auto [status, output] = Exec::output(
        "lsblk -npo PATH,PARTTYPENAME | grep 'EFI System' | awk '{print $1}'");

    if (status != 0) {
      m_Mesg = "failed to detect EFI partition, " + output;
      ERROR << m_Mesg << std::endl;
      return false;
    }

    std::vector<std::string> efi_s;
    std::string line;
    std::stringstream ss(output);
    while (std::getline(ss, line)) {
      efi_s.push_back(line);
    }

    if (efi_s.size() == 0) {
      m_Mesg =
          "no EFI partition detected, please make sure it has PartTypeName as "
          "'EFI System'";
      ERROR << m_Mesg << std::endl;
      return false;
    }

    if (efi_s.size() > 1) {
      std::string efis;
      for (auto const& i : efi_s) {
        efis += " " + i;
      }

      auto [status, output_] = Exec::output(
          ("zenity --list --column=EFI --text='Select Boot Parition' " + efis)
              .c_str());

      if (status != 0) {
        m_Mesg = "failed to load EFI paritition list";
        ERROR << m_Mesg << std::endl;
        return false;
      }

      output = output_;
    }

    if (std::filesystem::exists(output)) {
      m_Mesg = "Found EFI partition '" + output + "'";
      LOG << m_Mesg << std::endl;
      m_Data->bootDevice(output);

      return true;
    } else {
      m_Mesg =
          "InternalError, detected EFI partition '" + output + "' not exists";
      ERROR << m_Mesg << std::endl;
      return false;
    }
  } else {
    LOG << "Boot mode : legacy" << std::endl;
    std::string disk = m_Data->disk();
    while (isdigit(disk.back())) {
      disk.pop_back();
    }

    if (std::filesystem::exists(disk)) {
      m_Mesg = "Using '" + disk + "' for bootloader";
      LOG << m_Mesg << std::endl;
      m_Data->bootDevice(disk);
      return true;
    } else {
      m_Mesg = "InternalError! detected Primary disk not exists";
      ERROR << m_Mesg << std::endl;
      return false;
    }
  }
}