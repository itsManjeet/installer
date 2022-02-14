#include "efi.hh"

#include <filesystem>

bool EfiCheckup::process() {
  m_Data->isEfi(std::filesystem::exists("/sys/firmware/efi/efivars"));

  if (m_Data->isEfi()) {
    m_Mesg = "efi supported";
  } else {
    m_Mesg = "legcay bootloader";
  }

  return true;
}