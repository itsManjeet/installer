#include "efi.hh"

#include <filesystem>

#include "../../../logging.hh"

bool EfiCheckup::process() {
  m_Data->isEfi(std::filesystem::exists("/sys/firmware/efi/efivars"));

  if (m_Data->isEfi()) {
    m_Mesg = "efi supported";
  } else {
    m_Mesg = "legcay bootloader";
  }

  LOG << m_Mesg << std::endl;
  return true;
}