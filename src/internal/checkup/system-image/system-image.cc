#include "system-image.hh"

#include <filesystem>

#include "../../../logging.hh"
#include "config.h"
bool SystemImageCheckup::process() {
  if (std::filesystem::exists(ISO_PATH "rootfs.img")) {
    m_Data->systemImage(ISO_PATH "rootfs.img");
    m_Mesg = "found system image at '" ISO_PATH
             "rootfs.img"
             "'";
    LOG << m_Mesg << std::endl;
    return true;
  }

  m_Mesg = "no system image found, internal error";
  ERROR << m_Mesg << std::endl;
  return false;
}